package dentist

import (
	"github.com/JulietaAlfie/backendGo.git/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Dentist, error)
	GetByID(id int) (domain.Dentist, error)
	Create(d domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
	Update(id int, d domain.Dentist) (domain.Dentist, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Dentist, error) {
	dentists := s.r.GetAll()
	return dentists, nil
}

func (s *service) GetByID(id int) (domain.Dentist, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

func (s *service) Create(d domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}
func (s *service) Update(id int, d domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	if d.Lastname != "" {
		dentist.Lastname = d.Lastname
	}
	if d.Name != "" {
		dentist.Name = d.Name
	}
	if d.License != "" {
		dentist.License = d.License
	}
	dentist, err = s.r.Update(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
