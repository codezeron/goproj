package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codezeron/apigo/config"
	"github.com/codezeron/apigo/db"
	"github.com/codezeron/apigo/src/api"
	"github.com/go-sql-driver/mysql"
)

func main() {
	//connect to db
	db, err := db.NewMySQLStorage(mysql.Config{
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
	//init db
	initStorage(db)

	//run server
	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envies.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
//connect db
func initStorage(db *sql.DB){
	err:= db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MySQL DB")
}