package main

import (
	"fmt"
	_ "strings"

	"github.com/aws/aws-lambda-go/lambda"

	c "github.com/captaincrazybro/literalutil/console"

	"github.com/silinternational/gitops"

	sw "github.com/captaincrazybro/scrutinizer-webscraper"
)

func init() {
	gitops.Init()
	c.SetErrorPrefix("Error: ")
}

func main() {
	// starts the lambda
	fmt.Println("Lambda started!")
	lambda.Start(handleSchedule)

	// testing purposes
	//handleSchedule()
}

// handleSchedule function to call once every week
func handleSchedule() {
	repos, _ := sw.FetchScrutinizerRepos()

	// weekly audit depicting which github and bitbucket repositories are not registered to scrutinizer
	sz := compareRepos(repos)
	//fmt.Print(sz)
	uRepos := sz.Join("\n")

	data := `{"unregistered_repos":"` + uRepos + `"}`
	sw.SendEmail("ReposAuditTemplate", data.Tos())

}
