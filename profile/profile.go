package profile

import (
	"github.com/SsrCoder/onedrive-sdk-golang/client"
	"github.com/SsrCoder/onedrive-sdk-golang/microsoft"
)

type Profile struct {
	client    *client.Client
	msProfile *microsoft.Profile
}

func Me(client *client.Client) (resp *Profile, err error) {
	url := "https://graph.microsoft.com/v1.0/me/"
	var res microsoft.ProfileResp
	err = client.Get(url).BindJSON(&res).Do()
	resp = &Profile{
		client:    client,
		msProfile: &res.Profile,
	}
	return
}

func (p *Profile) DisplayName() string {
	return p.msProfile.DisplayName
}

func (p *Profile) Email() string {
	return p.msProfile.Mail
}
