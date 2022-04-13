package main

import (
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

	ClashRoyaleApiKey := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjQ2OTlhYzhkLTQ4YzAtNDE1Ni1iYWQ2LTM1NmVkODk1ODk1OCIsImlhdCI6MTY0OTgyOTUzNCwic3ViIjoiZGV2ZWxvcGVyL2YzNTQwYTY3LTFmOTAtMGQ5Yy1hMWYxLTdlYjBkYzBjYmI4OCIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyIzNC4xMzYuNjIuNDEiXSwidHlwZSI6ImNsaWVudCJ9XX0.Vno5yJ-I4pp0WkS1wh8aewylhKZjgsv-cHQ80IMr4dh0Wp7muyJ_GinJt2_pdMK4thFl_-BVEDAcKEKU0GrDlA"
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
	fmt.Println(string(responseData))

}
