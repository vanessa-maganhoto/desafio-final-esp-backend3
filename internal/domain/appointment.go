package domain

type Appointment struct {
	Id          int    `json:"id"`
	IdDentist   int    `json:"id_dentist" binding:"required"`
	IdPatient   int    `json:"id_patient" binding:"required"`
	ApptDate    string `json:"appt_date" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type AppointmentResponse struct {
	Id          int     `json:"id"`
	DentistD    Dentist `json:"Dentist"`
	PatientP    Patient `json:"Patient"`
	ApptDate    string  `json:"appt_date"`
	Description string  `json:"description"`
}
