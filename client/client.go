package client

import (
	"context"
	"fmt"
	"github.com/SsrCoder/onedrive-sdk-golang/consts"
	"github.com/SsrCoder/onedrive-sdk-golang/utils"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	ClientID     string   `json:"client_id"`
	Scope        []string `json:"scope"`
	ClientSecret string   `json:"client_secret"`
	RedirectUri  string   `json:"redirect_uri"`
	Code         string   `json:"code"`
	Token        token    `json:"token"`
	ProxyURL     string   `json:"proxy_url"`
}

func New(clientID, clientSecret, redirectUri string, scope []string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectUri:  redirectUri,
		Scope:        scope,
	}
}

func NewEmpty() *Client {
	return &Client{}
}

func (c *Client) AddScope(scope string) {
	c.Scope = append(c.Scope, scope)
}

func (c *Client) SetProxy(proxy string) {
	c.ProxyURL = proxy
}

func (c *Client) GetAuthUrl() string {
	param := url.Values{}
	param.Add("client_id", c.ClientID)
	param.Add("response_type", "code")
	param.Add("redirect_uri", c.RedirectUri)
	param.Add("scope", strings.Join(c.Scope, " "))
	return fmt.Sprintf("%s?%s", consts.MsaAuthServerUrl, param.Encode())
}

func (c *Client) GetAuthCode() string {
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
		_, _ = fmt.Fprint(writer, "<script>window.close();</script>")
		_ = server.Shutdown(context.TODO())
	})

	if err := utils.Open(c.GetAuthUrl()); err != nil {
		panic(err)
	}
	_ = server.ListenAndServe()

	return c.Code
}

func (c *Client) Authenticate() error {
	param := map[string]interface{}{
		"client_id":     c.ClientID,
		"redirect_uri":  c.RedirectUri,
		"client_secret": c.ClientSecret,
		"code":          c.GetAuthCode(),
		"grant_type":    "authorization_code",
	}

	return c.PostWWWForm(consts.MsaAuthTokenUrl, param).BindJSON(&c.Token).Do()
}

func (c *Client) defaultHeader() map[string]interface{} {
	header := make(map[string]interface{})
	if len(c.Token.AccessToken) != 0 {
		header["Authorization"] = fmt.Sprintf("bearer %s", c.Token.AccessToken)
	}
	return header
}

func (c *Client) Get(url string) *dataflow.DataFlow {
	g := gout.GET(url)
	if len(c.ProxyURL) != 0 {
		g = g.SetProxy(c.ProxyURL)
	}
	g = g.SetHeader(c.defaultHeader())
	return g
}

func (c *Client) PostJson(url string, data interface{}) *dataflow.DataFlow {
	g := gout.GET(url)
	if len(c.ProxyURL) != 0 {
		g = g.SetProxy(c.ProxyURL)
	}
	g = g.SetHeader(c.defaultHeader())
	g = g.SetJSON(data)
	return g
}

func (c *Client) PostForm(url string, data interface{}) *dataflow.DataFlow {
	g := gout.GET(url)
	if len(c.ProxyURL) != 0 {
		g = g.SetProxy(c.ProxyURL)
	}
	g = g.SetHeader(c.defaultHeader())
	g = g.SetForm(data)
	return g
}

func (c *Client) PostWWWForm(url string, data interface{}) *dataflow.DataFlow {
	g := gout.GET(url)
	if len(c.ProxyURL) != 0 {
		g = g.SetProxy(c.ProxyURL)
	}
	g = g.SetHeader(c.defaultHeader())
	g = g.SetWWWForm(data)
	return g
}
