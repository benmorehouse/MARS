package main

import(
	"fmt"
	"encoding/csv"
	"flag"
	"log"
	"os"
	"sync" // using waitGroup is great for keeping track of all the goroutines that you need
	"strings"
)
// Need to go and look at something that you can do with the csv file system management that they have in go


func fetchCSV(){
	csv_filename := flag.String("file","", "csv data file")
	_default := flag.Bool("default",false, "tells us whether they want to customize which rows to get or to just pass in defulat  configuration ")
	output_filename := flag.String("outputFile","output.csv","the output file to genarate") // needs to be checked
	stop_date := flag.String("stopDate","","the specific date to stop... must be in format 2019-09-09 (year, month, date))")
	start_date := flag.String("startDate","","the specific date to start... csv reads from oldest to newest ... must be in format 2019-09-09 (year, month, date))")

	flag.Parse()

	wg := &sync.WaitGroup{}

	if *csv_filename == ""{
		log.Fatal("You have not passed in a csv file")
	}else if *start_date ==""{
		log.Fatal("You have not passed in a start_date. This needs to be in the form of numbers:2019-09-09")
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

	columns , err := file_reader.Read() // pushes cursor to next line 

	if err != nil{
		log.Fatal("The file only has one row of data")
	}

	marshalldb.ColumnDescription = columns

	if *_default == false{
		err = marshalldb.AssignColumns() // need to pass in columns
		if err != nil{
			log.Fatal("Developer Error: database getColumns function returned non-nil")
		}
	}else{
		// we want index 17-23 is which class they were in 24-28 shows which professor that they are going for
		err = marshalldb.AssignDefault()
		if err != nil{
			log.Fatal("Input CSV file is not of correct length")
		}
	}
	// at this point we have the Assigned Columns updated within the database
	data , err := file_reader.Read()
	temp := data[0]
	tempField := strings.Fields(temp)

	for tempField[0] != *start_date{
		data , err := file_reader.Read()
		if err != nil{
			log.Fatal(err)
		}
		temp = data[0]
		tempField = strings.Fields(temp)
	}

	professors := make(map[string]int)

	for err == nil{
		temp := data[0]
		tempField := strings.Fields(temp)
		if tempField[0] == *stop_date{
			break
		}else{
			wg.Add(1)
			go marshalldb.ParseData(wg, data, professors)
			data , err = file_reader.Read()
		}
	}
	wg.Wait()

	// at this point we need to make a CSV file output

	outputFile , err := os.Create(*output_filename)

	if err != nil{
		log.Fatal("Cant write to ",*output_filename,":",err)
	}

	writer := csv.NewWriter(outputFile)

	if writer == nil{
		log.Fatal("writer is nil")
	}

	err = writer.WriteAll(marshalldb.PushData(professors))

	if err != nil{
		log.Fatal("Couldnt write the output csv file:",err)
	}

	fmt.Println("File writing finished")
}
