package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	var result DemoteAndPromoteList
	isOldMember := make(map[string]bool)
	isColeaderCandidate := make(map[string]bool)

	loc, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Clock: ", time.Now().In(loc))
	scheduler := gocron.NewScheduler(loc)
	scheduler.Every(1).Sunday().At("22:00").Do(getDemoteAndPromote, isOldMember, isColeaderCandidate, &result)

	fmt.Println("Start Scheduler")
	scheduler.StartBlocking()
}

func getDemoteAndPromote(isOldMember map[string]bool, isColeaderCandidate map[string]bool, result *DemoteAndPromoteList) {
	clanMembersData := getClanMembersData("#URLPR", os.Getenv("clashRoyaleApiKey"))

	fmt.Printf("Reason: %s\n", clanMembersData.Reason)
	fmt.Printf("Message: %s\n", clanMembersData.Message)
	fmt.Printf("Items: %v\n", clanMembersData.Items)
	fmt.Printf("Count: %d\n", len(clanMembersData.Items))

	var demoteList []string
	var promoteList []string

	for _, clanMember := range clanMembersData.Items {
		if clanMember.Donations < 70 && isOldMember[clanMember.Tag] {
			demoteList = append(demoteList, clanMember.Name)
		} else if (clanMember.Donations >= 350 && clanMember.Role == "member") || (clanMember.Donations >= 1000 && isColeaderCandidate[clanMember.Tag]) {
			promoteList = append(promoteList, clanMember.Name)
		}
	}

	isColeaderCandidate = map[string]bool{}
	isOldMember = map[string]bool{}
	for _, clanMember := range clanMembersData.Items {
		isOldMember[clanMember.Tag] = true
		if clanMember.Donations >= 1000 {
			isColeaderCandidate[clanMember.Tag] = true
		}
	}

	result.DemoteList = demoteList
	result.PromoteList = promoteList
	result.UpdateTime = time.Now()

	fmt.Printf("DemoteList: %v\n", demoteList)
	fmt.Printf("PromoteList: %v\n", promoteList)
	fmt.Println(result.UpdateTime)
}

func getClanMembersData(clanTag string, ClashRoyaleApiKey string) ClanMembersData {
	client := &http.Client{}

	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.clashroyale.com/v1/clans/%s/members", url.QueryEscape(clanTag)), nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	request.Header.Add("Authorization", "Bearer "+ClashRoyaleApiKey)

	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var clanMembersData ClanMembersData

	json.Unmarshal(responseData, &clanMembersData)

	return clanMembersData
}
