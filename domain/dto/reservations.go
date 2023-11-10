package dto

type ReservationsDto struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	WorkshopId uint   `json:"workshop_id"`
}
