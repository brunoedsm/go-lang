package main

import (	
	"log"
	"brunoedsm.com/v1/services"
)

func main() {

	//RESTful API start
	log.Fatal(services.ServeAPI("localhost:8181", "tcp:integ-db.azure.com:5024"))
}