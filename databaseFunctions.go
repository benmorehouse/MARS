package main

import(
	"fmt"
	"strconv"
	"os"
	"errors"
	"strings"
)

type db struct{
	NumColumns int
	DesignatedColumns []int
	ColumnsExist map[string]bool // a set that tells us all the possible columns the user could choose from 
	ColumnData []string
	ColumnCount []int
}

func (this *db) AssignColumns(columns []string)error{
	for i:=0;i<this.NumColumns;i++{
		this.ColumnsExist[columns[i]] = true
		fmt.Println(strconv.Itoa(i) + string(". ") + columns[i] + ", ")
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
		}else if index >= len(columns) || index < 0{
			fmt.Println("What you have entered is an invalid number: out_of_range")
		}

		if this.ColumnsExist[(columns[index])] == false{
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

func (this *db) AssignDefault(columns []string)error{
	if this.NumColumns != 29{
		return errors.New("Error within given CSV file input: column Size")
	}else{
		for i:=0;i<this.NumColumns;i++{
			this.ColumnsExist[columns[i]] = true
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

func (this *db) ParseData(data []string,channelCount int, c chan int){ // in this i need to pass in the file_reader, a channel, and return error
	defer wg.Done()
	if c == nil{
		fmt.Println("Developer error: channel memory no longer valid")
		c<-channelCount
		return
	}else if len(data) == 0{
		c<-channelCount
		return
	}else{
		for i:=0;i<len(this.DesignatedColumns)-1;i++{
			index := this.DesignatedColumns[i]
			if index >= len(data) || index < 0{
				continue
			}else if data[index] == ""{
				continue
			}else{
				this.ColumnCount[i]++
			}
		}
		c<-channelCount
	}
}

