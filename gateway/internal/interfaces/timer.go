package interfaces

import "time"

//go:generate moq -out ../../mocks/timer.go -pkg mocks . Timer
type Timer interface {
	Now() time.Time
}
