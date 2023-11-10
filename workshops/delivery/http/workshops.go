package http

import (
	"encoding/json"
	"fmt"
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

	fmt.Println(workshopList)

	if err != nil {
		log.Println(err)
	}

	if err := json.NewEncoder(w).Encode(workshopList); err != nil {
		log.Println(err)
	}

}
