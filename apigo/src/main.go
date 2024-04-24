package main

import (
	"log"

	"github.com/codezeron/apigo/prod/api"
)

func main(){
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil{
		log.Fatal(err)
	}
}