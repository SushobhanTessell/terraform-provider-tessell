package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const DefaultAPIAddress string = "https://api.dev.tessell-stage.cloud"

type Client struct {
	APIAddress          string
	HTTPClient          *http.Client
	AuthorizationToken  string
	AuthenticationToken string
	TenantId            string
	Auth                AuthStruct
}

func NewClient(apiAddress *string, emailId *string, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		APIAddress: DefaultAPIAddress,
		Auth: AuthStruct{
			EmailId:  *emailId,
			Password: *password,
		},
	}

	if apiAddress != nil {
		c.APIAddress = *apiAddress
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.AuthorizationToken = ar.AccessToken
	c.AuthenticationToken = ar.IdToken
	c.TenantId = ar.Tenant[0].TenantId

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.AuthorizationToken != "" {
		req.Header.Set("Authorization", c.AuthorizationToken)
	}
	if c.AuthenticationToken != "" {
		req.Header.Set("Authentication", c.AuthenticationToken)
	}
	if c.TenantId != "" {
		req.Header.Set("tenant-id", c.TenantId)
	}
	if req.Method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", body)
	}

	return body, err
}
