package sw

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	lu "github.com/captaincrazybro/literalutil"
	"github.com/gocolly/colly"
)

var _token lu.String = "none"

// Login logs into scrutinizer, retrieves the session cookies and returns them
func Login() lu.String {
	token, sessionId := getCSRPToken()
	_token = "none"
	sessionId = loginCheck(token, sessionId)
	fmt.Println(sessionId)

	return sessionId
}

// getCSRPToken retrieves the CSRP token from the login page button
func getCSRPToken() (lu.String, lu.String) {
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

	return waitUntilTokenIsSet(), findSessionId(c.Cookies(LoginPageURL))
}

// loginCheck POSTs to the /login_check URL and retrieves the cookies, returning them
func loginCheck(t lu.String, sI lu.String) lu.String {
	email := os.Getenv("SCRUTINIZER_USRNM")
	password := os.Getenv("SCRUTINIZER_PSSWD")

	if email == "" || password == "" {
		log.Fatalln("We forgot to initialize the email and/or password!")
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	bodyString := fmt.Sprintf("email=%s&password=%s&remember_me=1&_target_path=&_token=%s", strings.Replace(email, "@", "%40", 1), password, t)

	URL := Endpoint + "login_check"
	req, err := http.NewRequest("POST", URL, strings.NewReader(bodyString))
	//req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Set-Cookie", "SESS="+sI.ToS())
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Status)

	defer res.Body.Close()
	//bts, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	token := strings.Split(res.Header.Get("set-cookie"), "SESS=")[1]
	token = strings.Split(token, ";")[0]

	return lu.String(token)
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

// findSessionId finds the session id from the cookies array
func findSessionId(sz []*http.Cookie) lu.String {
	var s string
	for _, v := range sz {
		if strings.HasPrefix(fmt.Sprintf("%v", v), "SESS=") {
			s = fmt.Sprintf("%v", v)
			break
		}
	}
	return lu.String(s)
}
