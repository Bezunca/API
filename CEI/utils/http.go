package utils

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"net/http"
	"net/url"
)

type cookieJar struct {
	jar map[string] []*http.Cookie
}

func (p* cookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	//fmt.Printf("The URL is : %s\n", u.String())
	//fmt.Printf("The cookie being set is : %s\n", cookies)
	p.jar [u.Host] = cookies
}

func (p *cookieJar) Cookies(u *url.URL) []*http.Cookie {
	//fmt.Printf("The URL is : %s\n", u.String())
	//fmt.Printf("Cookie being returned is : %s\n", p.jar[u.Host])
	return p.jar[u.Host]
}

var jar = &cookieJar{make(map[string] []*http.Cookie)}

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
}

var client = http.Client{
	Transport: tr,
	Jar: jar,
}

func GetPage(url string) (*html.Node, error) {

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("REQUEST error: %v", err)
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("STATUS error: %v", resp.StatusCode)
	}

	read, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("READ error: %v", err)
	}

	data, err := html.Parse(read)
	if err != nil {
		return nil, fmt.Errorf("PARSE error: %v", err)
	}

	return data, nil
}

func PostPage(urlString string, payload url.Values) (*html.Node, error) {

	resp, err := client.PostForm(urlString, payload)
	if err != nil {
		return nil, fmt.Errorf("REQUEST error: %v", err)
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("STATUS error: %v", resp.StatusCode)
	}

	read, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("READ error: %v", err)
	}

	data, err := html.Parse(read)
	if err != nil {
		return nil, fmt.Errorf("PARSE error: %v", err)
	}

	return data, nil
}
