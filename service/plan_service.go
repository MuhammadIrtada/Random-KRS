package service

import (
	"fmt"
	"random-krs/entity"
	"strings"

	"golang.org/x/exp/slices"
)

type PlanService interface {
	Insert()
	List()
	FindPlan()
	Update()
	Delete()
	Show(kelass []entity.Kelas)
	Random(matkul []entity.Matkul, hasil [][]entity.Kelas) [][]entity.Kelas
	GetKelasFromMatkul() []entity.Kelas
}

type planService struct {
	collection    []entity.Plan
	matkul        []entity.Matkul
	hari          []string
	jam           [][]string
	matkulService MatkulService
	kelasService  KelasService
}

func NewPlanService(collection []entity.Plan, matkul []entity.Matkul, matkulService MatkulService, kelasService KelasService) PlanService {
	hari := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
	jam := [][]string{
		{"07.00", "08.39"},
		{"07.50", "08.39"},
		{"08.45", "10.24"},
		{"09.35", "10.24"},
		{"10.30", "11.19"},
		{"10.30", "12.09"},
		{"11.20", "12.09"},
		{"13.00", "13.49"},
		{"13.00", "14.39"},
		{"13.50", "14.39"},
		{"14.45", "15.34"},
		{"14.45", "16.24"},
		{"15.35", "16.24"},
		{"16.30", "18.09"},
		{"17.20", "18.09"},
	}

	return &planService{
		collection:    collection,
		matkul:        matkul,
		hari:          hari,
		jam:           jam,
		matkulService: matkulService,
		kelasService:  kelasService,
	}
}

func (service *planService) Insert() {

}
func (service *planService) List() {

}
func (service *planService) FindPlan() {

}
func (service *planService) Update() {

}
func (service *planService) Delete() {

}
func (service *planService) Show(kelass []entity.Kelas) {
	fmt.Println(strings.Repeat("=", 66))
	fmt.Print(strings.Repeat(" ", 14), "||")
	for _, h := range service.hari {
		fmt.Printf(" %-8s|", h)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 66))
	for _, j := range service.jam {
		max := 0
		fmt.Print(j[0], " - ", j[1], " ||")
		for _, h := range service.hari {
			kelasJadwals := service.JadwalPosition(h, j[0], j[1], kelass)

			cetak := ""
			if len(kelasJadwals) > 0 {
				cetak = kelasJadwals[0]
			}

			if max < len(kelasJadwals) {
				max = len(kelasJadwals)
			}
			fmt.Printf(" %-8s|", cetak)
		}
		for i := 1; i < max; i++ {
			fmt.Println()
			fmt.Print(strings.Repeat(" ", 14), "||")
			for _, h := range service.hari {
				kelasJadwals := service.JadwalPosition(h, j[0], j[1], kelass)
				cetak := ""

				if len(kelasJadwals) > 0 && i < len(kelasJadwals) {
					cetak = kelasJadwals[i]
				}
				fmt.Printf(" %-8s|", cetak)
			}
		}
		fmt.Println()
		fmt.Println(strings.Repeat("-", 66))
	}
}
func (service *planService) Random(matkul []entity.Matkul, hasil [][]entity.Kelas) [][]entity.Kelas {
	for i, namaKelas := range service.matkulService.GetAllNameKelas(matkul[0]) {
		kelasBaru := service.kelasService.FindKelas("nama kelas dari matkul", []string{matkul[0].Kode, namaKelas})
		if len(matkul) == 1 {
			hasil = append(hasil, []entity.Kelas{})

			hasil[len(hasil)-1] = append(hasil[len(hasil)-1], kelasBaru...)
			
			if i == len(service.matkulService.GetAllNameKelas(matkul[0]))-1 {				
				return hasil
			}
		} else {
			lenHasilLama := len(hasil)
			hasil = append(hasil, service.Random(matkul[1:], [][]entity.Kelas{})...)

			for i := lenHasilLama; i < len(hasil); i++ {
				isNabrak := false
				for _, kelasJadwal := range hasil[i] {
					for _, kelasTambah := range kelasBaru {
						if kelasJadwal.JamMulai == kelasTambah.JamMulai && kelasJadwal.JamSelesai == kelasTambah.JamSelesai && kelasJadwal.Hari == kelasTambah.Hari{
							isNabrak = true
						}
					}
				}
				if !isNabrak {
					hasil[i] = append(hasil[i], kelasBaru...)
					continue
				}
				hasil = slices.Delete(hasil, i, i+1)
				i--
			}
			
			
			if i == len(service.matkulService.GetAllNameKelas(matkul[0]))-1 {
				
				return hasil
			}
		}
	}
	return hasil
}

func (service *planService) GetKelasFromMatkul() []entity.Kelas {
	newKelas := []entity.Kelas{}
	for _, matkul := range service.matkul {
		newKelas = append(newKelas, matkul.Kelas...)
	}
	return newKelas
}

func (service *planService) JadwalPosition(hari, jamMulai, jamSelesai string, kelass []entity.Kelas) []string {

	kelasJadwal := []string{}
	for _, kelas := range kelass {
		if kelas.Hari == hari && kelas.JamMulai == jamMulai && kelas.JamSelesai == jamSelesai {
			kelasJadwal = append(kelasJadwal, service.matkulService.Singkatan(kelas.Matkul)+" "+kelas.NamaKelas)
		}
	}
	return kelasJadwal
}
