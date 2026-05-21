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

	if *x >= nMhs {
		fmt.Println("Data mahasiswa penuh!")
		return
	}
	fmt.Println("=== Tambah Data Mahasiswa ===")
	fmt.Print("Nama Mahasiswa     : ")
	fmt.Scan(&newMhs.name)
	fmt.Print("NIM Mahasiswa      : ")
	fmt.Scan(&newMhs.stdID)
	fmt.Print("Jurusan Mahasiswa  : ")
	fmt.Scan(&newMhs.major)
	fmt.Print("Angkatan Mahasiswa : ")
	fmt.Scan(&newMhs.batch)

	for i := 0; i < *x; i++ {
		if mhs[i].stdID == newMhs.stdID {
			fmt.Printf("NIM %s sudah terdaftar! Data tidak anda tambahkan.\n", newMhs.stdID)
			fmt.Println("+--------------------------------------+")
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
	fmt.Printf("[+] Data mahasiswa %s berhasil anda tambahkan.\n", newMhs.name)
	fmt.Println("+--------------------------------------+")
}

//Subprogram untuk mengubah data mahasiswa
func updateMhs(mhs *tabMhs, x *int) {
	var stdID, input string
	var left, mid, right, found int

	if *x == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println("=== Ubah Data Mahasiswa ===")
	fmt.Print("Masukkan NIM mahasiswa yang ingin anda ubah : ")
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

	fmt.Printf("Data ditemukan: %s | %s | %s | %s\n", mhs[found].name, mhs[found].stdID, mhs[found].major, mhs[found].batch)
	fmt.Println("Masukkan data baru (isi '-' untuk tidak mengubah field tersebut) :")

	fmt.Print("Nama baru    : ")
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

	fmt.Printf("[+] Data mahasiswa dengan NIM %s berhasil anda ubah.\n", stdID)
	fmt.Println("+--------------------------------------+")
}

// Subprogram untuk menghapus data mahasiswa
func deleteMhs(mhs *tabMhs, x *int) {
	var stdID, konfirmasi string
	var left, mid, right, found int

	if *x == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}
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
		fmt.Printf("Data mahasiswa dengan NIM %s tidak anda temukan.\n", stdID)
		fmt.Println("+--------------------------------------+")
		return
	}

	fmt.Printf("Data ditemukan: %s | %s | %s | %s\n", mhs[found].name, mhs[found].stdID, mhs[found].major, mhs[found].batch)
	fmt.Print("Apakah Anda yakin ingin menghapus data ini? (ya/tidak) : ")
	fmt.Scan(&konfirmasi)

	if konfirmasi != "ya" {
		fmt.Println("Penghapusan dibatalkan.")
		fmt.Println("+--------------------------------------+")
		return
	}

	// Geser data ke kiri untuk menutup celah
	for i := found; i < *x-1; i++ {
		mhs[i] = mhs[i+1]
	}
	mhs[*x-1] = student{}
	*x--

	fmt.Printf("[+] Data mahasiswa dengan NIM %s berhasil dihapus.\n", stdID)
	fmt.Println("+--------------------------------------+")
}

// Subprogram untuk menambahkan data jadwal kuliah
func addSchedule(schedule *tabSchedule, z *int) {
	var newSch schedule

	fmt.Println("=== Tambah Data Jadwal Kuliah ===")
	fmt.Print("Nama Mata Kuliah		: ")
	fmt.Scan(&newSch.course)
	fmt.Print("Kode Kelas			: ")
	fmt.Scan(&newSch.classCode)
	fmt.Print("Nama Dosen			: ")
	fmt.Scan(&newSch.lecturer)
	fmt.Print("Tanggal (DD-MM-YYYY) : ")
	fmt.Scan(&newSch.date)
	fmt.Print("Jam					: ")
	fmt.Scan(&newSch.jam)
	fmt.Print("Menit				: ")
	fmt.Scan(&newSch.menit)

	for i := 0; i < *z; i++ {
		if sch[i].classCode == newSch.classCode {
			fmt.Printf("Kode kelas %s sudah terdaftar! Data tidak anda tambahkan.\n", newSch.classCode)
			fmt.Println("+--------------------------------------+")
			return
		}
	}
	if *z >= nMhs {
		fmt.Println("Data jadwal sudah penuh!")
		return
	}

	sch[*z] = newSch
	*z++
	fmt.Printf("[+] Jadwal mata kuliah %s dengan kode kelas %s berhasil anda tambahkan.\n", newSch.course, newSch.classCode)
	fmt.Println("+--------------------------------------+")
}

// Subprogram untuk menambahkan data status kehadiran mahasiswa
func addLogPresent(mhs *tabMhs, x int, log *tabLog, y *int, sch *tabSchedule, z int) {
	var newLog log
	var pilihanMhs, pilihanSch int

	if x == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}
	if z == 0 {
		fmt.Println("Belum ada data jadwal kuliah.")
		return
	}
	if *y >= nLog {
		fmt.Println("Data log kehadiran sudah penuh!")
		return
	}

	fmt.Println("=== Daftar Mahasiswa ===")
	fmt.Println("+-----+----------------------+----------------+------------------+----------+")
	fmt.Printf("| %-3s | %-20s | %-14s | %-16s | %-8s |\n", "No", "Nama", "NIM", "Jurusan", "Angkatan")
	fmt.Println("+-----+----------------------+----------------+------------------+----------+")
	for i := 0; i < x; i++ {
		fmt.Printf("| %-3d | %-20s | %-14s | %-16s | %-8s |\n", i+1, mhs[i].name, mhs[i].stdID, mhs[i].major, mhs[i].batch)
	}
	fmt.Println("+-----+----------------------+----------------+------------------+----------+")

	fmt.Print("Pilih nomor mahasiswa : ")
	fmt.Scan(&pilihanMhs)
	if pilihanMhs < 1 || pilihanMhs > x {
		fmt.Println("Nomor mahasiswa tidak valid.")
		fmt.Println("+--------------------------------------+")
		return
	}
	selectedMhs := mhs[pilihanMhs-1]

	fmt.Println("\n=== Daftar Jadwal Kuliah ===")
	fmt.Println("+-----+----------------------+-------------+----------------------+------------+-------+")
	fmt.Printf("| %-3s | %-20s | %-11s | %-20s | %-10s | %-5s |\n", "No", "Mata Kuliah", "Kode Kelas", "Dosen", "Tanggal", "Jam")
	fmt.Println("+-----+----------------------+-------------+----------------------+------------+-------+")
	for i := 0; i < z; i++ {
		fmt.Printf("| %-3d | %-20s | %-11s | %-20s | %-10s | %02d:%02d |\n", i+1, sch[i].course, sch[i].classCode, sch[i].lecturer, sch[i].date, sch[i].jam, sch[i].menit)
	}
	fmt.Println("+-----+----------------------+-------------+----------------------+------------+-------+")

	fmt.Print("Pilih nomor jadwal : ")
	fmt.Scan(&pilihanSch)
	if pilihanSch < 1 || pilihanSch > z {
		fmt.Println("Nomor jadwal tidak valid.")
		fmt.Println("+--------------------------------------+")
		return
	}
	selectedSch := sch[pilihanSch-1]
	newLog.stdID = selectedMhs.stdID
	newLog.scheduleID = selectedSch.classCode
	fmt.Print("Sesi ke- : ")
	fmt.Scan(&newLog.session)

	for i := 0; i < *y; i++ {
		if log[i].stdID == newLog.stdID && log[i].scheduleID == newLog.scheduleID && log[i].session == newLog.session {
			fmt.Printf("Log kehadiran mahasiswa %s pada kelas %s sesi %d sudah tercatat!\n", newLog.stdID, newLog.scheduleID, newLog.session)
			fmt.Println("+--------------------------------------+")
			return
		}
	}

	fmt.Print("Status (Hadir/Sakit/Izin/Alpa) : ")
	fmt.Scan(&newLog.status)
	if newLog.status != "Hadir" && newLog.status != "Sakit" && newLog.status != "Izin" && newLog.status != "Alpa" {
		fmt.Println("Status tidak valid! Harus Hadir/Sakit/Izin/Alpa.")
		fmt.Println("+--------------------------------------+")
		return
	}
	log[*y] = newLog
	*y++
	fmt.Printf("[+] Log kehadiran mahasiswa %s (%s) pada kelas %s sesi %d dengan status %s berhasil anda tambahkan.\n", selectedMhs.name, selectedMhs.stdID, selectedSch.classCode, newLog.session, newLog.status)
	fmt.Println("+--------------------------------------+")
}

// Subprogram untuk mengubah data status kehadiran mahasiswa
func updateLogPresent(mhs *tabMhs, log *tabLog) {
	var newLog log
	var pilihanMhs, countLog, pilihanLog int
	var idxLog [nLog]int

	if x == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}
	if y == 0 {
		fmt.Println("Belum ada data log kehadiran.")
		return
	}

	fmt.Println("=== Daftar Mahasiswa ===")
	fmt.Println("+-----+----------------------+----------------+")
	fmt.Printf("| %-3s | %-20s | %-14s |\n", "No", "Nama", "NIM")
	fmt.Println("+-----+----------------------+----------------+")
	for i := 0; i < x; i++ {
		fmt.Printf("| %-3d | %-20s | %-14s |\n", i+1, mhs[i].name, mhs[i].stdID)
	}
	fmt.Println("+-----+----------------------+----------------+")
	fmt.Print("Pilih nomor mahasiswa : ")
	fmt.Scan(&pilihanMhs)
	if pilihanMhs < 1 || pilihanMhs > x {
		fmt.Println("Nomor mahasiswa tidak valid.")
		fmt.Println("+--------------------------------------+")
		return
	}
	selectedMhs := mhs[pilihanMhs-1]
	newLog.stdID = selectedMhs.stdID

	fmt.Printf("\n=== Log Kehadiran %s ===\n", selectedMhs.name)
	fmt.Println("+-----+-------------+----------+---------+")
	fmt.Printf("| %-3s | %-11s | %-8s | %-7s |\n", "No", "Kode Kelas", "Sesi", "Status")
	fmt.Println("+-----+-------------+----------+---------+")

	for i := 0; i < y; i++ {
		if log[i].stdID == selectedMhs.stdID {
			fmt.Printf("| %-3d | %-11s | %-8d | %-7s |\n", countLog+1, log[i].scheduleID, log[i].session, log[i].status)
			idxLog[countLog] = i
			countLog++
		}
	}
	fmt.Println("+-----+-------------+----------+---------+")

	if countLog == 0 {
		fmt.Printf("Belum ada log kehadiran untuk mahasiswa %s.\n", selectedMhs.name)
		fmt.Println("+--------------------------------------+")
		return
	}
	fmt.Print("Pilih nomor log yang ingin diubah : ")
	fmt.Scan(&pilihanLog)
	if pilihanLog < 1 || pilihanLog > countLog {
		fmt.Println("Nomor log tidak valid.")
		fmt.Println("+--------------------------------------+")
		return
	}

	targetIdx := idxLog[pilihanLog-1]
	fmt.Printf("Log dipilih: Kelas %s | Sesi %d | Status %s\n", log[targetIdx].scheduleID, log[targetIdx].session, log[targetIdx].status)

	fmt.Print("Status baru (Hadir/Sakit/Izin/Alpa) : ")
	fmt.Scan(&newLog.status)
	if newLog.status != "Hadir" && newLog.status != "Sakit" && newLog.status != "Izin" && newLog.status != "Alpa" {
		fmt.Println("Status tidak valid! Harus Hadir/Sakit/Izin/Alpa.")
		fmt.Println("+--------------------------------------+")
		return
	}

	log[targetIdx].status = newLog.status
	fmt.Printf("[+] Status kehadiran mahasiswa %s pada kelas %s sesi %d berhasil diubah menjadi %s.\n", selectedMhs.name, log[targetIdx].scheduleID, log[targetIdx].session, newLog.status)
	fmt.Println("+--------------------------------------+")
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
	var x, y, z, firstOption, secondOption int

	x = 0 // Variabel untuk menghitung jumlah data mahasiswa yang telah dimasukkan
	y = 0 // Variabel untuk menghitung jumlah data log kehadiran mahasiswa yang telah dimasukkan
	z = 0 // Variabel untuk menghitung jumlah data jadwal kuliah yang telah dimasukkan
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
			case 4:
				continue
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
