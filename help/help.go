package help

import "time"

func GetTime(timezone ...string) time.Time {
	if len(timezone) > 0 {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return time.Now()
		}
		return time.Now().In(loc)
	}
	return time.Now()
}
