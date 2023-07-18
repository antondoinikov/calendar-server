package event_v1

import (
	"context"
	"fmt"

	desc "github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/pkg/event_v1"
)

func (i *Implementation) CreateEvent(ctx context.Context, req *desc.CreateEventRequest) (*desc.CreateEventResponse, error) {
	fmt.Println(req.GetEvent().GetTitle())
	return &desc.CreateEventResponse{
		Id: 23,
	}, nil
}
