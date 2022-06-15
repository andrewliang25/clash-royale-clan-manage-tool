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
	isOldMember := make(map[string]bool)
	isColeaderCandidate := make(map[string]bool)

	loc, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Clock: ", time.Now().In(loc))
	scheduler := gocron.NewScheduler(loc)
	scheduler.Every(1).Sunday().At("22:00").Do(printDemoteAndPromote, isOldMember, isColeaderCandidate)

	fmt.Println("Start Scheduler")
	scheduler.StartBlocking()
}

func printDemoteAndPromote(isOldMember map[string]bool, isColeaderCandidate map[string]bool) {
	ClanMembersData := getClanMembersData("#URLPR", os.Getenv("clashRoyaleApiKey"))

	fmt.Printf("Reason: %s\n", ClanMembersData.Reason)
	fmt.Printf("Message: %s\n", ClanMembersData.Message)
	fmt.Printf("Items: %v\n", ClanMembersData.Items)
	fmt.Printf("Count: %d\n", len(ClanMembersData.Items))

	var DemoteList []string
	var PromoteList []string

	for _, ClanMember := range ClanMembersData.Items {
		if ClanMember.Donations < 70 && isOldMember[ClanMember.Tag] {
			DemoteList = append(DemoteList, ClanMember.Name)
		} else if (ClanMember.Donations >= 350 && ClanMember.Role == "member") || (ClanMember.Donations >= 1000 && isColeaderCandidate[ClanMember.Tag]) {
			PromoteList = append(PromoteList, ClanMember.Name)
		}
	}

	isColeaderCandidate = map[string]bool{}
	isOldMember = map[string]bool{}
	for _, ClanMember := range ClanMembersData.Items {
		isOldMember[ClanMember.Tag] = true
		if ClanMember.Donations >= 1000 {
			isColeaderCandidate[ClanMember.Tag] = true
		}
	}

	fmt.Printf("DemoteList: %v\n", DemoteList)
	fmt.Printf("PromoteList: %v\n", PromoteList)
}

func getClanMembersData(clanTag string, ClashRoyaleApiKey string) clanMembersData {
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

	var ClanMembersData clanMembersData

	json.Unmarshal(responseData, &ClanMembersData)

	return ClanMembersData
}
