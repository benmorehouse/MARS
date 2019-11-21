package main

import(
	"fmt"
	"strconv"
	"os"
	"errors"
	"strings"
	"sync"
)

type db struct{
	NumColumns int
	DesignatedColumns []int
	ColumnsExist map[string]bool // a set that tells us all the possible columns the user could choose from 
	ColumnCount []int
	ColumnDescription []string
}

func (this *db) AssignColumns()error{
	for i:=0;i<this.NumColumns;i++{
		this.ColumnsExist[this.ColumnDescription[i]] = true
		fmt.Println(strconv.Itoa(i) + string(". ") + this.ColumnDescription[i] + ", ")
	}

	fmt.Println("Pick which numbered columns you would like tallied")
	fmt.Println("Type done when you are finished")

	input := ""
	fmt.Scan(&input)

	if strings.ToLower(input) == "done"{
		os.Exit(1) // they entered done right away and we have nothing to do 
	}

	for strings.ToLower(input) != "done"{
		index , err := strconv.Atoi(input)

		if err != nil{
			fmt.Println("What you have entered is an invalid number: type_differentiation")
		}else if index >= len(this.ColumnDescription) || index < 0{
			fmt.Println("What you have entered is an invalid number: out_of_range")
		}

		if this.ColumnsExist[(this.ColumnDescription[index])] == false{
			// this means what they entered doesnt exist somehow
			fmt.Println("You have entered a column that is not available")
		}else{
			this.DesignatedColumns = append(this.DesignatedColumns,index)
		}
		fmt.Scan(&input)
	}

	for i:=0;i<len(this.DesignatedColumns);i++{
		this.ColumnCount = append(this.ColumnCount, 0)
	}

	return nil
}
/*
This assigns everything to default 

*/
func (this *db) AssignDefault()error{
	if this.NumColumns != 29{
		fmt.Println(this.NumColumns)
		return errors.New("Error within given CSV file input: column Size")
	}else{
		for i:=0;i<this.NumColumns;i++{
			this.ColumnsExist[this.ColumnDescription[i]] = true
		}
		for i:=17;i<29;i++{
			this.DesignatedColumns = append(this.DesignatedColumns,i)
		}
	}
	for i:=0;i<len(this.DesignatedColumns);i++{
		this.ColumnCount = append(this.ColumnCount, 0)
	}
	return nil
}
/*
this will parse each line of the file and total it up depending on what has been passed through as important
this is where i can change the column description and such 
*/

func (this *db) clearTable(sqlDatabase *sql.DB, tableName string){
	err = sqlDatabase.exec("delete from " + tableName)
	if err != nil{
		panic(err)
	}
}

func (this *db) ParseData(wg *sync.WaitGroup, row []string, sqlDatabase *sql.DB, query string)(error){
	if *wg == nil || len(data) == 0 || sqlDatabase == nil{
		return
	}else{
		defer wg.Done()
		//insert into data (first, last) values ("Ben","Morehouse"); // make this before and then pass it in 
		for i, val := range row{
			rowField := strings.Fields(val)
			if i == len(row)-1{
				insertRowInTableString += strings.Join(rowField,"") + ");"
				break
			}else{
				insertRowInTableString += strings.Join(rowField,"") + " , "
			}
		}
		err = sqlDatabase.exec(insertRowInTableString)
		if err != nil{
			this.clearTable(sqlDatabase, tableName)
			log.Fatal("Halted due to insert error")
		}
	}
}

/*
In this function we want to total up each of the columns that we have been counting and display them 
*/

func (this *db) PushData()[][]string{ // want it to a 2d array of strings
	var output [][]string
	var line []string
	for i , val := range this.ColumnCount{ // columnCount is not correct for last values
		line = append(line,this.ColumnDescription[this.DesignatedColumns[i]])
		line = append(line,strconv.Itoa(val))
		output = append(output, line)
		line = nil
	}
	// need to figure out how to get map range function working
/*
	professors.Range(f (key string, value int) bool{
		temp, exists := professors.Load(key)
		if !exists{
			return false
		}
		line = append(line, key)
		line = append(line,strconv.Itoa(value))
		output = append(output, line)
		line = nil
		return true
	}
*/
	return output
}

