package core_pgx_pool

import (
	"errors"
	"fmt"

	core_postgres_pool "github.com/MrLaplace19/ToDoList/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRow struct{
	pgx.Row
}

type pgxRows struct{
	pgx.Rows
}

func(r *pgxRow) Scan(dest ...any) error{
	if err := r.Row.Scan(dest...); err !=nil{
		return mapsErrors(err)
	}
	return nil
}

type pgxCommandTag struct{
	pgconn.CommandTag
}

func mapsErrors(err error) error{
	const(
		errViolatesForeignKeyCode = "23503"
	)

	if errors.Is(err, pgx.ErrNoRows){
		return core_postgres_pool.ErrNoRows
	}
	
	var pgxErr *pgconn.PgError
	if errors.As(err, &pgxErr){
		if pgxErr.Code == errViolatesForeignKeyCode{
			return fmt.Errorf(
				"%v: %w",
				err,
				core_postgres_pool.ErrViolatesForeignKey,
			)
		}
	}

	return fmt.Errorf(
		"%v: %w",
		err,
		core_postgres_pool.ErrUnknown,
	)
}