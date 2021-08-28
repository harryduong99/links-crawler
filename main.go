package main

import (
	"log"
	"song-chord-crawler/crawler"
	"song-chord-crawler/driver"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	driver.ConnectMongoDB("harry", "harry")
	crawler.Crawling()

	time.Sleep(10 * time.Second)
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
