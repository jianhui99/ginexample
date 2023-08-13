package utils

import (
	"sync"
	"time"
)

type TimeUtil struct{}

const (
	dateLayout = "2006-01-02"
)

var (
	timeUtil     *TimeUtil
	timeUtilOnce sync.Once
)

func GetTimeUtil() *TimeUtil {
	if timeUtil == nil {
		timeUtilOnce.Do(func() {
			timeUtil = &TimeUtil{}
		})
	}
	return timeUtil
}

func (u *TimeUtil) GetUnixFromDateString(dateString string) (int64, error) {
	t, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return 0, err
	}

	year, month, day := t.Date()
	newTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	return newTime.Unix(), nil
}
