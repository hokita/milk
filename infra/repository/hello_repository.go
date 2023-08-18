package infra

import domain "github.com/hokita/milk/domain/entity"

type HelloRepository struct{}

func (r *HelloRepository) Get() *domain.Hello {
	return &domain.Hello{
		Name: "hokita",
		Age:  30,
	}
}
