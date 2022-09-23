package dentist

import (
	
	"github.com/JulietaAlfie/backendGo.git/internal/domain"
	"errors"
	"fmt"
	"github.com/JulietaAlfie/backendGo.git/pkg/store"
)

type Repository interface {
	GetAll() []domain.Dentist
	GetByID(id int) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfaceDentist
}

func NewRepository(storage store.StoreInterfaceDentist) Repository {
	return &repository{storage}
}

func (r *repository) GetAll() []domain.Dentist {
	dentists, err := r.storage.ReadAll()
	if err != nil {
		return []domain.Dentist{}
	}
	return dentists
}

func (r *repository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	return dentist, nil

}

func (r *repository) Create(dentist domain.Dentist) (domain.Dentist, error) {
	if r.storage.Exists(dentist.License) {
		return domain.Dentist{}, errors.New("existing dentist license")
	}
	id, err := r.storage.Create(dentist)
	if err != nil {
		fmt.Println(err)
		return domain.Dentist{}, errors.New("an error occurred creating dentist")
	}
	dentist.Id = id
	return dentist, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, dentist domain.Dentist) (domain.Dentist, error) {
	if !r.storage.Exists(dentist.License) {
		return domain.Dentist{}, errors.New("existing dentist license")
	}
	err := r.storage.Update(dentist)
	if err != nil {
		return domain.Dentist{}, errors.New("an error occurred updating dentist")
	}
	return dentist, nil
}
