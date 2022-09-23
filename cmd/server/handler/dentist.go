package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/JulietaAlfie/backendGo.git/internal/dentist"
	"github.com/JulietaAlfie/backendGo.git/internal/domain"
	"github.com/JulietaAlfie/backendGo.git/pkg/web"

	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}
// StoreDentist godoc
// @Summary Store dentist
// @Tags Dentists
// @Description store dentist
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param dentist body domain.Dentist true "Dentist to store"
// @Success 200 {object} web.response
// @Failure 400 {object} web.response
// @Router /dentists [post]
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysDentistFields(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		dent, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, dent)
	}
}
// ListDentists godoc
// @Summary List dentists
// @Tags Dentists
// @Description get dentists
// @Produce  json
// @Success 200 {object} web.response
// @Failure 422 {object} web.errorResponse
// @Router /dentists [get]
func (h *dentistHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dentists, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, 422, errors.New("dentists could not be brought"))
			return
		}
		web.Success(c, 200, dentists)
	}
}
// Dentist godoc
// @Summary dentist
// @Tags Dentists
// @Description get dentists
// @Produce  json
// @Param id path int true "Dentist ID"
// @Success 200 {object} web.response
// @Failure 404 {object} web.response
// @Router /dentists/{id} [get]
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
// ModifyDentist godoc
// @Summary Modify dentist
// @Tags Dentists
// @Description modify dentist
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param dentist body domain.Dentist true "Dentist to store"
// @Success 200 {object} web.response
// @Failure 400 {object} web.response
// @Failure 401 {object} web.response
// @Failure 404 {object} web.errorResponse
// @Router /dentists/{id} [put]
func (h *dentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
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
		valid, err := validateEmptysDentistFields(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		dent, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, dent)
	}
}
// ModifyDentist godoc
// @Summary Modify dentist
// @Tags Dentists
// @Description modify dentist
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param dentist body domain.Dentist true "Dentist to store"
// @Success 200 {object} web.response
// @Router /dentists/{id} [patch]
func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Lastname string `json:"lastname,omitempty"`
		Name     string `json:"name,omitempty"`
		License  string `json:"license,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var req Request
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
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentist{
			Lastname: req.Lastname,
			Name:     req.Name,
			License:  req.License,
		}
		dent, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, dent)
	}
}
// DeleteDentist godoc
// @Summary Delete dentist
// @Tags Dentists
// @Description delete dentist
// @Param token header string true "token"
// @Param id path int true "Dentist ID"
// @Success 204 {object} web.response
// @Failure 400 {object} web.response
// @Failure 401 {object} web.response
// @Failure 404 {object} web.response
// @Router /dentists/{id} [delete]
func (h *dentistHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
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

func validateEmptysDentistFields(dentist *domain.Dentist) (bool, error) {
	switch {
	case dentist.Lastname == "":
		return false, errors.New("lastname was empty")
	case dentist.Name == "":
		return false, errors.New("name was empty")
	case dentist.License == "":
		return false, errors.New("license was empty")
	}
	return true, nil
}
