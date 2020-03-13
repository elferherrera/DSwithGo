package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

// Info struct with all the information
type Info struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

// NewData creates a field data for json
func NewData(n int) []Info {

	allData := make([]Info, n, n)
	for i := 0; i < n; i++ {
		data := Info{
			rand.Intn(100),
			Names[rand.Intn(len(Names))],
			Answers[rand.Intn(len(Answers))],
		}

		allData[i] = data

	}

	return allData
}

// PrettyPrint function to print formated info
func PrettyPrint(info Info) {
	infoJSON, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		fmt.Println("Unable to print info")
	}

	fmt.Println(string(infoJSON))
}

// SaveToFile stores data into a file
func SaveToFile(allData []Info, fileName string) error {

	formatedData, err := json.MarshalIndent(allData, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, formatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// LoadFile loads info from file
func LoadFile(fileName string) ([]Info, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Unable to read file %v\n", fileName)
		return nil, err
	}

	var infoStruct []Info
	err = json.Unmarshal(data, &infoStruct)
	if err != nil {
		fmt.Println("Unable to Unmarshal data")
		return nil, err
	}

	return infoStruct, nil
}
