package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {

	url := "http://164.92.251.245:8080/api/v1/products/"

	data, err := readCSVFromUrl(url)
	if err != nil {
		log.Fatal(err)
	}

	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		if idx == 8 {
			break
		}

		fmt.Println(row[0], row[1])
	}
}
