package store

import (
	"database/sql"
	"errors"

	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
)

// NewSqlStoreDentist estabelece conexão com a tabela dentists no banco de dados
func NewSqlStoreDentist(db *sql.DB) DentistInterface {
	return &SqlStore{
		db: db,
	}
}

// Create escreve a query para inserir um novo dentista no banco de dados dentro da tabela dentists e devolve possíveis erros
func (s *SqlStore) Create(dentist domain.Dentist) (domain.Dentist, error) {
	query := `INSERT INTO dentists (name, surname, registration)
	VALUES (?, ?, ?);`
	prepare, err := s.db.Prepare(query)
	if err != nil {
		return domain.Dentist{}, err
	}
	defer prepare.Close()

	res, err := prepare.Exec(dentist.Name, dentist.Surname, dentist.Registration)
	if err != nil {
		return domain.Dentist{}, nil
	}

	insertedId, _ := res.LastInsertId()

	dentist.Id = int(insertedId)

	_, err = res.RowsAffected()
	return dentist, err
}

// ReadAll escreve a query para buscar todos os dentistas no banco de dados dentro da tabela dentists e devolve possíveis erros
func (s *SqlStore) ReadAll() ([]domain.Dentist, error) {
	var (
		idRes int
		name string
		surname string
		registration string
		dentist domain.Dentist
		dentists []domain.Dentist
	)
	query := "SELECT id, name, surname, registration FROM dentists"
	rows, err := s.db.Query(query)
	if err != nil {
		return []domain.Dentist{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return []domain.Dentist{}, errors.New("")
	}

	for rows.Next() {
		err = rows.Scan(&idRes, &name, &surname, &registration)
		dentist.Id = idRes
		dentist.Name = name
		dentist.Surname = surname
		dentist.Registration = registration
		dentists = append(dentists, dentist)
	}

	return dentists, err
}

// Read escreve a query para buscar um dentista pelo seu ID no banco de dados dentro da tabela dentists e devolve possíveis erros
func (s *SqlStore) Read(id int) (domain.Dentist, error) {
	var (
		idRes int
		name string
		surname string
		registration string
		dentist domain.Dentist
	)
	query := `SELECT id, name, surname, registration FROM dentists
	WHERE id = ?;`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return domain.Dentist{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&idRes, &name, &surname, &registration)
		dentist.Id = idRes
		dentist.Name = name
		dentist.Surname = surname
		dentist.Registration = registration
	}

	return dentist, err
}

// Update escreve a query para atualizar um dentista no banco de dados dentro da tabela dentists e devolve possíveis erros
func (s *SqlStore) Update(dentist domain.Dentist) error {
	_, err := s.Read(dentist.Id)
	if err != nil {
		return err
	}
	query := `UPDATE dentists SET name = ?, surname = ?, registration = ?
	WHERE id = ?;`
	prepare, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := prepare.Exec(dentist.Name, dentist.Surname, dentist.Registration, dentist.Id)
	if err != nil {
		return nil
	}

	_, err = res.RowsAffected()
	return err
}

// Delete escreve a query para excluir um dentista pelo seu respectivo ID no banco de dados dentro da tabela dentists e devolve possíveis erros
func (s *SqlStore) Delete(id int) error {
	query := `DELETE FROM dentists WHERE id = ?;`
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
