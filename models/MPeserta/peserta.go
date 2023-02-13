package MPeserta

import (
	"fmt"
	"lms/config"
	"lms/config/helper"
	"lms/structs"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

type Peserta struct{ structs.Peserta }

func ViewPeserta(page int) ([]Peserta, error) {

	db := config.ConPg()
	defer db.Close()
	var pesertas []Peserta

	sqlStatement := `SELECT nrp,username,password,last_ip,created_at,updated_at FROM peserta`
	perPage := 2
	sqlStatement = fmt.Sprintf("%s LIMIT %d OFFSET %d", sqlStatement, perPage, (page-1)*perPage)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var peserta Peserta
		err = rows.Scan(&peserta.NRP, &peserta.Username, &peserta.Password, &peserta.LastIP, &peserta.CreatedAt, &peserta.UpdatedAt)
		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}
		pesertas = append(pesertas, peserta)
	}
	return pesertas, err
}

func UpdatePeserta(peserta Peserta) int64 {
	db := config.ConPg()
	defer db.Close()
	sqlStatement := `UPDATE peserta SET username=$2, password=$3, last_ip=$4, updated_at=current_timestamp WHERE nrp=$1`
	var paswd string = helper.Hash(peserta.Password)
	res, err := db.Exec(sqlStatement, peserta.NRP, peserta.Username, paswd, peserta.LastIP)
	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}
	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)
	return rowsAffected
}

func DeletePeserta(peserta Peserta) int64 {
	db := config.ConPg()
	defer db.Close()
	sqlStatement := `DELETE FROM peserta WHERE nrp=$1`
	res, err := db.Exec(sqlStatement, peserta.NRP)
	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}
	fmt.Printf("Total data yang terhapus %v", rowsAffected)
	return rowsAffected
}

func CreatePeserta(peserta Peserta) int64 {
	db := config.ConPg()
	defer db.Close()
	sqlStatement := `INSERT INTO peserta (nrp,username,password,last_ip,created_at) VALUES ($1,$2,$3,$4,current_timestamp) RETURNING nrp`
	var id int64
	err := db.QueryRow(sqlStatement, peserta.NRP, peserta.Username, peserta.Password, peserta.LastIP).Scan(&id)
	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}
	fmt.Printf("Insert data single record %v", id)
	return id
}
