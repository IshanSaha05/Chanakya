package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/IshanSaha05/pkg/models"
	"github.com/IshanSaha05/pkg/mongodb"
)

func getParliamentStateConstituencyNames(state string) ([]string, error) {
	// Create the url.
	fmt.Printf("\nMessage(Parliament): Scraping State --> %s || Constituency Names.", state)
	url := fmt.Sprintf("http://www.chanakyya.com/Chanakya/%s/%s.json?isNewDataFormat=true", state, state)

	// Get the response from the url.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check response status of the url.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("\nError(Parliament): Site \"%s\" not reachable.", url)
		return nil, nil
	}

	// Intialise variable to decode data.
	var parliamentStateConstituencyNamesReposneData models.ParliamentStateConstituencyName

	// Decode.
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&parliamentStateConstituencyNamesReposneData)
	if err != nil {
		return nil, err
	}

	// Constituency section does not exist check.
	if parliamentStateConstituencyNamesReposneData.ParliamentStateConstituencyMap == nil {
		fmt.Printf("\nError(Parliament): Parliament constituency section does not exist in json.")
		return nil, nil
	}

	var parliamentStateConstituencyNames []string

	// Convert from "xyz.json" to "xyz" constituency name pattern.
	for _, constituencyName := range parliamentStateConstituencyNamesReposneData.ParliamentStateConstituencyMap {
		extension := filepath.Ext(constituencyName)

		constituencyName = strings.TrimSuffix(constituencyName, extension)

		parliamentStateConstituencyNames = append(parliamentStateConstituencyNames, constituencyName)
	}

	fmt.Printf("\nMessage(Parliament): Finished Scraping State --> %s || Constituency Names.", state)

	return parliamentStateConstituencyNames, nil
}

func getParliamentStateConstituencyData(state string, constituency string) ([]models.ParliamentStateConstituencyData, error) {
	fmt.Printf("\nMessage(Parliament): Scraping State --> %s Constituency --> %s || Data.", state, constituency)

	// Create the url.
	url := fmt.Sprintf("http://www.chanakyya.com/Chanakya/%s/ParliamentData/%s.json?isNewDataFormat=true", state, constituency)

	// Get the response from the url.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check response status of the url.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("\nError(Parliament): Site \"%s\" not reachable.", url)
		return nil, nil
	}

	// Intialise variable to decode data.
	var responseData models.ParliamentStateConstituencyResponseData

	// Decode.
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	}

	// Data section exist check.
	if responseData.ParliamentData.ParliamentDataRaw == nil {
		fmt.Printf("\nError(Parliament): No data present for state --> %s and constituency --> %s.", state, constituency)
		return nil, nil
	}

	// Convert the data to store it.
	var datas []models.ParliamentStateConstituencyData

	for key, data := range responseData.ParliamentData.ParliamentDataRaw {
		for _, rawdata := range data.RawData {
			var temp models.ParliamentStateConstituencyData
			temp.State = state
			temp.Constituency = constituency
			temp.ElectionType = key
			temp.Candidate = rawdata.Candidate
			temp.NumberOfVotes = rawdata.NumberOfVotes
			temp.Party = rawdata.Party
			temp.VotePercentage = rawdata.VotePercentage

			datas = append(datas, temp)
		}
	}

	fmt.Printf("\nMessage(Parliament): Finished Scraping State --> %s Constituency --> %s || Data.", state, constituency)

	return datas, nil
}

func GetAndSaveAllParliamentData(mongoDBObject *mongodb.MongoDB, stateNames []string) error {
	// Loop through every state.
	for _, stateName := range stateNames {
		// Get the parliament consitiuency names list.
		parliamentStateConstituencyNames, err := getParliamentStateConstituencyNames(stateName)
		if err != nil {
			return err
		}

		// Loop through each parliament consitutency of the particular state.
		for _, constituencyName := range parliamentStateConstituencyNames {
			// Get the parliament consituency data of the particular state.
			parliamentStateConstituencyData, err := getParliamentStateConstituencyData(stateName, constituencyName)

			// Checks for the condition if the data exists or not.
			/*if err == fmt.Errorf("parliament: no data present for state --> %s and constituency --> %s", stateName, constituencyName) {
				continue
			}*/

			if err != nil {
				return err
			}

			// Insert the data into the db.
			fmt.Printf("\nMessage(Parliament): Inserting State --> %s Constituency --> %s || Data.", stateName, constituencyName)
			err = mongoDBObject.InsertIntoDBParliamentStateConstituencyData(parliamentStateConstituencyData)
			if err != nil {
				return err
			}
			fmt.Printf("\nMessage(Parliament): Finished Inserting State --> %s Constituency --> %s || Data.", stateName, constituencyName)
		}
	}

	return nil
}
