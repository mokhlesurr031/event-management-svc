package repository

import (
	"context"
	"fmt"

	"github.com/mokhlesurr031/event-management-svc/domain"
	"gorm.io/gorm"
)

func New(db *gorm.DB) domain.WorkshopRepository {
	return &WorkshopSQLStorage{
		db: db,
	}
}

type WorkshopSQLStorage struct {
	db *gorm.DB
}

func (e *WorkshopSQLStorage) WorkshopList(ctx context.Context, ctr *domain.WorkshopCriteria) (*domain.WorkshopList, error) {
	eventDetail := &domain.Event{}
	err := e.db.First(&eventDetail, "id=?", ctr.EventId).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(eventDetail)

	workshops := []*domain.Workshop{}
	if err := e.db.Find(&workshops, "event_id = ?", *ctr.EventId).Error; err != nil {
		return nil, err
	}

	result := &domain.WorkshopList{
		Id:       &eventDetail.Id,
		Title:    &eventDetail.Title,
		StartAt:  &eventDetail.StartAt,
		EndAt:    &eventDetail.EndAt,
		Workshop: workshops,
	}

	fmt.Println(result)

	return result, nil
}
