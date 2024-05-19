package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IshanSaha05/pkg/config"
	"github.com/IshanSaha05/pkg/mongodb"
	"github.com/IshanSaha05/pkg/scraper"
)

func main() {
	fmt.Println("Message: Connecting to MongoDB Server.")

	var MongoDBObject mongodb.MongoDb
	err := MongoDBObject.GetMongoClient()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	/*
		scraper.ScrapeStoreAssemblyData(&MongoDBObject)
		scraper.ScrapeStoreParliamentData(&MongoDBObject)
		scraper.ScrapeStoreDistrictData(&MongoDBObject)
	*/

	htmlContent, _ := scraper.GetHTMLContentAssemblyState(config.ScrapeSiteUrlAssembly)
	fmt.Println(htmlContent)

	fmt.Println("Message: Job Done Successfully.")
}
