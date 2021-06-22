package sw

import (
	"net/http"

	lu "github.com/captaincrazybro/literalutil"
	"github.com/gocolly/colly"
)

// Login logs into scrutinizer, retrieves the session cookies and returns them
func Login() lu.String {
	return ""
}

// getCSRPToken retrieves the CSRP token from the login page button
func getCSRPToken() lu.String {
	// creates a now colly collector
	c := colly.NewCollector()

	// handle html

	c.OnHTML("input[name]", func(e *colly.HTMLElement) {
		if e.Attr("name") == "_token" {

		}
	})
	return ""
}

// loginCheck POSTs to the /login_check URL and retrieves the cookies, returning them
func loginCheck(S lu.String) lu.String {
	URL := Endpoint + "/login_check"
	http.NewRequest("POST", URL, nil)

	/**
	Body: {
		"email": "the@email.tld",
		"password": "asdf123",
		"_token": "the token S",
		"remember_me": 1
	}
	*/

	return ""
}
