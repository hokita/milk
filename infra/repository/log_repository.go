package infra

import (
	"database/sql"
	"log"

	domain "github.com/hokita/milk/domain/entity"
)

type LogRepository struct {
	DB *sql.DB
}

func (r *LogRepository) Get() *domain.Log {
	rows, err := r.DB.Query("SELECT * FROM logs;")
	if err != nil {
		log.Fatalf("getRows db.Query error err:%v", err)
	}
	defer rows.Close()

	l := &domain.Log{}
	for rows.Next() {
		if err := rows.Scan(
			&l.ID,
			&l.LogType,
			&l.CheckinAt,
			&l.CreatedAt,
			&l.UpdatedAt,
		); err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
	}

	return l
}
