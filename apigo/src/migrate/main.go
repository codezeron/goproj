package main

import (
	"log"
	"os"

	"github.com/codezeron/apigo/config"
	"github.com/codezeron/apigo/db"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)
func main(){
	//connect to db
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:                 config.Envies.DBUser,
		Passwd:               config.Envies.DBPasswd,
		Addr:                 config.Envies.DBAddress,
		DBName:               config.Envies.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	//inicializar o driver
	driver, err:= mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}


	m, err := migrate.NewWithDatabaseInstance(
		"file://src/migrate/migrations", 
		"mysql", 
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	//comando para o cli
	cmd := os.Args[(len(os.Args)-1)]
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	default:
		log.Fatal("invalid command")
	}


}