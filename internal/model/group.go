package model

import (
	"time"
)

type Group struct {
	ID        string
	Name      string
	UserData  []string
	CreatedAt time.Time
	UserEmail string
}
