package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddNumber(ctx context.Context, number int) ([]int, error) {
	query := "INSERT INTO numbers (value) VALUES ($1)"
	selectQuery := "SELECT value FROM numbers ORDER BY value ASC"

	var numbers []int

	t, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, query, number)
	if err != nil {
		t.Rollback(ctx)

		return nil, err
	}

	rows, err := r.db.Query(ctx, selectQuery)
	if err != nil {
		t.Rollback(ctx)

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var number int

		err := rows.Scan(&number)
		if err != nil {
			rows.Close()

			t.Rollback(ctx)

			return nil, err
		}

		numbers = append(numbers, number)
	}

	err = t.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return numbers, nil
}
