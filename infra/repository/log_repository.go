package infra

import (
	"time"

	domain "github.com/hokita/milk/domain/entity"
)

type LogRepository struct{}

func (r *LogRepository) Get() *domain.Log {
	return &domain.Log{
		LogType: domain.Milk,
		Date:    time.Date(2023, 8, 18, 18, 37, 0, 0, time.Local),
	}
}
