package patient

import (
	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
)

type Service interface {
	//GetAll retorna todos os pacientes (patient) cadastrados
	GetAll() ([]domain.Patient, error)
	//GetByID retorna um paciente (patient) por id
	GetByID(id int) (domain.Patient, error)
	// Create insere um novo paciente
	Create(p domain.Patient) (domain.Patient, error)
	//Delete exclui um paciente
	Delete(id int) error
	//Update atualiza um paciente
	Update(id int, p domain.Patient) (domain.Patient, error)
}

type service struct {
	r Repository
}

// NewService cria um novo serviço
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Patient, error) {
	patients, err := s.r.GetAll()
	if err != nil {
		return []domain.Patient{}, err
	}
	return patients, err
}

func (s *service) GetByID(id int) (domain.Patient, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s *service) Create(p domain.Patient) (domain.Patient, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Patient) (domain.Patient, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.Surname != "" {
		p.Surname = u.Surname
	}
	if u.DocumentNumber != "" {
		p.DocumentNumber = u.DocumentNumber
	}
	if u.RegistrationDate != "" {
		p.RegistrationDate = u.RegistrationDate
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
