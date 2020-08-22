package onedrivesdk

import "github.com/guonaihong/gout"

func (c *client) Me() (resp interface{}, err error) {
	err = gout.GET("https://graph.microsoft.com/v1.0/me/").SetHeader(c.defaultHeader()).BindJSON(&resp).Do()
	return
}
