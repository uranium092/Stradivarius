package external

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/uranium092/stradivarius/backend/internal/models"
)

func DoRequestStock(baseURL string) (*models.StockResponsePopulation, error){
    //init clientHttp and create/custom Req
		client:=&http.Client{};
		req,err:=http.NewRequest("GET", baseURL, nil);
		if err!=nil {
			return nil, err;
		}
		req.Header.Add("Authorization", "Bearer "+os.Getenv("TOKEN"));
		
		//exec http req
		resp, err := client.Do(req);
		if err!=nil{
			return nil, err;
		}
		//close stream
		defer resp.Body.Close()

		if resp.StatusCode<200 || resp.StatusCode>300{
			return nil, fmt.Errorf("http error => %d", resp.StatusCode);
		}
		
		//decode JSON to struct
		var bodyResponse *models.StockResponsePopulation;
		err=json.NewDecoder(resp.Body).Decode(&bodyResponse);
		if err != nil{
			return nil, err;
		}
		return bodyResponse, nil;
}