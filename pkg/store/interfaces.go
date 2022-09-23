package store

import "github.com/JulietaAlfie/backendGo.git/internal/domain"

type StoreInterfaceDentist interface {
	Read(id int) (domain.Dentist, error)
	ReadAll() ([]domain.Dentist, error)
	Create(dentist domain.Dentist) (int, error)
	Update(dentist domain.Dentist) error
	Delete(id int) error
	Exists(license string) bool
}

type StoreInterfacePatient interface {
	Read(id int) (domain.Patient, error)
	ReadByDNI(dni int) (domain.Patient, error)
	ReadAll() ([]domain.Patient, error)
	Create(patient domain.Patient) (int, error)
	Update(patient domain.Patient) error
	Delete(id int) error
	Exists(dni int) bool
}

type StoreInterfaceAppointment interface {
	Read(id int) (domain.Appointment, error)
	ReadByDNI(dni int) (domain.Appointment, error)
	ReadAll() ([]domain.Appointment, error)
	Create(appointment domain.Appointment) (int, error)
	CreateByDniAndLicence(dni int, license string, date string, time string, description string) (domain.Appointment, error)
	Update(appointment domain.Appointment) error
	Delete(id int) error
}
