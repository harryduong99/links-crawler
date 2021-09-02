package main

import (
	"links-crawler/crawler"
	"links-crawler/driver"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var DbType string

func main() {
	loadEnv()
	loadDatabase()
	crawler.Crawling(DbType)

	time.Sleep(10 * time.Second)
}
func loadDatabase() {
	arg := os.Args[1:2]
	DbType = arg[0]
	log.Println(driver.GetDbDriverFactory(DbType))
	os.Exit(1)
	driver.GetDbDriverFactory(DbType).ConnectDatabase()
}
func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
