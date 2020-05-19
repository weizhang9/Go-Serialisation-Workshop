package pb

import (
	"time"

	pbtime "google.golang.org/protobuf/types/known/timestamppb"
)

// Timestamp converts time.Time to protobuf Timestamp
func Timestamp(t time.Time) *pbtime.Timestamp {
	var pbt pbtime.Timestamp
	pbt.Seconds = t.Unix()
	pbt.Nanos = int32(t.Nanosecond())
	return &pbt
}

// Time converts protobuf Timestamp to time.Time
func Time(pbt *pbtime.Timestamp) time.Time {
	return time.Unix(pbt.Seconds, int64(pbt.Nanos))
}

// Now return current time as protobuf Timestamp
func Now() *pbtime.Timestamp {
	now := time.Now()
	return Timestamp(now)
}
