package main

import(
	"database/sql"
	"errors"
)

type AttendanceField interface{
	GetData() (string)
	Exists() (bool, error)
	GetField() (string)
}

type GenFirstname struct{
	Firstname	string
	Exists		bool
}

func (g *GenFirstname) GetData() (string) {
	return g.Firstname
}

func (g *GenFirstname) Exists() (string) {
	return g.Exists
}

func (g *GenFirstname) GetField() (string) {
	return "firstname"
}

type GenLastname struct{
	Lastname	string
	Exists		bool
}

func (g *GenLastname) GetData() (string) {
	return g.Firstname
}

func (g *GenLastname) Exists() (string) {
	return g.Exists
}

func (g *GenLastname) GetField() (string) {
	return "lastname"
}

type GenClass struct{
	Class		string
	Exists		bool
}

func (g *GenClass) GetData() (string) {
	return g.Class
}

func (g *GenClass) Exists() (string) {
	return g.Exists
}

func (g *GenClass) GetField() (string) {
	return "class"
}

type GenProfessor struct{
	Professor	string
	Exists		bool
}

func (g *GenProfessor) GetData() (string) {
	return g.Professor
}

func (g *GenProfessor) Exists() (string) {
	return g.Exists
}

func (g *GenProfessor) GetField() (string) {
	return "professor"
}

type AttendanceSQL struct{
	Firstname	*GenFirstname
	lastname	*GenLastname
	Class		*GenClass
	Professor	*GenProfessor
}

// Called at the beginning of any instance to ensure existance.
func (a *App) CreateTableIfNotExists() error {

	conf := a.Conf
	c := a.Connection
	query := "Create table if not exists " + conf.DataTable
	query += `
	(
		firstname varchar(30), 
		lastname varchar(30), 
		class varchar(30), 
		professor varchar(30)
	);
	`

	result, err := c.Conn.ExecContext(*c.Context, query)
	if err != nil{
		log.Error(err)
		return err
	}

	return nil
}

// A function to add important information from csv file into database.
func (a *App) InsertAttendanceRow(s *AttendanceSQL) (error) {

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

	tableName := a.Conf.DataTable

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

// Will take in attendance field and return an array of all present results. 
// Good for the use of creating the map.
func (a *App) GetAllAsString(g *AttendanceField) ([]string, error) {

	// Need to make a query, and then execute the query, then return the rows.
	c := a.Connection
	if err := c.Conn.PingContext(*c.Context); err != nil {
		return nil, err
	}

	query := "select " + g.GetField + " from " + a.Conf.DataTable + ";"

	results, err := c.Conn.QueryContext(c.Context, query)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var gens []string
	var gen string
	for results.Next() {
		if err := results.Scan(gen); err == nil {
			gens = append(gens, gen)
		} else {
			log.Error(err)
		}
	}

	if len(gens) == 0 {
		err := errors.New("No Data Found")
		log.Error(err)
		return nil, err
	}

	return gens, nil
}

// Will get a map of all the values and how much they appear for each gen
func (a *App) GetGenMap(gs []string, g *AttendanceField) (map[string]int, error) {

	if len(gs) == 0 {
		err := errors.New("No gens given")
		log.Error(err)
		return nil, err
	}

	c := a.Connection
	if err := c.Conn.PingContext(*c.Context); err != nil {
		return nil, err
	}

	m := make(map[string]int)
	for _, gen := range gs {

		// Create a query, then execute, then populate map.
		query := "select count(" + g.GetField() + ") from "
		query += a.Conf.DataTable + ";"

		result, err := c.Conn.ExecContext(c.Context, query)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		if count, ok := result.(int); ok {
			m[gen] = count
		} else {
			err := errors.New("result being return has type indifference")
			log.Error(err)
			return nil, err
		}
	}

	return m, nil
}


