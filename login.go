package sw

import (
	"fmt"
	"io/ioutil"
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
	cookies = loginCheckColly(token, cookies)
	fmt.Println(cookies)

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

func loginCheckColly(t lu.String, cookies []*http.Cookie) []*http.Cookie {
	email := os.Getenv("SCRUTINIZER_USRNM")
	password := os.Getenv("SCRUTINIZER_PSSWD")

	if email == "" || password == "" {
		log.Fatalln("We forgot to initialize the email and/or password!")
	}

	urlValues := url.Values{}
	urlValues.Set("email", email)
	urlValues.Set("password", password)
	urlValues.Set("remember_me", "1")
	urlValues.Set("_token", t.ToS())
	bodyString := urlValues.Encode()

	URL := Endpoint + "login_check"

	// creates a new collector
	c := colly.NewCollector()
	/*c.SetCookies(sw.ReposPageURL, []*http.Cookie{{
		Name:    "SESS",
		Value:   sessionId.ToS(),
		Expires: time.Time{},
		Path:    "/",
	}})*/

	hdr := http.Header{}
	hdr.Set("Accept", "*/*")
	//hdr.Set("Accept-Encoding", "gzip, deflate, br")
	hdr.Set("Connection", "keep-alive")
	hdr.Set("User-Agent", c.UserAgent)
	c.SetCookies(URL, cookies)

	// handle html
	c.OnResponse(func(r *colly.Response) {
		//log.Println("Reponse Received", string(r.Body))
		ioutil.WriteFile("/home/lwenger/Documents/scrutinizer-webscraper/cmd/output.html", r.Body, 777)
	})

	// starts scraping
	err := c.Request("POST", URL, strings.NewReader(bodyString), nil, hdr)
	if err != nil {
		log.Fatalln(err)
	}

	// waits util done adding to repos variable

	return c.Cookies(URL)

}

// loginCheck POSTs to the /login_check URL and retrieves the cookies, returning them
func loginCheck(t lu.String, sI lu.String) []*http.Cookie {
	email := os.Getenv("SCRUTINIZER_USRNM")
	password := os.Getenv("SCRUTINIZER_PSSWD")

	if email == "" || password == "" {
		log.Fatalln("We forgot to initialize the email and/or password!")
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	bodyString := fmt.Sprintf("email=%s&password=%s&remember_me=1&_target_path=&_token=%s", email, password, t)

	URL := Endpoint + "login_check"
	req, err := http.NewRequest("POST", URL, strings.NewReader(bodyString))
	req.AddCookie(&http.Cookie{
		Name:     "SESS",
		Value:    sI.ToS(),
		Path:     "/",
		HttpOnly: true,
	})
	//fmt.Println(req.Cookies())
	//req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("cookie", "SESS="+sI.ToS())
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
		return nil
	}
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	fmt.Println(res.Header)

	//token := strings.Split(res.Header.Get("set-cookie"), "SESS=")[1]
	//token = strings.Split(token, ";")[0]

	return res.Cookies()
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
	return lu.String(s).Split("SESS=")[1].Split(";")[0]
}
