package repository

import (
	"context"

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

	return result, nil
}

func (e *WorkshopSQLStorage) WorkshopDetails(ctx context.Context, ctr *domain.WorkshopCriteria) (*domain.WorkshopDetails, error) {
	workshopDetail := &domain.Workshop{}

	if ctr.Id != nil {
		err := e.db.First(&workshopDetail, "id=?", ctr.Id).Error
		if err != nil {
			return nil, err
		}
		var count int64
		err = e.db.Model(&domain.Reservation{}).Where("workshop_id = ?", *ctr.Id).Count(&count).Error
		if err != nil {
			return nil, err
		}

		workshopDetails := &domain.WorkshopDetails{
			Id:               &workshopDetail.Id,
			Title:            &workshopDetail.Title,
			Description:      &workshopDetail.Description,
			StartAt:          &workshopDetail.StartAt,
			EndAt:            &workshopDetail.EndAt,
			TotaReservations: int(count),
		}

		return workshopDetails, nil
	}
	return nil, nil
}
