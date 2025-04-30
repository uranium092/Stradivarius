package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/uranium092/stradivarius/backend/internal/db"
)

type itemStock struct {
	Ticker string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo string `json:"target_to"`
	Company string `json:"company"`
	Action string `json:"action"`
	Brokerage string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo string `json:"rating_to"`
	Time string `json:"time"`
}

type stockResponse struct {
	Items []itemStock `json:"items"`
	NextPage string  `json:"next_page"`
}

type utilsForExec struct{
	query string
	args []interface{}
}

func createUtilsForExec(items []itemStock) (utilsForExec){
	baseQuery:="INSERT INTO stock (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, dateReleased) VALUES ";
	values:=[]interface{}{};
	for index,value:=range items{
		basePlaceholder:=index*9;

		baseQuery+=fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)", basePlaceholder+1, basePlaceholder+2, basePlaceholder+3, basePlaceholder+4, basePlaceholder+5, basePlaceholder+6, basePlaceholder+7, basePlaceholder+8, basePlaceholder+9);

		if index<len(items)-1{
			baseQuery+=",";
		}

		targetFrom, err:=strconv.ParseFloat( strings.ReplaceAll(value.TargetFrom,"$","") ,64);
		if err != nil{
			targetFrom=0;
		}
		targetTo, err:=strconv.ParseFloat( strings.ReplaceAll(value.TargetTo,"$","") ,64);
		if err != nil{
			targetTo=0;
		}
		values=append(values, value.Ticker, targetFrom, targetTo, value.Company, value.Action, value.Brokerage, value.RatingFrom, value.RatingTo, value.Time);
	}
	return utilsForExec{query: baseQuery, args: values};
}

func InitDataStock() error {
	conn:=db.GetConnection();
	if conn==nil{
		return errors.New("no DB connection present");
	}
	
	//verify stock_status
	sql:="SELECT done, next_page FROM stock_status LIMIT 1";
	var done bool;
	var nextPage string;
	err:=conn.QueryRow(context.Background(),sql).Scan(&done, &nextPage)
	if err != nil{
		return err;
	}
	if done {
		return nil;
	}
	
	//continue process of Stock population
	baseURLHttpReq:="https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list";
	errorsApiHttp:=0;
	var errorTransaction error;
	t,err:=conn.Begin(context.Background());

	defer func(){
		if errorTransaction!=nil{
			t.Rollback(context.Background());
		}
	}();
	

	if err!=nil{
		return err;
	}

	for {

		if errorsApiHttp>=3{
			break;
		}

		paramHttpReq:="";
		if nextPage != "" {
			paramHttpReq="?next_page="+nextPage;
		}

		//init clientHttp and create/custom Req
		client:=&http.Client{};
		req,err:=http.NewRequest("GET", baseURLHttpReq+paramHttpReq, nil);
		if err!=nil {
			errorsApiHttp++;
			continue;
		}
		req.Header.Add("Authorization", "Bearer "+os.Getenv("TOKEN"));
		
		//exec http req
		resp, err := client.Do(req);
		if err!=nil || resp.StatusCode<200 || resp.StatusCode>300{
			resp.Body.Close();
			errorsApiHttp++;
			continue;
		}
		
		//decode JSON to struct
		var bodyResponse stockResponse;
		decoded:=json.NewDecoder(resp.Body).Decode(&bodyResponse);
		if decoded != nil{
			resp.Body.Close();
			errorsApiHttp++;
			continue;
		}
		//close stream
		resp.Body.Close();

		dataForExec:=createUtilsForExec(bodyResponse.Items);

		_,errSQL:=t.Exec(context.Background(),dataForExec.query, dataForExec.args...);

		if errSQL!=nil{
			errorTransaction=errSQL;
			return fmt.Errorf("transaction error: %w", errSQL);
		}

		if bodyResponse.NextPage==""{
			break;
		}
		
		nextPage=bodyResponse.NextPage;
	}

	var statusToReturn error=nil;
	statusDone:=true;
	if errorsApiHttp>=3{
		statusToReturn=errors.New("stock server is not working");
		statusDone=false;
	}
	_,errSQL:=t.Exec(context.Background(),"UPDATE stock_status SET done=$1, next_page=$2",statusDone,nextPage);

	if errSQL!=nil{
		errorTransaction=errSQL;
		return fmt.Errorf("transaction error: %w", errSQL);
	}

	if err:=t.Commit(context.Background());err!=nil{
		return err;
	}

	return statusToReturn;
}