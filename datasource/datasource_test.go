package datasource

import (
	"net/http"
	"testing"
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
							http.Get(url)
							return "", nil
						},
					},
					// relies on local Postgres running by default
					URL: []byte("http://localhost:5432"),
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
