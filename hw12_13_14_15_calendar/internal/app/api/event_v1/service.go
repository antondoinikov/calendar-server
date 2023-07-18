package event_v1

import desc "github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/pkg/event_v1"

type Implementation struct {
	desc.UnimplementedEventV1Server
}

func NewEvent() *Implementation {
	return &Implementation{}
}
