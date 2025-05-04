package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uranium092/stradivarius/backend/internal/models"
	"github.com/uranium092/stradivarius/backend/internal/service"
)

func getRequestQueries(c *gin.Context) (*models.RequestQueries, error){
	page,err:=strconv.Atoi(c.DefaultQuery("page","1"));
		if err!=nil{
			return nil,err;
		}
		return &models.RequestQueries{Page:page, Search: c.DefaultQuery("search",""),Sort: c.DefaultQuery("sort","")},nil;
}

func SetupStockHanlder(router *gin.RouterGroup, stockService service.StockService) {
	router.GET("/stock/all", func(c *gin.Context) {
		queries,err:=getRequestQueries(c);
		if err!=nil{
			c.Status(http.StatusBadRequest);
			return;
		}
		stock, totalPages, err:=stockService.GetStock(queries, "all");
		if err!=nil{ 
			c.Status(http.StatusInternalServerError);
			return;
		}
		c.JSON(http.StatusOK, gin.H{"dataStock":stock, "totalPages":totalPages});
	});
	router.GET("/stock/recommendation", func(c *gin.Context) {
		queries,err:=getRequestQueries(c);
		if err!=nil{
			c.Status(http.StatusBadRequest);
			return;
		}
		stock, totalPages, err:=stockService.GetStock(queries, "recommendation");
		fmt.Println(err)
		if err!=nil{ 
			c.Status(http.StatusInternalServerError);
			return;
		}
		c.JSON(http.StatusOK, gin.H{"dataStock":stock, "totalPages":totalPages});
	})
}