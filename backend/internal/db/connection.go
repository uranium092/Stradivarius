package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitConnection() (*pgxpool.Pool,error){

	//parse stringConfig to Config{}
	config, err := pgxpool.ParseConfig(os.Getenv("URL_COCKROACHDB"))
	if err != nil {
    return nil, err;
	}
	//create pool based on config
	pool, err := pgxpool.NewWithConfig(context.Background(), config);
	if err != nil {
    return nil, err;
	}
	//try empty query to auth with DB
	if err:=pool.Ping(context.Background()); err!=nil{
		return nil,err;
	}
	return pool, nil;
}