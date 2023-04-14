package store

import (
	"database/sql"

	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
)

// NewSqlStorePatient estabelece conexão com a tabela patients no banco de dados
func NewSqlStorePatient(db *sql.DB) PatientInterface {
	return &SqlStore{
		db: db,
	}
}

// CreatePatient escreve a query para inserir um novo paciente no banco de dados dentro da tabela patients e devolve possíveis erros
func (s *SqlStore) CreatePatient(patient domain.Patient) (domain.Patient, error) {
	query := `INSERT INTO patients (name, surname, documentNumber, registrationDate)
	VALUES (?, ?, ?, ?);`
	prepare, err := s.db.Prepare(query)
	if err != nil {
		return domain.Patient{}, err
	}
	defer prepare.Close()

	res, err := prepare.Exec(patient.Name, patient.Surname, patient.DocumentNumber, patient.RegistrationDate)
	if err != nil {
		return domain.Patient{}, nil
	}

	insertedID, _ := res.LastInsertId()

	patient.Id = int(insertedID)

	_, err = res.RowsAffected()
	return patient, err
}

// ReadAllPatient escreve a query para buscar todos os dentistas no banco de dados dentro da tabela patients e devolve possíveis erros
func (s *SqlStore) ReadAllPatient() ([]domain.Patient, error) {
	var (
		idRes int
		name string
		surname string
		documentNumber string
		registrationDate string
		patient domain.Patient
		patients []domain.Patient
	)
	query := "SELECT id, name, surname, documentNumber, registrationDate FROM patients"
	rows, err := s.db.Query(query)
	if err != nil {
		return []domain.Patient{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&idRes, &name, &surname, &documentNumber, &registrationDate)
		patient.Id = idRes
		patient.Name = name
		patient.Surname = surname
		patient.DocumentNumber = documentNumber
		patient.RegistrationDate = registrationDate
		patients = append(patients, patient)
	}

	return patients, err
}

// ReadPatient escreve a query para buscar um paciente pelo seu ID no banco de dados dentro da tabela patients e devolve possíveis erros
func (s *SqlStore) ReadPatient(id int) (domain.Patient, error) {
	var (
		idRes int
		name string
		surname string
		documentNumber string
		registrationDate string
		patient domain.Patient
	)
	query := `SELECT id, name, surname, documentNumber, registrationDate FROM patients
	WHERE id = ?;`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return domain.Patient{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&idRes, &name, &surname, &documentNumber, &registrationDate)
		patient.Id = idRes
		patient.Name = name
		patient.Surname = surname
		patient.DocumentNumber = documentNumber
		patient.RegistrationDate = registrationDate
	}

	return patient, err
}

// UpdatePatient escreve a query para atualizar um paciente no banco de dados dentro da tabela patients e devolve possíveis erros
func (s *SqlStore) UpdatePatient(patient domain.Patient) error {
	_, err := s.Read(patient.Id)
	if err != nil {
		return err
	}
	query := `UPDATE patients SET name = ?, surname = ?, documentNumber = ?, registrationDate = ?
	WHERE id = ?;`
	prepare, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := prepare.Exec(patient.Name, patient.Surname, patient.DocumentNumber, patient.RegistrationDate, patient.Id)
	if err != nil {
		return nil
	}

	_, err = res.RowsAffected()
	return err
}

// DeletePatient escreve a query para excluir um paciente pelo seu respectivo ID no banco de dados dentro da tabela patients e devolve possíveis erros
func (s *SqlStore) DeletePatient(id int) error {
	query := `DELETE FROM patients WHERE id = ?;`
	prepare, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := prepare.Exec(id)
	if err != nil {
		return nil
	}

	_, err = res.RowsAffected()
	return err
}


