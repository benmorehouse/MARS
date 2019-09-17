package main

import(
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
final goal is to have a struct where it has counted each class that has been attended
*/
type db struct{
	num_columns int
	first_name string
	last_name string
	course string
	student_id string
	professor []string
	professor_attendence []int
}

type class struct{
	name string
	attendance int
}

type line struct{
	num_columns int
	first_name string
	last_name string

}

// Need to go and look at something that you can do with the csv file system management that they have in go
func fetchCSV(){
	csv_filename := flag.String("csv","data.csv", "csv data file")
	flag.Parse()
	file, err := os.Open(*csv_filename) // this will go and open the io.reader supplied by the flag
	file_reader := csv.NewReader(file) // file is io.reader that reads file
	column_description , err := file_reader.Read()


	//lines has all the information that we need now
	if err != nil{
		log.Fatal("Not able to read the inputted CSV file")
	}
	fmt.Println(column_description)
	fmt.Println(len(column_description))
	for i , val := range column_description{
		if strings.TrimSpace(strings.ToLower(val)) == "recipientlastname"{
			fmt.Println("Found the recipientlastname",i)
		}
	}



}

/* this is how we parse lines of code
func parse_lines(lines [][]string) []problem_set{
	returnVal := make([]problem_set,len(lines))
	for i, val := range lines{
		returnVal[i] = problem_set{
			question: val[0],
			answer: strings.TrimSpace(val[1]),
		}
	}
	return returnVal
}

*/

