package main

import (
	"fmt"
	"random-krs/entity"
	"random-krs/service"
)

func main() {
	// Membuat Mata Kuliah
	matkulService := service.NewMatkulService([]entity.Matkul{})
	matkulService.Insert("CIT62012", "Pengembangan Aplikasi Mobile", 4, 2020)
	matkulService.Insert("CIT62013", "Tata Kelola Teknologi Informasi", 3, 2020)
	matkulService.Insert("CIT62014", "Implementasi dan Evaluasi Sistem Informasi", 3, 2020)
	matkulService.Insert("CIT62015", "Administrasi Sistem", 3, 2020)
	matkulService.Insert("CIT62016", "Arsitektur dan Protokol Internet", 3, 2020)
	matkulService.Insert("CIT62017", "Teknologi Integrasi Sistem", 2, 2020)
	matkulService.Insert("CIT62018", "Keamanan Informasi", 3, 2020)
	matkulService.Insert("CIT60031", "Manajemen Proyek Teknologi Informasi", 3, 2020)
	// matkulService.Insert("CIT60034", "Sistem Informasi Geografi", 3, 2020)

	// Membuat Kelas
	kelasService := service.NewKelasService([]entity.Kelas{}, &matkulService)
	// kelasService.Insert("D", "Senin", "07.00", "09.29", "F2.9", matkulService.FindMatkul("nama", "Tata Kelola Teknologi Informasi")[0])
	// kelasService.Insert("A", "Senin", "08.40", "10.19", "G1.6", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	// kelasService.Insert("D", "Senin", "09.30", "11.59", "F3.5", matkulService.FindMatkul("nama", "Keamanan Informasi")[0])
	// kelasService.Insert("B", "Selasa", "08.40", "10.19", "G1.6", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	// kelasService.Insert("C", "Senin", "07.00", "08.39", "F4.12", matkulService.FindMatkul("nama", "Teknologi Integrasi Sistem")[0])
	// /*

	// PAM
	kelasService.Insert("A", "Senin", "09.30", "11.59", "F3.9", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	kelasService.Insert("B", "Selasa", "15.30", "18.00", "F3.13", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	kelasService.Insert("C", "Selasa", "14.30", "17.09", "F4.4", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	kelasService.Insert("D", "Selasa", "08.40", "10.19", "G1.2", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	kelasService.Insert("B", "Rabu", "16.20", "18.00", "G1.3", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	kelasService.Insert("D", "Rabu", "12.50", "15.19", "F3.5", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	kelasService.Insert("A", "Kamis", "08.40", "10.19", "G1.6", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	kelasService.Insert("C", "Jumat", "07.00", "08.39", "G1.2", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Mobile")[0])
	// TKTI
	kelasService.Insert("B", "Senin", "12.50", "15.19", "F3.13", matkulService.FindMatkul("nama", "Tata Kelola Teknologi Informasi")[0])
	kelasService.Insert("D", "Senin", "07.00", "09.29", "F2.9", matkulService.FindMatkul("nama", "Tata Kelola Teknologi Informasi")[0])
	kelasService.Insert("A", "Selasa", "12.50", "15.19", "F3.2", matkulService.FindMatkul("nama", "Tata Kelola Teknologi Informasi")[0])
	kelasService.Insert("C", "Kamis", "12.50", "15.19", "F4.6", matkulService.FindMatkul("nama", "Tata Kelola Teknologi Informasi")[0])
	// IESI
	kelasService.Insert("C", "Senin", "12.50", "15.19", "F2.1", matkulService.FindMatkul("nama", "Implementasi dan Evaluasi Sistem Informasi")[0])
	kelasService.Insert("B", "Selasa", "12.50", "15.19", "F2.8", matkulService.FindMatkul("nama", "Implementasi dan Evaluasi Sistem Informasi")[0])
	kelasService.Insert("A", "Rabu", "12.50", "15.19", "F3.1", matkulService.FindMatkul("nama", "Implementasi dan Evaluasi Sistem Informasi")[0])
	kelasService.Insert("D", "Kamis", "09.30", "11.59", "F3.6", matkulService.FindMatkul("nama", "Implementasi dan Evaluasi Sistem Informasi")[0])
	// ADSI
	kelasService.Insert("C", "Selasa", "07.00", "08.39", "G1.5", matkulService.FindMatkul("nama", "Administrasi Sistem")[0])
	kelasService.Insert("D", "Selasa", "16.20", "18.00", "G1.6", matkulService.FindMatkul("nama", "Administrasi Sistem")[0])
	kelasService.Insert("A", "Rabu", "10.20", "11.59", "G1.2", matkulService.FindMatkul("nama", "Administrasi Sistem")[0])
	kelasService.Insert("B", "Rabu", "07.00", "08.39", "F4.3", matkulService.FindMatkul("nama", "Administrasi Sistem")[0])
	kelasService.Insert("C", "Rabu", "07.00", "08.39", "F3.1", matkulService.FindMatkul("nama", "Administrasi Sistem")[0])
	kelasService.Insert("A", "Kamis", "14.30", "16.19", "F3.6", matkulService.FindMatkul("nama", "Administrasi Sistem")[0])
	kelasService.Insert("D", "Kamis", "07.00", "08.39", "F4.6", matkulService.FindMatkul("nama", "Administrasi Sistem")[0])
	kelasService.Insert("B", "Jumat", "07.00", "08.39", "G1.6", matkulService.FindMatkul("nama", "Administrasi Sistem")[0])
	// API
	kelasService.Insert("D", "Senin", "12.50", "14.29", "G1.5", matkulService.FindMatkul("nama", "Arsitektur dan Protokol Internet")[0])
	kelasService.Insert("C", "Selasa", "10.20", "11.59", "F2.5", matkulService.FindMatkul("nama", "Arsitektur dan Protokol Internet")[0])
	kelasService.Insert("D", "Selasa", "12.50", "14.29", "F2.6", matkulService.FindMatkul("nama", "Arsitektur dan Protokol Internet")[0])
	kelasService.Insert("B", "Rabu", "10.20", "11.59", "F3.9", matkulService.FindMatkul("nama", "Arsitektur dan Protokol Internet")[0])
	kelasService.Insert("A", "Kamis", "12.50", "14.29", "F3.15", matkulService.FindMatkul("nama", "Arsitektur dan Protokol Internet")[0])
	kelasService.Insert("B", "Kamis", "07.00", "08.39", "G1.6", matkulService.FindMatkul("nama", "Arsitektur dan Protokol Internet")[0])
	kelasService.Insert("C", "Kamis", "08.40", "10.19", "G1.5", matkulService.FindMatkul("nama", "Arsitektur dan Protokol Internet")[0])
	kelasService.Insert("A", "Jumat", "08.40", "10.19", "G1.4", matkulService.FindMatkul("nama", "Arsitektur dan Protokol Internet")[0])
	// TIS
	kelasService.Insert("C", "Senin", "07.00", "08.39", "F4.12", matkulService.FindMatkul("nama", "Teknologi Integrasi Sistem")[0])
	kelasService.Insert("D", "Selasa", "07.00", "08.39", "F3.3", matkulService.FindMatkul("nama", "Teknologi Integrasi Sistem")[0])
	kelasService.Insert("A", "Kamis", "10.20", "11.59", "F4.11", matkulService.FindMatkul("nama", "Teknologi Integrasi Sistem")[0])
	kelasService.Insert("B", "Kamis", "08.40", "10.19", "F3.13", matkulService.FindMatkul("nama", "Teknologi Integrasi Sistem")[0])
	// KI
	kelasService.Insert("A", "Senin", "12.50", "15.19", "F4.14", matkulService.FindMatkul("nama", "Keamanan Informasi")[0])
	kelasService.Insert("B", "Rabu", "12.50", "15.19", "F4.6", matkulService.FindMatkul("nama", "Keamanan Informasi")[0])
	kelasService.Insert("C", "Rabu", "12.50", "15.19", "F3.14", matkulService.FindMatkul("nama", "Keamanan Informasi")[0])
	kelasService.Insert("D", "Rabu", "09.30", "11.59", "F3.5", matkulService.FindMatkul("nama", "Keamanan Informasi")[0])
	// MPTI
	kelasService.Insert("A", "Selasa", "09.30", "11.59", "F2.6", matkulService.FindMatkul("nama", "Manajemen Proyek Teknologi Informasi")[0])
	kelasService.Insert("B", "Selasa", "12.50", "15.19", "F4.6", matkulService.FindMatkul("nama", "Manajemen Proyek Teknologi Informasi")[0])
	// SIG
	// kelasService.Insert("A", "Kamis", "09.30", "11.59", "F3.16", matkulService.FindMatkul("nama", "Sistem Informasi Geografi")[0])

	// **/

	// Menampilkan Semua Matkul
	// matkulService.Show(matkulService.List())
	// fmt.Printf("\n\n\n")

	// // Menampilkan Semua Kelas
	// kelasService.Show(kelasService.List())

	// Membuat Plan
	jam := [][]string{
		{"07.00", "08.39"},
		{"07.00", "09.29"},
		{"08.40", "10.19"},
		{"09.30", "11.59"},
		{"10.20", "11.59"},
		{"12.50", "14.29"},
		{"12.50", "15.19"},
		{"14.30", "16.19"},
		{"14.30", "17.09"},
		{"15.30", "18.00"},
		{"16.20", "18.00"},
	}
	planService := service.NewPlanService([]entity.Plan{}, matkulService.List(), jam, matkulService, kelasService)
	hasilRandom := planService.Random(matkulService.List(), [][]entity.Kelas{})

	// Menampilkan semua kelas pada jadwal
	// planService.Show(planService.GetKelasFromMatkul())

	// Menampilkan semua jadwal pada random format nama kelas
	// ShowNamaKelas(hasilRandom, matkulService, -1)

	// Pembuatan Filter Waktu
	// Rubah 0 menjadi 1 untuk menghindari jam dan hari tersebut
	filterLibur := [][]int{
		//   {"07.00", "08.39"}, -> 1
		//   {"07.00", "09.29"}, -> 2
		//   {"08.40", "10.19"}, -> 3
		//   {"09.30", "11.59"}, -> 4
		//   {"10.20", "11.59"}, -> 5
		//   {"12.50", "14.29"}, -> 6
		//   {"12.50", "15.19"}, -> 7
		//   {"14.30", "16.19"}, -> 8
		//   {"14.30", "17.09"}, -> 9
		//   {"15.30", "18.00"}, -> 10
		//   {"16.20", "18.00"}, -> 11
	//	 1, 2, 3, 4, 5, 6, 7, 8, 9,10,11
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // Senin
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // Selasa
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // Rabu
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // Kamis
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // Jumat
	}

	filtered := planService.Filter(filterLibur, hasilRandom)

	// Pembuatan Filter Berdasarkan Kelas yang dihindari
	// Ex :
	// filterByKelas := []string{
	// 	"ADDSI A", -> Analisis Dan Desain Sistem Informasi Kelas A
	//  "PAW C", -> Pengembangan Aplikasi Web Kelas C
	// }
	filterByKelas := []string{
		// "",
	}
	filteredByKelas := planService.FilterByKelas(filterByKelas, filtered)

	// Menampilkan semua jadwal yang sudah di filter dari random
	Show(filteredByKelas, planService, matkulService)
}

func Show(jadwal [][]entity.Kelas, planService service.PlanService, matkulService service.MatkulService) {
	if len(jadwal) == 0 {
		fmt.Println("======================")
		fmt.Println("JADWAL TIDAK DITEMUKAN")
		fmt.Println("======================")
		return
	}
	for i, v := range jadwal {
		fmt.Println("==========", i+1, "==========")
		planService.Show(v)
		ShowNamaKelas(jadwal, matkulService, i)
		fmt.Println("\n")
	}
}

func ShowNamaKelas(hasilRandom [][]entity.Kelas, matkulService service.MatkulService, i int) {
	showPerJadwal := func(hasilRandom []entity.Kelas) {
		for j, v := range hasilRandom {
			if j < len(hasilRandom)-1 {
				if hasilRandom[j].Matkul.Nama == hasilRandom[j+1].Matkul.Nama {
					continue
				}
			}
			fmt.Print(matkulService.Singkatan(v.Matkul), " ", v.NamaKelas, " | ")
		}
	}

	if i >= 0 {
		showPerJadwal(hasilRandom[i])
		return
	}

	for j := 0; j < len(hasilRandom); j++ {
		fmt.Print(j+1, " || ")
		showPerJadwal(hasilRandom[j])
		fmt.Println()
	}
}
