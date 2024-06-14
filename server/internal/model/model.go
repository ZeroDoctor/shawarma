package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Time time.Time

func (t *Time) Scan(src interface{}) error {
	if unix, ok := src.(int64); ok {
		*t = Time(time.Unix(unix, 0))
		return nil
	}

	return fmt.Errorf("unexcepted format [src=%T] wanted type int64", src)
}

func (t Time) Value() (driver.Value, error) {
	ti := time.Time(t)
	return ti.Unix(), nil
}
