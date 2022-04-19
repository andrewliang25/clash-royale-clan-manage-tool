package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	client := &http.Client{}

	clanTag := "#URLPR"
	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.clashroyale.com/v1/clans/%s/members", url.QueryEscape(clanTag)), nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	ClashRoyaleApiKey := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6ImQxZGRiMTQxLWZkYjMtNGVlZi05M2UxLTI0OTE4Mzk1ZTY2MCIsImlhdCI6MTY1MDM1MjIwNSwic3ViIjoiZGV2ZWxvcGVyL2YzNTQwYTY3LTFmOTAtMGQ5Yy1hMWYxLTdlYjBkYzBjYmI4OCIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyIzNC44My4xMTkuMjE4Il0sInR5cGUiOiJjbGllbnQifV19.Hhk4NLaMN2Dx9N69QabNGJ8DGfCooTS2cME9gGpUVQYYhyrwcNZHiyLpVxO2mUP7cMyH_6uL69UCgAh-Hgx5SA"
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

	// fmt.Println(string(responseData))

	type arena struct {
		Id   int
		Name string
	}

	type clanMemberInfo struct {
		Tag              string
		Name             string
		Role             string
		LastSeen         string
		ExpLevel         int
		Trophies         int
		Arena            arena
		ClanRank         int
		PreviousClanRank int
		Donations        int
		DonationsRecived int
		ClanChestPoints  int
	}

	type clanMembersData struct {
		Reason  string
		Message string
		Items   []clanMemberInfo
		Paging  interface{}
	}

	var ClanMembersData clanMembersData

	json.Unmarshal(responseData, &ClanMembersData)

	fmt.Printf("Reason: %s\n", ClanMembersData.Reason)
	fmt.Printf("Message: %s\n", ClanMembersData.Message)
	fmt.Printf("Items: %v\n", ClanMembersData.Items)
	fmt.Printf("Count: %d\n", len(ClanMembersData.Items))

	var DemoteList []string
	var PromoteList []string
	isColeaderCandidate := map[string]bool{}

	for _, ClanMember := range ClanMembersData.Items {
		if ClanMember.Donations < 70 {
			DemoteList = append(DemoteList, ClanMember.Name)
		} else if (ClanMember.Donations >= 350 && ClanMember.Role == "member") || (ClanMember.Donations >= 1000 && isColeaderCandidate[ClanMember.Name]) {
			PromoteList = append(PromoteList, ClanMember.Name)
		}
		delete(isColeaderCandidate, ClanMember.Name)
		if ClanMember.Donations >= 1000 {
			isColeaderCandidate[ClanMember.Name] = true
		}
	}

	fmt.Printf("DemoteList: %v\n", DemoteList)
	fmt.Printf("PromoteList: %v\n", PromoteList)

}
