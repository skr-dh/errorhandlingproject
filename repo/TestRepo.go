package repo

import (
	"context"
	"database/sql"
	"errorhandlingproject/entity"
)

type Input struct {
	dataToInsert []string
}

var connString = "user=postgres dbname=testdb password=secure-password host=localhost sslmode=disable"

func (d *Input) CreateData(ctx context.Context) *entity.TestData {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return &entity.TestData{
			Data: nil,
			Err: &entity.BaseError{
				Cause:   err,
				Status:  404,
				Message: "Unable to connect to db",
			},
		}
	} else {
		// connect to db and fetch data
		db.Exec("", "test")
		return &entity.TestData{
			// just creating a new array of string to represent test data
			Data: make([]string, 5),
			Err:  nil,
		}
	}

}
