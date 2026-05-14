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
	var newMhs student
	
	fmt.Println("=== Tambah Data Mahasiswa ===")
	fmt.Print("Nama Mahasiswa : ")
	fmt.Scan(&newMhs.name)
	fmt.Print("NIM Mahasiswa  : ")
	fmt.Scan(&newMhs.stdID)
	fmt.Print("Jurusan		  : ")
	fmt.Scan(&newMhs.major)
	fmt.Print("Angkatan		  : ")
	fmt.Scan(&newMhs.batch)

	for i := 0; i < *x; i++ {
		if mhs[i].stdID == newMhs.stdID {
			fmt.Printf("NIM %s sudah terdaftar!!! Data tidak ditambahkan.\n", newMhs.stdID)
			return
		}
	}

	pos := *x
	for pos > 0 && mhs[pos-1].stdID > newMhs.stdID {
		mhs[pos] = mhs[pos-1]
		pos--
	}
	mhs[pos] = newMhs
	*x++
	fmt.Printf("[+] Data mahasiswa %s berhasil ditambahkan.\n", newMhs.name)
	fmt.Println("+--------------------------------------+")

	if *x >= nMhs {
		fmt.Println("Data mahasiswa sudah penuh!")
		return
	}
}

//Subprogram untuk mengubah data mahasiswa
func updateMhs(mhs *tabMhs, x *int) {
	var stdID, input string
	var left, mid, right, found int

	fmt.Println("=== Ubah Data Mahasiswa ===")
	fmt.Print("Masukkan NIM mahasiswa yang ingin diubah : ")
	fmt.Scan(&stdID)

	left = 0
	right = *x - 1
	found = -1
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
		fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan.\n", stdID)
		fmt.Println("+--------------------------------------+")
		return
	}

	fmt.Printf("Data ditemukan: %s | %s | %s | %s\n",
		mhs[found].name, mhs[found].stdID, mhs[found].major, mhs[found].batch)
	fmt.Println("Masukkan data baru (tekan Enter dan isi '-' untuk tidak mengubah):")

	fmt.Print("Nama baru   : ")
	fmt.Scan(&input)
	if input != "-" {
		mhs[found].name = input
	}

	fmt.Print("Jurusan baru : ")
	fmt.Scan(&input)
	if input != "-" {
		mhs[found].major = input
	}

	fmt.Print("Angkatan baru : ")
	fmt.Scan(&input)
	if input != "-" {
		mhs[found].batch = input
	}
	fmt.Printf("[+] Data mahasiswa dengan NIM %s berhasil diubah.\n", stdID)
	fmt.Println("+--------------------------------------+")

	if *x == 0 {
		fmt.Println("Belum ada data mahasiswa!")
		return
	}
}

// Subprogram untuk menghapus data mahasiswa
func deleteMhs(mhs *tabMhs, x *int) {
	var stdID, konfirmasi string
	var left, mid, right, found int

	fmt.Println("=== Hapus Data Mahasiswa ===")
	fmt.Print("Masukkan NIM mahasiswa yang ingin anda hapus : ")
	fmt.Scan(&stdID)

	left = 0
	right = *x - 1
	found = -1
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
		fmt.Printf("Data mahasiswa dengan NIM %s tidak dapat ditemukan.\n", stdID)
		fmt.Println("+--------------------------------------+")
		return
	}

	fmt.Printf("Data mahasiswa ditemukan: %s | %s | %s | %s\n",
		mhs[found].name, mhs[found].stdID, mhs[found].major, mhs[found].batch)

	fmt.Print("Apakah Anda yakin ingin menghapus data ini? (ya/tidak) : ")
	fmt.Scan(&konfirmasi)

	if konfirmasi != "ya" {
		fmt.Println("Penghapusan dibatalkan.")
		fmt.Println("+--------------------------------------+")
		return
	}

	for i := found; i < *x-1; i++ {
		mhs[i] = mhs[i+1]
	}
	mhs[*x-1] = student{}
	*x--

	fmt.Printf("[+] Data mahasiswa dengan NIM %s berhasil anda hapus.\n", stdID)
	fmt.Println("+--------------------------------------+")

	if *x == 0 {
		fmt.Println("Data mahasiswa tidak ditemukan.")
		return
	}
}

// Subprogram untuk menambahkan data jadwal kuliah
func addSchedule(schedule *tabSchedule) {

}

// Subprogram untuk menambahkan data status kehadiran mahasiswa
func addLogPresent(mhs *tabMhs, x int, log *tabLog, y int, sch *tabSchedule, z int) {

}

// Subprogram untuk mengubah data status kehadiran mahasiswa
func updateLogPresent(mhs *tabMhs, log *tabLog) {

}

// Subprogram untuk mencari data mahasiswa berdasarkan status kehadiran menggunakan algoritma sequential search
func searchDataByPresent(mhs *tabMhs, x int, log *tabLog, y int, sch *tabSchedule, z int) {
	var findStatus string
	var i, j, k int
	var match, printedHeader, scheduleFound bool

	for {
		fmt.Println("Masukkan status kehadiran mahasiswa yang ingin dicari, misal Hadir/Sakit/Izin/Alpa (atau ketik 'keluar' untuk batalkan pencarian) : ")
		fmt.Print(">> ")
		fmt.Scan(&findStatus)
		if findStatus == "keluar" {
			fmt.Println("+--------------------------------------+")
			fmt.Println("Pencarian data mahasiswa dibatalkan. Kembali ke menu utama...")
			fmt.Println("+--------------------------------------+")
			return
		}
		match = false
		for i = 0; i < x; i++ {
			printedHeader = false
			for j = 0; j < y; j++ {
				if mhs[i].stdID == log[j].stdID && log[j].status == findStatus {
					match = true
					if !printedHeader {
						fmt.Printf("[+] Data mahasiswa dengan status kehadiran %s ditemukan.\n", findStatus)
						fmt.Printf("Nama : %s\nNIM : %s\nJurusan : %s\n", mhs[i].name, mhs[i].stdID, mhs[i].major)
						fmt.Println("Tercatat pada mata kuliah : ")
						printedHeader = true
						fmt.Println("+----------------------+-------------+----------------------+------------+-------+")
						fmt.Printf("| %-20s | %-11s | %-20s | %-10s | %-5s |\n", "Nama Mata Kuliah", "Kode Kelas", "Dosen", "Tanggal", "Jam")
						fmt.Println("+----------------------+-------------+----------------------+------------+-------+")
					}
					scheduleFound = false
					for k = 0; k < z && !scheduleFound; k++ {
						if log[j].scheduleID == sch[k].classCode {
							fmt.Printf("| %-20s | %-11s | %-20s | %-10s | %02d:%02d |\n", sch[k].course, sch[k].classCode, sch[k].lecturer, sch[k].date, sch[k].jam, sch[k].menit)
							scheduleFound = true
						}
					}
				}
			}
			if printedHeader {
				fmt.Println("+----------------------+-------------+----------------------+------------+-------+")
			}
		}
		if !match {
			fmt.Printf("Data mahasiswa dengan status kehadiran %s tidak ditemukan. Silakan coba kembali.\n", findStatus)
			fmt.Println("+--------------------------------------+")
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
			fmt.Println("+--------------------------------------+")
			fmt.Println("Pencarian data mahasiswa dibatalkan. Kembali ke menu utama...")
			fmt.Println("+--------------------------------------+")
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
			fmt.Println("+--------------------------------------+")
		} else {
			fmt.Printf("[+] Data mahasiswa dengan NIM %s ditemukan.\n", stdID)
			fmt.Printf("Nama : %s\nNIM : %s\nJurusan : %s\nAngkatan : %s\nIndeks Data : %d", mhs[found].name, mhs[found].stdID, mhs[found].major, mhs[found].batch, found+1)
			fmt.Printf("Nama : %s\nNIM : %s\nJurusan : %s\nAngkatan : %s\nIndeks Data : %d\n", mhs[found].name, mhs[found].stdID, mhs[found].major, mhs[found].batch, found+1)
			fmt.Println("+---------------------------------------+")
		}
	}
@@ -201,10 +201,6 @@
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
@@ -213,16 +209,15 @@
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
			case 4:
				continue
			}
		case 2:
			addSchedule(&schedule)
@@ -232,48 +227,48 @@
			fmt.Println("3. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu : ")
			fmt.Scan(&secondOption)
			if secondOption == 3 {
				continue
			}
			switch secondOption {
			case 1:
				addLogPresent(&mhs, x, &log, y, &schedule, z)
			case 2:
				updateLogPresent(&mhs, &log)
			case 3:
				continue
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
				searchDataByPresent(&mhs, x, &log, y, &schedule, z)
			case 2:
				sortByStdID(&mhs, x)
				searchDataByStdID(&mhs, x)
			case 3:
				continue
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
			case 3:
				continue
			}
		case 6:
			statisticPresent(&mhs, &log)
		case 7:
			fmt.Println("Terima kasih telah menggunakan Aplikasi SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa.\nSampai jumpa!")
			return
		}
	}
}