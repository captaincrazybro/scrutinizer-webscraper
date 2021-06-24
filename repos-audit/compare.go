package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/silinternational/gitops"

	lu "github.com/captaincrazybro/literalutil"
	sw "github.com/captaincrazybro/scrutinizer-webscraper"
)

// compareRepos compares the repositories from github and bitbucket to the scrutinizer repos
// telling which ones are registered on scrutinizer
func compareRepos(rs []string) []gitops.Repository {
	ghProv := gitops.GetProvider(sw.GitHub)
	bbProv := gitops.GetProvider(sw.BitBucket)

	// converts scrutinizer repo strings to ScrutRepo struct
	rz := []sw.ScrutRepo{}
	for _, v := range rs {
		repo, err := sw.ScrutRepo{}.FromString(v)
		if err != nil {
			log.Fatalln(err)
		}
		rz = append(rz, repo)
	}

	// gets repos from github and puts them into an array
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

	// gets repos from bitbucket and puts them into an array
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

	// does the comparing
	repos := []gitops.Repository{}

	for _, v := range ghRepos {
		if !arrIncludesRepo(rz, v, ghProv) && !isIgnored(v, ghProv) {
			repos = append(repos, v)
		}
	}

	for _, v := range bbRepos {
		if !arrIncludesRepo(rz, v, bbProv) && !isIgnored(v, bbProv) {
			repos = append(repos, v)
		}
	}

	return repos
}

func arrIncludesRepo(a []sw.ScrutRepo, r gitops.Repository, p gitops.Provider) bool {
	for _, v := range a {
		provName := lu.STernary(p.GetName() == sw.GitHub, "g", "b")
		if v.Provider == provName && v.Owner == p.GetOwner() && v.Name == r.GetName() {
			return true
		}
	}
	return false
}

func isIgnored(r gitops.Repository, p gitops.Provider) bool {
	filePtr, err := os.Open(sw.IgnoredReposFileName)
	if err != nil {
		log.Fatalf("%s file doesn't exist\n", sw.IgnoredReposFileName)
	}

	rdr := bufio.NewReader(filePtr)
	line, _, fileError := rdr.ReadLine()
	prov := lu.STernary(p.GetName() == sw.GitHub, "g", "b")
	fmtR := fmt.Sprintf("/%s/%s/%s", prov, p.GetOwner(), r.GetName())

	for fileError == nil {
		s := string(line)

		if s == fmtR {
			return true
		}

		line, _, fileError = rdr.ReadLine()
	}

	return false
}
