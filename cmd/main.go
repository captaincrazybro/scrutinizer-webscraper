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
	cookies := sw.Login()
	Sz := getRepos(cookies)
	fmt.Println(Sz)
	repos = lu.Array{}
}

func getRepos(cookies []*http.Cookie) lu.Array {
	// creates a new collector
	c := colly.NewCollector()

	// handle html
	c.OnHTML("a[title]", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("title"), sw.BBOrgName+"/") || strings.HasPrefix(e.Attr("title"), sw.GHOrgName+"/") {
			repos = append(repos, lu.String(e.Attr("title")).Split("/")[1])
		}
	})

	// prepare request
	hdr := http.Header{}
	hdr.Set("Accept", "*/*")
	hdr.Set("Connection", "keep-alive")
	hdr.Set("User-Agent", c.UserAgent)
	c.SetCookies(sw.ReposPageURL, cookies)

	// starts scraping
	err := c.Request("GET", sw.ReposPageURL, nil, nil, hdr)
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

	Sz := lu.Array{}
	for _, v := range A {
		Sz = append(Sz, v)
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
