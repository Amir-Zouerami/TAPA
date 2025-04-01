package types

import (
	"fmt"
	"strings"
	"time"
)

type SqliteTime time.Time

func (st *SqliteTime) Scan(value any) error {
	switch v := value.(type) {
	case time.Time:
		*st = SqliteTime(v)
	case string:
		// Handles SQLite's verbose timestamp format (e.g. "2025-03-31 15:39:35.132068083 +0330 +0330 m=+7.845060104")
		mainPart := strings.Split(v, " +")[0]
		formats := []string{
			"2006-01-02 15:04:05.999999999", // With nanoseconds
			"2006-01-02 15:04:05",           // Without nanoseconds
		}

		for _, format := range formats {
			if t, err := time.Parse(format, mainPart); err == nil {
				*st = SqliteTime(t)
				return nil
			}
		}
		return fmt.Errorf("failed to parse time: %s", v)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}

func (st SqliteTime) MarshalJSON() ([]byte, error) {
	t := time.Time(st)
	if t.IsZero() {
		return []byte(`null`), nil
	}
	return []byte(`"` + t.Format(time.RFC3339Nano) + `"`), nil
}

func NewSqliteTime(t time.Time) SqliteTime {
	return SqliteTime(t)
}
