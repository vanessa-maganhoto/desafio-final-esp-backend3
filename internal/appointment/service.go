package appointment

import (
	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
)

type Service interface {
	//GetById retorna uma consulta (appointment) por id
	GetByID(id int) (domain.AppointmentResponse, error)
	// Create cria uma nova consulta
	Create(p domain.Appointment) (domain.AppointmentResponse, error)
	//Delete exclui uma consulta
	Delete(id int) error
	//Update atualiza uma consulta
	Update(id int, p domain.Appointment) (domain.AppointmentResponse, error)
	// GetByRg busca uma consulta pelo Rg do paciente
	GetByRg(rg string) (domain.AppointmentResponse, error)
}

type service struct {
	r Repository
}

// NewService cria um novo serviço
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.AppointmentResponse, error) {
	a, err := s.r.GetByID(id)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	return a, nil
}

func (s *service) Create(a domain.Appointment) (domain.AppointmentResponse, error) {
	appt, err := s.r.Create(a)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	return appt, nil
}
func (s *service) Update(id int, u domain.Appointment) (domain.AppointmentResponse, error) {
	var (
		appointment domain.Appointment
	)

	a, err := s.r.GetByID(id)

	appointment.Id = a.Id
	appointment.IdDentist = a.DentistD.Id
	appointment.IdPatient = a.PatientP.Id
	appointment.ApptDate = a.ApptDate
	appointment.Description = a.Description

	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	if u.IdDentist != 0 {
		appointment.IdDentist = u.IdDentist
	}
	if u.IdPatient != 0 {
		appointment.IdPatient = u.IdPatient
	}
	if u.ApptDate != "" {
		appointment.ApptDate = u.ApptDate
	}
	if u.Description != "" {
		appointment.Description = u.Description
	}
	appt, err := s.r.Update(id, appointment)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	return appt, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetByRg(rg string) (domain.AppointmentResponse, error) {
	a, err := s.r.GetByRg(rg)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	return a, nil
}