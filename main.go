package main

import (
	"fmt"
	"libraryManagement/api"
	"libraryManagement/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	err := db.ConnectPostgres(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

//	err = db.DropTables() // drop tables
	err = db.CreateTables() // create new tables
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(":3000")

	fmt.Println("server starting at :3000")
	server.RunServer()
}
