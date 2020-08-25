package onedrivesdk

import (
	"github.com/SsrCoder/onedrive-sdk-golang/client"
	"github.com/SsrCoder/onedrive-sdk-golang/filesystem"
	"github.com/SsrCoder/onedrive-sdk-golang/profile"
)

func NewClient(clientID, clientSecret, redirectUri string, scope []string) *client.Client {
	return client.New(clientID, clientSecret, redirectUri, scope)
}

func NewEmptyClient() *client.Client {
	return client.NewEmpty()
}

func NewFileSystem(client *client.Client) *filesystem.FileSystem {
	return filesystem.New(client)
}

func Me(client *client.Client) (*profile.Profile, error) {
	return profile.Me(client)
}
