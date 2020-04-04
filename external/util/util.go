package util

import "time"

// WIBTimezone return time Jakarta timezone
func WIBTimezone(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	strTime := t.In(loc).Format(time.RFC3339)
	tz, _ := time.Parse(time.RFC3339, strTime)
	return tz
}
