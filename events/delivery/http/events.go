package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
		r.Get("/details", handler.EventDetails)

	})
}

func (e *EventHandler) EventList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	perPage := r.URL.Query().Get("perPage")
	currentPage := r.URL.Query().Get("currentPage")

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

func (e *EventHandler) EventDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	eventId := r.URL.Query().Get("eventId")
	id64, err := strconv.ParseUint(eventId, 10, 64)
	if err != nil {
		// Handle the error, maybe return a bad request response
		http.Error(w, "Invalid eventId", http.StatusBadRequest)
		return
	}
	id := uint(id64)
	eventDetails := &domain.EventDetails{}
	eventDetails.Id = &id

	ctx := r.Context()

	eventData, err := e.EventsUseCase.EventDetails(ctx, eventDetails)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(eventData)
}
