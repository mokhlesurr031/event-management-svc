package repository

import (
	"context"
	"fmt"

	"github.com/mokhlesurr031/event-management-svc/domain"
	"gorm.io/gorm"
)

func New(db *gorm.DB) domain.EventRepository {
	return &EventSQLStorage{
		db: db,
	}
}

type EventSQLStorage struct {
	db *gorm.DB
}

func (e *EventSQLStorage) EventList(ctx context.Context, ctr *domain.EventCriteria) ([]*domain.Event, error) {
	eventList := make([]*domain.Event, 0)

	fmt.Println("called", eventList)
	if err := e.db.WithContext(ctx).Find(&eventList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return eventList, nil
}
