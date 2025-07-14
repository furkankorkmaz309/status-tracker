package models

import "time"

type Checks struct {
	ID          int       `json:"id"`
	ServiceID   int       `json:"service_id"`
	ServiceName string    `json:"service_name"`
	CheckedAt   time.Time `json:"checked_at"`
	StatusCode  int       `json:"status_code"`
}
