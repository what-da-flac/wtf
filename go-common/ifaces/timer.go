package ifaces

import "time"

type Timer interface {
	Now() time.Time
}
