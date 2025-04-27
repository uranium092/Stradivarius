package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var poolDb *pgxpool.Pool;

func InitConnection() error{
	config, err := pgxpool.ParseConfig(os.Getenv("URL_COCKROACHDB"))
	if err != nil {
    return err;
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config);
	if err != nil {
    return err;
	}
	if err:=pool.Ping(context.Background()); err!=nil{
		return err;
	}
	poolDb=pool;
	return nil;
}

func GetConnection() *pgxpool.Pool{
	return poolDb;
}