package main

import(
	"database/sql"
	"errors"
)

type attendanceField interface{
	Exists() (bool, error)
}

type GenFirstname struct{
	
}

type AttendanceSQL struct{
	Firstname	string
	lastname	string
	Class		string
	Professor	string
}

// Called at the beginning of any instance to ensure existance.
func (a *App) CreateTableIfNotExists() {

	conf := a.Conf
	query := "Create table if not exists " + conf.DataTable
	query += `
	(
		firstname varchar(30), 
		lastname varchar(30), 
		class varchar(30), 
		professor varchar(30)
	);
	`
}

func (a *App) InsertAttendance(tableName string, s *AttendanceSQL) (error) {

	if s == nil {
		err := errors.New("Attendance is nil")
		log.Error(err)
		return err
	}

	c := a.Connection

	if err := c.Conn.PingContext(*c.Context); err != nil {
		log.Error(err)
		return err
	}

	attendanceQuery := "insert into " + tableName
	attendanceQuery += " (firstname, lastname, class, professor)"
	attendanceQuery += "values (" + s.Firstname + ", "
	attendanceQuery += s.Lastname + ", " + s.Class + ", "
	attendanceQuery += s.Professor + ");"

	result, err := c.Conn.ExecContext(*c.Context, attendanceQuery)
	if err != nil{
		log.Error(err)
		return err
	}

	return nil
}

func (a *App) GetAllProfessors() ([]string, error) {

	// Need to make a query, and then execute the query, then return the rows.
	c := a.Connection
	if err := c.Conn.PingContext(*c.Context); err != nil {
		return nil, err
	}

	query := "select professors from " + a.Conf.DataTable + ";"

	results, err := c.Conn.QueryContext(c.Context, query)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var professors []string
	var professor string
	for results.Next() {
		if err := results.Scan(professor); err == nil {
			professors = append(professors, professor)
		} else {
			log.Error(err)
		}
	}

	if len(professors) == 0 {
		err := errors.New("No professors found")
		log.Error(err)
		return nil, err
	}

	return professors, nil
}

func (a *App) CountProfessorAttendance(professors []string) (map[string]int, error) {

	if len(professors) == 0 {
		err := errors.New("No professors given")
		log.Error(err)
		return nil, err
	}

	c := a.Connection
	if err := c.Conn.PingContext(*c.Context); err != nil {
		return nil, err
	}

	m := make(map[string]int)
	for _, professor := range professors {

		// Create a query, then execute, then populate map.
		query := "select count(" + professor + ") from "
		query += a.Conf.DataTable + ";"

		result, err := c.Conn.ExecContext(c.Context, query)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		if count, ok := result.(int); ok {
			m[professor] = count
		} else {
			err := errors.New("result being return has type indifference")
			log.Error(err)
			return nil, err
		}
	}

	return m, nil
}

func (a *App) GetAllClasses() ([]string, error) {

	// Need to make a query, and then execute the query, then return the rows.
	c := a.Connection
	if err := c.Conn.PingContext(*c.Context); err != nil {
		return nil, err
	}

	query := "select classes from " + a.Conf.DataTable + ";"

	results, err := c.Conn.QueryContext(c.Context, query)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var classes []string
	var class string
	for results.Next() {
		if err := results.Scan(class); err == nil {
			professors = append(class, classes)
		} else {
			log.Error(err)
		}
	}

	if len(class) == 0 {
		err := errors.New("No classes found")
		log.Error(err)
		return nil, err
	}

	return classes, nil
}


