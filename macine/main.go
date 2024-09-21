package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VinDecoderResponse struct {
	Results []struct {
		Make  string `json:"Make"`
		Model string `json:"Model"`
		Year  string `json:"ModelYear"`
		Name  string `json:"VehicleName"`
	} `json:"Results"`
}

func checkVin(vin string) (*VinDecoderResponse, error) {
	url := fmt.Sprintf("https://vpic.nhtsa.dot.gov/api/vehicles/DecodeVinValues/%s?format=json", vin)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var vinResponse VinDecoderResponse
	if err := json.Unmarshal(body, &vinResponse); err != nil {
		return nil, err
	}

	return &vinResponse, nil
}


func main() {
	for i := 1; i <= 7; i++ {
		for j := 1; j <= 7; j++ {
			if j == 1 || (i == 1 && j < 7) || (i == 7 && j < 7) || (j == 7 && i > 1) && i < 7 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			}
			fmt.Println()
		}
}

