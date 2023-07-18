package event_v1

import (
	"context"
	"fmt"

	desc "github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) DeleteEvent(ctx context.Context, req *desc.DeleteEventRequest) (*emptypb.Empty, error) {
	fmt.Println(req.GetId())
	return &emptypb.Empty{}, nil
}
