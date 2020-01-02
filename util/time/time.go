package time

import "time"

const day = 86400

func IsToday(timestamp int64) bool {
	unixTime := time.Unix(timestamp/1000, 0)
	now := time.Now().Unix()
	limit := now - (now % day)

	return unixTime.Unix() >= limit
}
