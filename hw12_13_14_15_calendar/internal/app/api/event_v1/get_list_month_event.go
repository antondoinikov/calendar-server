package event_v1

import (
	"context"
	"fmt"

	desc "github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) GetListMonthEvent(ctx context.Context, req *desc.GetListMonthEventRequest) (*desc.GetListMonthEventResponse, error) {
	fmt.Println(req.GetDateStartMonth())
	return &desc.GetListMonthEventResponse{
		Event: []*desc.InfoEvent{
			{
				Id:              1,
				Title:           "qwert",
				TimePreEvent:    nil,
				DateEventStart:  timestamppb.Now(),
				DateEventFinish: timestamppb.Now(),
				IdUser:          2,
				MessageEvent:    "12345rert",
			},
			{
				Id:              2,
				Title:           "qwert",
				TimePreEvent:    nil,
				DateEventStart:  timestamppb.Now(),
				DateEventFinish: timestamppb.Now(),
				IdUser:          2,
				MessageEvent:    "12345rert",
			},
		},
	}, nil
}
