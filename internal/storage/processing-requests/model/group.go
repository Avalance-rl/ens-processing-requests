package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Group struct {
	ID        string           `db:"id"`
	Name      string           `db:"email"`
	UserData  []string         `db:"user_data"`
	CreatedAt pgtype.Timestamp `db:"created_at"`
	UserEmail string           `db:"user_email"`
}
