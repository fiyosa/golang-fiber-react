package helper

import (
	"strconv"
	"time"
)

func Str2Bool(data string) (bool, error) {
	boolValue, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}
	return boolValue, nil
}

func Int2Str(data int) string {
	return strconv.Itoa(data)
}

func Str2Int(data string) int {
	i, err := strconv.Atoi(data)
	if err != nil {
		return -1
	}
	return i
}

func Time2Str(date time.Time) string {
	layout := "2006-01-02 15:04:05"
	return date.Format(layout) // yyyy-MM-dd hh:mm:ss
}
