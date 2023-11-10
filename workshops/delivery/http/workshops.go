package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mokhlesurr031/event-management-svc/domain"
)

type WorkshopHandler struct {
	WorkshopUseCase domain.WorkshopUseCase
}

func NewHttpHandler(r *chi.Mux, workshopUseCase domain.WorkshopUseCase) {
	handler := &WorkshopHandler{
		WorkshopUseCase: workshopUseCase,
	}

	r.Route("/api/v1/workshops", func(r chi.Router) {
		r.Get("/list", handler.WorkshopList)
		r.Get("/details", handler.WorkshopDetails)
	})
}

func (e *WorkshopHandler) WorkshopList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	eventId := r.URL.Query().Get("eventId")
	id64, err := strconv.ParseUint(eventId, 10, 64)
	if err != nil {
		// Handle the error, maybe return a bad request response
		http.Error(w, "Invalid eventId", http.StatusBadRequest)
		return
	}
	id := uint(id64)
	workshops := &domain.WorkshopCriteria{}
	workshops.EventId = &id

	ctx := r.Context()

	workshopList, err := e.WorkshopUseCase.WorkshopList(ctx, workshops)

	if err != nil {
		log.Println(err)
	}

	if err := json.NewEncoder(w).Encode(workshopList); err != nil {
		log.Println(err)
	}

}

func (e *WorkshopHandler) WorkshopDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	workshopId := r.URL.Query().Get("workshopId")
	id64, err := strconv.ParseUint(workshopId, 10, 64)
	if err != nil {
		// Handle the error, maybe return a bad request response
		http.Error(w, "Invalid eventId", http.StatusBadRequest)
		return
	}
	id := uint(id64)

	workshopDetails := &domain.WorkshopCriteria{}
	workshopDetails.Id = &id

	ctx := r.Context()

	eventData, err := e.WorkshopUseCase.WorkshopDetails(ctx, workshopDetails)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(eventData)

}
