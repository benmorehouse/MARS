package main

import(
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

type AppConf struct{
	DBName			string `json:"DBName"`
	DBUser			string `json:"DBUser"`
	DBPass			string `json:"DBPass"`
	DBPort			string `json:"DBPort"`
	DBIP			string `json:"DBIP"`
	DataTable		string `json:"DataTable"`
	OutFile			string `json:"OutFile"`
	InFile			string `json:"InFile"`
	IsDefault		bool   `json:"IsDefault"`
	StartDate		string `json:"StartDate"`
	StopDate		string `json:"StopDate"`
}

func loadConfig() (*AppConf, error) {

	defaultConfig := &AppConf {
		DBName: "MARSdb",
		DBUser: "benmorehouse",
		DBPass: "Moeller12!", // this is just my password.
		DBPort: "3306",
		DBIP: "127.0.0.1",
		DataTable: "attendance",
		InFile: "input",
		OutFile: "output",
		IsDefault: true,
		StartDate: getToday(),
		StopDate: getToday(),
	}

	jsonFile, err := os.Open("conf.json")
	defer jsonFile.Close()
	if err != nil{
		log.Error(err)
		return defaultConfig, err
	}

	config := AppConf{}
	confData, err := ioutil.ReadAll(jsonFile)
	if err != nil{
		log.Error(err)
		return defaultConfig, err
	}

	if err = json.Unmarshal(confData, &config); err != nil{
		log.Error(err)
		return defaultConfig, err
	}

	return config, nil
}

// A simple function to return a pretty printed date for default.
func getToday() (string) {

	t := time.Now()
	s := strings.Fields(t)
	return s[0]
}


