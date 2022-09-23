package appointment

import (
	"errors"
	"fmt"

	"github.com/JulietaAlfie/backendGo.git/internal/domain"
	"github.com/JulietaAlfie/backendGo.git/pkg/store"
)

type Repository interface {
	GetAll() []domain.Appointment
	GetByID(id int) (domain.Appointment, error)
	GetByDNI(dni int) (domain.Appointment, error)
	Create(appointment domain.Appointment) (domain.Appointment, error)
	CreateByDniAndLicence(dni int, license string, date string, time string, description string) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfaceAppointment
}

func NewRepository(storage store.StoreInterfaceAppointment) Repository {
	return &repository{storage}
}

func (r *repository) GetAll() []domain.Appointment {
	appointments, err := r.storage.ReadAll()
	if err != nil {
		return []domain.Appointment{}
	}
	return appointments
}

func (r *repository) GetByID(id int) (domain.Appointment, error) {
	appointment, err := r.storage.Read(id)
	if err != nil {
		fmt.Println(err)
		return domain.Appointment{}, errors.New("appointment not found")
	}
	return appointment, nil

}

func (r *repository) GetByDNI(dni int) (domain.Appointment, error) {
	appointment, err := r.storage.ReadByDNI(dni)
	if err != nil {
		fmt.Println(err)
		return domain.Appointment{}, errors.New("appointment not found")
	}
	return appointment, nil
}

func (r *repository) Create(appointment domain.Appointment) (domain.Appointment, error) {
	id, err := r.storage.Create(appointment)
	if err != nil {
		fmt.Println(err)
		return domain.Appointment{}, errors.New("error creating appointment")
	}
	appointment.Id = id
	return appointment, nil
}

func (r *repository) CreateByDniAndLicence(dni int, license string, date string, time string, description string) (domain.Appointment, error) {
	var appointment domain.Appointment
	app, err := r.storage.CreateByDniAndLicence(dni, license, date, time, description)
	if err != nil {
		return domain.Appointment{}, errors.New("error creating appointment")
	}
	appointment = app
	appointment.Date = date
	appointment.Description = description
	appointment.Time = time
	return appointment, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, appointment domain.Appointment) (domain.Appointment, error) {
	err := r.storage.Update(appointment)
	if err != nil {
		return domain.Appointment{}, errors.New("an error occurred updating appointment")
	}
	return appointment, nil
}
