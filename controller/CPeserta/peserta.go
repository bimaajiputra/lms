package CPeserta

import (
	"fmt"
	"lms/config/helper"
	"lms/models/MPeserta"
	"lms/responses"
	"lms/responses/formaterror"
	"net/http"

	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func ViewPeserta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := r.URL.Query()
	var page int
	if query.Get("page") == "" {
		page = 1
	} else {
		page = int(helper.StringkeInt(query.Get("page")))
	}
	pesertas, err := MPeserta.ViewPeserta(page)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, responses.Sukses(pesertas))
}

func CreatePeserta(w http.ResponseWriter, r *http.Request) {
	peserta := MPeserta.Peserta{}
	peserta.NRP = uint32(helper.StringkeInt(r.FormValue("nrp")))
	peserta.Username = r.FormValue("username")
	peserta.Password = r.FormValue("password")
	peserta.LastIP = r.FormValue("last_ip")
	r.ParseForm()
	insertRows := MPeserta.CreatePeserta(peserta)
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, insertRows))
	responses.JSON(w, http.StatusCreated, responses.Sukses(insertRows))
}

func UpdatePeserta(w http.ResponseWriter, r *http.Request) {
	peserta := MPeserta.Peserta{}
	peserta.NRP = uint32(helper.StringkeInt(r.FormValue("nrp")))
	peserta.Username = r.FormValue("username")
	peserta.Password = r.FormValue("password")
	peserta.LastIP = r.FormValue("last_ip")
	r.ParseForm()
	updatedRows := MPeserta.UpdatePeserta(peserta)
	msg := fmt.Sprintf("Buku telah berhasil diupdate. Jumlah yang diupdate %v rows/record", updatedRows)
	responses.JSON(w, http.StatusOK, responses.Sukses(msg))
}

func DeletePeserta(w http.ResponseWriter, r *http.Request) {
	peserta := MPeserta.Peserta{}
	peserta.NRP = uint32(helper.StringkeInt(r.FormValue("nrp")))
	r.ParseForm()
	deletedRows := MPeserta.DeletePeserta(peserta)
	msg := fmt.Sprintf("buku sukses di hapus. Total data yang dihapus %v", deletedRows)
	responses.JSON(w, http.StatusOK, responses.Sukses(msg))
}

/* func uploadGambar(w http.ResponseWriter, r *http.Request) string {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("gambar_buku")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	f, err := os.OpenFile("assets/buku/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return ""
	}
	defer f.Close()
	io.Copy(f, file)

	fmt.Println(w, "Successfully Uploaded File\n")

	return handler.Filename
} */
