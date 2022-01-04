package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"terraform-provider-tessell/internal/models"
)

func (c *Client) GetDatabaseByName(name string) (*models.Database, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/databases/%s", c.APIAddress, name), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	database := models.Database{}
	err = json.Unmarshal(body, &database)
	if err != nil {
		return nil, err
	}
	return &database, nil
}

func (c *Client) GetDatabases() (*[]models.Database, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/databases", c.APIAddress), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	databases := models.Databases{}
	err = json.Unmarshal(body, &databases)
	if err != nil {
		return nil, err
	}

	return &databases.Databases, nil
}

func (c *Client) CreateDatabase(payload models.DatabaseCreationPayload) (*models.DatabaseCreationResponse, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/databases", c.APIAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	databaseCreationResponse := models.DatabaseCreationResponse{}
	err = json.Unmarshal(body, &databaseCreationResponse)
	if err != nil {
		return nil, err
	}

	return &databaseCreationResponse, nil
}

func (c *Client) CloneDatabase(sourceDmmName string, payload models.DatabaseCloningPayload) (*models.DatabaseCreationResponse, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dmms/%s/data-management/clone", c.APIAddress, sourceDmmName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	databaseCreationResponse := models.DatabaseCreationResponse{}
	err = json.Unmarshal(body, &databaseCreationResponse)
	if err != nil {
		return nil, err
	}

	return &databaseCreationResponse, nil
}

func (c *Client) UpdateDatabase(name string, payload models.DatabaseUpdationPayload) (*models.DatabaseCreationResponse, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/databases/%s", c.APIAddress, name), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	databaseCreationResponse := models.DatabaseCreationResponse{}
	err = json.Unmarshal(body, &databaseCreationResponse)
	if err != nil {
		return nil, err
	}

	return &databaseCreationResponse, nil
}

func (c *Client) DeleteDatabase(name string, payload models.DatabaseDeletionPayload) (*models.DatabaseBackupDeletionResponse, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/databases/%s", c.APIAddress, name), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	databaseBackupDeletionResponse := models.DatabaseBackupDeletionResponse{}
	err = json.Unmarshal(body, &databaseBackupDeletionResponse)
	if err != nil {
		return nil, err
	}

	return &databaseBackupDeletionResponse, nil
}
