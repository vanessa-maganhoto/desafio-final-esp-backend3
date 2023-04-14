package store

import (
	"database/sql"

	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
)

// NewSqlStoreAppointment estabelece conexão com a tabela appointments no banco de dados
func NewSqlStoreAppointment(db *sql.DB) AppointmentInterface {
	return &SqlStore{
		db: db,
	}
}

// CreateAppointment escreve a query para inserir uma nova consulta no banco de dados dentro da appointments e devolve possíveis erros
func (s *SqlStore) CreateAppointment(appointment domain.Appointment) (domain.AppointmentResponse, error) {
	var (
		dentist         domain.Dentist
		patient         domain.Patient
		appointmentResp domain.AppointmentResponse
	)

	query := `INSERT INTO appointments (idDentist, idPatient, apptDate, description)
	VALUES (?, ?, ?, ?);`
	prepare, err := s.db.Prepare(query)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	defer prepare.Close()

	res, err := prepare.Exec(appointment.IdDentist, appointment.IdPatient, appointment.ApptDate, appointment.Description)
	if err != nil {
		return domain.AppointmentResponse{}, nil
	}

	insertedId, _ := res.LastInsertId()

	appointment.Id = int(insertedId)

	dentist, err = NewSqlStoreDentist(s.db).Read(appointment.IdDentist)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	patient, err = NewSqlStorePatient(s.db).ReadPatient(appointment.IdPatient)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}

	appointmentResp.Id = appointment.Id
	appointmentResp.DentistD = dentist
	appointmentResp.PatientP = patient
	appointmentResp.ApptDate = appointment.ApptDate
	appointmentResp.Description = appointment.Description

	_, err = res.RowsAffected()
	if err != nil {
		return domain.AppointmentResponse{}, err
	}

	return appointmentResp, err
}

// ReadAppointment escreve a query para buscar uma consulta pelo seu ID no banco de dados dentro da tabela appointments e devolve possíveis erros
func (s *SqlStore) ReadAppointment(id int) (domain.AppointmentResponse, error) {
	var (
		idRes           int
		idDentist       int
		idPatient       int
		dentist         domain.Dentist
		patient         domain.Patient
		apptDate        string
		description     string
		appointment     domain.Appointment
		appointmentResp domain.AppointmentResponse
	)

	query := `SELECT id, idDentist, idPatient, apptDate, description
	FROM appointments
	WHERE id = ?;`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&idRes, &idDentist, &idPatient, &apptDate, &description); err != nil {
			return domain.AppointmentResponse{}, err
		}
		appointment.Id = idRes
		appointment.IdDentist = idDentist
		appointment.IdPatient = idPatient
		appointment.ApptDate = apptDate
		appointment.Description = description
	}

	dentist, err = NewSqlStoreDentist(s.db).Read(appointment.IdDentist)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	patient, err = NewSqlStorePatient(s.db).ReadPatient(appointment.IdPatient)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}

	appointmentResp.Id = appointment.Id
	appointmentResp.DentistD = dentist
	appointmentResp.PatientP = patient
	appointmentResp.ApptDate = appointment.ApptDate
	appointmentResp.Description = appointment.Description

	return appointmentResp, err
}

// UpdateAppointment escreve a query para atualizar uma consulta no banco de dados dentro da tabela appointments e devolve possíveis erros
func (s *SqlStore) UpdateAppointment(appointment domain.Appointment) (domain.AppointmentResponse, error) {
	var (
		dentist         domain.Dentist
		patient         domain.Patient
		appointmentResp domain.AppointmentResponse
	)

	_, err := s.ReadAppointment(appointment.Id)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	query := `UPDATE appointments SET idDentist = ?, idPatient = ?, apptDate = ?, description = ?
	WHERE id = ?;`
	prepare, err := s.db.Prepare(query)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}

	res, err := prepare.Exec(appointment.IdDentist, appointment.IdPatient, appointment.ApptDate, appointment.Description, appointment.Id)
	if err != nil {
		return domain.AppointmentResponse{}, nil
	}

	dentist, err = NewSqlStoreDentist(s.db).Read(appointment.IdDentist)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	patient, err = NewSqlStorePatient(s.db).ReadPatient(appointment.IdPatient)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}

	appointmentResp.Id = appointment.Id
	appointmentResp.DentistD = dentist
	appointmentResp.PatientP = patient
	appointmentResp.ApptDate = appointment.ApptDate
	appointmentResp.Description = appointment.Description

	_, err = res.RowsAffected()
	return appointmentResp, err
}

// DeleteAppointment escreve a query para excluir uma consulta pelo seu respectivo ID no banco de dados dentro da tabela appointments e devolve possíveis erros
func (s *SqlStore) DeleteAppointment(id int) error {
	query := `DELETE FROM appointments WHERE id = ?;`
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

// ReadAppointmentByRgPatient escreve a query para buscar a consulta de um paciente pelo seu número de RG devolve possíveis erros
func (s *SqlStore) ReadAppointmentByRgPatient(rg string) (domain.AppointmentResponse, error) {
	var (
		idRes           int
		idDentist       int
		idPatient       int
		dentist         domain.Dentist
		patient         domain.Patient
		apptDate        string
		description     string
		appointment     domain.Appointment
		appointmentResp domain.AppointmentResponse
	)

	query := `SELECT a.id, a.idDentist, a.idPatient, a.apptDate, a.description
	FROM patients p
	JOIN appointments a 
	ON p.id = a.idPatient
	WHERE p.documentNumber = ?;`

	rows, err := s.db.Query(query, rg)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	defer rows.Close()

	// if !rows.Next() {
	// 	return domain.AppointmentResponse{}, errors.New("")
	// }

	for rows.Next() {
		if err = rows.Scan(&idRes, &idDentist, &idPatient, &apptDate, &description); err != nil {
			return domain.AppointmentResponse{}, err
		}
		appointment.Id = idRes
		appointment.IdDentist = idDentist
		appointment.IdPatient = idPatient
		appointment.ApptDate = apptDate
		appointment.Description = description
	}

	dentist, err = NewSqlStoreDentist(s.db).Read(appointment.IdDentist)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}
	patient, err = NewSqlStorePatient(s.db).ReadPatient(appointment.IdPatient)
	if err != nil {
		return domain.AppointmentResponse{}, err
	}

	appointmentResp.Id = appointment.Id
	appointmentResp.DentistD = dentist
	appointmentResp.PatientP = patient
	appointmentResp.ApptDate = appointment.ApptDate
	appointmentResp.Description = appointment.Description

	return appointmentResp, err

}
