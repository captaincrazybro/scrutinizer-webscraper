package main

import (
	//"fmt"
	"log"
	_ "strings"

	"github.com/silinternational/gitops"

	sw "github.com/captaincrazybro/scrutinizer-webscraper"
)

func init() {
	gitops.Init()
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
	//repos, _ := sw.FetchScrutinizerRepos()

	// weekly audit depicting which github and bitbucket repositories are not registered to scrutinizer
	//sz := compareRepos(repos)
	//fmt.Print(sz)
	// TODO: send email
	sw.SendEmail()

}
