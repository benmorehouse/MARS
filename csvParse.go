package main

import(
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

/*
final goal is to have a struct where it has counted each class that has been attended
*/

type class struct{
	name string
	attendance int
}

// Need to go and look at something that you can do with the csv file system management that they have in go
func parseCSV(){
	csv_filename := flag.String("csv","data.csv", "csv data file")
	flag.Parse()
	file, err := os.Open(*csv_filename) // this will go and open the io.reader supplied by the flag
	file_reader := csv.NewReader(file) // file is io.reader that reads file
	lines , err := file_reader.ReadAll()

	//lines has all the information that we need now
	if err != nil{
		log.Fatal("Not able to read the inputted CSV file")
	}
	fmt.Println(lines)
}

func attendance_count([][]string)(
