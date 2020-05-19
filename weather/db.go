package weather

import (
	"fmt"
	"sync"
	"time"
)

const (
	station = "GHCND:USW00094728"
)

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

var (
	dbLock sync.RWMutex
	/*
		https://www.wunderground.com/history/monthly/us/ny/new-york-city/KLGA/
		date/2020-3
	*/
	db = []Record{
		{date(2020, 3, 1), station, Value{35.5, "f"}, Value{0, "in"}},
		{date(2020, 3, 2), station, Value{48.2, "f"}, Value{0, "in"}},
		{date(2020, 3, 3), station, Value{52.4, "f"}, Value{0.01, "in"}},
		{date(2020, 3, 4), station, Value{50.5, "f"}, Value{0.28, "in"}},
		{date(2020, 3, 5), station, Value{44.8, "f"}, Value{0, "in"}},
	}
)

func dateEqual(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// GetRecord gets a record for a date
func GetRecord(date time.Time) (Record, error) {
	dbLock.RLock()
	defer dbLock.RUnlock()

	for i := range db {
		if dateEqual(db[i].Time, date) {
			return db[i], nil
		}
	}

	ts := date.Format("2006-01-02")
	return Record{}, fmt.Errorf("can't find record for %s", ts)
}

// AddRecord adds a record to the database
func AddRecord(r Record) int {
	dbLock.Lock()
	defer dbLock.Unlock()

	db = append(db, r)
	return len(db)
}
