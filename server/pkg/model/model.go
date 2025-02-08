package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type StringList []string

func (sl *StringList) Scan(src interface{}) error {
	if str, ok := src.(string); ok {
		*sl = strings.Split(str, ",")
		return nil
	}

	return fmt.Errorf("unexcepted format [StringList=%T] wanted type string", src)
}

func (sl StringList) Value() (driver.Value, error) {
	return strings.Join(sl, ","), nil
}

type UUID uuid.UUID

func (u *UUID) Scan(src interface{}) error {
	if str, ok := src.(string); ok {
		id, err := uuid.Parse(str)
		if err != nil {
			return err
		}

		*u = UUID(id)
	}

	return nil
}

func (u UUID) Value() (driver.Value, error) {
	return uuid.UUID(u).String(), nil
}

type Time time.Time

func (t *Time) Scan(src interface{}) error {
	if unix, ok := src.(int64); ok {
		*t = Time(time.Unix(unix, 0))
		return nil
	}

	return nil
}

func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}
