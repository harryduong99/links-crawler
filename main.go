package main

import (
	"log"
	"os"
	"song-chord-crawler/crawler"
	"song-chord-crawler/driver"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	crawler.Crawling()

	// time.Sleep(100 * time.Second)
}
func loadDatabase() {
	arg := os.Args[1:2]
	driver.GetDbDriverFactory(arg[0]).ConnectDatabase()
}
func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
