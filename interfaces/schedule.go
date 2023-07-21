package interfaces

import "time"

type Schedule interface {
	CheckDoNow() bool
	Run(now time.Time, client Client)
}
