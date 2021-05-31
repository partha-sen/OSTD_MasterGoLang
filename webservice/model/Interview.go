package model

import "time"

type interview struct {
	id         string
	opening_id string
	date       time.Time
	person     string
}
