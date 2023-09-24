package app

import (
	"context"

	"github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/internal/service/event"
	memstorage "github.com/antondoinikov/calendar-server/hw12_13_14_15_calendar/internal/storage/memorystorage"
)

type App struct {
	Service event.Service
}

type Logger interface { // TODO
}

type Storage interface { // TODO
}

func New(logger Logger, storage Storage) *App {
	return &App{}
}

func (a *App) InitService() {
	a.Service.Storage = &memstorage.Memory{}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
