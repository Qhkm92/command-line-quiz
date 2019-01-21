package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"os"
)

func main() {

	const sourcePath string = "./name.csv"
	ReadCsvFile(sourcePath)
		
}

//function to read csv file
func ReadCsvFile(filepath string){

	f, _ := os.Open(filepath)
	defer f.Close()
	//Read CSV file
	r := csv.NewReader(strings.NewReader(f))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}