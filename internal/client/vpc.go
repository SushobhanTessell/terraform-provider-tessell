package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"terraform-provider-tessell/internal/models"
)

func (c *Client) GetVPC(name string, cloudType string, region string, subscriptionName string) (*models.VPC, error) {
	params := url.Values{}
	params.Add("cloudType", cloudType)
	params.Add("region", region)
	params.Add("subscriptionName", subscriptionName)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/network/governance/vpcs/%s?%s", c.APIAddress, name, params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	vpc := models.VPC{}
	err = json.Unmarshal(body, &vpc)
	if err != nil {
		return nil, err
	}
	return &vpc, nil
}

func (c *Client) GetVPCs() (*models.VPCs, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/network/governance/vpcs", c.APIAddress), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	vpcs := models.VPCs{}
	err = json.Unmarshal(body, &vpcs.VPCs)
	if err != nil {
		return nil, err
	}

	return &vpcs, nil
}

func (c *Client) CreateVPC(payload models.VPCCreationPayload) (*models.VPC, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/network/governance/vpcs", c.APIAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	vpc := models.VPC{}
	err = json.Unmarshal(body, &vpc)
	if err != nil {
		return nil, err
	}

	return &vpc, nil
}

func (c *Client) UpdateVPC(name string) (*models.VPC, error) {
	return nil, fmt.Errorf("VPC updation is currently not supported")
}

func (c *Client) DeleteVPC(name string) (*models.VPCDeletionResponse, error) {
	return nil, fmt.Errorf("VPC deletion is currently not supported")
}
