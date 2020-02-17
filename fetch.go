package main

import(
	"bufio"
	"encoding/csv"
	log "github.com/sirupsen/logrus"
)


func (a *App) CountAttendance() (map[string]map[string]int, error) {

	student := &GenFirstname{}
	class := &GenClass{}
	professor := &GenProfessor{}
	m := make(map[string]map[string]int)

	totalStudents, err := a.GetAllAsString(first)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	buffer := make[string]int
	buffer["count"] = len(totalStudents)
	m[student.GetField()] = buffer

	classes, err := a.GetAllAsString(class)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	buffer, err := a.GetGenMap(classes, class)
	m[class.GetField()] = buffer

	professors, err := a.GetAllAsString(professor)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	buffer, err := a.GetGenMap(professors, professor)
	m[professor.GetField()] = buffer
	return m, nil
}

func generateAttendance(map[string]map[string]int) [][]string {

	var outer [][]string
	for _, gen := range m { // for each category we go through.
		for gen, count := range gen {
			var inner []string
			inner = append(inner, gen)
			inner = append(inner, count)
			outer = append(outer, inner)
		}
	}

	return outer
}

// Point of this function is to generate a 2 x N csv file for the results.
func (a *App) GenerateOutFile(m map[string]map[string]int) error {

	data := generateAttendance(m)
	oFile, err := os.Create(a.Conf.OutFile)
	if err != nil {
		log.Error(err)
		return err
	}

	writer := csv.NewWriter(oFile)
	if err := writer.WriteAll(data); err != nil {
		log.Error(err)
		return err
	}

	writer.Flush()
	if writer.Error() != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (a *App) GenerateStdOut(m map[string]map[string]int) {

	data := generateAttendance(m)
	fmt.Println("\n\n\n" + data + "\n\n\n")
}


