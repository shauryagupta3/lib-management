package main

import (
	"context"
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
	defer dbConnection.conn.Close(context.Background())

	//server connection
	server := NewApiServer(":3000", *dbConnection)
	server.runServer()
}
