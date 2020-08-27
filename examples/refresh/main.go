package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	onedrivesdk "github.com/SsrCoder/onedrive-sdk-golang"
)

func main() {
	client := onedrivesdk.NewEmptyClient()
	bytes, err := ioutil.ReadFile("client.json")
	if err != nil {
		panic(err)
	}
	client.SetProxy("ssrcoder.com:7890")
	err = json.Unmarshal(bytes, &client)
	if err != nil {
		panic(err)
	}
	err = client.Refresh()
	if err != nil {
		panic(err)
	}
	me, err := onedrivesdk.Me(client)
	if err != nil {
		panic(err)
	}
	fmt.Println("email:", me.Email())
}
