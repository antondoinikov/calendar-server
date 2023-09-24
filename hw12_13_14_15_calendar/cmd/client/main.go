package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/pkg/event_v1"
	desc "github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/pkg/event_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const address = "localhost:50051"

func main() {
	ctx := context.Background()
	con, err := grpc.Dial(address, grpc.WithInsecure()) //nolint
	fmt.Println("Start...")
	if err != nil {
		log.Fatalf("failed to grpc connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewEventV1Client(con)
	res, err := client.CreateEvent(ctx, &desc.CreateEventRequest{
		Event: &desc.Event{
			Title:           "Anton",
			DateEventStart:  timestamppb.Now(),
			DateEventFinish: timestamppb.Now(),
			MessageEvent:    "first message",
			TimePreEvent:    timestamppb.New(time.Unix(3, 0)),
		},
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("--Create Event--")
	log.Println("Id: ", res.GetId())
}
