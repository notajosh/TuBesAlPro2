# 🎓 SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa

SiPresensi adalah aplikasi berbasis antarmuka baris perintah (*Command Line Interface* / CLI) yang dirancang untuk mengelola, mencatat, dan memantau kehadiran mahasiswa secara efisien. Aplikasi ini dibangun menggunakan bahasa pemrograman **Go (Golang)** sebagai pemenuhan Tugas Besar mata kuliah Algoritma dan Pemrograman, program studi S1 Informatika, Fakultas Informatika, Telkom University.

Aplikasi ini mengimplementasikan konsep struktur data statis (Array of Structs) beserta algoritma pencarian (*Searching*) dan pengurutan (*Sorting*) secara mendalam tanpa mengandalkan *package* bawaan Golang untuk proses logikanya.

---

## ✨ Fitur Utama

Aplikasi SiPresensi dilengkapi dengan serangkaian fitur untuk mempermudah operasional pencatatan akademik:

1. **Kelola Data Mahasiswa (CRUD)**
   * Tambah, lihat, perbarui, dan hapus data mahasiswa.
   * Validasi integritas data (misal: NIM wajib 12 digit, input tidak boleh kosong).

2. **Kelola Jadwal Kuliah (CRUD)**
   * Pencatatan jadwal mata kuliah, nama dosen, kelas, hari, dan waktu pelaksanaan.

3. **Manajemen Presensi Mahasiswa**
   * Pencatatan log kehadiran (Hadir/Izin/Sakit/Alpa) per pertemuan.
   * Modifikasi riwayat presensi yang otomatis mensinkronisasi akumulasi kehadiran mahasiswa terkait.

4. **Algoritma Pencarian (Searching)**
   * **Sequential Search:** Digunakan untuk mencari indeks secara linear (seperti pada proses *update/delete*).
   * **Binary Search:** Diimplementasikan untuk mencari data berdasarkan NIM secara lebih cepat (dengan *auto-sort*).
   * **Modified Binary Search:** Digunakan untuk mencari dan mengelompokkan log riwayat presensi berdasarkan status.

5. **Algoritma Pengurutan (Sorting)**
   * **Insertion Sort & Selection Sort:** Tersedia sebagai opsi untuk mengurutkan data mahasiswa.
   * **Tuple Assignment:** Menerapkan pertukaran nilai secara langsung (*idiomatic Go*) tanpa variabel pembantu (`temp`).
   * Pengurutan dinamis berdasarkan Nama (A-Z / Z-A) atau Metrik Presensi (Total Hadir, Izin, Sakit, atau Alpa).

6. **Dashboard & Statistik Akademik**
   * *Progress Bar* visual untuk memantau kapasitas penyimpanan Array (Jadwal & Mahasiswa).
   * Kalkulasi progres pertemuan (contoh: `1 / 16 Pertemuan Tercatat`).
   * Menampilkan data akumulasi mahasiswa beserta peringatan untuk mahasiswa dengan tingkat "Alpa" tertinggi.

---

## 📂 Struktur Repositori

* `TuBesSiPresensi.go` : Berkas utama yang berisi seluruh *logic* antarmuka, CRUD, *Searching*, *Sorting*, dan algoritma utama.
* `dummy.go` : Berkas injeksi data awal (*dummy data*) yang terpisah untuk memudahkan proses demo aplikasi tanpa harus menginput data satu per satu.

---

## 🚀 Cara Instalasi dan Menjalankan Program

### Prasyarat
Pastikan komputer kamu sudah terinstal **Go** (Golang). Jika belum, silakan unduh dan instal melalui [situs resmi Go](https://go.dev/).

### Langkah-langkah Eksekusi
1. Lakukan *clone* repositori ini ke dalam direktori lokal kamu:
   ```bash
   git clone https://github.com/notajosh/TuBesAlPro2.git
   ```
2. Pindah ke dalam direktori repositori:
   ```
   cd TuBesAlPro2
   ```
3. Kemudian jalankan semua program Go nya:
   ```
   go run .
   ```

## 🧑‍💻 Pengembang
 * Joshua Prasetyo Sadewo Paundu (103012530050)
 * Muh. Fachri Haikal (103012500379)
