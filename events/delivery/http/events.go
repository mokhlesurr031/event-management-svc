package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mokhlesurr031/event-management-svc/domain"
)

type EventHandler struct {
	EventsUseCase domain.EventUseCase
}

func NewHttpHandler(r *chi.Mux, eventUseCase domain.EventUseCase) {
	handler := &EventHandler{
		EventsUseCase: eventUseCase,
	}

	r.Route("/api/v1/events", func(r chi.Router) {
		r.Get("/list", handler.EventList)
	})
}

func (e *EventHandler) EventList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	perPage := r.URL.Query().Get("perPage")
	currentPage := r.URL.Query().Get("currentPage")

	if perPage == "" && currentPage == "" {
		fmt.Println("AAAA")
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, "perPage", perPage)
	ctx = context.WithValue(ctx, "currentPage", currentPage)

	events := &domain.EventCriteria{}
	eventList, err := e.EventsUseCase.EventList(ctx, events)

	if err != nil {
		log.Println(err)
	}

	if err := json.NewEncoder(w).Encode(eventList); err != nil {
		log.Println(err)
	}

}
