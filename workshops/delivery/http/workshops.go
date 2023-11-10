package http

import (
	"encoding/json"
	"log"
	"net/http"

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

	ctx := r.Context()
	workshops := &domain.WorkshopCriteria{}

	eventList, err := e.WorkshopUseCase.WorkshopList(ctx, workshops)

	if err != nil {
		log.Println(err)
	}

	if err := json.NewEncoder(w).Encode(eventList); err != nil {
		log.Println(err)
	}

}
