package http

import (
	"encoding/json"
	"fmt"
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

	workshopList, err := e.WorkshopUseCase.WorkshopList(ctx, workshops)

	fmt.Println(workshopList)

	if err != nil {
		log.Println(err)
	}

	if err := json.NewEncoder(w).Encode(workshopList); err != nil {
		log.Println(err)
	}

}
