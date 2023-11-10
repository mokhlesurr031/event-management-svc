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
	res.Events = eventList[perPageInt*currentPageInt-perPageInt : perPageInt*currentPageInt]
	res.Pagination = pagination

	return res, nil
}

func (e *EventSQLStorage) EventDetails(ctx context.Context, ctr *domain.EventDetails) (*domain.EventDetails, error) {
	eventDetail := &domain.Event{}

	if ctr.Id != nil {
		err := e.db.First(&eventDetail, "id=?", ctr.Id).Error
		if err != nil {
			return nil, err
		}
		fmt.Println(eventDetail)
		var count int64
		err = e.db.Model(&domain.Workshop{}).Where("event_id = ?", *ctr.Id).Count(&count).Error
		if err != nil {
			return nil, err
		}

		eventDetails := &domain.EventDetails{
			Id:             &eventDetail.Id,
			Title:          &eventDetail.Title,
			StartAt:        &eventDetail.StartAt,
			EndAt:          &eventDetail.EndAt,
			TotalWorkshops: int(count),
		}
		return eventDetails, nil
	}
	return nil, nil
}
