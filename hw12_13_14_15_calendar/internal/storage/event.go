package storage

import (
	"context"
	"time"

	"github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/internal/model"
)

type Event struct {
	ID    string
	Title string
	// TODO
}

type Storage interface {
	CreateEvent(ctx context.Context, event *model.Event) (int64, error)
	DeleteEvent(ctx context.Context, num_event int64) error
	GetListEvent(ctx context.Context) ([]*model.InfoEvent, error)
	GetListMonthEvent(ctx context.Context, req time.Time) ([]*model.InfoEvent, error)
	GetListWeekEvent(ctx context.Context, req time.Time) ([]*model.InfoEvent, error)
	UpdateEvent(ctx context.Context, req *model.UpdateEvent) error
}
