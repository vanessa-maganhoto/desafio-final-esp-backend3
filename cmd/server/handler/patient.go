package handler

import (
	"errors"
	"strconv"

	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
	"github.com/ctd/esp-backend-desafio-2.git/internal/patient"
	"github.com/ctd/esp-backend-desafio-2.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	p patient.Service
}

// NewAppointmentHandler cria um novo controller de patient
func NewPatientHandler(p patient.Service) *patientHandler {
	return &patientHandler{
		p: p,
	}
}

//GetAll retorna todos os pacientes (patient) cadastrados
func (h *patientHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, err := h.p.GetAll()
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
		}
		web.Success(c, 200, patients)
	}
}

//GetByID retorna um paciente (patient) por id
func (h *patientHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.p.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, 200, patient)
	}
}

// validateEmptysPatient valida se os campos não estão vazios
func validateEmptysPatient(patient *domain.Patient) (bool, error) {
	if patient.Name == "" || patient.Surname == "" || patient.DocumentNumber == "" || patient.RegistrationDate == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post insere um novo paciente
func (h *patientHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		err := c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysPatient(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.p.Create(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

//Delete exclui um paciente
func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.p.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

//Put atualiza um paciente
func (h *patientHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.p.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var patient domain.Patient
		err = c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysPatient(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.p.Update(id, patient)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// Patch atualiza um paciente ou algum de seus campos
func (h *patientHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name             string `json:"name,omitempty"`
		Surname          string `json:"surname,omitempty"`
		DocumentNumber   string `json:"document_number,omitempty"`
		RegistrationDate string `json:"registration_date,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.p.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Patient{
			Name:             r.Name,
			Surname:          r.Surname,
			DocumentNumber:   r.DocumentNumber,
			RegistrationDate: r.RegistrationDate,
		}
		d, err := h.p.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}
