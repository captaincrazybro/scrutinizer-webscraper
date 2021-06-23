package main

import (
	"fmt"
	"net/http"
	"strings"
	"sort"
	"strconv"

	sw "github.com/captaincrazybro/scrutinizer-webscraper"

	"github.com/gocolly/colly"
)

func main() {
	var (
		repos, str []string
		sum, count float64
	)

	//Get cookies needed to log in and start a new collector
	cookies := sw.Login()
	collector := colly.NewCollector()

	//Get list of repositories and their quality scores
	collector.OnHTML("a[title]", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("title"), sw.BBOrgName+"/") || 
		strings.HasPrefix(e.Attr("title"), sw.GHOrgName+"/") {
			str = strings.Fields(e.Text)
			repos = append(repos, str...)
		}
	})
	collector.OnHTML("div[class]", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("class"), "span2") {
			str = strings.Fields(e.Text)
			if num, err := strconv.ParseFloat(str[0], 64); err == nil {
				sum += num
				count++
			} else {
				fmt.Errorf("%s\n", err)
			}
		}
	})

	//Prepare request and scrape website
	hdr := http.Header{}
	hdr.Set("Accept", "*/*")
	hdr.Set("Connection", "keep-alive")
	hdr.Set("User-Agent", collector.UserAgent)
	collector.SetCookies(sw.ReposPageURL, cookies)
	err := collector.Request("GET", sw.ReposPageURL, nil, nil, hdr)
	if err != nil {
		fmt.Errorf("%s\n", err)
	}

	sort.Strings(repos)
	fmt.Printf("%s\n%f\n", repos, sum / count)
}