package store

import (
	"database/sql"

	"github.com/JulietaAlfie/backendGo.git/internal/domain"
)

type sqlStorePatient struct {
	db *sql.DB
}

func NewSqlStorePatient(db *sql.DB) StoreInterfacePatient {
	return &sqlStorePatient{
		db: db,
	}
}

func (s *sqlStorePatient) ReadAll() ([]domain.Patient, error) {
	list := []domain.Patient{}

	rows, err := s.db.Query("select * from patients;")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var patient domain.Patient
		err := rows.Scan(&patient.Id, &patient.Name, &patient.Lastname, &patient.Residence, &patient.DNI, &patient.DischargeDate)
		if err != nil {
			return []domain.Patient{}, err
		}
		list = append(list, patient)
	}
	return list, nil
}

func (s *sqlStorePatient) Read(id int) (domain.Patient, error) {
	var patient domain.Patient 
	row := s.db.QueryRow("select * from patients where id = ?", id)
	err := row.Scan(&patient.Id, &patient.Name, &patient.Lastname, &patient.Residence, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *sqlStorePatient) ReadByDNI(dni int) (domain.Patient, error) {
	var patient domain.Patient 
	row := s.db.QueryRow("select * from patients where dni = ?", dni)
	err := row.Scan(&patient.Id, &patient.Name, &patient.Lastname, &patient.Residence, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *sqlStorePatient) Create(patient domain.Patient) (int, error) {
	query := "insert into patients (name, lastname, residence, dni, discharge_date) values (?, ?, ?, ?, ?)"
	st, err := s.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	res, err := st.Exec(&patient.Name, &patient.Lastname, &patient.Residence, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *sqlStorePatient) Update(patient domain.Patient) error {
	stmt, err := s.db.Prepare("UPDATE patients SET name = ?, lastname = ?, residence = ?, dni = ?, discharge_date = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(patient.Name, patient.Lastname, patient.Residence, patient.DNI, patient.DischargeDate, patient.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStorePatient) Delete(id int) error {
	stmt := "delete from patients where id = ?"
	_, err := s.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStorePatient) Exists(dni int) bool {
	var id int
	row := s.db.QueryRow("select id from patients where dni = ?", dni)
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}

	return false
}
