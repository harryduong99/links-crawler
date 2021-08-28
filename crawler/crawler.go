package crawler

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"song-chord-crawler/crawler/extractlinks"
	"song-chord-crawler/models"
	"song-chord-crawler/repository/linksRepo"
	"time"
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
	queue      = make(chan string)
	hasVisited = make(map[string]bool)
)

var Result = []string{}

func Crawling() {

	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Printf("Missing Url")
		os.Exit(1)
	}

	baseURL := arg[0]
	go func() {
		queue <- baseURL
	}()
	time.Sleep(2 * time.Second)
	fmt.Println(len(queue))

	for href := range queue {
		if !hasVisited[href] { // should check is same domain too
			go func() {
				insertLinkToDb(href)
			}()
			CrawlUrl(href)
		}
	}

}

func CrawlUrl(href string) {
	hasVisited[href] = true
	Result = append(Result, href)
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
		go func() {
			queue <- absoluteUrl // this will make the for in range queue keep running
		}()
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

func insertLinkToDb(href string) bool {
	link := models.Link{
		Url: href,
	}

	if linksRepo.IsLinkExist(href) {
		log.Printf("%s is existing in database!", href)
		return false
	}

	err := linksRepo.StoreLink(link)
	if err != nil {
		log.Printf("storing href: %s failed!", href)
		return false
	}

	return true
}

func isSameDomain() {

}
