package schedules

import "time"

type ScheduleTime struct {
	Hour   int
	Minute int
}

func checkDoNow(t ScheduleTime) bool {
	now := time.Now()
	return t.Hour == now.Hour() && t.Minute == now.Minute()
}
