package onedrivesdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/SsrCoder/onedrive-sdk-golang/consts"
	"github.com/SsrCoder/onedrive-sdk-golang/utils"
	"github.com/guonaihong/gout"
)

type client struct {
	ClientID     string   `json:"client_id"`
	Scope        []string `json:"scope"`
	ClientSecret string   `json:"client_secret"`
	RedirectUri  string   `json:"redirect_uri"`
	Code         string   `json:"code"`
	Token        token    `json:"token"`
}

type token struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewClient(clientID, clientSecret, redirectUri string, scope []string) *client {
	return &client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectUri:  redirectUri,
		Scope:        scope,
	}
}

func (c *client) AddScope(scope string) {
	c.Scope = append(c.Scope, scope)
}

func (c *client) GetAuthUrl() string {
	param := url.Values{}
	param.Add("client_id", c.ClientID)
	param.Add("response_type", "code")
	param.Add("redirect_uri", c.RedirectUri)
	param.Add("scope", strings.Join(c.Scope, " "))
	return fmt.Sprintf("%s?%s", consts.MsaAuthServerUrl, param.Encode())
}

func (c *client) GetAuthCode() string {
	if len(c.Code) != 0 {
		return c.Code
	}

	defer func() {
		_ = recover()
	}()

	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		c.Code = req.URL.Query().Get("code")
		fmt.Fprint(writer, "<script>window.close();</script>")
		_ = server.Shutdown(context.TODO())
	})

	if err := utils.Open(c.GetAuthUrl()); err != nil {
		panic(err)
	}
	_ = server.ListenAndServe()

	return c.Code
}

func (c *client) Authenticate() error {
	param := map[string]interface{}{
		"client_id":     c.ClientID,
		"redirect_uri":  c.RedirectUri,
		"client_secret": c.ClientSecret,
		"code":          c.GetAuthCode(),
		"grant_type":    "authorization_code",
	}

	return gout.POST(consts.MsaAuthTokenUrl).SetWWWForm(param).BindJSON(&c.Token).Do()
}

func (c *client) defaultHeader() map[string]interface{} {
	return map[string]interface{}{
		"Authorization": fmt.Sprintf("bearer %s", c.Token.AccessToken),
	}
}
