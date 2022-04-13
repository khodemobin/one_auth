package domain

type Activity struct {
	Action    string
	IP        string
	Path      string
	Operating string
	Version   string
	Headers   string
}

type ActivityRepository interface {
	CreateActivity(ac Activity) error
}
