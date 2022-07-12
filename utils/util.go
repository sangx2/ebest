package utils

import "time"

func GetDateString() string {
	return time.Now().Format("2006-01-02")
}
