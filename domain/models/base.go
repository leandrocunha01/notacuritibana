package models

import "time"

type Base struct {
	Id        int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
