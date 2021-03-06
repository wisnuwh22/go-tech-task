[![Go Report Card](https://goreportcard.com/badge/github.com/wisnuwh22/go-tech-task)](https://goreportcard.com/report/github.com/wisnuwh22/go-tech-task)

__Language :__ [English](README.md) | Bahasa Indonesia

# Go Technical Task
API untuk menyarankan resep makan siang

## Manajemen Waktu
Tidak ada batasan waktu untuk mengerjakan *task* ini. Anda bebas mengatur waktu untuk menyelesaikan *requirement* yang kami minta.

## Penilaian
Kriteria penilaian kami akan memperhatikan hal - hal berikut:
- Bagaimana struktur aplikasinya. 
- *Code quality (Clean code)*.
- Kualitas dari *test* (*Unit test*).
- Pengertian pada masalah.
- Penggunaan `git`.
- Implementasi dan eksekusi akhir.
- *Commits*, ini akan membantu kami untuk mengerti, bagaimana alur kerja dan keputusan anda selama mengerjakan *task* ini.

## User Story
Sebagai *User*, saya ingin melakukan *request* ke *API* yang akan menentukan dari sekumpulan resep, apa yang dapat saya persiapkan untuk makan siang hari ini, berdasarkan bahan - bahan di kulkas saya. Sehingga saya dapat memutuskan apa yang akan saya makan.

__Kriteria Utama__
- Ketika saya melalukan request ke `/lunch` endpoint, saya harus mendapatkan *response* resep dalam bentuk `JSON`, yang dapat saya persiapkan berdasarkan bahan - bahan yang ada di kulkas saya.
- Ketika ada bahan saya, yang sudah melewati tanggal `use-by`, saya harus tidak mendapatkan resep yang mengandung bahan tersebut.
- Ketika ada bahan saya, yang sudah melewati tanggal `best-before`, tetapi masih belum melewatin tanggal `use-by`, resep yang mengandung bahan yang tidak segar, harus terletak pada bagian bawah dari *response*.

__Kriteria Tambahan__
- Aplikasi HARUS memiliki unit / integration tests (contohnya `testing` package).
- Aplikasi HARUS diselesaikan dengan pendekatan `OOP`.
- Aplikasi HARUS mendapatkan nilai A+ dari [Go Report Card](https://goreportcard.com/) website.
- Semua dependencies HARUS diinstal melalui `Go Modules` (tidak perlu untuk commit dependencies, cukup file `go.mod` saja).
- Gunakan versi `Go 1.11` ke atas.
- Semua instruksi untuk instalasi, cara build, testing dan menjalankan HARUS tersedia pada file `README.md` yang berada di folder utama aplikasi.

## Framework
Kami merekemondasikan anda untuk menggunakan *default package* dari Go dalam membuat aplikasi ini.

## Application Data
Untuk tujuan *task* ini, Aplikasi harus dengan mudah membaca data dari 2 *JSON file* yang kami sediakan. Konten untuk *JSON file* ini dapat anda lihat [disini](ingredients/data.json) dan [disini](recipes/data.json).
 
## Submission
Aplikasi harus di *commit* ke __public repository__ di `GitHub` or `BitBucket` (`<lastname>-<firstname>-techtask-go`) dan mohon informasikan *link repository* anda kepada kami.

## Bonus
Konfigurasikan sebuah *environment* `Docker`, sehingga dapat melakukan test dan menjalankan aplikasi dengan cepat. Aplikasi harus terinstall dalam satu perintah.
