package model

import (
	"time"
)

type ProductInspections struct {
	Id        string    `json:"id" xorm:"not null pk UUID"`
	Serial    string    `json:"serial" xorm:"not null unique TEXT"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
}
