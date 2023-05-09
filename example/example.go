package main

import (
	"encoding/json"
	"log"
	"os"

	"emaildomainstats"
)

func main() {
	results, err := emaildomainstats.ProcessCSVFile("customer_data.csv", emaildomainstats.WithColumns([]int{1, 2, 3}))
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("./example/results.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
