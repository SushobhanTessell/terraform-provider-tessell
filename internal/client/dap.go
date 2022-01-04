package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"terraform-provider-tessell/internal/models"
)

func (c *Client) GetDAPByName(name string, dmmName string) (*models.DAP, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dmms/%s/availability-policies/%s", c.APIAddress, dmmName, name), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	dap := models.DAP{}
	err = json.Unmarshal(body, &dap)
	if err != nil {
		return nil, err
	}
	return &dap, nil
}

func (c *Client) GetDAPs(dmmName string) (*models.DAPs, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dmms/%s/availability-policies", c.APIAddress, dmmName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	daps := models.DAPs{}
	err = json.Unmarshal(body, &daps)
	if err != nil {
		return nil, err
	}

	return &daps, nil
}

func (c *Client) CreateDAP(dmmName string, payload models.DAPCreationUpdationPayload) (*models.DAP, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dmms/%s/availability-policies", c.APIAddress, dmmName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	dap := models.DAP{}
	err = json.Unmarshal(body, &dap)
	if err != nil {
		return nil, err
	}

	return &dap, nil
}

func (c *Client) UpdateDAP(name string, dmmName string, payload models.DAPCreationUpdationPayload) (*models.DAP, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/dmms/%s/availability-policies/%s", c.APIAddress, dmmName, name), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	dap := models.DAP{}
	err = json.Unmarshal(body, &dap)
	if err != nil {
		return nil, err
	}

	return &dap, nil
}

func (c *Client) DeleteDAP(name string, dmmName string) (*models.APIStatus, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/dmms/%s/availability-policies/%s", c.APIAddress, dmmName, name), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	dapDeletionResponse := models.APIStatus{}
	err = json.Unmarshal(body, &dapDeletionResponse)
	if err != nil {
		return nil, err
	}

	return &dapDeletionResponse, nil
}
