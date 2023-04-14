package patient

import (
	"errors"

	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
	"github.com/ctd/esp-backend-desafio-2.git/pkg/store"
)

type Repository interface {
	//GetAll retorna todos os pacientes (patient) cadastrados
	GetAll() ([]domain.Patient, error)
	//GetByID retorna um paciente (patient) por id
	GetByID(id int) (domain.Patient, error)
	// Create insere um novo paciente
	Create(p domain.Patient) (domain.Patient, error)
	//Update atualiza um paciente
	Update(id int, p domain.Patient) (domain.Patient, error)
	//Delete exclui um paciente
	Delete(id int) error
}

type repository struct {
	storage store.PatientInterface
}

//NewRepository cria um novo repositório
func NewRepository(storage store.PatientInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetAll() ([]domain.Patient, error) {
	patients, err := r.storage.ReadAllPatient()
	if err != nil {
		return []domain.Patient{}, errors.New("patients not found")
	}
	return patients, err
}

func (r *repository) GetByID(id int) (domain.Patient, error) {
	Patient, err := r.storage.ReadPatient(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return Patient, nil

}

func (r *repository) Create(p domain.Patient) (domain.Patient, error) {
	patient, err := r.storage.CreatePatient(p)
	if err != nil {
		return domain.Patient{}, errors.New("error creating patient")
	}
	return patient, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, p domain.Patient) (domain.Patient, error) {
	err := r.storage.UpdatePatient(p)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return p, nil
}
