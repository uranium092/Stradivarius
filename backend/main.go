package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uranium092/stradivarius/backend/internal/db"
	"github.com/uranium092/stradivarius/backend/internal/repository"
	"github.com/uranium092/stradivarius/backend/internal/service"
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
	conn,err:=db.InitConnection();
	if err != nil{
		log.Fatalf("Error on initConnection with DB => %v",err);
	}

	//init layers
	stockRepository:=repository.NewStockRepository(conn);
	stockService:=service.NewStockService(stockRepository);

	//stock initialization
	res:=stockService.InitDataStock();
	if res != nil{
		log.Fatalf("Error populating Stock. Restart to continue. Info => %v",res);
	}
	
}