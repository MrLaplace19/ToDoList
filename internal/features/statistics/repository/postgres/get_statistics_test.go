package statistics_postgres_repository

import (
	"context"
	"errors"
	corepool "github.com/MrLaplace19/ToDoList/internal/core/repository/postgres/pool"
	"reflect"
	"testing"
	"time"
)

type testPool struct {
	corepool.Pool
	args     []any
	query    string
	rowError error
}

func (p *testPool) OptionalTimeOut() time.Duration { return time.Second }
func (p *testPool) Query(_ context.Context, query string, args ...any) (corepool.Rows, error) {
	p.query = query
	p.args = args
	return testRows{err: p.rowError}, nil
}

type testRows struct{ err error }

func (testRows) Close()            {}
func (r testRows) Err() error      { return r.err }
func (testRows) Next() bool        { return false }
func (testRows) Scan(...any) error { return nil }

func TestStatisticsDateBounds(t *testing.T) {
	from := time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC)
	to := from.AddDate(0, 1, 0)
	id := 12
	for _, tc := range []struct {
		name     string
		id       *int
		from, to *time.Time
		args     []any
	}{
		{"end only", nil, nil, &to, []any{&to}},
		{"full range", &id, &from, &to, []any{&id, &from, &to}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pool := &testPool{}
			_, err := NewStatisticsRepository(pool).GetTasks(context.Background(), tc.id, tc.to, tc.from)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(pool.args, tc.args) {
				t.Fatalf("SQL arguments = %v, want %v", pool.args, tc.args)
			}
		})
	}
}
func TestStatisticsIterationError(t *testing.T) {
	expected := errors.New("connection interrupted")
	_, err := NewStatisticsRepository(&testPool{rowError: expected}).GetTasks(context.Background(), nil, nil, nil)
	if !errors.Is(err, expected) {
		t.Fatalf("error = %v, want %v", err, expected)
	}
}
