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
	jar map[string][]*http.Cookie
}

func (p *cookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	//fmt.Printf("The URL is : %s\n", u.String())
	//fmt.Printf("The cookie being set is : %s\n", cookies)
	p.jar[u.Host] = cookies
}

func (p *cookieJar) Cookies(u *url.URL) []*http.Cookie {
	//fmt.Printf("The URL is : %s\n", u.String())
	//fmt.Printf("Cookie being returned is : %s\n", p.jar[u.Host])
	return p.jar[u.Host]
}

var jar = &cookieJar{make(map[string][]*http.Cookie)}

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}

var client = http.Client{
	Transport: tr,
	Jar:       jar,
}

func GetPage(url string) *html.Node {

	resp, err := client.Get(url)
	Check(err)

	if resp.StatusCode != http.StatusOK {
		Check(fmt.Errorf("STATUS error: %v", resp.StatusCode))
	}

	read, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	Check(err)

	data, err := html.Parse(read)
	Check(err)

	return data
}

func PostPage(urlString string, payload url.Values) *html.Node {

	resp, err := client.PostForm(urlString, payload)
	Check(err)

	if resp.StatusCode != http.StatusOK {
		Check(fmt.Errorf("STATUS error: %v", resp.StatusCode))
	}

	read, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	Check(err)

	data, err := html.Parse(read)
	Check(err)

	return data
}
