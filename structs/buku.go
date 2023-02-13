package structs

import "time"

type Buku struct {
	// gorm.Model
	IDBuku          int64  `json:"id_bukuid,string,omitempty"`
	ISBN            string `json:"isbn"`
	IDKategoriJenis int64  `json:"id_kategori_buku"`
	//KategoriJenis   string `json:"kategori_buku"`
	Details struct {
		KategoriJenis string `json:"jenis_buku"`
		Deskripsi     string `json:"deskripsi"`
	} `json:"Details"`
	Judul          string        `json:"judul_buku"id,string,omitempty`
	IDPenulisBuku  int64         `json:"id_penulis_buku"`
	PenulisBuku    Penulis_Buku  `json:"penulis_buku"`
	IDPenerbitBuku int64         `json:"id_penerbit_buku"`
	PenerbitBuku   Penerbit_Buku `json:"penerbit_buku"`
	ThnTerbit      string        `json:"tahun_terbit"`
	StokBuku       int64         `json:"stok_buku"`
	RakBuku        string        `json:"rak_buku"`
	DeskripsiBuku  string        `json:"deskripsi_buku"`
	Gambarbuku     string        `json:"gambar_buku"`
	Kondisibuku    string        `json:"kondisi_buku"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type Detail_buku struct {
	// gorm.Model
	IDDetailBuku int64  `json:"id"`
	IDBuku       int64  ` json:"id_buku"`
	Buku         []Buku `json:"buku"`
	Gambarbuku   string `json:"gambar_buku"`
	Kondisibuku  string `json:"kondisi_buku"`
}
type Jenis_Buku struct {
	IDJenis   int64  `json:"id"`
	JenisBuku string `json:"jenis_buku"`
	Deskripsi string `json:"deskripsi"`
}

type Penulis_Buku struct {
	// gorm.Model
	IDPenulis     int64  `json:"id"`
	PenulisBuku   string `gorm:"size:255;not null;unique" json:"penulis_buku"`
	AlamatPenulis string `gorm:"size:255;null;" json:"alamat"`
	EmailPenulis  string `gorm:"size:255;null;unique" json:"email_penulis"`
	Deskripsi     string `gorm:"type:text;null;" json:"deskripsi"`
}

type Penerbit_Buku struct {
	// gorm.Model
	IDPenerbit     int64  `gorm:"primary_key;auto_increment" json:"id"`
	PenerbitBuku   string `gorm:"size:255;not null;unique" json:"penerbit_buku"`
	AlamatPenerbit string `gorm:"size:255;null" json:"alamat_penerbit"`
	TelpPenerbit   string `gorm:"size:255;null" json:"telp_penerbit"`
	EmailPenerbit  string `gorm:"size:255;null;unique" json:"email_penerbit"`
	Deskripsi      string `gorm:"type:text;null;" json:"deskripsi_penerbit"`
}
