package util

import (
	"fmt"
	"time"
)

// Now return current time using single variable for easy mock
var Now = time.Now

// WIBTimezone return time Jakarta timezone
func WIBTimezone(t time.Time) time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err)
	}
	strTime := t.In(loc).Format(time.RFC3339)
	tz, _ := time.Parse(time.RFC3339, strTime)
	return tz
}

// IsErrorRecordNotFound return bool
func IsErrorRecordNotFound(err error) bool {
	if err.Error() == "record not found" {
		return true
	}

	return false
}
