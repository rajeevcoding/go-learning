package logging

import "time"

type BaseLogger struct {
	Utc bool
}

func (bl *BaseLogger) timestamp() string {

	if bl.Utc {
		return time.Now().UTC().Format("2006-01-02 15:04:05")
	}

	return time.Now().Format("2006-01-02 15:04:05")
}
