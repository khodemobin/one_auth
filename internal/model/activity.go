package model

type Activity struct {
	ID        uint   `gorm:"primarykey" faker:"-"`
	Action    string `db:"action" faker:"-"`
	IP        string `db:"ip" faker:"-"`
	Path      string `db:"path" faker:"-"`
	Operation string `db:"operation" faker:"-"`
	Version   string `db:"version" faker:"-"`
	Headers   string `db:"headers" faker:"-"`
}

func (Activity) TableName() string {
	return "activies"
}
