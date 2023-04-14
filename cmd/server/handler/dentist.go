package handler

import (
	"errors"
	"strconv"

	"github.com/ctd/esp-backend-desafio-2.git/internal/domain"
	"github.com/ctd/esp-backend-desafio-2.git/internal/dentist"
	"github.com/ctd/esp-backend-desafio-2.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

// NewAppointmentHandler cria um novo controller de dentist
func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

//GetAll retorna todos os dentistas (dentist) cadastrados
func (h *dentistHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dentists, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
		}
		web.Success(c, 200, dentists)
	}
}

//GetByID retorna um dentista (dentist) por id
func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentist, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		web.Success(c, 200, dentist)
	}
}

// validateEmptys valida se os campos não estão vazios
func validateEmptys(dentist *domain.Dentist) (bool, error) {
	if (dentist.Name == "" || dentist.Surname == "" || dentist.Registration == "") {
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post insere um novo dentista
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

//Delete exclui um dentista
func (h *dentistHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

//Put atualiza um dentista
func (h *dentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var dentist domain.Dentist
		err = c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// Patch atualiza um dentista ou algum de seus campos
func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name           string `json:"name,omitempty"`
		Surname        string `json:"surname,omitempty"`
		Registration   string `json:"registration,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentist{
			Name:          r.Name,
			Surname:       r.Surname,
			Registration:  r.Registration,
		}
		d, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}
