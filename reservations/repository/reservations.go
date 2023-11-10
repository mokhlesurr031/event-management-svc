package repository

import (
	"context"

	"github.com/mokhlesurr031/event-management-svc/domain"
	"gorm.io/gorm"
)

func New(db *gorm.DB) domain.ReservationRepository {
	return &ReservationSQLStorage{
		db: db,
	}
}

type ReservationSQLStorage struct {
	db *gorm.DB
}

func (rsvn *ReservationSQLStorage) ReservationsCreate(ctx context.Context, ctr *domain.Reservation) (*domain.Reservation, error) {
	if err := rsvn.db.Create(ctr).Error; err != nil {
		return nil, err
	}

	createdID := ctr.Id

	reservationResp := domain.Reservation{
		Id:         createdID,
		Name:       ctr.Name,
		Email:      ctr.Email,
		WorkshopId: ctr.WorkshopId,
	}

	return &reservationResp, nil

}
