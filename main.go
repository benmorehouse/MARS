package main

import(
	log "github.com/sirupsen/logrus"
)

type App struct {
	Conf *AppConf
	Cxn  *Connection
}

func NewApp() (*App, error) {

	a := new(App)
	// It is okay if loading the config returns an error. 
	// I just leave a warning and use a default config.
	if err := a.loadConfig(); err != nil {
		log.Warning(err)
	}

	if err := a.Connect(); err != nil {
		log.Error(err)
		return nil, err
	}

	return a, nil
}

func main(){

	a, err := NewApp()
	if err != nil {
		log.Error(err)
		return
	}

	err := a.Feed()
	if err != nil {
		log.Error(err)
		return
	}

}


