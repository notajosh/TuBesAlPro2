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
func searchDataByPresent(mhs *tabMhs, x int, log *tabLog, y int, sch *tabSchedule, z int) {
	var findStatus string
	var i, j, k int
	var match, foundOnLog bool

	for {
		fmt.Println("Masukkan status kehadiran mahasiswa yang ingin dicari, misal Hadir/Sakit/Izin/Alpa (atau ketik 'keluar' untuk batalkan pencarian) : ")
		fmt.Print(">> ")
		fmt.Scan(&findStatus)
		if findStatus == "keluar" {
			fmt.Println("--------------------------------------")
			fmt.Println("Pencarian data mahasiswa dibatalkan. Kembali ke menu utama...")
			fmt.Println("--------------------------------------")
			return
		}
		match = false
		for i = 0; i < x; i++ {
			foundOnLog = false
			for j = 0; j < y && !foundOnLog; j++ {
				if mhs[i].stdID == log[j].stdID && log[j].status == findStatus {
					match = true
					foundOnLog = true
					fmt.Printf("[+] Data mahasiswa dengan status kehadiran %s ditemukan.\n", findStatus)
					fmt.Printf("Nama : %s\nNIM : %s\nJurusan : %s\nAngkatan : %s\nIndeks Data : %d\n", mhs[i].name, mhs[i].stdID, mhs[i].major, mhs[i].batch, i+1)
				}
			}
		}
		if !match {
			fmt.Printf("Data mahasiswa dengan status kehadiran %s tidak ditemukan. Silakan coba kembali.\n", findStatus)
			fmt.Println("--------------------------------------")
		}
	}
}

// Subprogram untuk mencari data mahasiswa berdasarkan NIM menggunakan algoritma binary search
func searchDataByStdID(mhs *tabMhs, n int) {
	var stdID string
	var left, mid, right, found int
	for {
		left = 0
		right = n - 1
		found = -1
		fmt.Println("Masukkan NIM mahasiswa yang ingin dicari (atau ketik 'keluar' untuk batalkan pencarian) : ")
		fmt.Print(">> ")
		fmt.Scan(&stdID)
		if stdID == "keluar" {
			fmt.Println("--------------------------------------")
			fmt.Println("Pencarian data mahasiswa dibatalkan. Kembali ke menu utama...")
			fmt.Println("--------------------------------------")
			return
		}
		for left <= right && found == -1 {
			mid = (left + right) / 2
			if mhs[mid].stdID == stdID {
				found = mid
			} else if mhs[mid].stdID < stdID {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if found == -1 {
			fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan. Silakan coba kembali.\n", stdID)
			fmt.Println("--------------------------------------")
		} else {
			fmt.Printf("[+] Data mahasiswa dengan NIM %s ditemukan.\n", stdID)
			fmt.Printf("Nama : %s\nNIM : %s\nJurusan : %s\nAngkatan : %s\nIndeks Data : %d", mhs[found].name, mhs[found].stdID, mhs[found].major, mhs[found].batch, found+1)
			fmt.Println("---------------------------------------")
		}
	}
}

// Subprogram untuk mengurutkan data mahasiswa berdasarkan NIM menggunakan algoritma selection sort untuk subprogram searchDataByStdID yang menggunakan algoritma binary search
func sortByStdID(mhs *tabMhs, n int) {
	var j, k, minIdx int
	for j = 0; j < n-1; j++ {
		minIdx = j
		for k = j + 1; k < n; k++ {
			if mhs[k].stdID < mhs[minIdx].stdID {
				minIdx = k
			}
		}
		mhs[j], mhs[minIdx] = mhs[minIdx], mhs[j]
	}
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
	// Inisialisasi variabel yang akan digunakan
	var mhs tabMhs
	var schedule tabSchedule
	var log tabLog
	var x, y, firstOption, secondOption int

	x = 0 // Variabel untuk menghitung jumlah data mahasiswa yang telah dimasukkan
	y = 0 // Variabel untuk menghitung jumlah data log kehadiran mahasiswa yang telah dimasukkan
	welcomeMsg := "Selamat datang di Aplikasi SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa"
	for {
		fmt.Println("||============================================================================================||")
		fmt.Printf("|| %-5s ||\n", welcomeMsg)
		fmt.Println("||============================================================================================||")
		fmt.Println("Layanan Menu Aplikasi")
		fmt.Println("\t1. Kelola Data Mahasiswa")
		fmt.Println("\t2. Tambah Jadwal Kuliah")
		fmt.Println("\t3. Kelola Kehadiran Mahasiswa")
		fmt.Println("\t4. Cari Data Mahasiswa")
		fmt.Println("\t5. Urutkan Data Mahasiswa")
		fmt.Println("\t6. Statistik Kehadiran Mahasiswa")
		fmt.Println("\t7. Keluar")
		fmt.Print("Pilih layanan menu : ")
		fmt.Scan(&firstOption)
		if firstOption == 7 {
			fmt.Println("Terima kasih telah menggunakan Aplikasi SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa.\nSampai jumpa!")
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
				addLogPresent(&mhs, x, &log, &y)
			case 2:
				updateLogPresent(&mhs, &log)
			}
		case 4:
			fmt.Println("1. Cari Data Mahasiswa Berdasarkan Status Kehadiran")
			fmt.Println("2. Cari Data Mahasiswa Berdasarkan NIM")
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
				sortByStdID(&mhs, x)
				searchDataByStdID(&mhs, x)
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
