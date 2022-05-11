package common

import "time"

func UnixTimestamp() int64 {
	return time.Now().Unix()
}
