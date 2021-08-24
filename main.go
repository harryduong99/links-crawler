package main 

import (
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	"fmt"
	"crypto/tls"
	"github.com/steelx/extractlinks"
	"net/url"
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
)

 
func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Printf("Missing Url")
		os.Exit(1)
	}

	baseURL := arg[0]
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
		// if (link.Href == "https://hopamviet.vn/" || link.Href == "https://hopamviet.vn") {
			// fmt.Printf("Not this one!")
			// return
		// }
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
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
}
