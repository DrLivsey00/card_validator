package models

type Card struct {
	Number          string `json:"card_number"`
	ExpirationMonth int    `json:"card_expiration_month"`
	ExpirationYear  int    `json:"card_expiration_year"`
}
