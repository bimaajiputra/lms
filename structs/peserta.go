package structs

import (
	"time"
)

type Peserta struct {
	NRP       uint32    `json:"nrp"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LastIP    string    `json:"last_ip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uint32    `json:"id"`
}
