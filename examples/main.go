package main

import (
	"fmt"

	onedrivesdk "github.com/SsrCoder/onedrive-sdk-golang"
)

func main() {
	clientID := "**************"
	clientSecret := "**************"
	redirectUri := "http://localhost:8080"
	scopes := []string{
		"User.Read",
	}

	client := onedrivesdk.NewClient(clientID, clientSecret, redirectUri, scopes)
	if err := client.Authenticate(); err != nil {
		panic(err)
	}

	resp, err := client.Me()
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
