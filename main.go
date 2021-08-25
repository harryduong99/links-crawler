package main

import (
	// "io/ioutil"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/steelx/extractlinks"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	netClient = &http.Client{
		Transport: transport,
	}
	queue = make(chan string)
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Printf("Missing Url")
		os.Exit(1)
	}

	// baseURL := arg[0]
	go func() {
		queue <- agr[0]
	}()

	for href := range queue {
		CrawlUrl(href)
	}
	fmt.Printf("base url: %v\n", baseURL)
	// baseUrl := "https://hopamviet.vn/"

	CrawlUrl(baseURL)

}

func CrawlUrl(href string) {
	fmt.Printf("Url: %v \n", href)
	resp, err := netClient.Get(href)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	links, err := extractlinks.All(resp.Body)
	for _, link := range links {
		fmt.Println(link.Href)
		absoluteUrl := toFixedUrl(link.Href, href)
		go func(link) {
			queue <- absoluteUrl
		}()
		// CrawlUrl(toFixedUrl(link.Href, href))

	}
}

func toFixedUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil || uri.Scheme == "mailto" || uri.Scheme == "tel" {
		return base
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
