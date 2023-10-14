package services

import "time"

func TimeNowUnix() int64 {
	now := time.Now()
	return now.Unix()
}