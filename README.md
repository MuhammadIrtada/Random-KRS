# Random-KRS
To get all possibility KRS schedule from selected matkul

1. Install golang di atas versi 1.18, Jika belum bisa didownload di https://go.dev/dl/
2. Download Repository tersebut
3. Untuk tambah matkul dengan :
- matkulService.Insert("[Kode Matkul (unique)]", "[Nama Matkul (unique)]", [SKS], [Tahun Kurikulum])
- ex: matkulService.Insert("CIT61006", "Pengembangan Aplikasi Web", 4, 2020)
- ![image](https://user-images.githubusercontent.com/88333614/215084823-d1c1db8b-8d54-42ee-8f27-99e54c4f7693.png)
4. Untuk tambah kelas dengan :
- kelasService.Insert("[Nama Kelas]", "[Hari]", "[Jam Mulai]", "[Jam Selesai]", "[Ruangan]", matkulService.FindMatkul("nama", "[Nama Matkul]")[0])
- ex: kelasService.Insert("A", "Senin", "10.30", "12.09", "G1.3", matkulService.FindMatkul("nama", "Pengembangan Aplikasi Web")[0])
- ![image](https://user-images.githubusercontent.com/88333614/215085256-666d39e8-c47c-4745-a3dd-62ef1dcb8003.png)
5. Uncommand code berikut untuk melihat semua matkul
- ![image](https://user-images.githubusercontent.com/88333614/215086375-c27a64f4-7c07-4abf-82c2-a1e39fb1fd03.png)
- ![image](https://user-images.githubusercontent.com/88333614/215085799-1e9f7632-77f8-41db-bd48-b9d671ebdaa5.png)
6. Uncommand code berikut untuk melihat semua kelas
- ![image](https://user-images.githubusercontent.com/88333614/215086846-d9437495-d90b-47e3-8c20-d8dfa15fed7a.png)
- ![image](https://user-images.githubusercontent.com/88333614/215086613-19f6b992-f005-4cc2-b013-8914d1b986f4.png)
7. Sesuaikan jam dengan jadwal yang digunakan
- ![image](https://user-images.githubusercontent.com/88333614/215087131-8a9a0db5-cbab-44eb-9343-f0273d5e5f10.png)
8. Uncommand code berikut untuk menampilkan semua kelas pada jadwal
- ![image](https://user-images.githubusercontent.com/88333614/215088527-c5dffb1e-8677-4bc1-bf3d-fdb4c6b05633.png)
- ![image](https://user-images.githubusercontent.com/88333614/215088245-ceea0582-f397-405d-a990-394a10e0c49a.png)
9. Pada bagian filter, jika ingin di suatu hari dan jam menjadi libur rubah angka 0 menjadi 1
- Contoh hari jumat libur :
- ![image](https://user-images.githubusercontent.com/88333614/215090337-c0d59c30-6d7c-4149-bb3d-04d2b130770a.png)
10. Jika ingin filter suatu kelas dapat dengan :
- Masukkan singkatan nama matkul dengan huruf depan setiap kata kemudian nama kelas yang dihindari
- ![image](https://user-images.githubusercontent.com/88333614/215095614-bf83cbba-d4d4-4021-bd17-2bb307265069.png)
- Jika ingin menghindari suatu kelas masukkan kata kunci "without"
- Jika ingin memunculkan matkul yang diinginkan masukkan kata kunci "with"
- ![image](https://user-images.githubusercontent.com/88333614/215416486-3f30aa9f-0e38-45d5-bf67-0162b582b5b1.png)
11. Jalankan di cmd/terminal dengan mengetikkan "go run main.go"
