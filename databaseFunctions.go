package main

import(
	"fmt"
	"strconv"
	"os"
	"errors"
)

func (this *db) AssignColumns(columns []string)error{
	fmt.Println("Pick which numbered columns you would like tallied")
	fmt.Println("Type done when you are finished")

	input := ""
	fmt.Scan(&input)

	if input == "done"{
		os.Exit(1) // they entered done right away and we have nothing to do 
	}

	for input != "done"{
		index , err := strconv.Atoi(input)

		if err != nil{
			fmt.Println("What you have entered is an invalid number: type_differentiation")
		}else if index >= len(columns) || indexHelper < 0{
			fmt.Println("What you have entered is an invalid number: out_of_range")
		}

		if this.ColumnsExist[strconv.Atoi(columns[index])] == false{
			// this means what they entered doesnt exist somehow
			fmt.Println("You have entered a column that is not available")
		}else{
			this.DesignatedColumns = append(this.DesignatedColumns,int(input))
		}
		fmt.Scan(&input)
	}

	return nil
}

func (this *db) AssignDefault()error{
	if this.NumColumns != 29{
		return errors.New("Error within given CSV file input: column Size")
	}else{
		for i:=17;i<29;i++{
			this.DesignatedColumns = append(this.DesignatedColumns,i)
		}
	}
	return nil
}

