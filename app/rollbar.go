package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kenzo0107/go-scrape/config"
)

const (
	rollbarAPIInviteTeam = "https://api.rollbar.com/api/1/team/%s/invites"
	rollbarAPIAllTeams   = "https://api.rollbar.com/api/1/teams?access_token=%s"
)

type rollbarInvite struct {
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
}

// Response ... rollbar api response
type Response struct {
	Message string `json:"message"`
	Result  Result `json:"result"`
	Err     int    `json:"err"`
}

// Result ... "result" in rollbar api response
type Result struct {
	ID         uint64 `json:"id"`
	FromUserID uint64 `json:"from_user_id"`
	TeamID     uint64 `json:"team_id"`
	ToEmail    string `json:"to_email"`
	Status     string `json:"status"`
}

type Team struct {
	ID          uint64 `json:"id"`
	AccountID   uint64 `json:"account_id"`
	Name        string `json:"name"`
	AccessLevel string `json:"access_level"`
}

type ResponseAllTeam struct {
	Err    uint64 `json:"err"`
	Result []Team `json:result`
}

// GetAllTeamOfRollbar ... Get All Teams of Rollbar
func GetAllTeamOfRollbar() ([]Team, error) {
	url := fmt.Sprintf(rollbarAPIAllTeams, config.Secrets.RollbarReadAccessToken)
	r, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", body)
	}

	response := ResponseAllTeam{}
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return response.Result, nil
}

// InviteToRollbar ... Invite user to Rollbar
func InviteToRollbar(inviteEmail, teamID string) error {
	rollbarInvite := &rollbarInvite{
		AccessToken: config.Secrets.RollbarWriteAccessToken,
		Email:       inviteEmail,
	}
	b, err := json.Marshal(rollbarInvite)
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}

	url := fmt.Sprintf(rollbarAPIInviteTeam, teamID)
	r, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", body)
	}

	response := Response{}
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
