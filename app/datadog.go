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

const datadogAPIURLFormat = "https://api.datadoghq.com/api/v1/user?api_key=%s&application_key=%s"

// AccessRole ... Datadog AccessRole (ro|st|adm)
type AccessRole string

const (
	ReadOnly AccessRole = "ro"
	Standard            = "st"
	Admin               = "adm"
)

var AccessRoles [3]AccessRole

type datadogInvite struct {
	Handle     string     `json:"handle"`
	Name       string     `json:"email"`
	AccessRole AccessRole `json:"access_role"`
}

// ResponseDatadog ... rollbar api response
type ResponseDatadog struct {
	User   User     `json:"user"`
	Errors []string `json:"errors"`
}

// User ... "user" in datadog api response
type User struct {
	Disabled   bool   `json:"disabled"`
	Handle     string `json:"handle"`
	Name       string `json:"name"`
	IsAdmin    bool   `json:"is_admin"`
	Role       string `json:"role"`
	AccessRole string `json:"access_role"`
	Verified   bool   `json:"verified"`
	Email      string `json:"email"`
	Icon       string `json:"icon"`
}

func init() {
	AccessRoles = [3]AccessRole{ReadOnly, Standard, Admin}
}

// InviteToDatadog ... invite email to datadog
func InviteToDatadog(inviteEmail, accessRole string) error {
	datadogAPIURL := fmt.Sprintf(datadogAPIURLFormat, config.Secrets.DatadogAPIKey, config.Secrets.DatadogAppKey)

	datadogInvite := &datadogInvite{
		Handle:     inviteEmail,
		AccessRole: AccessRole(accessRole),
	}

	b, err := json.Marshal(datadogInvite)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	r, err := http.Post(
		datadogAPIURL,
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

	response := ResponseDatadog{}
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
