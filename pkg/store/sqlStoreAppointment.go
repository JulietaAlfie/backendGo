package store

import (
	"database/sql"

	"github.com/JulietaAlfie/backendGo.git/internal/domain"
)

type sqlStoreAppointment struct {
	db *sql.DB
}

func NewSqlStoreAppointment(db *sql.DB) StoreInterfaceAppointment {
	return &sqlStoreAppointment{
		db: db,
	}
}

func (s *sqlStoreAppointment) ReadAll() ([]domain.Appointment, error) {
	list := []domain.Appointment{}

	rows, err := s.db.Query("select * from appointments;")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(&appointment.Id, &appointment.Patient, &appointment.Dentist, &appointment.Date, &appointment.Time, &appointment.Description)
		if err != nil {
			return []domain.Appointment{}, err
		}
		list = append(list, appointment)
	}
	return list, nil
}

func (s *sqlStoreAppointment) Read(id int) (domain.Appointment, error) {
	var appointment domain.Appointment
	row := s.db.QueryRow("select t.id, t.patient_id, p.name, p.lastname, p.residence, p.dni, p.discharge_date, t.dentist_id, o.name, o.lastname, o.license, t.date, t.time, t.description from appointments t	inner join dentists o on t.dentist_id = o.id 	inner join patients p on t.patient_id = p.id where t.id= ?", id)
	err := row.Scan(&appointment.Id, &appointment.Patient.Id, &appointment.Patient.Name, &appointment.Patient.Lastname, &appointment.Patient.Residence, &appointment.Patient.DNI, &appointment.Patient.DischargeDate, &appointment.Dentist.Id, &appointment.Dentist.Name, &appointment.Dentist.Lastname, &appointment.Dentist.License, &appointment.Date, &appointment.Time, &appointment.Description)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

// ReadByDNI devuelve un turno por dni del paciente
// func (s *sqlStoreTurno) ReadByDNI(dni int) ([]domain.Turno, error) {
// 	turnos := []domain.Turno{}

// 	rows, err := s.db.Query("SELECT turnos.id, turnos.paciente_id, turnos.odontologo_id, fecha, hora, descripcion FROM turnos INNER JOIN pacientes ON turnos.paciente_id = pacientes.id WHERE pacientes.dni = ?", dni)
// 	if err != nil {
// 		return turnos, err
// 	}

// 	for rows.Next() {
// 		turno := domain.Turno{}

// 		err := rows.Scan(&turno.Id, &turno.Paciente.Id, &turno.Odontologo.Id, &turno.Fecha, &turno.Hora, &turno.Descripcion)
// 		if err != nil {
// 			return []domain.Turno{}, err
// 		}

// 		paciente := s.db.QueryRow("SELECT * FROM pacientes WHERE id = ?", turno.Paciente.Id)
// 		err = paciente.Scan(&turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI, &turno.Paciente.FechaAlta)
// 		if err != nil {
// 			return []domain.Turno{}, err
// 		}
// 		odontologo := s.db.QueryRow("SELECT * FROM odontologos WHERE id = ?", turno.Odontologo.Id)
// 		err = odontologo.Scan(&turno.Odontologo.Id, &turno.Odontologo.Apellido, &turno.Odontologo.Nombre, &turno.Odontologo.Matricula)
// 		if err != nil {
// 			return []domain.Turno{}, err
// 		}

// 		turnos = append(turnos, turno)
// 	}

// 	return turnos, nil
// }

func (s *sqlStoreAppointment) ReadByDNI(dni int) (domain.Appointment, error) {
	var appointment domain.Appointment
	row := s.db.QueryRow("select t.id, t.patient_id, p.name, p.lastname, p.residence, p.dni, p.discharge_date, t.dentist_id, o.name, o.lastname, o.license, t.date, t.time, t.description from appointments t 	inner join dentists o on t.dentist_id = o.id inner join patients p on t.patient_id = p.id where dni= ?", dni)
	err := row.Scan(&appointment.Id, &appointment.Patient.Id, &appointment.Patient.Name, &appointment.Patient.Lastname, &appointment.Patient.Residence, &appointment.Patient.DNI, &appointment.Patient.DischargeDate, &appointment.Dentist.Id, &appointment.Dentist.Name, &appointment.Dentist.Lastname, &appointment.Dentist.License, &appointment.Date, &appointment.Time, &appointment.Description)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *sqlStoreAppointment) Create(appointment domain.Appointment) (int, error) {
	query := "insert into appointments (patient_id, dentist_id, date, time, description) values (?, ?, ?, ?, ?)"
	st, err := s.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	res, err := st.Exec(&appointment.Patient.Id, &appointment.Dentist.Id, &appointment.Date, &appointment.Time, &appointment.Description)
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

// CreateByDNIAndLicense agrega un nuevo appointment
func (s *sqlStoreAppointment) CreateByDniAndLicence(dni int, license string, date string, time string, description string) (domain.Appointment, error) {
	var appointment domain.Appointment
	patient := s.db.QueryRow("select * from patients where dni = ?", dni)
	err := patient.Scan(&appointment.Patient.Id, &appointment.Patient.Name, &appointment.Patient.Lastname, &appointment.Patient.Residence, &appointment.Patient.DNI, &appointment.Patient.DischargeDate)
	if err != nil {
		return domain.Appointment{}, err
	}
	dentist := s.db.QueryRow("select * from dentists where license = ?", license)
	err2 := dentist.Scan(&appointment.Dentist.Id, &appointment.Dentist.Lastname, &appointment.Dentist.Name, &appointment.Dentist.License)
	if err2 != nil {
		return domain.Appointment{}, err
	}
	query := "insert into appointments (patient_id, dentist_id, date, time, description) values (?, ?, ?, ?, ?)"
	st, err := s.db.Prepare(query)
	if err != nil {
		return domain.Appointment{}, err
	}
	res, err := st.Exec(appointment.Patient.Id, appointment.Dentist.Id, date, time, description)
	if err != nil {
		return domain.Appointment{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Appointment{}, err
	}
	appointment.Id = int(id)
	_, err = res.RowsAffected()
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

// Update actualiza un appointment
func (s *sqlStoreAppointment) Update(appointment domain.Appointment) error {
	stmt, err := s.db.Prepare("UPDATE appointments SET patient_id = ?, dentist_id = ?, date = ?, time = ?, description = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(&appointment.Patient.Id, &appointment.Dentist.Id, &appointment.Date, &appointment.Time, &appointment.Description, appointment.Id)
	if err != nil {
		return err
	}
	return nil
}

// Delete elimina un appointment
func (s *sqlStoreAppointment) Delete(id int) error {
	stmt := "delete from appointments where id = ?"
	_, err := s.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
