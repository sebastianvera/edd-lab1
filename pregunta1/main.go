package main

import (
	"fmt"
	"os"

	"github.com/sebastianvera/edd-lab1/pregunta1/scraper"
)

func main() {
	csv, err := scraper.GetCSV()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Do something with the CSV
	fmt.Println(csv)
}
