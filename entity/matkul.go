package entity

import "time"

type Matkul struct {
	Kode           string
	Nama           string
	Sks            int
	TahunKurikulum int
	Kelas          []Kelas
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
