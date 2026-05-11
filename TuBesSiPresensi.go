package main

import "fmt"

// Pembuatan konstanta untuk jumlah mahasiswa dan log kehadiran mahasiswa
const nMhs int = 520
const nLog int = 9999

// Pembuatan struct untuk data mahasiswa
type student struct {
	name, stdID, major, batch string
}

// Pembuatan struct untuk data jadwal kuliah
type schedule struct {
	course, classCode, lecturer, date string
	jam, menit                        int
}

// Pembuatan struct untuk data log kehadiran mahasiswa
type log struct {
	stdID, scheduleID, status string
	session                   int
}

// Pembuatan tipe data untuk array mahasiswa, jadwal kuliah, dan log kehadiran mahasiswa
type tabMhs [nMhs]student
type tabSchedule [nMhs]schedule
type tabLog [nLog]log

// Subprogram untuk menambahkan data mahasiswa
func addMhs(mhs *tabMhs, x *int) {

}

//Subprogram untuk mengubah data mahasiswa
func updateMhs(mhs *tabMhs, x *int) {

}

// Subprogram untuk menghapus data mahasiswa
func deleteMhs(mhs *tabMhs, x *int) {

}

// Subprogram untuk menambahkan data jadwal kuliah
func addSchedule(schedule *tabSchedule) {

}

// Subprogram untuk menambahkan data status kehadiran mahasiswa
func addLogPresent(mhs *tabMhs, log *tabLog) {

}

// Subprogram untuk mengubah data status kehadiran mahasiswa
func updateLogPresent(mhs *tabMhs, log *tabLog) {

}

// Subprogram untuk mencari data mahasiswa berdasarkan status kehadiran menggunakan algoritma sequential search
func searchDataByPresent(mhs *tabMhs, log *tabLog) {

}

// Subprogram untuk mencari data mahasiswa berdasarkan nama menggunakan algoritma binary search
func searchDataByStdID(mhs *tabMhs) {

}

// Subprogram untuk mengurutkan data mahasiswa berdasarkan status kehadiran menggunakan algoritma selection sort
func sortingDataByPresent(mhs *tabMhs, log *tabLog) {

}

// Subprogram untuk mengurutkan data mahasiswa berdasarkan nama menggunakan algoritma insertion sort
func sortingDataByName(mhs *tabMhs) {

}

// Subprogram untuk menampilkan statistik kehadiran mahasiswa
func statisticPresent(mhs *tabMhs, log *tabLog) {

}

func main() {
	var mhs tabMhs
	var schedule tabSchedule
	var log tabLog
	var x, firstOption, secondOption int

	for {
		fmt.Println("Selamat datang di Aplikasi SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa")
		fmt.Println("1. Kelola Data Mahasiswa")
		fmt.Println("2. Tambah Jadwal Kuliah")
		fmt.Println("3. Kelola Kehadiran Mahasiswa")
		fmt.Println("4. Cari Data Mahasiswa")
		fmt.Println("5. Urutkan Data Mahasiswa")
		fmt.Println("6. Statistik Kehadiran Mahasiswa")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih layanan menu : ")
		fmt.Scan(&firstOption)
		if firstOption == 7 {
			fmt.Println("Terima kasih telah menggunakan Aplikasi SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa. \nSampai jumpa!")
			return
		}
		switch firstOption {
		case 1:
			fmt.Println("1. Tambah Data Mahasiswa")
			fmt.Println("2. Ubah Data Mahasiswa")
			fmt.Println("3. Hapus Data Mahasiswa")
			fmt.Println("4. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu : ")
			fmt.Scan(&secondOption)
			if secondOption == 4 {
				continue
			}
			switch secondOption {
			case 1:
				addMhs(&mhs, &x)
			case 2:
				updateMhs(&mhs, &x)
			case 3:
				deleteMhs(&mhs, &x)
			}
		case 2:
			addSchedule(&schedule)
		case 3:
			fmt.Println("1. Tambah Data Kehadiran Mahasiswa")
			fmt.Println("2. Ubah Data Kehadiran Mahasiswa")
			fmt.Println("3. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu : ")
			fmt.Scan(&secondOption)
			if secondOption == 3 {
				continue
			}
			switch secondOption {
			case 1:
				addLogPresent(&mhs, &log)
			case 2:
				updateLogPresent(&mhs, &log)
			}
		case 4:
			fmt.Println("1. Cari Data Mahasiswa Berdasarkan Status Kehadiran")
			fmt.Println("2. Cari Data Mahasiswa Berdasarkan Nama")
			fmt.Println("3. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu : ")
			fmt.Scan(&secondOption)
			if secondOption == 3 {
				continue
			}
			switch secondOption {
			case 1:
				searchDataByPresent(&mhs, &log)
			case 2:
				searchDataByStdID(&mhs)
			}
		case 5:
			fmt.Println("1. Urutkan Data Mahasiswa Berdasarkan Status Kehadiran")
			fmt.Println("2. Urutkan Data Mahasiswa Berdasarkan Nama")
			fmt.Println("3. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu : ")
			fmt.Scan(&secondOption)
			if secondOption == 3 {
				continue
			}
			switch secondOption {
			case 1:
				sortingDataByPresent(&mhs, &log)
			case 2:
				sortingDataByName(&mhs)
			}
		case 6:
			statisticPresent(&mhs, &log)
		}
	}
}
