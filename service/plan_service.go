package service

import (
	"fmt"
	"random-krs/entity"
	"strings"
)

type PlanService interface {
	Insert()
	List()
	FindPlan()
	Update()
	Delete()
	Show(kelass []entity.Kelas)
	Random(list []entity.Kelas) [][]entity.Kelas
	GetKelasFromMatkul() []entity.Kelas
}

type planService struct {
	collection    []entity.Plan
	matkul        []entity.Matkul
	hari          []string
	jam           [][]string
	matkulService MatkulService
}

func NewPlanService(collection []entity.Plan, matkul []entity.Matkul, matkulService MatkulService) PlanService {
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
				
				if len(kelasJadwals) > 0 && i < len(kelasJadwals){
					cetak = kelasJadwals[i]
				}
				fmt.Printf(" %-8s|", cetak)
			}
		}
		fmt.Println()
		fmt.Println(strings.Repeat("-", 66))
	}
}
func (service *planService) Random(list []entity.Kelas) [][]entity.Kelas{
	listRandom := [][]entity.Kelas{}
	service.Random(list[1:])

	return listRandom
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
