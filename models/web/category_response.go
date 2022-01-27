package web

import (
	"database/sql"
	"time"
)

type CategoryResponse struct {
	Name       string         `json:"name"`
	Image      sql.NullString `json:"image"`
	CreateTime time.Time      `json:"create_time"`
	UpdateTime time.Time      `json:"update_time"`
	Active     int8           `json:"active"`
}
