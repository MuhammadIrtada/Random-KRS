package service

import (
	"fmt"
	"random-krs/entity"
	"random-krs/helper"
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
	Filter(filter [][]int, jadwals [][]entity.Kelas) [][]entity.Kelas
	FilterByKelas(mode string, filters []string, jadwals [][]entity.Kelas) [][]entity.Kelas
}

type planService struct {
	collection    []entity.Plan
	matkul        []entity.Matkul
	hari          []string
	jam           [][]string
	matkulService MatkulService
	kelasService  KelasService
}

func NewPlanService(collection []entity.Plan, matkul []entity.Matkul, jam [][]string,matkulService MatkulService, kelasService KelasService) PlanService {
	hari := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}

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
						if ((kelasJadwal.JamMulai == kelasTambah.JamMulai && kelasJadwal.JamSelesai == kelasTambah.JamSelesai && kelasJadwal.Hari == kelasTambah.Hari) || (kelasJadwal.Hari == kelasTambah.Hari && helper.StringToHours(kelasTambah.JamMulai) >= helper.StringToHours(kelasJadwal.JamMulai) && kelasTambah.JamMulai < kelasJadwal.JamSelesai) || (kelasJadwal.Hari == kelasTambah.Hari && helper.StringToHours(kelasTambah.JamMulai) <= helper.StringToHours(kelasJadwal.JamMulai) && kelasTambah.JamSelesai > kelasJadwal.JamMulai)) {
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

func (service *planService) Filter(filter [][]int, jadwals [][]entity.Kelas) [][]entity.Kelas {
	jadwalsBaru := [][]entity.Kelas{}

	for _, jadwal := range jadwals {
		isLibur := false
		for _, kelas := range jadwal {
			for i, hari := range filter {
				for j, jam := range hari {
					if kelas.JamMulai == service.jam[j][0] && kelas.JamSelesai == service.jam[j][1] && jam == 1 && kelas.Hari == service.hari[i] {
						isLibur = true
					}
				}
			}
		}
		if !isLibur {
			jadwalsBaru = append(jadwalsBaru, jadwal)
		}
	}
	return jadwalsBaru
}

func (service *planService) FilterByKelas(mode string, filters []string, jadwals [][]entity.Kelas) [][]entity.Kelas {
	jadwalsBaru := [][]entity.Kelas{}
	uniqueKelasFromJadwal := [][]string{}

	removeDuplicateValues := func (slice []string) []string {
		keys := make(map[string]bool)
		list := []string{}

		for _, entry := range slice {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}
		return list
	}

	switch mode {
	case "with":
		newFilters := []string{}
		for _, filter := range filters {
			newFilters = append(newFilters, strings.Split(filter, " ")[0])
		}
		newFilters = removeDuplicateValues(newFilters)

		for i, jadwal := range jadwals {
			uniqueKelas := []string{}
			count := 0

			for _, kelas := range jadwal {
				uniqueKelas = append(uniqueKelas, service.matkulService.Singkatan(kelas.Matkul)+" "+kelas.NamaKelas)
			}
			uniqueKelasFromJadwal = append(uniqueKelasFromJadwal, removeDuplicateValues(uniqueKelas))
			
			for _, kelasJadwal := range uniqueKelasFromJadwal[i] {
				for _, filter := range  filters{
					if kelasJadwal == filter {
						count++
					}
				}
			}

			if count == len(newFilters) {
				jadwalsBaru = append(jadwalsBaru, jadwal)
			}
		}
		return jadwalsBaru
	case "without":
		for _, jadwal := range jadwals {
			isSama := false
			for _, kelas := range jadwal {
				for _, filter := range filters {
					if service.matkulService.Singkatan(kelas.Matkul)+" "+kelas.NamaKelas == strings.ToUpper(filter) {
						isSama = true
					}
				}
			}
			if !isSama {
				jadwalsBaru = append(jadwalsBaru, jadwal)
			}
		}
		return jadwalsBaru
	}
	return nil
}
