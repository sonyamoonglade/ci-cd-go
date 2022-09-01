package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	ci_cd_test "gihub.com/sonyamoonglade/ci-cd-golang"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) GetUser(ctx context.Context) (*ci_cd_test.User, error) {

	q := fmt.Sprintf("SELECT * FROM users")
	rows, err := r.db.QueryxContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u ci_cd_test.User

	for rows.Next() {
		err = rows.StructScan(&u)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}
	}

	return &u, nil
}
