package main

import (
	"log"
	"os"
)

const SlackAPITokenEnvKey = "SLACK_API_TOKEN"

func GetAPIToken() string {
	var env = os.Getenv(SlackAPITokenEnvKey)
	if env == "" {
		log.Print("SLACK_API_TOKEN is not found. Need to set as an environment variable.")
		return ""
	}
	return env
}
