package main

import(
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type db struct{
	NumColumns int
	DesignatedColumns []int
	ColumnsExist map[string]bool // a set that tells us all the possible columns the user could choose from 
	ColumnData []string
}

// Need to go and look at something that you can do with the csv file system management that they have in go
func fetchCSV(){
	csv_filename := flag.String("file","", "csv data file")
	_default := flag.Bool("default",false, "tells us whether they want to customize which rows to get or to just pass in defulat  configuration ")
	flag.Parse()

	if *csv_filename == ""{
		log.Fatal("You have not passed in a csv file")
	}

	file, err := os.Open(*csv_filename)

	if err != nil{
		log.Fatal("Not able to open the inputted CSV file")
	}

	file_reader := csv.NewReader(file) // file is io.reader that reads file

	column_description , err := file_reader.Read() // reads the first line of the csv file

	if err != nil{
		log.Fatal("Not able to read the inputted CSV file")
	}

	var marshalldb = db{
		NumColumns: len(column_description),
		ColumnsExist: make(map[string]bool),
	}

	columns , err := file_reader.Read()

	if err != nil{
		log.Fatal("The file only has one row of data")
	}


	for i:=0;i<marshalldb.NumColumns;i++{
		if i == marshalldb.NumColumns - 1{
			fmt.Println(columns[i])
			marshalldb.ColumnsExist[columns[i]] = true
			break
		}
		marshalldb.ColumnsExist[columns[i]] = true
		fmt.Print(string(i) + string(". ") + columns[i] + ", ")
	}

	if *_default == false{
		err = marshalldb.AssignColumns(columns) // need to pass in columns
		if err != nil{
			log.Fatal("Developer Error: database getColumns function returned non-nil")
		}
	}else{
		// we want index 17, 18 19-23 is which class they were in 24-28 shows which professor that they are going for
		err = marshalldb.AssignDefault()
		if err != nil{
			log.Fatal("Input CSV file is not of correct length")
		}
	}


}

