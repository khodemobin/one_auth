package model

import (
	"database/sql"
)

type Activity struct {
	ID        uint
	Action    string
	IP        string
	Path      string
	Operation string
	Version   string
	Headers   string
	CreatedAt sql.NullTime `gorm:"autoCreateTime" faker:"-"`
}

func (Activity) TableName() string {
	return "activies"
}
