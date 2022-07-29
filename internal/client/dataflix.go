package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) GetDataflixByName(id string) (*model.TessellAmDataflixDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dataflix/%s", c.APIAddress, id), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	q.Add("id", fmt.Sprintf("%v", id))
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellAmDataflixDTO := model.TessellAmDataflixDTO{}
	err = json.Unmarshal(body, &tessellAmDataflixDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellAmDataflixDTO, statusCode, nil
}

func (c *Client) GetDataflixes(name string, loadAcls bool, owners []string) (*model.TessellDataflixResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dataflix", c.APIAddress), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	q.Add("name", fmt.Sprintf("%v", name))
	q.Add("load-acls", fmt.Sprintf("%v", loadAcls))
	q.Add("owners", strings.Join(owners, ","))
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellDataflixResponse := model.TessellDataflixResponse{}
	err = json.Unmarshal(body, &tessellDataflixResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellDataflixResponse, statusCode, nil
}

func (c *Client) DataflixPollForStatus(id string, status string, timeout int, interval int) error {
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

func (c *Client) DataflixPollForStatusCode(id string, statusCodeRequired int, timeout int, interval int) error {
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
