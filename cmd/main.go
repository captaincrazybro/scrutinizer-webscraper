package main

import (
	"fmt"
	"gitops"
	"log"
	"strings"

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
	repos, avg := sw.FetchScrutinizerRepos()

	// weekly audit depicting which github and bitbucket repositories are not registered to scrutinizer
	sz := compareRepos(repos)
	fmt.Print(sz)

	// weekly average score report
	sw.SendEmail(repos, avg) //Doesn't work right now, this is just a template
	err := sw.SendDataToGA(avg)
	if err != nil {
		log.Fatalln(err)
	}
}

// parseStructRepoString removes the "/"s from either side of the string
func parseStructRepoString(s string) string {
	s = strings.TrimPrefix(s, "/")
	return strings.TrimSuffix(s, "/")
}

// compareRepos compares the repositories from github and bitbucket to the scrutinizer repos
// telling which ones are registered on scrutinizer
func compareRepos(rs []string) []gitops.Repository {
	ghProv := gitops.GetProvider(sw.GitHub)
	bbProv := gitops.GetProvider(sw.BitBucket)

	rz := []sw.ScrutRepo{}
	for _, v := range rs {
		repo, err := sw.ScrutRepo{}.FromString(v)
		if err != nil {
			log.Fatalln(err)
		}
		rz = append(rz, repo)
	}

	ghListedRepos, err := ghProv.ListRepos(1)
	if err != nil {
		log.Fatalln(err)
	}
	ghRepos := ghListedRepos

	for i := 2; len(ghListedRepos) != 0; i++ {
		ghListedRepos, err = ghProv.ListRepos(i)
		if err != nil {
			log.Fatalln(err)
		}
		ghRepos = append(ghRepos, ghListedRepos...)
	}

	bbListedRepos, err := bbProv.ListRepos(1)
	if err != nil {
		log.Fatalln(err)
	}
	bbRepos := bbListedRepos

	for i := 2; len(bbListedRepos) != 0; i++ {
		bbListedRepos, err = bbProv.ListRepos(i)
		if err != nil {
			log.Fatalln(err)
		}
		bbRepos = append(bbRepos, bbListedRepos...)
	}

	repos := []gitops.Repository{}

	for _, v := range ghRepos {
		if arrIncludesRepo(rz, v, ghProv) {
			repos = append(repos, v)
		}
	}

	for _, v := range bbRepos {
		if arrIncludesRepo(rz, v, bbProv) {
			repos = append(repos, v)
		}
	}

	return repos
}

func arrIncludesRepo(a []sw.ScrutRepo, r gitops.Repository, p gitops.Provider) bool {
	for _, v := range a {
		if v.Provider == p.GetName() && v.Owner == p.GetOwner() && v.Name == r.GetName() {
			return true
		}
	}
	return false
}
