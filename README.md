# Random-KRS
To get all possibility KRS schedule from selected matkul

1. Install golang di atas versi 1.18, Jika belum bisa didownload di https://go.dev/dl/
2. Download Repository tersebut
3. Untuk tambah matkul dengan :
- matkulService.Insert("[Kode Matkul (unique)]", "[Nama Matkul (unique)]", [SKS], [Tahun Kurikulum])
- ex: matkulService.Insert("CIT61006", "Pengembangan Aplikasi Web", 4, 2020)
4. Untuk tambah kelas dengan :
- kelasService.Insert("[Nama Kelas]", "[Hari]", "[Jam Mulai]", "[Jam Selesai]", "[Ruangan]", matkulService.FindMatkul("nama", "[Nama Matkul]")[0])
- ex: kelasService.Insert("A", "Senin", "10.30", "12.09", "G1.3", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
5. Pada bagian filter, jika ingin di suatu hari dan jam menjadi libur rubah angka 0 menjadi 1
6. Jalankan di cmd/terminal dengan mengetikkan "go run main.go"
