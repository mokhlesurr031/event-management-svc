package domain

import "time"

type Reservations struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      time.Time `json:"email"`
	WorkshopId time.Time `json:"workshop_id"`
}

type ReservationsCriteria struct {
	Id         *uint      `json:"id"`
	Name       *string    `json:"name"`
	Email      *time.Time `json:"email"`
	WorkshopId *time.Time `json:"workshop_id"`
}
