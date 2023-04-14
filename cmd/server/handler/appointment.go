package handler

import (
	"errors"
	"strconv"

	"github.com/ctd/esp-backend-desafio-2.git/internal/appointment"
	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
	"github.com/ctd/esp-backend-desafio-2.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	a appointment.Service
}

// NewAppointmentHandler cria um novo controller de appointments
func NewAppointmentHandler(a appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		a: a,
	}
}

//GetById retorna uma consulta (appointment) por id
func (h *appointmentHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		appointment, err := h.a.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		web.Success(c, 200, appointment)
	}
}

// validateEmptysappointment valida se os campos não estão vazios
func validateEmptysappointment(appointment *domain.Appointment) (bool, error) {
	var idDentistString = strconv.Itoa(appointment.IdDentist)
	var idPatientString = strconv.Itoa(appointment.IdPatient)
	if idDentistString == "" || idPatientString == "" || appointment.ApptDate == "" || appointment.Description == "" {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post cria uma nova consulta
func (h *appointmentHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment domain.Appointment
		err := c.ShouldBindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysappointment(&appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		d, err := h.a.Create(appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

//Delete exclui uma consulta
func (h *appointmentHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.a.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

//Put atualiza uma consulta
func (h *appointmentHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.a.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var appointment domain.Appointment
		err = c.ShouldBindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysappointment(&appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.a.Update(id, appointment)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// Patch atualiza uma consulta ou algum de seus campos
func (h *appointmentHandler) Patch() gin.HandlerFunc {
	type Request struct {
		IdDentist   int    `json:"id_dentist,omitempty"`
		IdPatient   int    `json:"id_patient,omitempty"`
		ApptDate    string `json:"appt_date,omitempty"`
		Description string `json:"description,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.a.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}

		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Appointment{
			IdDentist:   r.IdDentist,
			IdPatient:   r.IdPatient,
			ApptDate:    r.ApptDate,
			Description: r.Description,
		}
		d, err := h.a.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// GetByRg busca uma consulta pelo Rg do paciente
func (h *appointmentHandler) GetByRg() gin.HandlerFunc {
	return func(c *gin.Context) {

		rgParam := c.Param("rg")
		 _, err := strconv.Atoi(rgParam)

		if err != nil {
			web.Failure(c, 400, errors.New("invalid rg"))
			return
		}
		appointment, err := h.a.GetByRg(rgParam)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		web.Success(c, 200, appointment)
	}
}
