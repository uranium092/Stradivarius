package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uranium092/stradivarius/backend/internal/models"
)

type StockRepository interface {
	GetConnection() *pgxpool.Pool;
	GetStockStatus() (bool, string, error)
	InsertStockItems(items []models.ItemStock, tr pgx.Tx) error
}

type stockRepository struct {
	db *pgxpool.Pool
}

func (conn *stockRepository) GetStockStatus() (bool, string, error){
	var done bool;
	var nextPage string;
	err:=conn.db.QueryRow(context.Background(),"SELECT done, next_page FROM stock_status LIMIT 1").Scan(&done, &nextPage);
	if err!=nil {
		return false,"",err;
	}
	return done,nextPage,nil;
}

func (conn *stockRepository) InsertStockItems(items []models.ItemStock, tr pgx.Tx) error{
	baseQuery:="INSERT INTO stock (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, dateReleased) VALUES ";
	values:=[]interface{}{};

	for index,value:=range items{
				// base N to auto-increment
				basePlaceholder:=index*9;
				// auto-increment N $1, $2, ...
				baseQuery+=fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)", basePlaceholder+1, basePlaceholder+2, basePlaceholder+3, basePlaceholder+4, basePlaceholder+5, basePlaceholder+6, basePlaceholder+7, basePlaceholder+8, basePlaceholder+9);

				if index<len(items)-1{
					baseQuery+=",";
				}
				targetFrom, _:=strconv.ParseFloat(strings.ReplaceAll(value.TargetFrom,"$",""),64);
				targetTo, _:=strconv.ParseFloat(strings.ReplaceAll(value.TargetTo,"$",""),64);

				// add args (matching position con N...)
				values=append(values, value.Ticker, targetFrom, targetTo, value.Company, value.Action, value.Brokerage, value.RatingFrom, value.RatingTo, value.Time);
	}
	_,err:=tr.Exec(context.Background(),baseQuery,values...);
	if err!=nil{
		return err;
	}
	return nil;
}

func (conn *stockRepository) GetConnection() *pgxpool.Pool{
	return conn.db;
}

func NewStockRepository(db *pgxpool.Pool) StockRepository {
	return &stockRepository{db: db}
}