package store

import (
	"database/sql"

	"github.com/JulietaAlfie/backendGo.git/internal/domain"
)

type sqlStoreDentist struct {
	db *sql.DB
}

func NewSqlStoreDentist(db *sql.DB) StoreInterfaceDentist {
	return &sqlStoreDentist{
		db: db,
	}
}

func (s *sqlStoreDentist) ReadAll() ([]domain.Dentist, error) {
	list := []domain.Dentist{}

	rows, err := s.db.Query("select * from dentists;")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var dentist domain.Dentist
		err := rows.Scan(&dentist.Id, &dentist.Lastname, &dentist.Name, &dentist.License)
		if err != nil {
			return []domain.Dentist{}, err
		}
		list = append(list, dentist)
	}
	return list, nil
}

func (s *sqlStoreDentist) Read(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	row := s.db.QueryRow("select * from dentists where id = ?", id)
	err := row.Scan(&dentist.Id, &dentist.Lastname, &dentist.Name, &dentist.License)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *sqlStoreDentist) Create(dentist domain.Dentist) (int, error) {
	query := "insert into dentists (lastname, name, license) values (?, ?, ?)"
	st, err := s.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	res, err := st.Exec(&dentist.Lastname, &dentist.Name, &dentist.License)
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

func (s *sqlStoreDentist) Update(dentist domain.Dentist) error {
	stmt, err := s.db.Prepare("UPDATE dentists SET lastname = ?, name = ?, license = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(dentist.Lastname, dentist.Name, dentist.License, dentist.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreDentist) Delete(id int) error {
	stmt := "delete from dentists where id = ?"
	_, err := s.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreDentist) Exists(license string) bool {
	var id int
	row := s.db.QueryRow("select id from dentists where license = ?", license)
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}

	return false
}
