package handler

import (
	"net/http"
	"strconv"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/uranium092/stradivarius/backend/internal/apperrors"
	"github.com/uranium092/stradivarius/backend/internal/models"
	"github.com/uranium092/stradivarius/backend/internal/service"
)

func getRequestQueries(c *gin.Context) (models.RequestQueries, error){
	page,err:=strconv.Atoi(c.DefaultQuery("page","1"));
		if err!=nil{
			return models.RequestQueries{},err;
		}
		return models.RequestQueries{Page:page, Search: c.DefaultQuery("search",""),Sort: c.DefaultQuery("sort","")},nil;
}

func SetupStockHanlder(router *gin.RouterGroup, stockService service.StockService) {
	router.GET("/stock/all", func(c *gin.Context) {
		queries,err:=getRequestQueries(c);
		if err!=nil{
			c.Status(http.StatusBadRequest); // 400 status code
			return;
		}
		data, err:=stockService.GetStock(queries, "all");
		if err!=nil{
			if errors.Is(err, apperrors.ErrBadRequest){
				c.Status(http.StatusBadRequest); // 400 status code
				return;
			}
			c.Status(http.StatusInternalServerError); // 500 status code
			return;
		}
		c.JSON(http.StatusOK, data);
	});
	router.GET("/stock/recommendation", func(c *gin.Context) {
		queries,err:=getRequestQueries(c);
		if err!=nil{
			c.Status(http.StatusBadRequest); // 400 status code
			return;
		}
		data, err:=stockService.GetStock(queries, "recommendation");
		if err!=nil{
			if errors.Is(err, apperrors.ErrBadRequest){
				c.Status(http.StatusBadRequest); // 400 status code
				return;
			}
			c.Status(http.StatusInternalServerError); // 500 status code
			return;
		}
		c.JSON(http.StatusOK, data);
	})
}