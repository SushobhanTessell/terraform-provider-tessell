package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"terraform-provider-tessell/internal/models"
)

func (c *Client) GetDatabaseBackupByName(backupName string, dmmName string) (*models.DatabaseBackup, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dmms/%s/availability-machine/backups/%s", c.APIAddress, dmmName, backupName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	databaseBackup := models.DatabaseBackup{}
	err = json.Unmarshal(body, &databaseBackup)
	if err != nil {
		return nil, err
	}
	return &databaseBackup, nil
}

func (c *Client) GetDatabaseBackups(dmmName string) (*models.DatabaseBackups, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dmms/%s/availability-machine/backups", c.APIAddress, dmmName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	databaseBackups := models.DatabaseBackups{}
	err = json.Unmarshal(body, &databaseBackups.DatabaseBackups)
	if err != nil {
		return nil, err
	}

	return &databaseBackups, nil
}

func (c *Client) CreateDatabaseBackup(dmmName string, payload models.DatabaseBackupCreationPayload) (*models.DatabaseBackupCreationResponse, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dmms/%s/availability-machine/backups", c.APIAddress, dmmName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	databaseBackupCreationResponse := models.DatabaseBackupCreationResponse{}
	err = json.Unmarshal(body, &databaseBackupCreationResponse)
	if err != nil {
		return nil, err
	}

	return &databaseBackupCreationResponse, nil
}

func (c *Client) DeleteDatabaseBackup(backupName string, dmmName string) (*models.DatabaseBackupDeletionResponse, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/dmms/%s/availability-machine/backups/%s", c.APIAddress, dmmName, backupName), nil)
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
