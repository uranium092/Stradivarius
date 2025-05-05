package service

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/jackc/pgx/v5"
	"github.com/uranium092/stradivarius/backend/internal/external"
	"github.com/uranium092/stradivarius/backend/internal/models"
	"github.com/uranium092/stradivarius/backend/internal/repository"
)

type StockService interface {
	InitDataStock() error;
	GetStock(queries models.RequestQueries, mode string) (models.StockResponse, error);
}

type stockService struct {
	repo repository.StockRepository
}

func (service *stockService) InitDataStock() error{
	//verify stock_status
	stockStatus, err:=service.repo.GetStockStatus();
	if err!=nil{
		return err;
	}
  if stockStatus.Done{
		return nil;
	}

	nextPage:=stockStatus.NextPage; // continue stock population with last report of nextPage

	//continue process of Stock population
	baseURLHttpReq:="https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list";
	errorsAPI:=0;
	var errorTransaction error;
	tr, err:=service.repo.GetConnection().Begin(context.Background());
	if err!=nil{
		return err;
	}

	//abort changes in DB if something crash with transaction
	defer func(){
		if errorTransaction!=nil{
			tr.Rollback(context.Background());
		}
	}();

	for{
		//controlled errors: API Response
		if errorsAPI>=3{
			break;
		}
		paramHttpReq:="";
		if nextPage != "" { //determine next_page query
			paramHttpReq="?next_page="+nextPage;
		}
    bodyResponse,err:=external.DoRequestStock(baseURLHttpReq+paramHttpReq); //exec Http Request
		if err!=nil {
			errorsAPI++; //inc errors API counter and try with the same page
			continue;
		}
    errQuery:=service.repo.InsertStockItems(bodyResponse.Items, tr);
		if errQuery!=nil {
			errorTransaction=errQuery; //mark error for this transaction. Defer gonna make Rollback.
			return fmt.Errorf("transaction error: %w", errQuery);
		}
		if bodyResponse.NextPage==""{ //finish stock population
			break;
		}
		nextPage=bodyResponse.NextPage;
	}

		//validate if API failed
		var status error=nil;
		if errorsAPI>=3{
			status=errors.New("stock server is not working");
		}
		//save status and stock population progress
		_,errQuery:=tr.Exec(context.Background(),"UPDATE stock_status SET done=$1, next_page=$2",status==nil,nextPage);
		if errQuery!=nil{
			errorTransaction=errQuery;
			return fmt.Errorf("transaction error: %w", errQuery);
		}
		//save all or nothing
		if err:=tr.Commit(context.Background());err!=nil{
			return err;
		}
	
		return status;
}

func (service *stockService) GetStock(queries models.RequestQueries, mode string) (models.StockResponse,error){
	var rows pgx.Rows;
	var err error;
	if mode=="all"{ //validate what search type It gonna use
		rows,err=service.repo.GetAllStock(queries);
		if err!=nil{
			return models.StockResponse{}, fmt.Errorf("error retrieving entire Stock -> %w",err);
		}
	}else if mode=="recommendation"{
		rows,err=service.repo.GetRecommendation(queries);
		if err!=nil{
			return models.StockResponse{}, fmt.Errorf("error retrieving recommendation Stock -> %w",err);
		}
	}
	stock :=[]models.ItemStock{}; //save the resulting records of prev search
	var count int; //amount records; useful for pagination
	for rows.Next(){
		var row models.ItemStock;
		err=rows.Scan(&count, &row.Id, &row.Ticker, &row.TargetFrom, &row.TargetTo, &row.Company, &row.Action, &row.Brokerage, &row.RatingFrom, &row.RatingTo, &row.Time);
		if err!=nil{
			return models.StockResponse{}, fmt.Errorf("error processing row -> %w",err);
		}
		stock=append(stock, row); //save each record
	}
	//assuming 25 records per page, we find total pages, rounding to the nearest integer(page), if last page contains < 25 records
	totalPages:=math.Ceil( float64(count)/float64(25) );
	return models.StockResponse{TotalPages: totalPages, DataStock: stock}, nil;
}

func NewStockService(r repository.StockRepository) StockService {
	return &stockService{repo: r}
}
