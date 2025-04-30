package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uranium092/stradivarius/backend/internal/db"
	"github.com/uranium092/stradivarius/backend/internal/services"
)

func main() {
	//load env vars
	if os.Getenv("GO_ENV") == "DEV" {
		err:=godotenv.Load(".ENV");
		if err != nil{
			log.Fatalf("Error on .ENV file: %v", err)
		 }
	}

	//connect to cockroachDB
	err:=db.InitConnection();
	if err != nil{
		log.Fatalf("Error on initConnection with DB => %v",err);
	}

	//stock initialization
	res:=services.InitDataStock();
	if res != nil{
		log.Fatalf("Error populating Stock. Restart to continue. Info => %v",res);
	}

}