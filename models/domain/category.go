package domain

import (
	"database/sql"
	"time"
)

type Category struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Image      sql.NullString `json:"image"`
	Slug       string         `json:"slug"`
	CreateTime time.Time      `json:"create_time"`
	UpdateTime time.Time      `json:"update_time"`
	Active     int8           `json:"active"`
}
