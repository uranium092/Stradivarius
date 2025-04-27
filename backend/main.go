package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uranium092/stradivarius/backend/db"
)

func main() {
	//load env vars
	if os.Getenv("GO_ENV") == "DEV" {
		err:=godotenv.Load(".ENV");
		if err != nil{
			log.Fatalf("Error on .ENV file: %s", err)
		 }
	}

	//connect to cockroachDB
	err:=db.InitConnection();
	if err != nil{
		log.Fatalf("Error on initConnection with DB => %s",err);
	}
}