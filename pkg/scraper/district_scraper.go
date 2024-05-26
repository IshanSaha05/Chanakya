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

func getDistrictStateDistrictNames(state string) ([]string, error) {
	// Create the url.
	fmt.Printf("\nMessage(District): Scraping State --> %s || District Names.", state)
	url := fmt.Sprintf("http://www.chanakyya.com/Chanakya/%s/%s.json?isNewDataFormat=true", state, state)

	// Get the response from the url.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check response status of the url.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("\nError(District): Site \"%s\" not reachable.", url)
		return nil, nil
	}

	// Initialise variable to decode data.
	var districtStateDistrictNamesResponseData models.DistrictStateDistrictName

	// Decode.
	decoder := json.NewDecoder((resp.Body))
	err = decoder.Decode(&districtStateDistrictNamesResponseData)
	if err != nil {
		return nil, err
	}

	// District section does not exist check.
	if districtStateDistrictNamesResponseData.DistrictStateDistrictMap == nil {
		fmt.Printf("Error(District): District section does not exist in json.")
		return nil, nil
	}

	var districtStateDistrictNames []string

	// Convert from "xyz.json" to "xyz" district name pattern.
	for _, districtName := range districtStateDistrictNamesResponseData.DistrictStateDistrictMap {
		extension := filepath.Ext(districtName)

		districtName = strings.TrimSuffix(districtName, extension)

		districtStateDistrictNames = append(districtStateDistrictNames, districtName)
	}

	fmt.Printf("\nMessage(District): Finished Scraping State --> %s || District Names.", state)

	return districtStateDistrictNames, nil
}

func getDistrictStateDistrictData(state string, district string) ([]models.DistrictStateDistrictData, error) {
	fmt.Printf("\nMessage(District): Scraping State --> %s District --> %s || Data.", state, district)

	// Create the url.
	url := fmt.Sprintf("http://www.chanakyya.com/Chanakya/%s/DistrictData/%s.json?isNewDataFormat=true", state, district)

	// Get the response from the url.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status of the url.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error(District): Site \"%s\" not reachable.", url)
		return nil, nil
	}

	// Initialise variable to decode data.
	var responseData models.DistrictStateDistrictResponseData

	// Decode.
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	}

	// Data section exist check.
	if responseData.DistrictData.DistrictDataRaw == nil {
		fmt.Printf("Error(District): no data present for state --> %s and district --> %s.", state, district)
		return nil, nil
	}

	// Convert the data to store it.
	var datas []models.DistrictStateDistrictData

	for key, data := range responseData.DistrictData.DistrictDataRaw {
		for _, rawdata := range data.RawData {
			var temp models.DistrictStateDistrictData
			temp.State = state
			temp.District = district
			temp.ElectionType = key
			temp.Party = rawdata.PartyName
			temp.AssemblySegmentsLeading = rawdata.AssemblySegmentsLeading
			temp.NumberOfVotes = rawdata.NumberOfVotes
			temp.VotePercentage = rawdata.VotePercentage

			datas = append(datas, temp)
		}
	}

	fmt.Printf("\nMessage(District): Finished Scraping State --> %s District --> %s || Data.", state, district)

	return datas, nil
}

func GetAndSaveAllDistrictData(mongoDBObject *mongodb.MongoDB, stateNames []string) error {
	// Loop through every state.
	for _, stateName := range stateNames {
		// Get the district names list.
		districtStateDistrictNames, err := getDistrictStateDistrictNames(stateName)
		if err != nil {
			return err
		}

		// Loop through each district of the particular state.
		for _, districtName := range districtStateDistrictNames {
			// Get the district data of the particular state.
			districtStateDistrictData, err := getDistrictStateDistrictData(stateName, districtName)

			// Checks for the condition if the data exists or not.
			/*if err == fmt.Errorf("district: no data present for state --> %s and district --> %s", stateName, districtName) {
				continue
			}*/

			if err != nil {
				return err
			}

			// Insert the data into the db.
			fmt.Printf("\nMessage(District): Inserting State --> %s District --> %s || Data.", stateName, districtName)
			err = mongoDBObject.InsertIntoDBDistrictStateDistrictData(districtStateDistrictData)
			if err != nil {
				return err
			}
			fmt.Printf("\nMessage(District): Finished Inserting State --> %s District --> %s || Data.", stateName, districtName)
		}
	}

	return nil
}
