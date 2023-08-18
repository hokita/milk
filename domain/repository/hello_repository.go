package domain

import domain "github.com/hokita/milk/domain/entity"

type HelloRepository interface {
	Get() domain.Hello
}
