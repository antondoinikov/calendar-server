package event_v1

import (
	"context"

	desc "github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) GetListEvent(ctx context.Context, req *emptypb.Empty) (*desc.GetListEventResponse, error) {
	return &desc.GetListEventResponse{
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
