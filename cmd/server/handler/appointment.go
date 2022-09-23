package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/JulietaAlfie/backendGo.git/internal/appointment"
	"github.com/JulietaAlfie/backendGo.git/internal/domain"
	"github.com/JulietaAlfie/backendGo.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	s appointment.Service
}

func NewAppointmentHandler(s appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		s: s,
	}
}

func (h *appointmentHandler) Post() gin.HandlerFunc {
	type Request struct {
		PatientId   int    `json:"patient" binding:"required"`
		DentistId   int    `json:"dentist" binding:"required"`
		Date        string `json:"date" binding:"required"`
		Time        string `json:"time" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	return func(c *gin.Context) {
		var appointment domain.Appointment
		err := c.ShouldBindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysAppointment(&appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		app, err := h.s.Create(appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, app)
	}
}

func (h *appointmentHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		appointment, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		web.Success(c, 200, appointment)
	}
}

func (h *appointmentHandler) Put() gin.HandlerFunc {
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
		valid, err := validateEmptysAppointment(&appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		app, err := h.s.Update(id, appointment)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, app)
	}
}

func (h *appointmentHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Patient     domain.Patient `json:"patient,omitempty"`
		Dentist     domain.Dentist `json:"dentist,omitempty"`
		Date        string         `json:"date,omitempty"`
		Time        string         `json:"time,omitempty"`
		Description string         `json:"description,omitempty"`
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
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Appointment{
			Patient:     req.Patient,
			Dentist:     req.Dentist,
			Date:        req.Date,
			Time:        req.Time,
			Description: req.Description,
		}
		app, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, app)
	}
}

func (h *appointmentHandler) Delete() gin.HandlerFunc {
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

func (h *appointmentHandler) PostByDniAndLicence() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			Date        string `json:"date" binding:"required"`
			Time        string `json:"time" binding:"required"`
			Description string `json:"description" binding:"required"`
		}
		dniParam, _ := strconv.Atoi(c.Param("dni"))
		licenseParam := c.Param("license")
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		tur, err := h.s.CreateByDniAndLicence(dniParam, licenseParam, req.Date, req.Time, req.Description)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, tur)
	}
}
func (h *appointmentHandler) GetByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			fmt.Println(err)
			web.Failure(c, 400, errors.New("invalid dni"))
			return
		}
		appointment, err := h.s.GetByDNI(dni)
		if err != nil {
			fmt.Println(err)
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		web.Success(c, 200, appointment)
	}
}

func validateEmptysAppointment(appointment *domain.Appointment) (bool, error) {
	switch {
	case appointment.Patient == domain.Patient{}:
		return false, errors.New("patient was empty")
	case appointment.Dentist == domain.Dentist{}:
		return false, errors.New("dentist was empty")
	case appointment.Date == "":
		return false, errors.New("date was empty")
	case appointment.Time == "":
		return false, errors.New("time was empty")
	case appointment.Description == "":
		return false, errors.New("description was empty")
	}
	return true, nil
}
