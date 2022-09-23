package patient

import "github.com/JulietaAlfie/backendGo.git/internal/domain"

type Service interface {
	GetAll() ([]domain.Patient, error)
	GetByID(id int) (domain.Patient, error)
	Create(pac domain.Patient) (domain.Patient, error)
	Delete(id int) error
	Update(id int, pac domain.Patient) (domain.Patient, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Patient, error) {
	patients := s.r.GetAll()
	return patients, nil
}

func (s *service) GetByID(id int) (domain.Patient, error) {
	pac, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return pac, nil
}

func (s *service) Create(pac domain.Patient) (domain.Patient, error) {
	pac, err := s.r.Create(pac)
	if err != nil {
		return domain.Patient{}, err
	}
	return pac, nil
}
func (s *service) Update(id int, pac domain.Patient) (domain.Patient, error) {
	pacien, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	if pac.Name != "" {
		pacien.Name = pac.Name
	}
	if pac.Lastname != "" {
		pacien.Lastname = pac.Lastname
	}
	if pac.Residence != "" {
		pacien.Residence = pac.Residence
	}
	if pac.DNI != 0 {
		pacien.DNI = pac.DNI
	}
	if pac.DischargeDate != "" {
		pacien.DischargeDate = pac.DischargeDate
	}
	pacien, err = s.r.Update(id, pacien)
	if err != nil {
		return domain.Patient{}, err
	}
	return pacien, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
