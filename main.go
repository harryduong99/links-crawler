package main

import (
	"log"
	"os"
	"song-chord-crawler/crawler"
	"song-chord-crawler/driver"

	"github.com/joho/godotenv"
)

var DbType string

func main() {
	loadEnv()
	loadDatabase()
	crawler.Crawling(DbType)

	// time.Sleep(100 * time.Second)
}
func loadDatabase() {
	arg := os.Args[1:2]
	DbType = arg[0]
	driver.GetDbDriverFactory(DbType).ConnectDatabase()
}
func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
