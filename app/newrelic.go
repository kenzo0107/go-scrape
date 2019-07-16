package app

import (
	"log"
	"strings"

	"github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/matchers"

	"github.com/kenzo0107/go-scrape/config"
)

var newrelicURLLogin string

func init() {
	teamID := config.Secrets.NewrelicTeamID
	newrelicURLLogin = "https://login.newrelic.com/login?return_to=https%3A%2F%2Faccount.newrelic.com%2Faccounts%2F" + teamID + "%2Fusers%2Fnew&account_id=" + teamID
}

func InviteToNewrelic(inviteEmail string) error {
	gomega.RegisterFailHandler(func(message string, callerSkip ...int) {
		log.Printf("Assertion Error %s", message)
	})

	d := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",
			"--disable-gpu",
			"--no-sandbox",
		}),
	)

	if err := d.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer d.Stop()

	p, err := d.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
		return err
	}

	// Access to Login page.
	if err := p.Navigate(newrelicURLLogin); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
		return err
	}

	log.Println("success to access login page")

	username := p.FindByID("login_email")
	password := p.FindByID("login_password")

	username.Fill(config.Secrets.NewrelicLoginEmail)
	password.Fill(config.Secrets.NewrelicLoginPassword)

	if err := p.FindByID("login").Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
		return err
	}

	log.Println("success to login")

	// wait for loading class dom
	gomega.Eventually(p.FindByClass("User-userInfoContainer"), "10s").Should(matchers.BeFound())

	name := p.Find("#user-management-ui-app > div:nth-child(2) > div.User > div.User-userInfoContainer > div:nth-child(1) > div:nth-child(2) > div > input")
	email := p.Find("#user-management-ui-app > div:nth-child(2) > div.User > div.User-userInfoContainer > div:nth-child(2) > div:nth-child(2) > div > input")

	v := strings.Split(inviteEmail, "@")
	inviteUsername := v[0]
	name.Fill(inviteUsername)
	email.Fill(inviteEmail)

	if err := p.FindByXPath("//*[@id=\"user-management-ui-app\"]/div[2]/div[2]/div[5]/div/button[2]").Click(); err != nil {
		log.Fatalf("Failed to add user:%v", err)
		return err
	}

	gomega.Eventually(p.FindByClass("Notification-Success"), "10s").Should(matchers.BeFound())

	log.Println("success to invite")

	p.Screenshot("c.png")

	return nil
}
