package sw

import (
	"fmt"
	"net/http"
	"strings"
	"sort"
	"strconv"

	"github.com/gocolly/colly"
)

func ScrapeScrutinizer() ([]string, float64) {
	var (
		repos, str []string
		sum, count float64
	)

	//Get cookies needed to log in and start a new collector
	cookies := Login()
	collector := colly.NewCollector()

	//Get list of repositories and their quality scores
	collector.OnHTML("a[title]", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("title"), BBOrgName+"/") || 
		strings.HasPrefix(e.Attr("title"), GHOrgName+"/") {
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
			}
		}
	})

	//Prepare request and scrape website
	hdr := http.Header{}
	hdr.Set("Accept", "*/*")
	hdr.Set("Connection", "keep-alive")
	hdr.Set("User-Agent", collector.UserAgent)
	collector.SetCookies(ReposPageURL, cookies)
	err := collector.Request("GET", ReposPageURL, nil, nil, hdr)
	if err != nil {
		fmt.Errorf("%s\n", err)
	}

	sort.Strings(repos)
	return repos, sum / count
}