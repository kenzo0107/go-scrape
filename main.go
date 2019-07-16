package main

import (
	"log"

	"github.com/kenzo0107/go-scrape/app"
)

const (
	inviteEmail = "hogemoge@example.com"
	role        = "st"
)

func main() {
	// NewRelic
	if err := app.InviteToNewrelic(inviteEmail); err != nil {
		log.Println(err)
	}
	log.Println("Success to invite to NewRelic")

	// // Datadog
	// if err := app.InviteToDatadog(inviteEmail, role); err != nil {
	// 	log.Println(err)
	// }
	// log.Println("Success to invite to Datadog")

	// // Rollbar
	// if err := app.InviteToRollbar(inviteEmail, teamID); err != nil {
	// 	log.Println(err)
	// }
	// log.Println("Success to invite to Datadog")
}
