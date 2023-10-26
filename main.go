package main

import (
	"log"
	"survivor_fantasy/db"
	"survivor_fantasy/web"
)

func main() {
	log.Printf("Survivor fantasy server starting")

	con, err := db.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	web.Run(con)

}
