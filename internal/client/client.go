package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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
		APIAddress: *apiAddress,
		Auth: AuthStruct{
			EmailId:  *emailId,
			Password: *password,
		},
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

func (c *Client) doRequest(req *http.Request) ([]byte, int, error) {
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
		return nil, res.StatusCode, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("%s", body)
	}

	return body, res.StatusCode, err
}
