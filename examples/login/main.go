package main

import (
	"encoding/json"
	"fmt"
	onedrivesdk "github.com/SsrCoder/onedrive-sdk-golang"
	"io/ioutil"
)

func main() {
	clientID := "ead3c269-a606-4a52-93f6-934c550814d9"
	clientSecret := "C4Vi7C.t~Sv87~~xnT6msjH-mS3M1l6ohu"

	redirectUri := "http://localhost:8080"
	scopes := []string{
		"User.Read",
		"Files.ReadWrite.All",
		"offline_access",
	}

	client := onedrivesdk.NewClient(clientID, clientSecret, redirectUri, scopes)

	if err := client.Authenticate(); err != nil {
		panic(err)
	}

	profile, err := onedrivesdk.Me(client)
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", profile.DisplayName())
	fmt.Println("email:", profile.Email())

	bytes, err := json.Marshal(client)
	if err == nil {
		if err = ioutil.WriteFile("client.json", bytes, 0777); err != nil {
			panic(err)
		}
	}
}
