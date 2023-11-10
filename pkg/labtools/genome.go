package labtools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func SaveDNAs(dnas []dna.DNA, filename string) error {
	// Save a population DNA to File
	dnasJson, err := json.Marshal(dnas)
	if err != nil {
		log.Println("labtools.SavePopulationDNA: Error marshalling DNA to JSON")
		return err
	}

	if !strings.HasSuffix(filename, ".json") {
		filename += ".json"
	}

	err = ioutil.WriteFile(filename, dnasJson, 0644)
	if err != nil {
		log.Println("labtools.SavePopulationDNA: Error writing DNA to file")
		return err
	}

	return nil
}

func LoadDNAs(filename string) ([]dna.DNA, error) {
	// Load a population DNA from File
	var dnas []dna.DNA

	if !strings.HasSuffix(filename, ".json") {
		filename += ".json"
	}

	dnasJson, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("labtools.LoadPopulationDNA: Error reading DNA from file")
		return nil, err
	}

	err = json.Unmarshal(dnasJson, &dnas)
	if err != nil {
		log.Println("labtools.LoadPopulationDNA: Error unmarshalling DNA from JSON")
		return nil, err
	}

	return dnas, nil
}
