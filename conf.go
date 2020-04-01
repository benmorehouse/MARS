package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type AppConf struct {
	DBName    string `json:"DBName"`
	DBUser    string `json:"DBUser"`
	DBPass    string `json:"DBPass"`
	DBPort    string `json:"DBPort"`
	DBIP      string `json:"DBIP"`
	DataTable string `json:"DataTable"`
	OutFile   string `json:"OutFile"`
	InFile    string `json:"InFile"`
	IsDefault bool   `json:"IsDefault"`
	StartDate string `json:"StartDate"`
	StopDate  string `json:"StopDate"`
}

func (a *App) loadConfig() error {

	handleError := func(a *App) {

		t := time.Now()
		s := strings.Fields(t)
		today := s[0]

		defaultConfig := &AppConf{
			DBName:    "MARSdb",
			DBUser:    "benmorehouse",
			DBPass:    "Moeller12!", // this is just my password.
			DBPort:    "3306",
			DBIP:      "127.0.0.1",
			DataTable: "attendance",
			InFile:    "input",
			OutFile:   "output",
			IsDefault: true,
			StartDate: today,
			StopDate:  today,
		}

		log.Warning("Using native default configuration.")
		a.Conf = defaultConfig
	}

	// NOTE: need to create conf.json as default and use a.confFileName
	jsonFile, err := os.Open("conf.json")
	defer jsonFile.Close()
	if err != nil {
		handleError(a)
		return err
	}

	config := AppConf{}
	confData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		handleError(a)
		return err
	}

	if err = json.Unmarshal(confData, &config); err != nil {
		handleError(a)
		return err
	}

	a.Conf = config
	return nil
}

// A simple function to return a pretty printed date for default.
