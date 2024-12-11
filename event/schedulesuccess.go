package event

type ScheduleSuccess struct {
}

func (e *ScheduleSuccess) Trigger() string {
	return ""
}
