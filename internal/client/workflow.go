package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"terraform-provider-tessell/internal/models"
	"time"
)

func (c *Client) GetWorkflowStatus(taskId string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/workflows/%s", c.APIAddress, taskId), nil)
	if err != nil {
		return "", err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	workflow := models.Workflow{}
	err = json.Unmarshal(body, &workflow)
	if err != nil {
		return "", err
	}
	return workflow.Status, nil
}

func (c *Client) waitTillWorkflowCompleted(taskId string, waitTime int) (string, error) {
	loopCount := -5
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return "", err
	}
	sleepCycleDuration, err := time.ParseDuration("60s")
	if err != nil {
		return "", err
	}

	loops := waitTime / int(sleepCycleDuration.Seconds())

	for {
		workflowStatus, err := c.GetWorkflowStatus(taskId)
		if err != nil {
			return "", err
		}
		if workflowStatus != "PAUSED" && workflowStatus != "RUNNING" && workflowStatus != "SCHEDULED" {
			return workflowStatus, nil
		}
		loopCount = loopCount + 1
		if loopCount > loops {
			return workflowStatus, fmt.Errorf("TIMED_OUT")
		}
		if loopCount > 1 && loopCount < loops-2 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}

	}
}

func (c *Client) WaitTillWorkflowCompleted(taskId string, waitTime int, task string) error {
	workflowStatus, err := c.waitTillWorkflowCompleted(taskId, waitTime)
	if err != nil {
		return fmt.Errorf("operation timed out waiting for %s to complete; Last known Workflow Status: %s; Workflow ID: %s", task, workflowStatus, taskId)
	}
	if workflowStatus != "COMPLETED" {
		return fmt.Errorf("%s did not complete successfully; Workflow Status: %s; Workflow ID: %s", task, workflowStatus, taskId)
	}
	return nil
}
