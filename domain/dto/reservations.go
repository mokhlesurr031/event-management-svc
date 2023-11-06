package dto

type Reservations struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	WorkshopId uint   `json:"workshop_id"`
}
