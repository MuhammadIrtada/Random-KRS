package service

import (
	"fmt"
	"random-krs/entity"
	"strings"
	"time"
)

type KelasService interface {
	Insert(nama, hari, jamMulai, jamSelesai, ruangKelas string, matkul entity.Matkul)
	List() []entity.Kelas
	FindKelas(filter string, data interface{}) []entity.Kelas
	Update()
	Delete()
	Show(kelass []entity.Kelas)
}

type kelasService struct {
	collection    []entity.Kelas
	matkulService MatkulService
}

func NewKelasService(collection []entity.Kelas, matkulService *MatkulService) KelasService {
	return &kelasService{
		collection:    []entity.Kelas{},
		matkulService: *matkulService,
	}
}

func (service *kelasService) Insert(nama, hari, jamMulai, jamSelesai, ruangKelas string, matkul entity.Matkul) {
	newKelas := entity.Kelas{
		Id:         len(service.collection) + 1,
		NamaKelas:  nama,
		Hari:       hari,
		JamMulai:   jamMulai,
		JamSelesai: jamSelesai,
		RuangKelas: ruangKelas,
		Matkul:     matkul,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// kelass := service.FindKelas("kode matkul", matkul.Kode)
	matkulTarget := service.matkulService.FindMatkul("kode", matkul.Kode)
	// lenKelas := len(kelass)

	// if lenKelas == 0 {
	// 	kelass = append(kelass, newKelas)
	// }

	// for i, kelas := range  kelass{
	// 	if kelas.NamaKelas == nama && lenKelas > 0{
	// 		break;
	// 	}

	// 	if i != len(kelass) - 1 {
	// 		continue
	// 	}
		kelasBaru := append(matkulTarget[0].Kelas, newKelas)
		data := entity.Matkul{
			Kelas: kelasBaru,
		}
		service.matkulService.Update(matkul.Kode, data)
	// }
	
	service.collection = append(service.collection, newKelas)
}
func (service *kelasService) List() []entity.Kelas{
	return service.collection
}
func (service *kelasService) FindKelas(filter string, data interface{}) []entity.Kelas {
	kelas := []entity.Kelas{}

	category := func(filter string, kelas entity.Kelas) interface{} {
		switch strings.ToLower(filter) {
		case "kode matkul":
			return kelas.Matkul.Kode
		case "nama matkul":
			return kelas.Matkul.Nama
		case "hari":
			return kelas.Hari
		}
		return nil
	}

	for _, k := range service.collection {
		if category(filter, k) == data {
			kelas = append(kelas, k)
		}
		continue
	}

	return kelas
}
func (service *kelasService) Update() {

}
func (service *kelasService) Delete() {

}
func (service *kelasService) Show(kelass []entity.Kelas) {
	fmt.Println(strings.Repeat("=", 129))
	fmt.Println("No", "||", "Hari", strings.Repeat(" ", 2), "||", "Jam Mulai", "||", "Jam Selesai", "||", "Nama Kelas", "||", "Nama Matkul", strings.Repeat(" ", 39), "||", "Kode", strings.Repeat(" ", 3), "||", "SKS")
	fmt.Println(strings.Repeat("=", 129))
	for i, kelas := range kelass {
		fmt.Printf("%-3d|| %-7s || %-9s ||  %-10s || %5s%5s || %-51s || %6s || %2d\n", i+1, kelas.Hari, kelas.JamMulai, kelas.JamSelesai, kelas.NamaKelas, " ", kelas.Matkul.Nama, kelas.Matkul.Kode, kelas.Matkul.Sks)
		fmt.Println(strings.Repeat("-", 129))
	}
}
