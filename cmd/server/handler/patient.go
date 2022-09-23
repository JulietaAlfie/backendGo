package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/JulietaAlfie/backendGo.git/internal/domain"
	"github.com/JulietaAlfie/backendGo.git/internal/patient"
	"github.com/JulietaAlfie/backendGo.git/pkg/web"

	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.Service
}

func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}
// StorePatient godoc
// @Summary Store patient
// @Tags Patients
// @Description store patient
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param patient body domain.Patient true "Patient to store"
// @Success 200 {object} web.response
// @Failure 400 {object} web.response
// @Router /patients [post]
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
		p, err := h.s.Create(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}
// ListPatients godoc
// @Summary List patient
// @Tags Patients
// @Description get patient
// @Produce  json
// @Success 200 {object} web.response
// @Failure 422 {object} web.errorResponse
// @Router /patients [get]
func (h *patientHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		patients, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, 422, errors.New("patients could not be brought"))
			return
		}
		web.Success(c, 200, patients)
	}
}
// Patient godoc
// @Summary patient
// @Tags Patients
// @Description get patient
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} web.response
// @Failure 404 {object} web.response
// @Router /patients/{id} [get]
func (h *patientHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, 200, patient)
	}
}

// ModifyPatient godoc
// @Summary Modify patient
// @Tags Patients
// @Description modify patient
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param patient body domain.Patient true "Patient to store"
// @Success 200 {object} web.response
// @Failure 400 {object} web.response
// @Failure 401 {object} web.response
// @Failure 404 {object} web.errorResponse
// @Router /patients/{id} [put]
func (h *patientHandler) Put() gin.HandlerFunc {
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
		pat, err := h.s.Update(id, patient)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, pat)
	}
}
// ModifyPatient godoc
// @Summary Modify patient
// @Tags Patients
// @Description modify patient
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param patient body domain.Patient true "Patient to store"
// @Success 200 {object} web.response
// @Router /patients/{id} [patch]
func (h *patientHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Lastname      string `json:"lastname,omitempty"`
		Name          string `json:"name,omitempty"`
		Residence     string `json:"residence,omitempty"`
		DNI           int    `json:"dni,omitempty"`
		DischargeDate string `json:"discharge_date,omitempty"`
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
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Patient{
			Lastname:      req.Lastname,
			Name:          req.Name,
			Residence:     req.Residence,
			DNI:           req.DNI,
			DischargeDate: req.DischargeDate,
		}
		pat, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, pat)
	}
}

// DeletePatient godoc
// @Summary Delete patient
// @Tags Patients
// @Description delete patient
// @Param token header string true "token"
// @Param id path int true "Patient ID"
// @Success 204 {object} web.response
// @Failure 400 {object} web.response
// @Failure 401 {object} web.response
// @Failure 404 {object} web.response
// @Router /patients/{id} [delete]
func (h *patientHandler) Delete() gin.HandlerFunc {
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

func validateEmptysPatient(patient *domain.Patient) (bool, error) {
	switch {
	case patient.Lastname == "":
		return false, errors.New("lastname was empty")
	case patient.Name == "":
		return false, errors.New("name was empty")
	case patient.Residence == "":
		return false, errors.New("residence was empty")
	case patient.DNI == 0:
		return false, errors.New("dni was empty")
	case patient.DischargeDate == "":
		return false, errors.New("discharge_date was empty")
	}
	return true, nil
}
