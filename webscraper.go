package sw

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	lu "github.com/captaincrazybro/literalutil"

	"github.com/gocolly/colly"
)

// FetchScrutinizerRepos fetches registered scrutinizer repos and retrieves the average quality score
func FetchScrutinizerRepos() ([]string, float64) {
	var (
		repos, str []string
		sum, count float64
	)

	//Get cookies needed to log in and start a new collector
	cookies := login()
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

var _token lu.String = "none"

// login logs into scrutinizer, retrieves the session cookies and returns them
func login() []*http.Cookie {
	token, cookies := getCSRPToken()
	_token = "none"
	cookies = loginCheck(token, cookies)

	return cookies
}

// getCSRPToken retrieves the CSRP token from the login page button
func getCSRPToken() (lu.String, []*http.Cookie) {
	// creates a now colly collector
	c := colly.NewCollector()

	// handle html
	c.OnHTML("input[name]", func(e *colly.HTMLElement) {
		if e.Attr("name") == "_token" {
			_token = lu.String(e.Attr("value"))
		}
	})

	// visits the site
	err := c.Visit(LoginPageURL)
	if err != nil {
		log.Fatalln(err)
	}

	return waitUntilTokenIsSet(), c.Cookies(LoginPageURL)
}

// loginCheck posts to /login_check to retrieve session id and other cookies
func loginCheck(t lu.String, cookies []*http.Cookie) []*http.Cookie {
	envVars := []string{ScrutinizerUsrnm, ScrutinizerPsswd}
	GetEnvVariables(envVars)

	// creates a new collector
	c := colly.NewCollector()

	// prepares request
	urlValues := url.Values{}
	urlValues.Set("email", envVars[0])
	urlValues.Set("password", envVars[1])
	urlValues.Set("remember_me", "1")
	urlValues.Set("_token", t.ToS())
	bodyString := urlValues.Encode()

	URL := Endpoint + "login_check"

	hdr := http.Header{}
	hdr.Set("Accept", "*/*")
	hdr.Set("Connection", "keep-alive")
	hdr.Set("User-Agent", c.UserAgent)
	c.SetCookies(URL, cookies)

	// starts scraping
	err := c.Request("POST", URL, strings.NewReader(bodyString), nil, hdr)
	if err != nil {
		log.Fatalln(err)
	}

	return c.Cookies(URL)

}

// waitUntilTokenIsSet waits until the token is set and then returns it
func waitUntilTokenIsSet() lu.String {
	for i := 0; _token == "none"; i++ {
		if i > 10 {
			return _token
		}
		time.Sleep(time.Second)
	}
	return _token
}
