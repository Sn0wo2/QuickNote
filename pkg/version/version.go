package version

import (
	"time"
)

var (
	version = "unknown"
	commit  = "unknown"
	date    = "unknown"
)

func GetVersion() string {
	return version
}

func GetCommit() string {
	return commit
}

func GetDate() string {
	return date
}

func GetDateTime() time.Time {
	t, _ := time.Parse(time.RFC3339, date)

	return t
}

func SetVersion(ver string) {
	version = ver
}

func SetCommit(comm string) {
	commit = comm
}

func SetDate(dat string) {
	date = dat
}
