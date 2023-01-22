package entity

import "time"

type Kelas struct {
	Id         int
	NamaKelas  string
	Hari       string
	JamMulai   string
	JamSelesai string
	RuangKelas string
	Matkul     Matkul
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
