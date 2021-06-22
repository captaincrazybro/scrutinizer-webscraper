package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	sw "github.com/captaincrazybro/scrutinizer-webscraper"

	"github.com/gocolly/colly"

	lu "github.com/captaincrazybro/literalutil"
)

var repos = lu.Array{}

func main() {
	fmt.Println("Scrutinizer Webscraper started!")
	// TODO add support for lambda
	sessionId := sw.Login()
	Sz := getRepos(sessionId)
	fmt.Println(Sz)
}

func getRepos(token lu.String) []lu.String {
	// creates a new collector
	c := colly.NewCollector()
	c.SetCookies(sw.ReposPageURL, []*http.Cookie{{
		Name:    "SESS",
		Value:   token.ToS(),
		Expires: time.Time{},
	}})

	// handle html
	c.OnHTML("a[title]", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("title"), sw.BBOrgName+"/") || strings.HasPrefix(e.Attr("title"), sw.GHOrgName+"/") {
			repos = append(repos, lu.String(e.Attr("title")).Split("/")[1])
		}
	})

	// starts scraping
	err := c.Visit(sw.ReposPageURL)
	if err != nil {
		return nil
	}

	// waits util done adding to repos variable
	A := waitUtilReposIsSet()
	A.Sort(func(a, b interface{}) int {
		if (a).(lu.String) > (b).(lu.String) {
			return 1
		} else if (b).(lu.String) > (a).(lu.String) {
			return -1
		} else {
			return 0
		}
	})

	Sz := []lu.String{}
	for _, v := range A {
		Sz = append(Sz, (v).(lu.String))
	}

	return Sz
}

// waitUtilReposIsSet waits until the repos variable is set
func waitUtilReposIsSet() lu.Array {
	for i := 0; repos.Len() == 0; i++ {
		if i > 10 {
			return lu.Array{}
		}
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second * 5)
	return repos
}
