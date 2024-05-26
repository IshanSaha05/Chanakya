package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IshanSaha05/pkg/config"
	"github.com/IshanSaha05/pkg/models"
	"github.com/IshanSaha05/pkg/mongodb"
)

func GetStateNames() ([]string, error) {
	// Getting the url saved in the config package.
	url := config.AssemblyStateNamesSite

	// Getting the response from the url.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Checking if the status code of the url is ok.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: Site not reachable.\nStatus Code: %d", resp.StatusCode)
	}

	// Initialising the variable to decode and store the response from the body.
	var responseData []models.StateName

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	}

	// Converting the decoded data into slice of string.
	var stateNames []string
	for _, stateName := range responseData {
		stateNames = append(stateNames, stateName.StateName)
	}

	return stateNames, nil
}

func GetAllAndSave(mongoDBObject mongodb.MongoDB) (error, error, error, error) {
	// Storing the data in the db.
	mongoDBObject.SetMongoDatabase("Election-Database-Chanakyya")
	mongoDBObject.SetMongoCollection("Election-Collection-Chanakyya")

	// First getting all the state names.
	stateNames, err := GetStateNames() // --> Done
	if err != nil {
		return err, nil, nil, nil
	}

	fmt.Printf("\nMessage(Assembly): Starting Scraping.")
	err1 := GetAndSaveAllAssemblyData(&mongoDBObject, stateNames)
	fmt.Printf("\nMessage(Assembly): Starting Scraping.")

	fmt.Printf("\nMessage(Parliament): Starting Scraping.")
	err2 := GetAndSaveAllParliamentData(&mongoDBObject, stateNames)
	fmt.Printf("\nMessage(Parliament): Finsihed Scraping.")

	fmt.Printf("\nMessage(District): Starting Scraping.")
	err3 := GetAndSaveAllDistrictData(&mongoDBObject, stateNames)
	fmt.Printf("\nMessage(District): Finsihed Scraping.")

	//return nil, err1, nil, nil
	//return nil, nil, err2, nil
	//return nil, nil, nil, err3
	return nil, err1, err2, err3
}
