package usecase

import (
	"context"
	"errorhandlingproject/entity"
)

type repo interface {
	CreateData(ctx context.Context) *entity.TestData
}

type InputUseCase struct {
	repo repo
}

func (r *InputUseCase) CreateData(ctx context.Context) ([]string, error) {
	data := r.repo.CreateData(ctx)
	if data.Err != nil {
		return nil, data.Err
	} else {
		// if there is some business logic, else return the data that is required, here  am just returning an array
		return data.Data, nil
	}
}
