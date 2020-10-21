package datasource

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/zeebo/errs"
)

func TestDatasource(t *testing.T) {
	t.Run("test checker func", func(t *testing.T) {
		var sourceTests = []struct {
			source         *Source
			expectedErr    error
			checkerName    string
			expectedReturn string
		}{
			{
				source: &Source{
					Checkers: map[string]CheckerFunc{
						"default": func(url string) (string, error) {
							return "up", nil
						},
					},
				},
				checkerName:    "default",
				expectedErr:    nil,
				expectedReturn: "up",
			},
			{
				source: &Source{
					Checkers: map[string]CheckerFunc{
						"postgres": func(url string) (string, error) {
							db, err := sql.Open("postgres", url)
							if err != nil {
								fmt.Printf("failed to open PG connection: %s", err)
								return "down", errs.Wrap(err)
							}
							defer db.Close()

							err = db.Ping()
							if err != nil {
								fmt.Printf("failed to ping postgres: %+v\n", err)
								return "unauthorized", errs.New("database is down")
							}

							return "up", nil
						},
					},
					// relies on local Postgres running by default
					URL: []byte("postgres://meroxa:meroxa@localhost:5431/meroxa?sslmode=disable"),
				},
				checkerName:    "postgres",
				expectedErr:    nil,
				expectedReturn: "up",
			},
		}
		for _, tt := range sourceTests {
			result, err := tt.source.Check(tt.source.Checkers[tt.checkerName])

			if result != tt.expectedReturn {
				t.Fail()
			}

			if err != tt.expectedErr {
				t.Fail()
			}
		}
	})
}
