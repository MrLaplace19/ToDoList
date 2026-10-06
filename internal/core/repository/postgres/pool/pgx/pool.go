package core_pgx_pool

import (
	"context"
	"fmt"
	"time"

	core_postgres_pool "github.com/MrLaplace19/ToDoList/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
	optTimeOut time.Duration
}

func NewPgxConnectionPool(cfg Config, ctx context.Context) (*Pool, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DataBase)

	pgxConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("parse pgx config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)

	if err != nil {
		return nil, fmt.Errorf("pgxpool create pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pool ping: %w", err)
	}
	return &Pool{
		Pool:       pool,
		optTimeOut: cfg.TimeOut,
	}, nil
}

func (p *Pool) OptionalTimeOut() time.Duration {
	return p.optTimeOut
}

func (p *Pool) Query(
	ctx context.Context,
	sql string,
	args ...any,
	) (core_postgres_pool.Rows, error){
		rows, err := p.Pool.Query(ctx, sql, args...)
		if err != nil {
			return nil, err
		}
		return pgxRows{rows}, nil
	}

	
func (p *Pool)	QueryRow(
	ctx context.Context,
	sql string,
	args ...any,
	) core_postgres_pool.Row{
		return  p.Pool.QueryRow(ctx, sql, args...)
	}


func (p *Pool)	Exec(
	ctx context.Context,
	sql string,
	arguments ...any,
	) (core_postgres_pool.CommandTag, error){
		commandTag, err := p.Pool.Exec(ctx, sql, arguments...)
		if err != nil{
			return nil, err
		}

		return pgxCommandTag{commandTag}, nil
	}	