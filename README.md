# Random-KRS
To get all possibility KRS schedule from selected class

1. Install golang, Jika belum bisa didownload di https://go.dev/dl/
2. Download Repository tersebut
3. Untuk tambah matkul dengan :
- matkulService.Insert("[Kode Matkul (unique)]", "[Nama Matkul (unique)]", [SKS], [Tahun Kurikulum], nil)
- ex: matkulService.Insert("CIT61006", "Pengembangan Aplikasi Web", 4, 2020, nil)
4. Untuk tambah kelas dengan :
- kelasService.Insert("[Nama Kelas]", "[Hari]", "[Jam Mulai]", "[Jam Selesai]", "[Ruangan]", matkulService.FindMatkul("nama", "[Nama Matkul]")[0])
- ex: kelasService.Insert("A", "Senin", "10.30", "12.09", "G1.3", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
5. Jalankan di cmd/terminal dengan mengetikkan "go run main.go"
6. Pada bagian filter, jika ingin di suatu hari dan jam menjadi libur rubah angka 0 menjadi 1
