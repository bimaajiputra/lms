package CHome

import (
	"net/http"

	"lms/responses"
)

func Home(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["status"] = "sukses"
	m["pesan"] = "Ini adalah halaman sementara dashboard, untuk update selanjutnya bakal diisi seperti atribut fitur ini"
	m["fitur"] = "hitung jumlah peminjaman dalam tahun ini, total denda, jumlah data anggota"
	responses.JSON(w, http.StatusOK, m)
}
