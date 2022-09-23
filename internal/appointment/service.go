package appointment

import (
	"fmt"

	"github.com/JulietaAlfie/backendGo.git/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Appointment, error)
	GetByID(id int) (domain.Appointment, error)
	GetByDNI(dni int) (domain.Appointment, error)
	Create(appointment domain.Appointment) (domain.Appointment, error)
	CreateByDniAndLicence(dni int, license string, date string, time string, description string) (domain.Appointment, error)
	Delete(id int) error
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Appointment, error) {
	appointments := s.r.GetAll()
	return appointments, nil
}

func (s *service) GetByID(id int) (domain.Appointment, error) {
	appointment, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) GetByDNI(dni int) (domain.Appointment, error) {
	appointment, err := s.r.GetByDNI(dni)
	if err != nil {
		fmt.Println(err)
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) Create(appointment domain.Appointment) (domain.Appointment, error) {
	appointment, err := s.r.Create(appointment)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) CreateByDniAndLicence(dni int, license string, date string, time string, description string) (domain.Appointment, error) {

	appointment, err := s.r.CreateByDniAndLicence(dni, license, date, time, description)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil

}

func (s *service) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	appointmentDB, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	if appointment.Description != "" {
		appointmentDB.Description = appointment.Description
	}
	if appointment.Date != "" {
		appointmentDB.Date = appointment.Date
	}
	if appointment.Time != "" {
		appointmentDB.Time = appointment.Time
	}
	appointmentDB, err = s.r.Update(id, appointmentDB)
	if err != nil {
		return domain.Appointment{}, err
	}

	return appointmentDB, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
