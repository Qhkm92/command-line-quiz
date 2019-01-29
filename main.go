package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"

)

// 2 main things we want to be in main function, 1. present quiz 2. accepts user input 3. check for correctness
func main() {

	//read csv as string
	csvFileName := flag.String("csv", "problem.csv", "a csv file in the format of 'question, answer'")

	//parse string
	flag.Parse()

	//open csv file
	//error handling
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("failed to open the CSV file: %\n", *csvFileName))
	}

	//read csv file
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	//print to check values -> fmt.Println(lines)
	problems := parseLine(lines)
	fmt.Println(problems)

	//prints every question to the terminal 
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problems #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
}

func parseLine(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

//Create a problem struct from slice
type problem struct {
	q string
	a string 
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}







//function to read csv file
// func ReadCsvFile(filepath string){

// 	f, _ := os.Open(filepath)
// 	defer f.Close()
// 	//Read CSV file
// 	r := csv.NewReader(strings.NewReader(f))

// 	for {
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(record)
// 	}
// }