package sw

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	lu "github.com/captaincrazybro/literalutil"
	"github.com/gocolly/colly"
)

var _token lu.String = "none"

// Login logs into scrutinizer, retrieves the session cookies and returns them
func Login() []*http.Cookie {
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
	email := os.Getenv("SCRUTINIZER_USRNM")
	password := os.Getenv("SCRUTINIZER_PSSWD")

	if email == "" || password == "" {
		log.Fatalln("We forgot to initialize the email and/or password!")
	}

	// creates a new collector
	c := colly.NewCollector()

	// prepares request
	urlValues := url.Values{}
	urlValues.Set("email", email)
	urlValues.Set("password", password)
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
