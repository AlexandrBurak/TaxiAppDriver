package model

type Register struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required,e164"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	TaxiType string `json:"taxi_type" binding:"required"`
}
