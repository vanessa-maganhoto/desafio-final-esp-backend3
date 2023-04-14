package appointment

import (
	"errors"

	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
	"github.com/ctd/esp-backend-desafio-2.git/pkg/store"
)

type Repository interface {
	//GetById retorna uma consulta (appointment) por id
	GetByID(id int) (domain.AppointmentResponse, error)
	// Create cria uma nova consulta
	Create(p domain.Appointment) (domain.AppointmentResponse, error)
	//Update atualiza uma consulta
	Update(id int, a domain.Appointment) (domain.AppointmentResponse, error)
	//Delete exclui uma consulta
	Delete(id int) error
	// GetByRg busca uma consulta pelo Rg do paciente
	GetByRg(rg string) (domain.AppointmentResponse, error)
}

type repository struct {
	storage store.AppointmentInterface
}

//NewRepositoryAppointment cria um novo repositório
func NewRepositoryAppointment(storage store.AppointmentInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.AppointmentResponse, error) {
	appointment, err := r.storage.ReadAppointment(id)
	if err != nil {
		return domain.AppointmentResponse{}, errors.New("appointment not found")
	}
	return appointment, nil

}

func (r *repository) Create(a domain.Appointment) (domain.AppointmentResponse, error) {
	appt, err := r.storage.CreateAppointment(a)
	if err != nil {
		return domain.AppointmentResponse{}, errors.New("error creating Appointment")
	}
	return appt, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, a domain.Appointment) (domain.AppointmentResponse, error) {
	appt, err := r.storage.UpdateAppointment(a)
	if err != nil {
		return domain.AppointmentResponse{}, errors.New("error updating Appointment")
	}
	return appt, nil
}

func (r *repository) GetByRg(rg string) (domain.AppointmentResponse, error) {
	appointment, err := r.storage.ReadAppointmentByRgPatient(rg)
	if err != nil {
		return domain.AppointmentResponse{}, errors.New("appointment not found")
	}
	return appointment, nil

}
