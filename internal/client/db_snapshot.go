package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) CreateTessellServiceBackupRequest(id string, payload model.CreateBackupTaskPayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/availability-machines/%s/backups", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	q.Add("id", fmt.Sprintf("%v", id))
	req.URL.RawQuery = q.Encode()

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) DeleteBackupRequest(availabilityMachineId string, id string) (*model.ApiStatus, int, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/availability-machines/%s/backups/%s", c.APIAddress, availabilityMachineId, id), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	q.Add("availabilityMachineId", fmt.Sprintf("%v", availabilityMachineId))
	q.Add("id", fmt.Sprintf("%v", id))
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	apiStatus := model.ApiStatus{}
	err = json.Unmarshal(body, &apiStatus)
	if err != nil {
		return nil, statusCode, err
	}

	return &apiStatus, statusCode, nil
}

func (c *Client) GetBackup(availabilityMachineId string, id string) (*model.TessellDmmDataflixBackupDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines/%s/backups/%s", c.APIAddress, availabilityMachineId, id), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	q.Add("availabilityMachineId", fmt.Sprintf("%v", availabilityMachineId))
	q.Add("id", fmt.Sprintf("%v", id))
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellDmmDataflixBackupDTO := model.TessellDmmDataflixBackupDTO{}
	err = json.Unmarshal(body, &tessellDmmDataflixBackupDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellDmmDataflixBackupDTO, statusCode, nil
}

func (c *Client) DBSnapshotPollForStatus(id string, status string, timeout int, interval int) error {
	//loopCount := -5
	loopCount := 0
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	sleepCycleDuration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return err
	}

	//loops := timeout / int(sleepCycleDuration.Seconds())
	loops := timeout/int(sleepCycleDuration.Seconds()) + 5

	for {
		response, _, err := c.GetTessellService(id)
		if err != nil {
			return err
		}
		switch *response.Status {
		case status:
			return nil
		case "FAILED":
			return fmt.Errorf("received status FAILED while polling")
		}

		loopCount = loopCount + 1
		if loopCount > loops {
			return fmt.Errorf("timed out with last seen status '%s'", *response.Status)
		}
		//if loopCount > 1 && loopCount < loops-2 {
		if loopCount > 6 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}
	}
}

func (c *Client) DBSnapshotPollForStatusCode(id string, statusCodeRequired int, timeout int, interval int) error {
	loopCount := -5
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	sleepCycleDuration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return err
	}

	loops := timeout / int(sleepCycleDuration.Seconds())

	for {
		_, statusCode, err := c.GetTessellService(id)
		if err != nil {
			if statusCode == statusCodeRequired {
				return nil
			}
			return fmt.Errorf("error while polling: %s", err.Error())
		}

		loopCount = loopCount + 1
		if loopCount > loops {
			return fmt.Errorf("timed out")
		}
		if loopCount > 1 && loopCount < loops-2 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}
	}
}
