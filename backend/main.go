package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uranium092/stradivarius/backend/internal/db"
	"github.com/uranium092/stradivarius/backend/internal/handler"
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

	// configure gin Mode based on env
	if os.Getenv("GO_ENV") != "DEV"{
		gin.SetMode(gin.ReleaseMode);
	}
	
	// define router and groups
	router:=gin.Default();
	
	// config CORS for environment type
	if os.Getenv("GO_ENV") == "DEV"{
		corsConfig:=cors.DefaultConfig();
		corsConfig.AllowOrigins = []string{"http://localhost:5173"};
		router.Use(cors.New(corsConfig))
	}

	// set handler for requests with /api
	handler.SetupStockHanlder(router.Group("/api"),stockService);

	router.Run(os.Getenv("PORT"));
}