package main

import (
	"encoding/json"
	"fmt"
	onedrivesdk "github.com/SsrCoder/onedrive-sdk-golang"
	"github.com/SsrCoder/onedrive-sdk-golang/filesystem"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("client.json")
	if err != nil {
		panic(err)
	}

	client := onedrivesdk.NewEmptyClient()
	if err := json.Unmarshal(bytes, &client); err != nil {
		panic(err)
	}

	fs := onedrivesdk.NewFileSystem(client)

	root, err := fs.Root()
	if err != nil {
		panic(err)
	}

	if !root.IsFolder() {
		return
	}

	root.Walk(func(file *filesystem.File) {
		fmt.Println(file.Path())
	})
}
