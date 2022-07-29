package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AuthStruct struct {
	EmailId  string `json:"emailId"`
	Password string `json:"password"`
}

type AuthResponse struct {
	EmailId     string `json:"emailId"`
	AccessToken string `json:"accessToken"`
	IdToken     string `json:"idToken"`
	Tenant      []struct {
		TenantId string `json:"tenant"`
	} `json:"tenantUserAttributes"`
}

func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.EmailId == "" || c.Auth.Password == "" || c.APIAddress == "" {
		return nil, fmt.Errorf("all of 'email_id', 'password' and 'api_address' are required")
	}
	rb, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/iam/users/login", c.APIAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) SignOut() error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signout", c.APIAddress), strings.NewReader(string("")))
	if err != nil {
		return err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Signed out user" {
		return errors.New(string(body))
	}

	return nil
}
