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
	GetAllStock(queries *models.RequestQueries) (pgx.Rows, error)
	GetRecommendation(queries *models.RequestQueries) (pgx.Rows, error)
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
	if len(items)==0{
		return nil;
	}
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

func (conn *stockRepository) buildSQLClause(queries *models.RequestQueries, mode string) (string, []interface{}){

	//WHERE (search)
	baseSQLClause:="";
	baseSQLArgs:=[]interface{}{};
	if queries.Search!=""{
		expression:=" WHERE";
		if mode=="recommendation"{
			expression=" AND";
		}
		baseSQLClause+=fmt.Sprintf(" %s ticker ILIKE $1 OR company ILIKE $2 OR action ILIKE $3 OR brokerage ILIKE $4 OR rating_to ILIKE $5",expression);
		param:="%"+queries.Search+"%";
		baseSQLArgs=append(baseSQLArgs, param, param, param, param, param);
	}
	
	// ORDER BY (sort)
	if queries.Sort!=""{
		parts:=strings.Split(queries.Sort,"$"); // tableName$direction -> [tableName, direction]
		parts[0]=strings.ToLower(parts[0]); // tableName
		parts[1]=strings.ToUpper(parts[1]); // direction
		baseSQLClause+=fmt.Sprintf(" ORDER BY %s %s, id", parts[0], parts[1]);
	}else if mode=="recommendation"{
		baseSQLClause+=" ORDER BY total_rating DESC, id"
	}

	//OFFSET-LIMIT (page)
	offset:=(queries.Page-1)*25;
	limit:=25;
	baseSQLClause+=fmt.Sprintf(" OFFSET $%d LIMIT $%d", len(baseSQLArgs)+1, len(baseSQLArgs)+2);
	baseSQLArgs=append(baseSQLArgs, offset, limit);

	return baseSQLClause, baseSQLArgs;
}

func (conn *stockRepository) GetAllStock(queries *models.RequestQueries) (pgx.Rows,error){
	baseQuery:="SELECT COUNT(*) OVER(),* FROM stock";
	clause, args:=conn.buildSQLClause(queries, "all");
	rows,err:=conn.db.Query(context.Background(),baseQuery+clause,args...);
	if err!=nil{
		return nil, fmt.Errorf("error on Query -> %w",err);
	}
	return rows, nil;
}

func (conn *stockRepository) GetRecommendation(queries *models.RequestQueries) (pgx.Rows,error){
	baseQuery:="SELECT COUNT(*) OVER(), id, ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, datereleased FROM (SELECT gen_rating(rating_to, target_from, target_to, action) AS total_rating,* FROM STOCK WHERE (rating_to ILIKE '%buy%') AND (rating_to NOT ILIKE '%spe%') AND (target_from>0 AND target_to>0))as sub WHERE total_rating>5";

	clause, args:=conn.buildSQLClause(queries, "recommendation");
	fmt.Println(baseQuery+clause, args);
	rows,err:=conn.db.Query(context.Background(),baseQuery+clause,args...);
	if err!=nil{
		return nil, fmt.Errorf("error on Query -> %w",err);
	}
	return rows, nil;
}

func NewStockRepository(db *pgxpool.Pool) StockRepository {
	return &stockRepository{db: db}
}