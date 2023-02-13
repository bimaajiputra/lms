package MAuth

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"lms/config"
	"lms/config/auth"
	"lms/config/helper"
	"lms/structs"
)

type Peserta struct{ structs.Peserta }

func (u *Peserta) Validasi() error {
	if u.Username == "" {
		return errors.New("Username masih kosong")
	}
	if u.Password == "" {
		return errors.New("Password masih kosong")
	}
	return nil
}

func (u *Peserta) Persiapan(action string) {
	u.NRP = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	if strings.ToLower(action) == "seed" || strings.ToLower(action) == "tambah" || strings.ToLower(action) == "update" {
		u.Password = helper.Hash(u.Password)
	}
}

func (u *Peserta) ProsesLogin() (map[string]string, error) {
	var err error
	db := config.ConPg()
	defer db.Close()

	var peserta Peserta
	sqlStatement := `SELECT id,nrp,username,password FROM peserta where username=$1`
	row := db.QueryRow(sqlStatement, u.Username)
	err = row.Scan(&peserta.ID, &peserta.NRP, &peserta.Username, &peserta.Password)

	if err != nil {
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	match, err := helper.VerifikasiPassword(peserta.Password, u.Password)
	if match == false {
		return nil, err
	}
	//log.Println(match)
	return auth.BuatToken(peserta.ID)
}
