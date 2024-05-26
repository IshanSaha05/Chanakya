package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IshanSaha05/pkg/mongodb"
	"github.com/IshanSaha05/pkg/scraper"
)

func main() {
	fmt.Println("Message: Connecting to MongoDB Server.")

	var mongoDBObject mongodb.MongoDB
	err := mongoDBObject.GetMongoClient()
	if err != nil {
		log.Fatalf("Error: %s", err)
		os.Exit(1)
	}

	err, err1, err2, err3 := scraper.GetAllAndSave(mongoDBObject)
	if err != nil {
		log.Fatalf("Error(StateName): %s", err)
		os.Exit(1)
	}
	if err1 != nil {
		log.Fatalf("Error(Assembly): %s", err)
		os.Exit(1)
	}
	if err2 != nil {
		log.Fatalf("Error(Parliament): %s", err)
		os.Exit(1)
	}
	if err3 != nil {
		log.Fatalf("Error(District): %s", err)
		os.Exit(1)
	}

	fmt.Println("Message: Job Done Successfully.")
}
