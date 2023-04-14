package store

import "github.com/ctd/esp-backend-desafio-2.git/internal/domain"

type DentistInterface interface {
	//ReadAll retorna todos os dentistas (dentist) cadastrados
	ReadAll() ([]domain.Dentist, error)
	//Read retorna um dentista (dentist) por id
	Read(id int) (domain.Dentist, error)
	// Create insere um novo dentista
	Create(dentist domain.Dentist) (domain.Dentist, error)
	//Update atualiza um dentista
	Update(dentist domain.Dentist) error
	//Delete exclui um dentista
	Delete(id int) error
}

type PatientInterface interface {
	//ReadAllPatient retorna todos os pacientes (patient) cadastrados
	ReadAllPatient() ([]domain.Patient, error)
	//ReadPatient retorna um paciente (patient) por id
	ReadPatient(id int) (domain.Patient, error)
	// CreatePatient insere um novo paciente
	CreatePatient(patient domain.Patient) (domain.Patient, error)
	//UpdatePatient atualiza um paciente
	UpdatePatient(patient domain.Patient) error
	//DeletePatient exclui um paciente
	DeletePatient(id int) error
}

type AppointmentInterface interface {
	//ReadAppointment retorna uma consulta (appointment) por id
	ReadAppointment(id int) (domain.AppointmentResponse, error)
	// CreateAppointment cria uma nova consulta
	CreateAppointment(Appointment domain.Appointment) (domain.AppointmentResponse, error)
	//UpdateAppointment atualiza uma consulta
	UpdateAppointment(Appointment domain.Appointment) (domain.AppointmentResponse, error)
	//DeleteAppointment exclui uma consulta
	DeleteAppointment(id int) error
	// GetByRg busca uma consulta pelo Rg do paciente
	ReadAppointmentByRgPatient(rg string) (domain.AppointmentResponse, error)
}