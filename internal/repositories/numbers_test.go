package repositories

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var testdsl = "dbname=test_numbers user=igor password=igor"

func initalizeDB(ctx context.Context, db *pgxpool.Pool) error {
	query := `
		CREATE TABLE IF NOT EXISTS numbers (
			id SERIAL PRIMARY KEY,
			value INTEGER NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		);
	`
	_, err := db.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func trancataTable(ctx context.Context, db *pgxpool.Pool) error {
	query := "TRUNCATE TABLE numbers;"

	_, err := db.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func TestRepositoryGetAll(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{
			name: "When add 1",
			args: args{value: 1},
			want: []int{1},
		},
		{
			name: "When add 10",
			args: args{value: 10},
			want: []int{1, 10},
		},
		{
			name: "When add 3",
			args: args{value: 3},
			want: []int{1, 3, 10},
		},
	}

	ctx := context.Background()

	db, err := pgxpool.New(ctx, testdsl)
	if err != nil {
		assert.NoError(t, err)
	}

	defer db.Close()

	err = initalizeDB(ctx, db)
	if err != nil {
		assert.NoError(t, err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: db,
			}
			got, err := r.AddNumber(ctx, tt.args.value)

			if tt.wantErr != nil {
				assert.Equal(t, err.Error(), tt.wantErr.Error())
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				assert.Equal(t, got, tt.want)
			}
		})
	}

	err = trancataTable(ctx, db)
	if err != nil {
		assert.NoError(t, err)
	}
}
