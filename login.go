package sw

import (
	"net/http"
	"encoding/json"
	"log"
	"io/ioutils"
	"strings"

	lu "github.com/captaincrazybro/literalutil"
	"github.com/gocolly/colly"
)

var email, password, token, remember_me string

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
	if email == "" {
		log.Fatalln("We forgot to initialize the email!")
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	bts, err := json.Marshall(body{email, password, token, remember_me})
	if err != nil {
		log.Fatalln(err)
	}

	URL := Endpoint + "/login_check"
	req, err := http.NewRequest("POST", URL, bts)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(req.Body)) 
	if err != nil {
		log.Fatalln(err)
	}

	cookies := res.Header.Get("Cookies")

	token := strings.Split(cookies, "SESS=")[1]
	token = strings.Split(token, ";")[0]

	return token
}

type body struct {
	email, password, token, remember_me string
}