package adapter

import (
	"blazperic/santa/port"
	"database/sql"
)

type DbAdapter struct {
	database *sql.DB
}

func CreateDatabaseAdapter(db *sql.DB) *DbAdapter {
	return &DbAdapter{database: db}
}

func (d *DbAdapter) AddSanta(ime string, isChosen bool) error {

	sqlStatement := "INSERT INTO santas (ime, isChosen) VALUES (?, ?);"
	statement, err := d.database.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(ime, isChosen)
	if err != nil {
		return err
	}
	return nil
}

func (d *DbAdapter) GetAllSantas() ([]*port.SantasDTO, error) {
	sqlStatement := "SELECT id, ime, isChosen FROM santas;"

	statement, err := d.database.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	var santas []*port.SantasDTO

	for rows.Next() {
		var santa port.SantasDTO
		err = rows.Scan(&santa.Id, &santa.Ime, &santa.IsChosen)
		if err != nil {
			return nil, err
		}
		// santas = append(santas, &Santas{
		// 	Id:       santa.Id,
		// 	Ime:      santa.Ime,
		// 	IsChosen: santa.IsChosen})
		santas = append(santas, &santa)
	}

	return santas, nil

}

func (d *DbAdapter) ChoiceSanta() (*port.SantasDTO, error) {
	sqlStatement := "SELECT id, ime, isChosen FROM santas WHERE isChosen = false LIMIT 1;"

	statement, err := d.database.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	var santa port.SantasDTO

	for rows.Next() {
		err = rows.Scan(&santa.Id, &santa.Ime, &santa.IsChosen)
		if err != nil {
			return nil, err
		}
	}

	sqlStatement = "UPDATE santas SET isChosen = true WHERE id = ?;"

	statement, err = d.database.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(santa.Id)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	return &santa, nil

}
