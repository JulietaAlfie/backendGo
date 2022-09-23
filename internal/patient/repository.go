package patient

import (
	"errors"
	"fmt"
	"github.com/JulietaAlfie/backendGo.git/internal/domain"
	"github.com/JulietaAlfie/backendGo.git/pkg/store"
)

type Repository interface {
	GetAll() []domain.Patient
	GetByID(id int) (domain.Patient, error)
	GetByDNI(dni int) (domain.Patient, error)
	Create(od domain.Patient) (domain.Patient, error)
	Update(id int, od domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfacePatient
}

func NewRepository(storage store.StoreInterfacePatient) Repository {
	return &repository{storage}
}

func (r *repository) GetAll() []domain.Patient {
	patients, err := r.storage.ReadAll()
	if err != nil {

		fmt.Println(patients, err)
		return []domain.Patient{}
	}
	return patients
}

func (r *repository) GetByID(id int) (domain.Patient, error) {
	patient, err := r.storage.Read(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil

}

func (r *repository) GetByDNI(dni int) (domain.Patient, error) {
	patient, err := r.storage.ReadByDNI(dni)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil

}

func (r *repository) Create(pac domain.Patient) (domain.Patient, error) {
	if r.storage.Exists(pac.DNI) {
		return domain.Patient{}, errors.New("that dni already exists")
	}
	id, err := r.storage.Create(pac)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return domain.Patient{}, errors.New("error creating patient")
	}
	pac.Id = id
	return pac, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, pac domain.Patient) (domain.Patient, error) {
	if !r.storage.Exists(pac.DNI) {
		return domain.Patient{}, errors.New("existing identity document")
	}
	err := r.storage.Update(pac)
	if err != nil {
		return domain.Patient{}, errors.New("error updating patient")
	}
	return pac, nil
}
