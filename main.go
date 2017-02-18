package main

import (
	"github.com/nlopes/slack"
	"log"
	"os"
)

func main() {
	api := slack.New(GetAPIToken())
	os.Exit(run(api))
}

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Launch slackbot")
			case *slack.MessageEvent:
				log.Printf("Message: %s\n", ev)
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello World!", ev.Channel))
			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1
			}
		}
	}
}
