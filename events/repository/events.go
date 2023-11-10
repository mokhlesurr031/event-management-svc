package repository

import (
	"context"
	"fmt"
	"strconv"

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

func (e *EventSQLStorage) EventList(ctx context.Context, ctr *domain.EventCriteria) (*domain.EventCriteriaPagination, error) {
	pagination := domain.Pagination{}
	eventList := make([]*domain.Event, 0)

	perPage := ctx.Value("perPage")
	perPageStr, _ := perPage.(string)
	perPageInt, _ := strconv.Atoi(perPageStr)

	currentPage := ctx.Value("currentPage")
	currentPageStr, _ := currentPage.(string)
	currentPageInt, _ := strconv.Atoi(currentPageStr)

	if err := e.db.WithContext(ctx).Find(&eventList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	pagination.Total = len(eventList)
	pagination.CurrentPage = currentPageInt
	pagination.PerPage = perPageInt
	pagination.TotalPages = pagination.Total / pagination.PerPage
	if pagination.Total%pagination.PerPage > 0 {
		pagination.TotalPages++
	}
	fmt.Println(pagination)

	res := &domain.EventCriteriaPagination{}
	res.Events = eventList
	res.Pagination = pagination

	fmt.Println(eventList)

	return res, nil
}
