package main

import(
	log "github.com/sirupsen/logrus"
)


func CountAttendance() error {
	student := &GenFirstname{}
	class := &GenClass{}
	professor := &GenProfessor{}
	m := make(map[string]map[string]int)

	totalStudents, err := a.GetAllAsString(first)
	if err != nil {
		log.Error(err)
		return err
	}

	buffer := make[string]int
	buffer["count"] = len(totalStudents)
	m[student.GetField()] = buffer

	classes, err := a.GetAllAsString(class)
	if err != nil {
		log.Error(err)
		return err
	}

	buffer, err := a.GetGenMap(classes, class)
	m[class.GetField()] = buffer

	professors, err := a.GetAllAsString(professor)
	if err != nil {
		log.Error(err)
		return err
	}

	buffer, err := a.GetGenMap(professors, professor)
	m[professor.GetField()] = buffer

}

// Point of this function is to generate a 2 x N csv file for the results.
func (a *App) GenerateOutFile() error {

}
