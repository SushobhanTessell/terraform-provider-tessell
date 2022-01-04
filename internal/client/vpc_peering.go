package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"terraform-provider-tessell/internal/models"
)

func (c *Client) GetVPCPeering(name string, vpcName string, cloudType string, region string, subscriptionName string) (*models.VPCPeering, error) {
	params := url.Values{}
	params.Add("cloudType", cloudType)
	params.Add("region", region)
	params.Add("subscriptionName", subscriptionName)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/network/governance/vpc/%s/vpc-peerings/%s?%s", c.APIAddress, vpcName, name, params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	vpcPeering := models.VPCPeering{}
	err = json.Unmarshal(body, &vpcPeering)
	if err != nil {
		return nil, err
	}
	return &vpcPeering, nil
}

func (c *Client) GetVPCPeerings(vpcName string, cloudType string, region string, subscriptionName string) (*models.VPCPeerings, error) {
	params := url.Values{}
	params.Add("cloudType", cloudType)
	params.Add("region", region)
	params.Add("subscriptionName", subscriptionName)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/network/governance/vpc/%s/vpc-peerings?%s", c.APIAddress, vpcName, params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	vpcPeerings := models.VPCPeerings{}
	err = json.Unmarshal(body, &vpcPeerings.VPCPeerings)
	if err != nil {
		return nil, err
	}

	return &vpcPeerings, nil
}

func (c *Client) CreateVPCPeering(vpcName string, payload models.VPCPeeringCreationPayload) (*models.VPCPeering, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/network/governance/vpc/%s/vpc-peerings", c.APIAddress, vpcName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	vpcPeering := models.VPCPeering{}
	err = json.Unmarshal(body, &vpcPeering)
	if err != nil {
		return nil, err
	}

	return &vpcPeering, nil
}

func (c *Client) DeleteVPCPeering(name string, vpcName string, cloudType string, region string, subscriptionName string) (*models.APIStatus, error) {
	params := url.Values{}
	params.Add("cloudType", cloudType)
	params.Add("region", region)
	params.Add("subscriptionName", subscriptionName)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/network/governance/vpc/%s/vpc-peerings/%s?%s", c.APIAddress, vpcName, name, params.Encode()), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	vpcPeeringDeletionResponse := models.APIStatus{}
	err = json.Unmarshal(body, &vpcPeeringDeletionResponse)
	if err != nil {
		return nil, err
	}

	return &vpcPeeringDeletionResponse, nil
}
