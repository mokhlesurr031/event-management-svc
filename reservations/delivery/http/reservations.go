package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/mokhlesurr031/event-management-svc/domain"
)

type ReservationHandler struct {
	ReservationUseCase domain.ReservationUseCase
}

func NewHttpHandler(r *chi.Mux, reservationUseCase domain.ReservationUseCase) {
	handler := &ReservationHandler{
		ReservationUseCase: reservationUseCase,
	}

	r.Route("/api/v1/reservations", func(r chi.Router) {
		r.Post("/", handler.ReservationsCreate)
	})
}

type ReqReservation struct {
	domain.Reservation
}

func (rsvn *ReservationHandler) ReservationsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := ReqReservation{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}

	ctx := r.Context()
	reservation := domain.Reservation(req.Reservation)
	res, err := rsvn.ReservationUseCase.ReservationsCreate(ctx, &reservation)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
