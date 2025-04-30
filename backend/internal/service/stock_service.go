package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/uranium092/stradivarius/backend/internal/external"
	"github.com/uranium092/stradivarius/backend/internal/repository"
)

type StockService interface {
	InitDataStock() error;
}

type stockService struct {
	repo repository.StockRepository
}

func (service *stockService) InitDataStock() error{
	//verify stock_status
	done, nextPage, err:=service.repo.GetStockStatus();
	if err!=nil{
		return err;
	}
  if done{
		return nil;
	}

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
		if nextPage != "" {
			paramHttpReq="?next_page="+nextPage;
		}
    bodyResponse,err:=external.DoRequestStock(baseURLHttpReq+paramHttpReq);
		if err!=nil {
			errorsAPI++;
			continue;
		}
    errQuery:=service.repo.InsertStockItems(bodyResponse.Items, tr);
		if errQuery!=nil {
			errorTransaction=errQuery;
			return fmt.Errorf("transaction error: %w", errQuery);
		}
		if bodyResponse.NextPage==""{
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

func NewStockService(r repository.StockRepository) StockService {
	return &stockService{repo: r}
}
