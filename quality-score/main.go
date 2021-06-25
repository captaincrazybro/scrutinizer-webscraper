package main

import (
	"log"

	sw "github.com/captaincrazybro/scrutinizer-webscraper"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("Error: ")
}

func main() {
	// starts the lambda
	//fmt.Println("Lambda started!")
	//lambda.Start(handleSchedule)

	// testing purposes
	handleSchedule()
}

// handleSchedule function to call once every week
func handleSchedule() {
	_, avg := sw.FetchScrutinizerRepos()

	err := sw.SendDataToGA(avg)
	if err != nil {
		log.Println(err)
	}

	// TODO: send email

}
