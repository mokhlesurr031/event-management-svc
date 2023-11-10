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

func (e *WorkshopSQLStorage) WorkshopList(ctx context.Context, ctr *domain.WorkshopCriteria) ([]*domain.Workshop, error) {
	workshopList := make([]*domain.Workshop, 0)

	fmt.Println("called", workshopList)
	if err := e.db.WithContext(ctx).Find(&workshopList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return workshopList, nil
}
