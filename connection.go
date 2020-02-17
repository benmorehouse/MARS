package main

import(
	"database/sql"
	"context"

	log "github.com/sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

// create table attendance( firstname varchar(30), lastname varchar(30), class varchar(30), professor varchar(30) )

type Connection struct{
	Conn		*sql.Conn //unexported connection to database
	Context		*context.Context
}

func (a *App) Connect() (error) {

	cxn := a.Conf
	connectDB := cxn.DBUser + ":" + cxn.DBPass + "@tcp("
	connectDB += cxn.DBIP + ":" + cxn.DBPort + ")"
	connectDB += cxn.DBName

	db, err := sql.Open("mysql", connectDB)
	if err != nil {
		log.Error(err)
		return err
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}

	connection := Connection {
		Conn: &conn,
		Context: context.Background()
	}

	a.Cxn = connection
}


