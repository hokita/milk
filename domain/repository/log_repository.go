package domain

import domain "github.com/hokita/milk/domain/entity"

type LogRepository interface {
	Get() domain.Log
}
