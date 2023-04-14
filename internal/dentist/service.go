package dentist

import (
	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
)

type Service interface {
	//GetAll retorna todos os dentistas (dentist) cadastrados
	GetAll() ([]domain.Dentist, error)
	//GetByID retorna um dentista (dentist) por id
	GetByID(id int) (domain.Dentist, error)
	// Create insere um novo dentista
	Create(d domain.Dentist) (domain.Dentist, error)
	//Delete exclui um dentista
	Delete(id int) error
	//Update atualiza um dentista
	Update(id int, d domain.Dentist) (domain.Dentist, error)
}

type service struct {
	r Repository
}

// NewService cria um novo serviço
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Dentist, error) {
	dentists, err := s.r.GetAll()
	if err != nil {
		return []domain.Dentist{}, err
	}
	return dentists, err
}

func (s *service) GetByID(id int) (domain.Dentist, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

func (s *service) Create(d domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}
func (s *service) Update(id int, u domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	if u.Name != "" {
		d.Name = u.Name
	}
	if u.Surname != "" {
		d.Surname = u.Surname
	}
	if u.Registration != "" {
		d.Registration = u.Registration
	}
	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
