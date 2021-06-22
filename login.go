package sw

import (
	lu 	"github.com/captaincrazybro/literalutil"
	"github.com/gocolly/colly"
)

// Login logs into scrutinizer, retrieves the session cookies and returns them
func Login() lu.String {

	return ""
}

// getCSRPToken retrieves the CSRP token from the login page button
func getCSRPToken() lu.String {

	return ""
}

// loginCheck POSTs to the /login_check URL and retrieves the cookies, returning them
func loginCheck(c Collector) string {
	URL := Endpoint + "/login_check"
	siteCookies := c.Cookies(r.Request.URL.String())

	return siteCookies
}
