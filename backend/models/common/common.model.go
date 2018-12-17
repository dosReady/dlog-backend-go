package common

import "time"

type Base struct {
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}
