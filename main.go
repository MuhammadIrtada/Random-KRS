package main

import (
	"fmt"
	"random-krs/entity"
	"random-krs/service"
)

func main() {
	// Membuat Mata Kuliah
	matkulService := service.NewMatkulService([]entity.Matkul{})
	matkulService.Insert("CIT61006", "Pengembangan Aplikasi Web", 4, 2020, nil)
	matkulService.Insert("CIT61008", "Analisis dan Desain Sistem Informasi", 4, 2020, nil)
	matkulService.Insert("CSD60001", "Pemrograman Basis Data", 3, 2020, nil)
	matkulService.Insert("CIT61007", "Algoritma dan Struktur Data", 3, 2020, nil)
	matkulService.Insert("CIT61009", "Jaringan Komputer Dasar", 4, 2020, nil)
	matkulService.Insert("UBU60003", "Kewirausahaan", 2, 2020, nil)
	matkulService.Insert("COM60052", "Etika Profesi", 2, 2020, nil)

	// Membuat Kelas
	kelasService := service.NewKelasService([]entity.Kelas{}, &matkulService)
	// /*

	// PAW
	kelasService.Insert("A", "Senin", "10.30", "12.09", "G1.3", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("B", "Senin", "13.00", "14.39", "G1.5", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("C", "Senin", "14.45", "16.24", "G1.5", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("D", "Senin", "14.45", "16.24", "F4.14", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("E", "Senin", "16.30", "18.09", "F4.14", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("E", "Selasa", "13.00", "14.39", "F4.3", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("B", "Rabu", "16.30", "18.09", "F4.10", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("C", "Rabu", "13.00", "14.39", "F4.3", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("C", "Kamis", "14.45", "15.34", "F3.8", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("D", "Kamis", "15.35", "16.24", "F4.12", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("E", "Kamis", "11.20", "12.09", "F4.2", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("A", "Jumat", "10.30", "11.19", "F3.1", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	kelasService.Insert("B", "Jumat", "17.20", "18.09", "F2.2", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
	// ADSI
	kelasService.Insert("C", "Senin", "08.45", "10.24", "F2.8", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("E", "Senin", "14.45", "16.24", "G1.2", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("A", "Selasa", "07.00", "08.39", "F4.3", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("C", "Selasa", "14.45", "16.24", "G1.6", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("D", "Selasa", "14.45", "16.24", "F3.1", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("B", "Rabu", "13.00", "14.39", "F3.12", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("D", "Rabu", "08.45", "10.24", "G1.2", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("E", "Rabu", "08.45", "10.24", "F4.2", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("A", "Kamis", "07.50", "08.39", "F3.12", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("B", "Kamis", "16.30", "18.09", "G1.3", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("D", "Kamis", "14.45", "15.34", "F4.12", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("E", "Kamis", "13.50", "14.39", "F2.1", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("A", "Jumat", "07.00", "08.39", "G1.6", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("B", "Jumat", "09.35", "10.24", "F3.12", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	kelasService.Insert("C", "Jumat", "17.20", "18.09", "F4.5", matkulService.FindMatkul("nama", "Analisis dan Desain Sistem Informasi")[0])
	// PBD
	kelasService.Insert("B", "Senin", "14.45", "16.24", "F3.7", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	kelasService.Insert("A", "Selasa", "16.30", "18.09", "F2.2", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	kelasService.Insert("C", "Selasa", "13.00", "14.39", "F2.2", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	kelasService.Insert("D", "Selasa", "13.00", "14.39", "F3.2", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	kelasService.Insert("E", "Rabu", "14.45", "16.24", "F4.4", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	kelasService.Insert("D", "Kamis", "13.50", "14.39", "F3.2", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	kelasService.Insert("E", "Kamis", "13.00", "13.49", "F4.5", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	kelasService.Insert("A", "Jumat", "14.45", "15.34", "F3.17", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	kelasService.Insert("B", "Jumat", "07.50", "08.39", "F3.16", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	kelasService.Insert("C", "Jumat", "13.50", "14.39", "F3.11", matkulService.FindMatkul("nama", "Pemrograman Basis Data")[0])
	// ASD
	kelasService.Insert("B", "Senin", "16.30", "18.09", "G1.6", matkulService.FindMatkul("nama", "Algoritma dan Struktur Data")[0])
	kelasService.Insert("C", "Senin", "10.30", "12.09", "G1.5", matkulService.FindMatkul("nama", "Algoritma dan Struktur Data")[0])
	kelasService.Insert("D", "Senin", "08.45", "10.24", "G1.6", matkulService.FindMatkul("nama", "Algoritma dan Struktur Data")[0])
	kelasService.Insert("A", "Rabu", "16.30", "18.09", "G1.4", matkulService.FindMatkul("nama", "Algoritma dan Struktur Data")[0])
	kelasService.Insert("C", "Rabu", "10.30", "12.09", "G1.6", matkulService.FindMatkul("nama", "Algoritma dan Struktur Data")[0])
	kelasService.Insert("B", "Kamis", "13.00", "14.39", "G1.3", matkulService.FindMatkul("nama", "Algoritma dan Struktur Data")[0])
	kelasService.Insert("E", "Kamis", "08.45", "10.24", "G1.3", matkulService.FindMatkul("nama", "Algoritma dan Struktur Data")[0])
	kelasService.Insert("A", "Jumat", "08.45", "10.24", "G1.2", matkulService.FindMatkul("nama", "Algoritma dan Struktur Data")[0])
	kelasService.Insert("D", "Jumat", "07.00", "08.39", "G1.3", matkulService.FindMatkul("nama", "Algoritma dan Struktur Data")[0])
	// Jarkom
	kelasService.Insert("D", "Senin", "13.00", "14.39", "F4.2", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	kelasService.Insert("A", "Selasa", "13.00", "14.39", "F3.8", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	kelasService.Insert("E", "Selasa", "14.45", "16.24", "G1.3", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	kelasService.Insert("A", "Rabu", "07.00", "08.39", "G1.4", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	kelasService.Insert("C", "Rabu", "07.00", "08.39", "F3.7", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	kelasService.Insert("D", "Rabu", "13.00", "14.39", "G1.5", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	kelasService.Insert("E", "Rabu", "10.30", "12.09", "F4.2", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	kelasService.Insert("B", "Kamis", "08.45", "10.24", "F3.5", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	kelasService.Insert("C", "Kamis", "07.00", "08.39", "G1.4", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	kelasService.Insert("B", "Jumat", "14.45", "16.24", "G1.3", matkulService.FindMatkul("nama", "Jaringan Komputer Dasar")[0])
	// Kwu
	kelasService.Insert("A", "Selasa", "16.30", "18.09", "F3.17", matkulService.FindMatkul("nama", "Kewirausahaan")[0])
	kelasService.Insert("B", "Selasa", "13.00", "14.39", "F4.14", matkulService.FindMatkul("nama", "Kewirausahaan")[0])
	kelasService.Insert("C", "Rabu", "13.00", "14.39", "F3.16", matkulService.FindMatkul("nama", "Kewirausahaan")[0])
	// Etprof
	kelasService.Insert("C", "Senin", "10.30", "12.09", "F3.7", matkulService.FindMatkul("nama", "Etika Profesi")[0])
	kelasService.Insert("B", "Rabu", "10.30", "12.09", "F4.3", matkulService.FindMatkul("nama", "Etika Profesi")[0])
	kelasService.Insert("A", "Kamis", "16.30", "18.09", "F4.5", matkulService.FindMatkul("nama", "Etika Profesi")[0])

	// **/

	// Menampilkan Semua Matkul
	// matkulService.Show(matkulService.List())
	// fmt.Printf("\n\n\n")

	// // Menampilkan Semua Kelas
	// kelasService.Show(kelasService.List())

	// Membuat Plan
	planService := service.NewPlanService([]entity.Plan{}, matkulService.List(), matkulService, kelasService)
	hasilRandom := planService.Random(matkulService.List(), [][]entity.Kelas{})

	// Menampilkan semua kelas pada jadwal
	// planService.Show(planService.GetKelasFromMatkul())

	// Menampilkan semua jadwal pada random
	for i, v := range hasilRandom {
		fmt.Println("==========", i+1, "==========")
		planService.Show(v)
	}
	
	// Menampilkan semua jadwal pada random format nama kelas
	// for i := 0; i < len(hasilRandom); i++ {
	// 	fmt.Print(i+1, " || ")
	// 	for _, v := range hasilRandom[i] {
	// 		fmt.Print(matkulService.Singkatan(v.Matkul), " ", v.NamaKelas, " | ")
	//	}
	// 	fmt.Println()
	// }
}
