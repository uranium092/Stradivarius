package external

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/uranium092/stradivarius/backend/internal/models"
)

func DoRequestStock(baseURL string) (models.StockResponse, error){
    //init clientHttp and create/custom Req
		client:=&http.Client{};
		req,err:=http.NewRequest("GET", baseURL, nil);
		if err!=nil {
			return models.StockResponse{}, err;
		}
		req.Header.Add("Authorization", "Bearer "+os.Getenv("TOKEN"));
		
		//exec http req
		resp, err := client.Do(req);
		if err!=nil{
			return models.StockResponse{}, err;
		}
		//close stream
		defer resp.Body.Close()

		if resp.StatusCode<200 || resp.StatusCode>300{
			return models.StockResponse{}, err;
		}
		
		//decode JSON to struct
		var bodyResponse models.StockResponse;
		decoded:=json.NewDecoder(resp.Body).Decode(&bodyResponse);
		if decoded != nil{
			return models.StockResponse{}, err;
		}
		return bodyResponse, nil;
}