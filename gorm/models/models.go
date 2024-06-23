package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Roles []string

type Job struct {
	Title    string
	Location string
}

type Teacher struct {
	gorm.Model
	Name         string `gorm:"size:256"`
	Email        string `gorm:"size:256"`
	Age          uint8  `gorm:"check:age>30"`
	WorkingYears uint8
	Birthday     int64 `gorm:"serializer:unixtime;type:time"`
	StuNumber    sql.NullString
	Roles        Roles `gorm:"serializer:json"`
	JobInfo      Job   `gorm:"embedded;embeddedPrefix:job_"`
	JobInfo2     Job   `gorm:"type:bytes;serializer:gob"`
}
