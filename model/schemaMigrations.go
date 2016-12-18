package model

type SchemaMigrations struct {
	Version int `json:"version" xorm:"not null pk INTEGER"`
}
