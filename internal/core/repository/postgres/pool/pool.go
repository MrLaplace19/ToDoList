package core_postgres_pool

import (
	"context"
	"fmt"
	"time"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface{
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Close()
	OptionalTimeOut() time.Duration
}

type ConnectionPool struct{
	*pgxpool.Pool
	OptTimeOut time.Duration
}

func NewConnectionPool(cfg Config, ctx context.Context) (*ConnectionPool, error){
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host,cfg.Port, cfg.DataBase)
	
	pgxConfig, err := pgxpool.ParseConfig(connectionString)
	if err !=nil{
		return nil, fmt.Errorf("parse pgx config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)

	if err!=nil{
		return nil, fmt.Errorf("pgxpool create pool: %w", err)
	} 

	if err:= pool.Ping(ctx); err !=nil{
		return nil, fmt.Errorf("pool ping: %w", err)
	}
	return &ConnectionPool{
		Pool: pool,
		OptTimeOut: cfg.TimeOut,
	}, nil
}

func (p *ConnectionPool) OptionalTimeOut() time.Duration{
	return p.OptTimeOut
}