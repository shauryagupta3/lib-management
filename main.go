package main

import (
	"fmt"
	"os"
)

func main() {
	// postgres db connection
	dbConnection, err := initPostgres()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	// creating db tables
	if err := dbConnection.createTables();err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	
	// server connection
	server := NewApiServer(":3000", *dbConnection)
	server.runServer()


	defer dbConnection.conn.Close()
}
