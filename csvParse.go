package main

import(
	"encoding/csv"
	"flag"
	"log"
	"os"
	"sync" // using waitGroup is great for keeping track of all the goroutines that you need
	"fmt"
)
// Need to go and look at something that you can do with the csv file system management that they have in go


func fetchCSV(){
	csv_filename := flag.String("file","", "csv data file")
	_default := flag.Bool("default",false, "tells us whether they want to customize which rows to get or to just pass in defulat  configuration ")
	flag.Parse()

	wg := &sync.WaitGroup{}

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

	if *_default == false{
		err = marshalldb.AssignColumns(columns) // need to pass in columns
		if err != nil{
			log.Fatal("Developer Error: database getColumns function returned non-nil")
		}
	}else{
		// we want index 17, 18 19-23 is which class they were in 24-28 shows which professor that they are going for
		err = marshalldb.AssignDefault(columns)
		if err != nil{
			log.Fatal("Input CSV file is not of correct length")
		}
	}
	// at this point we have the Assigned Columns updated within the database
	data , err := file_reader.Read()
	for err == nil{
		wg.Add(1)
		go marshalldb.ParseData(wg, data)
		data , err = file_reader.Read()
	}
	// at this point the length of the channel is 0
	wg.Wait()
	fmt.Println("Test")
}
