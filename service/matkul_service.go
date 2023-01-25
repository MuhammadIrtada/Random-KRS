package service

import (
	"fmt"
	"random-krs/entity"
	"reflect"
	"strings"
	"time"
)

type MatkulService interface {
	Insert(kode, nama string, sks, tahunKurikulum int)
	List() []entity.Matkul
	FindMatkul(filter string, data interface{}) []entity.Matkul
	Update(kode string, data entity.Matkul)
	Delete()
	Show(matkuls []entity.Matkul)
	Singkatan(matkul entity.Matkul) string
	GetAllNameKelas(matkul entity.Matkul) []string
}

type matkulService struct {
	collection []entity.Matkul
}

func NewMatkulService(collection []entity.Matkul) MatkulService {
	return &matkulService{
		collection: collection,
	}
}

func (service *matkulService) Insert(kode, nama string, sks, tahunKurikulum int) {
	newMatkul := entity.Matkul{
		Kode:           kode,
		Nama:           nama,
		Sks:            sks,
		TahunKurikulum: tahunKurikulum,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	service.collection = append(service.collection, newMatkul)
}
func (service *matkulService) List() []entity.Matkul {
	return service.collection
}
func (service *matkulService) FindMatkul(filter string, data interface{}) []entity.Matkul {
	matkuls := []entity.Matkul{}

	category := func(filter string, matkul entity.Matkul) interface{} {
		switch strings.ToLower(filter) {
		case "kode":
			return matkul.Kode
		case "nama":
			return matkul.Nama
		}
		return nil
	}

	for _, matkul := range service.collection {
		if category(filter, matkul) == data {
			matkuls = append(matkuls, matkul)
		}
		continue
	}
	return matkuls
}
func (service *matkulService) Update(kode string, data entity.Matkul) {
	for i := 0; i < len(service.collection); i++ {
		if service.collection[i].Kode != kode {
			continue
		}

		values := reflect.ValueOf(&service.collection[i]).Elem()

		for i := 0; i < values.NumField(); i++ {
			if reflect.ValueOf(data).Field(i).IsZero() {
				continue
			}
			values.Field(i).Set(reflect.ValueOf(data).Field(i))
			break
		}
		break
	}
}
func (service *matkulService) Delete() {

}
func (service *matkulService) Show(matkuls []entity.Matkul) {
	fmt.Println(strings.Repeat("=", 111))
	fmt.Println("No", "||", "Nama Matkul", strings.Repeat(" ", 39), "||", "Kode", strings.Repeat(" ", 3), "||", "SKS", "||", "Tahun Kurikulum", "||", "Jumlah Kelas")
	fmt.Println(strings.Repeat("=", 111))
	for i, matkul := range matkuls {
		fmt.Printf("%-3d|| %-51s || %-5s ||  %-2d || %5s%-10d ||%6d\n", i+1, matkul.Nama, matkul.Kode, matkul.Sks, " ", matkul.TahunKurikulum, len(service.GetAllNameKelas(matkul)))
		fmt.Println(strings.Repeat("-", 111))
	}
}

func (service *matkulService) Singkatan(matkul entity.Matkul) string {
	kata := strings.Split(matkul.Nama, " ")
	var singkatan string

	for _, k := range kata {
		singkatan += strings.ToUpper(k[0:1])
	}

	return singkatan
}

func (service *matkulService) GetAllNameKelas(matkul entity.Matkul) []string {
	uniqueKelas := []string{}

	for _, kelas := range matkul.Kelas {
		lenKelas := len(uniqueKelas)

		if lenKelas == 0 {
			uniqueKelas = append(uniqueKelas, kelas.NamaKelas)
			continue
		}

		for i, k := range uniqueKelas {
			if kelas.NamaKelas == k && lenKelas > 0 {
				break
			}

			if i != len(uniqueKelas)-1 {
				continue
			}
			uniqueKelas = append(uniqueKelas, kelas.NamaKelas)
		}
	}

	return uniqueKelas
}
