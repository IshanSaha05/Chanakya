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

func getAssemblyStateData(state string) ([]models.AssemblyStateData, error) {
	fmt.Printf("\nMessage(Assembly): Scraping State --> %s || Data.", state)

	// Create the url.
	url := fmt.Sprintf("http://www.chanakyya.com/Chanakya/%s/%s.json?isNewDataFormat=true", state, state)

	// Get the response from the url.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check resposne status of the url.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error(Assembly): Site \"%s\" not reachable.", url)
		return nil, nil
	}

	// Initialise variable to decode data.
	var responseData models.AssemblyStateResponseData

	// Decode.
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	}

	// Data section exist check.
	if responseData.StateLevelData.StateDataRaw == nil {
		fmt.Printf("Error(Assembly): No data present for state --> %s.", state)
		return nil, nil
	}

	// Convert the data to store it.
	var datas []models.AssemblyStateData

	for key, data := range responseData.StateLevelData.StateDataRaw {
		for _, rawdata := range data {
			var temp models.AssemblyStateData
			temp.State = state
			temp.ElectionType = key
			temp.Party = rawdata.Party
			temp.AssemblySegmentsLeading = rawdata.AssemblySegmentsLeading
			temp.VotePercentage = rawdata.VotePercentage

			datas = append(datas, temp)
		}
	}

	fmt.Printf("\nMessage(Assembly): Finished Scraping State --> %s || Data.", state)

	return datas, nil
}

func getAssemblyStateConstituencyNames(state string) ([]string, error) {
	fmt.Printf("\nMessage(Assembly): Scraping State --> %s || Constituency Names.", state)

	// Create the url.
	url := fmt.Sprintf("http://www.chanakyya.com/Chanakya/%s/%s.json?isNewDataFormat=true", state, state)

	// Get the response from the url.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check response status of the url.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error(Assembly): Site \"%s\" not reachable.", url) // Need to change.
		return nil, nil
	}

	// Initialise variable to decode data.
	var assemblyStateConstituencyNamesResponseData models.AssemblyStateConstituencyNames

	// Decode.
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&assemblyStateConstituencyNamesResponseData)
	if err != nil {
		return nil, err
	}

	// Constituency section does not exist check.
	if assemblyStateConstituencyNamesResponseData.ConstituencyName == nil {
		fmt.Printf("Error(Assembly): Parliament constituency section does not exist in json.")
		return nil, nil
	}

	var assemblyStateConstituencyNames []string

	// Convert from "xyz.json" to "xyz" constituency name pattern.
	for _, constituencyName := range assemblyStateConstituencyNamesResponseData.ConstituencyName {
		extension := filepath.Ext(constituencyName)

		constituencyName = strings.TrimSuffix(constituencyName, extension)

		assemblyStateConstituencyNames = append(assemblyStateConstituencyNames, constituencyName)
	}

	fmt.Printf("\nMessage(Assembly): Finished Scraping State --> %s || Constituency Names.", state)

	return assemblyStateConstituencyNames, nil
}

func getAssemblyStateConstituencyData(state string, constituency string) ([]models.AssemblyStateConstituencyData, error) {
	fmt.Printf("\nMessage(Assembly): Scraping State --> %s Constituency --> %s || Data.", state, constituency)

	// Create the url.
	url := fmt.Sprintf("http://www.chanakyya.com/Chanakya/%s/AssemblyData/%s.json?isNewDataFormat=true", state, constituency)

	// Get the respone from the url.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check response status of the url.
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error(Assembly): Site \"%s\" not reachable.", url)
		return nil, nil
	}

	// Initialise variable to decode data.
	var responseData models.AssemblyStateConstituencyResponseData

	// Decode.
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&responseData)
	if err != nil {
		return nil, err
	}

	// Data section exist check.
	if responseData.AssemblyData.AssemblyDataRaw == nil {
		fmt.Printf("Error(Assembly): No data present for state --> %s and constituency --> %s.", state, constituency)
		return nil, nil
	}

	// Convert the data to store it.
	var datas []models.AssemblyStateConstituencyData

	for key, data := range responseData.AssemblyData.AssemblyDataRaw {
		for _, rawdata := range data.RawData {
			var temp models.AssemblyStateConstituencyData
			temp.State = state
			temp.Constituency = constituency
			temp.ElectionType = key
			temp.Candidate = rawdata.Candidate
			temp.Party = rawdata.Party
			temp.NumberOfVotes = rawdata.NumberOfVotes
			temp.VotePercentage = rawdata.VotePercentage

			datas = append(datas, temp)
		}
	}

	fmt.Printf("\nMessage(Assembly): Finished Scraping State --> %s Constituency --> %s || Data.", state, constituency)

	return datas, nil
}

func GetAndSaveAllAssemblyData(mongoDBObject *mongodb.MongoDB, stateNames []string) error {
	// Loop through every state.
	for _, stateName := range stateNames {
		// Get the Assembly State data.
		assemblyStateData, err := getAssemblyStateData(stateName)
		// Checks for the condition if the data exists or not.
		/*if err == fmt.Errorf("assembly: no data present for state --> %s", stateName) {
			continue
		}*/

		if err != nil {
			return err
		}

		// Insert the data into the db.
		fmt.Printf("\nMessage(Assembly): Inserting State --> %s || Data.", stateName)
		err = mongoDBObject.InsertIntoDBAssemblyStateData(assemblyStateData)
		if err != nil {
			return err
		}
		fmt.Printf("\nMessage(Assembly): Finished Inserting State --> %s || Data.", stateName)

		// Get the assembly constituency names list.
		assemblyStateConstituencyNames, err := getAssemblyStateConstituencyNames(stateName)
		if err != nil {
			return err
		}

		// Loop through each assembly constituency of the particular state.
		for _, constituencyName := range assemblyStateConstituencyNames {
			// Get the assembly constituency data of the particular state.
			assemblyStateConstituencyData, err := getAssemblyStateConstituencyData(stateName, constituencyName)
			/*if err == fmt.Errorf("assembly: no data present for state --> %s and constituency --> %s", stateName, constituencyName) {
				continue
			}*/

			if err != nil {
				return err
			}

			// Insert the data into the db.
			fmt.Printf("\nMessage(Assembly): Inserting State --> %s Constituency --> %s || Data.", stateName, constituencyName)
			err = mongoDBObject.InsertIntoDBAssemblyStateConstituencyData(assemblyStateConstituencyData)
			if err != nil {
				return err
			}
			fmt.Printf("\nMessage(Assembly): Finished Inserting State --> %s Constituency --> %s || Data.", stateName, constituencyName)
		}

	}

	return nil
}
