package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"time"
	"strings"

)

// 2 main things we want to be in main function, 1. present quiz 2. accepts user input 3. check for correctness
func main() {

	//read csv as string
	csvFileName := flag.String("csv", "problem.csv", "a csv file in the format of 'question, answer'")

	//add time limit
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	//parse string
	flag.Parse()

	if *timeLimit == nil {
		flag.PrintDefaults()
		os.Exit(1)
	}

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
	//fmt.Println(problems)

	//set timer 
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
//wait message from this channel 

	//prints every question to the terminal 
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problems #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <- timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <- answerCh:
		if answer == p.a {
			correct++
			}
		}
	}

	
}

func parseLine(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	fmt.Println(ret)
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
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