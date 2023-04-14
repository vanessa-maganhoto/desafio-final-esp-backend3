package dentist

import (
	"errors"

	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
	"github.com/ctd/esp-backend-desafio-2.git/pkg/store"
)

type Repository interface {
	//GetAll retorna todos os dentistas (dentist) cadastrados
	GetAll() ([]domain.Dentist, error)
	//GetByID retorna um dentista (dentist) por id
	GetByID(id int) (domain.Dentist, error)
	// Create insere um novo dentista
	Create(d domain.Dentist) (domain.Dentist, error)
	//Update atualiza um dentista
	Update(id int, d domain.Dentist) (domain.Dentist, error)
	//Delete exclui um dentista
	Delete(id int) error
}

type repository struct {
	storage store.DentistInterface
}

//NewRepository cria um novo repositório
func NewRepository(storage store.DentistInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetAll() ([]domain.Dentist, error) {
	dentists, err := r.storage.ReadAll()
	if err != nil {
		return []domain.Dentist{}, errors.New("dentists not found")
	}
	return dentists, err
}

func (r *repository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	return dentist, nil

}

func (r *repository) Create(d domain.Dentist) (domain.Dentist, error) {
	dentist, err := r.storage.Create(d)
	if err != nil {
		return domain.Dentist{}, errors.New("error creating dentist")
	}
	return dentist, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, d domain.Dentist) (domain.Dentist, error) {
	err := r.storage.Update(d)
	if err != nil {
		return domain.Dentist{}, errors.New("error updating dentist")
	}
	return d, nil
}
