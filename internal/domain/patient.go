package domain

type Patient struct {
	Id               int    `json:"id"`
	Name             string `json:"name" binding:"required"`
	Surname          string `json:"surname" binding:"required"`
	DocumentNumber   string `json:"document_number" binding:"required"`
	RegistrationDate string `json:"registration_date" binding:"required"`
}
