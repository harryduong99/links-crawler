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
	"sync"
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
	queue          = make(chan string)
	hasVisited     = make(map[string]bool)
	chunksToInsert = []models.Link{}
	w              = sync.WaitGroup{}
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

	for href := range queue {
		if !hasVisited[href] && isSameDomain(href, baseURL) { // should check is same domain too
			// w.Add(1)
			go func() {
				insertLinkToDb(href)
			}()
			CrawlUrl(href)
		}
	}

	// w.Wait()
	fmt.Println("Main: Completed")

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

	if err != nil {
		log.Fatal("get all links failed!")
	}

	for _, link := range links {
		// fmt.Println(link.Href)
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

func insertLinkToDb(href string) bool {
	// defer w.Done()

	link := models.Link{
		Url:     href,
		Crawled: false,
		Domain:  getDomain(href),
	}

	if linksRepo.IsLinkExist(href) {
		log.Printf("%s is existing in database!", href)
		return false
	}

	chunksToInsert = append(chunksToInsert, link)
	log.Println(len(chunksToInsert))

	if len(chunksToInsert) == 50 {
		err := linksRepo.StoreLinks(chunksToInsert)
		if err != nil {
			log.Printf("storing href: %s failed!", href)
			return false
		}
		chunksToInsert = nil

		return true
	}

	return false
}

func isSameDomain(href, baseUrl string) bool {
	uri, err := url.Parse(href)
	if err != nil {
		return false
	}

	parentUri, err := url.Parse(baseUrl)

	if err != nil {
		return false
	}

	if uri.Host != parentUri.Host {
		return false
	}

	return true
}

func getDomain(href string) string {

	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}

	return uri.Host
}
