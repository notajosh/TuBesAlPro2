package main

// Import packages yang dibutuhkan
import (
	"bufio"   // Untuk membaca input string dengan spasi
	"fmt"     // Untuk input/output
	"os"      // Untuk clear screen dan pause terminal
	"os/exec" // Untuk menjalankan perintah clear screen di terminal
	"runtime" // Untuk mendeteksi OS pada saat clear screen
	"strings" // Untuk manipulasi string, seperti membuat border tabel dinamis
)

// DEKLARASI KONSTANTA DAN TIPE DATA SERTA BENTUKAN DALAM VARIABEL GLOBAL
// Kapasitas di bawah ini silakan disesuaikan dengan kebutuhan
const NMHS int = 50  // Max Capacity untuk Mahasiswa
const NSCH int = 10  // Max Capacity untuk Jadwal
const NATT int = 100 // Max Capacity untuk Log Kehadiran

// Tipe bentukan serta tipe data untuk menyimpan data mahasiswa
type Student struct {
	Name, StdID, Class string
	TH, TI, TS, TA     int
}

// Tipe bentukan serta tipe data untuk menyimpan data jadwal kuliah
type Schedule struct {
	SubjectCode, SubjectName, LectureCode, LectureName, Class, Day, Time string
}

// Tipe bentukan serta tipe data untuk menyimpan data kehadiran mahasiswa
type Attendance struct {
	StdID, SubjectCode, Status string
	Meeting                    int
}

type tabStudent [NMHS]Student       // Tipe data array untuk menyimpan data mahasiswa
type tabSchedule [NSCH]Schedule     // Tipe data array untuk menyimpan data jadwal kuliah
type tabAttendance [NATT]Attendance // Tipe data array untuk menyimpan data presensi mahasiswa

// INISIALISASI UI BERBASIS CLI

// Subprogram untuk membersihkan layar terminal
func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Subprogram untuk menunggu input Enter dari pengguna untuk melanjutkan ke langkah berikutnya
func pauseTerminal() {
	var enter string
	fmt.Print("Tekan Enter untuk melanjutkan...")
	fmt.Scanln(&enter)
}

// Subprogram pembantu untuk mencetak header aplikasi
func printAppHeader(isMainMenu bool, subMenuTitle string, nMhs int, nSch int, K tabAttendance, nK int) {
	fmt.Println("||==================================================================================================||")

	interiorWidth := 98 // Total ruang kosong di dalam kotak

	if isMainMenu {
		// Header untuk Menu Utama dan Dashboard Grafik kapasitas data
		title := "Selamat datang di Aplikasi SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa"
		padLeft := (interiorWidth - len(title)) / 2
		padRight := interiorWidth - len(title) - padLeft
		fmt.Printf("||%s%s%s||\n", strings.Repeat(" ", padLeft), title, strings.Repeat(" ", padRight))
		fmt.Println("||--------------------------------------------------------------------------------------------------||")

		// Kalkulasi Bar dan Persentase untuk Dashboard Grafik Kapasitas Data Mahasiswa
		mhsFilled := int((float64(nMhs) / float64(NMHS)) * 30)
		mhsBar := strings.Repeat("■", mhsFilled) + strings.Repeat(".", 30-mhsFilled)
		mhsPercent := (float64(nMhs) / float64(NMHS)) * 100
		mhsText := fmt.Sprintf("Kapasitas Mahasiswa : [%s] %5.1f%% Terpakai", mhsBar, mhsPercent)

		// Kalkulasi Bar dan Persentase untuk Dashboard Grafik Kapasitas Data Jadwal
		schFilled := int((float64(nSch) / float64(NSCH)) * 30)
		schBar := strings.Repeat("■", schFilled) + strings.Repeat(".", 30-schFilled)
		schPercent := (float64(nSch) / float64(NSCH)) * 100
		schText := fmt.Sprintf("Kapasitas Jadwal    : [%s] %5.1f%% Terpakai", schBar, schPercent)

		// Kalkulasi Progres Pertemuan Terjauh
		var maxMeeting int = 0
		for i := 0; i < nK; i++ {
			if K[i].Meeting > maxMeeting {
				maxMeeting = K[i].Meeting
			}
		}
		attText := fmt.Sprintf("Progres Pertemuan   : %d / 16 Pertemuan Tercatat", maxMeeting)

		// Print Dashboard Grafik Kapasitas Data Mahasiswa, Jadwal, dan Progres Pertemuan
		fmt.Printf("||    %-90s    ||\n", mhsText)
		fmt.Printf("||    %-90s    ||\n", schText)
		fmt.Printf("||    %-90s    ||\n", attText)
		fmt.Println("||--------------------------------------------------------------------------------------------------||")
	} else {
		// Header untuk Sub Menu
		title := "SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa"
		padLeft := (interiorWidth - len(title)) / 2
		padRight := interiorWidth - len(title) - padLeft
		fmt.Printf("||%s%s%s||\n", strings.Repeat(" ", padLeft), title, strings.Repeat(" ", padRight))
		fmt.Println("||--------------------------------------------------------------------------------------------------||")
	}

	// Print Title Sub Menu dengan Center Text format
	padLeftSub := (interiorWidth - len(subMenuTitle)) / 2
	padRightSub := interiorWidth - len(subMenuTitle) - padLeftSub
	fmt.Printf("||%s%s%s||\n", strings.Repeat(" ", padLeftSub), subMenuTitle, strings.Repeat(" ", padRightSub))
	fmt.Println("||==================================================================================================||")
}

// CRUD DATA MAHASISWA

// Subprogram untuk menambahkan data mahasiswa
func addMhs(M *tabStudent, n *int, reader *bufio.Reader) {
	if *n >= NMHS { // Cek jika kapasitas data mahasiswa sudah penuh (Dibalik untuk Early Return)
		fmt.Println("Kapasitas data mahasiswa sudah penuh.")
		return
	}

	fmt.Println("\n--- Tambah Data Mahasiswa ---")

	// Validasi inputan Nama agar tidak boleh kosong
	for {
		fmt.Print("Masukkan Nama Mahasiswa: ")
		tempName := readStringWithSpace(reader)
		if tempName != "" {
			M[*n].Name = tempName
			break
		}
		fmt.Println("[!] Error: Nama Mahasiswa tidak boleh kosong! Silakan ulangi.\n")
	}

	// Validasi inputan NIM agar tepat 12 digit
	for {
		fmt.Print("Masukkan NIM Mahasiswa (Wajib 12 Digit): ")
		tempNIM := readStringWithSpace(reader)
		if len(tempNIM) == 12 {
			M[*n].StdID = tempNIM
			break
		} else {
			fmt.Printf("[!] Error: NIM Anda %d digit. Harus tepat 12 digit! Silakan ulangi.\n\n", len(tempNIM))
		}
	}

	// Validasi inputan Kelas agar tidak boleh kosong
	for {
		fmt.Print("Masukkan Kelas Mahasiswa: ")
		tempClass := readStringWithSpace(reader)
		if tempClass != "" {
			M[*n].Class = tempClass
			break
		}
		fmt.Println("[!] Error: Kelas Mahasiswa tidak boleh kosong! Silakan ulangi.\n")
	}

	fmt.Printf("\nYakin menambahkan mahasiswa %s (%s)? (y/n): ", M[*n].Name, M[*n].StdID)
	confirm := readStringWithSpace(reader)
	if confirm == "y" || confirm == "Y" {
		M[*n].TH, M[*n].TI, M[*n].TS, M[*n].TA = 0, 0, 0, 0
		*n++
		fmt.Println("Data mahasiswa berhasil ditambahkan.")
	} else {
		fmt.Println("Penambahan data mahasiswa dibatalkan.")
	}
}

// Subprogram untuk membaca data mahasiswa dengan Tabel Dinamis (Auto-Resize)
func readMhs(M tabStudent, n int) {
	if n == 0 { // Cek jika belum ada data mahasiswa yang tersedia
		fmt.Println("\nBelum ada data mahasiswa yang tersedia untuk ditampilkan.")
		return
	}

	fmt.Print("\nApakah Anda yakin ingin melihat data seluruh mahasiswa? (y/n): ")
	var confirm string
	fmt.Scanln(&confirm)
	if confirm != "y" && confirm != "Y" {
		fmt.Println("Menampilkan data mahasiswa dibatalkan.")
		return
	}

	// Cari teks terpanjang di setiap kolom untuk menentukan lebar kolom yang dinamis
	maxName := 20 // Jatah minimal untuk kolom Nama
	maxNim := 15  // Jatah minimal untuk kolom NIM
	maxClass := 8 // Jatah minimal untuk kolom Kelas

	// Loop untuk mencari panjang teks terpanjang di setiap kolom
	for i := 0; i < n; i++ {
		if len(M[i].Name) > maxName {
			maxName = len(M[i].Name)
		}
		if len(M[i].StdID) > maxNim {
			maxNim = len(M[i].StdID)
		}
		if len(M[i].Class) > maxClass {
			maxClass = len(M[i].Class)
		}
	}
	// Buat border line berdasarkan lebar kolom yang sudah dihitung
	borderLine := fmt.Sprintf("+%s+%s+%s+-------+-------+-------+-------+",
		strings.Repeat("-", maxName+2),
		strings.Repeat("-", maxNim+2),
		strings.Repeat("-", maxClass+2),
	)

	// Cetak header tabel dengan lebar kolom yang sudah dihitung
	fmt.Println("\n--- Data Mahasiswa ---")
	fmt.Println(borderLine)

	// Menggunakan %-*s untuk menyuntikkan variabel ukuran kolom (maxName, maxNim, dll)
	fmt.Printf("| %-*s | %-*s | %-*s | %-5s | %-5s | %-5s | %-5s |\n",
		maxName, "Nama", maxNim, "NIM", maxClass, "Kelas", "Hadir", "Izin", "Sakit", "Alpa")
	fmt.Println(borderLine)

	// Loop untuk mencetak setiap data mahasiswa dengan format yang sudah disesuaikan dengan lebar kolom
	for i := 0; i < n; i++ {
		fmt.Printf("| %-*s | %-*s | %-*s | %-5d | %-5d | %-5d | %-5d |\n",
			maxName, M[i].Name, maxNim, M[i].StdID, maxClass, M[i].Class, M[i].TH, M[i].TI, M[i].TS, M[i].TA)
	}
	fmt.Println(borderLine)

	// Footer dengan total mahasiswa terdaftar
	footerWidth := maxName + maxNim + maxClass + 6
	fmt.Printf("| %-*s | %-29d |\n", footerWidth-1, "Total Mahasiswa Terdaftar", n)

	// Bottom border untuk menutup tabel
	bottomBorder := fmt.Sprintf("+%s+-------------------------------+", strings.Repeat("-", footerWidth+2))
	fmt.Println(bottomBorder)
}

// Subprogram untuk memperbarui data mahasiswa
func updateMhs(M *tabStudent, n int, reader *bufio.Reader) {
	if n == 0 { // Cek jika belum ada data mahasiswa yang tersedia untuk diperbarui
		fmt.Println("\nBelum ada data mahasiswa untuk diperbarui.")
		return
	}

	var findStdID string
	fmt.Println("\n--- Perbarui Data Mahasiswa ---")
	fmt.Print("Masukkan NIM mahasiswa yang ingin diperbarui: ")
	fmt.Scanln(&findStdID)

	// Panggil fungsi pembantu untuk mencari indeks
	idx := studentIdxSearch(*M, n, findStdID)
	if idx == -1 { // Jika pencarian menggunakan loop selesai tetapi data dengan NIM input tidak ditemukan
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", findStdID)
		return
	}

	// Pencarian data mahasiswa berdasarkan NIM yang dimasukkan (Logika dieksekusi langsung menggunakan idx)
	fmt.Printf("Data ditemukan:\nNama: %s\nNIM: %s\nKelas: %s\n", M[idx].Name, M[idx].StdID, M[idx].Class)
	fmt.Println("Masukkan data baru (tekan Enter untuk mempertahankan data lama):")

	fmt.Print("Nama Mahasiswa: ")
	newName := readStringWithSpace(reader)

	fmt.Print("Kelas Mahasiswa: ")
	newClass := readStringWithSpace(reader)

	fmt.Printf("\nYakin menyimpan perubahan data %s? (y/n): ", M[idx].StdID)
	confirm := readStringWithSpace(reader)

	if confirm == "y" || confirm == "Y" {
		if newName != "" {
			M[idx].Name = newName
		}
		if newClass != "" {
			M[idx].Class = newClass
		}
		fmt.Println("Data mahasiswa berhasil diperbarui.")
	} else {
		fmt.Println("Perubahan dibatalkan.")
	}
}

// Subprogram untuk menghapus data mahasiswa
func deleteMhs(M *tabStudent, n *int) {
	if *n == 0 { // Cek jika belum ada data mahasiswa yang tersedia untuk dihapus
		fmt.Println("\nBelum ada data mahasiswa yang tersedia untuk dihapus.")
		return
	}

	var findStdID string
	fmt.Println("\n--- Hapus Data Mahasiswa ---")
	fmt.Print("Masukkan NIM mahasiswa yang ingin dihapus: ")
	fmt.Scanln(&findStdID)

	// Panggil fungsi pembantu untuk mencari indeks
	idx := studentIdxSearch(*M, *n, findStdID)
	if idx == -1 { // Data mahasiswa dengan NIM input tidak ditemukan
		fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan. Tidak ada data yang dihapus.\n", findStdID)
		return
	}

	// Print data ditemukan dan konfirmasi penghapusan data mahasiswa jika ditemukan
	fmt.Printf("Data mahasiswa ditemukan:\nNama: %s\nNIM: %s\nKelas: %s\n", M[idx].Name, M[idx].StdID, M[idx].Class)
	fmt.Printf("\nApakah Anda yakin untuk menghapus data mahasiswa %s? (y/n): ", M[idx].StdID)
	var confirm string
	fmt.Scanln(&confirm)

	if confirm == "y" || confirm == "Y" {
		// Proses menghapus data mahasiswa dengan menggeser data setelahnya ke kiri
		for j := idx; j < *n-1; j++ {
			M[j] = M[j+1]
		}
		*n--
		fmt.Println("Data mahasiswa berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan data mahasiswa dibatalkan.")
	}
}

// PRESENSI MAHASISWA

// Subprogram untuk menandai kehadiran mahasiswa
func markAttendance(M *tabStudent, nMhs int, J tabSchedule, nJ int, K *tabAttendance, nK *int) {
	if nMhs == 0 { // Cek jika belum ada data mahasiswa yang tersedia untuk mencatat presensi
		fmt.Println("\nBelum ada data mahasiswa yang tersedia untuk mencatat presensi.")
		return
	}
	if nJ == 0 { // Cek jika belum ada data jadwal yang tersedia untuk mencatat presensi
		fmt.Println("\nBelum ada data jadwal yang tersedia untuk mencatat presensi.")
		return
	}

	var targetClass, subjectCode string
	var meeting int
	fmt.Println("\n--- Catat Presensi Mahasiswa ---")
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&targetClass)
	fmt.Print("Masukkan Kode Mata Kuliah: ")
	fmt.Scanln(&subjectCode)
	fmt.Print("Masukkan Pertemuan Ke-: ")
	fmt.Scanln(&meeting)

	// Panggil fungsi pembantu untuk mencari indeks jadwal
	scheduleIdx := scheduleIdxSearch(J, nJ, subjectCode, targetClass)
	if scheduleIdx == -1 { // Jadwal dengan Kelas dan Kode Mata Kuliah tidak ditemukan
		fmt.Printf("Jadwal dengan Kelas %s dan Kode Mata Kuliah %s tidak ditemukan.\n", targetClass, subjectCode)
		return
	}

	// Print jadwal ditemukan dan lanjutkan proses pencatatan presensi jika jadwal ditemukan
	fmt.Printf("Jadwal ditemukan!\nMata Kuliah: %s\nDosen: %s (%s)\nHari: %s\nJam: %s\n", J[scheduleIdx].SubjectName, J[scheduleIdx].LectureName, J[scheduleIdx].LectureCode, J[scheduleIdx].Day, J[scheduleIdx].Time)
	fmt.Println("Masukkan status kehadiran untuk setiap mahasiswa: Hadir (H), Izin (I), Sakit (S), Alpa (A)")

	var classFound bool = false // Flag mahasiswa dengan kelas yang sesuai ditemukan atau tidak

	for i := 0; i < nMhs; i++ {
		if M[i].Class == targetClass {
			classFound = true
			var status, statusFull string

			fmt.Printf("Mahasiswa: %s (%s) - Status Kehadiran: ", M[i].Name, M[i].StdID)
			fmt.Scanln(&status)

			switch status {
			case "H", "h", "Hadir", "hadir", "HADIR":
				statusFull = "Hadir"
				M[i].TH++
			case "I", "i", "Izin", "izin", "IZIN":
				statusFull = "Izin"
				M[i].TI++
			case "S", "s", "Sakit", "sakit", "SAKIT":
				statusFull = "Sakit"
				M[i].TS++
			case "A", "a", "Alpa", "alpa", "ALPA":
				statusFull = "Alpa"
				M[i].TA++
			default:
				fmt.Println("Input invalid, status kehadiran dianggap Alpa.")
				statusFull = "Alpa"
				M[i].TA++
			}

			if *nK < NATT { // Batas kapasitas aman untuk log presensi
				K[*nK].StdID = M[i].StdID
				K[*nK].SubjectCode = subjectCode
				K[*nK].Status = statusFull
				K[*nK].Meeting = meeting
				*nK++
			}
		}
	}

	if classFound { // Cek jika ada mahasiswa dengan kelas yang sesuai ditemukan untuk mencatat presensi
		fmt.Printf("Presensi pertemuan ke-%d untuk kelas %s pada mata kuliah %s berhasil dicatat.\n", meeting, targetClass, J[scheduleIdx].SubjectName)
	} else {
		fmt.Printf("Tidak ditemukan mahasiswa yang terdaftar di kelas %s. Presensi tidak dapat dicatat.\n", targetClass)
	}
}

// Subprogram untuk membaca data kehadiran mahasiswa
func readAttendance(K tabAttendance, nK int) {
	if nK == 0 { // Cek jika belum ada data kehadiran mahasiswa yang tersedia untuk ditampilkan
		fmt.Println("\nBelum ada data kehadiran yang tersedia untuk ditampilkan.")
		return
	}

	// Cari teks terpanjang di setiap kolom untuk menentukan lebar kolom yang dinamis
	maxNim := 15    // Jatah minimal untuk kolom NIM
	maxSubj := 15   // Jatah minimal untuk kolom Kode MK
	maxStatus := 9  // Jatah minimal untuk kolom Status Kehadiran
	maxMeeting := 9 // Jatah minimal untuk kolom Pertemuan

	// Loop untuk mencari panjang teks terpanjang di setiap kolom
	for i := 0; i < nK; i++ {
		if len(K[i].StdID) > maxNim {
			maxNim = len(K[i].StdID)
		}
		if len(K[i].SubjectCode) > maxSubj {
			maxSubj = len(K[i].SubjectCode)
		}
		if len(K[i].Status) > maxStatus {
			maxStatus = len(K[i].Status)
		}
	}

	// Buat border line berdasarkan lebar kolom yang sudah dihitung
	borderLine := fmt.Sprintf("+%s+%s+%s+%s+",
		strings.Repeat("-", maxNim+2),
		strings.Repeat("-", maxSubj+2),
		strings.Repeat("-", maxStatus+2),
		strings.Repeat("-", maxMeeting+2),
	)

	// Cetak header tabel dengan lebar kolom yang sudah dihitung
	fmt.Println("\n--- Data Kehadiran Mahasiswa ---")
	fmt.Println(borderLine)

	// Menggunakan %-*s untuk menyuntikkan variabel ukuran kolom (maxNim, maxSubj, dll)
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s |\n",
		maxNim, "NIM", maxSubj, "Kode MK", maxStatus, "Status", maxMeeting, "Pertemuan")
	fmt.Println(borderLine)

	// Loop untuk mencetak setiap data kehadiran mahasiswa dengan format yang sudah disesuaikan dengan lebar kolom
	for i := 0; i < nK; i++ {
		fmt.Printf("| %-*s | %-*s | %-*s | %-*d |\n",
			maxNim, K[i].StdID, maxSubj, K[i].SubjectCode, maxStatus, K[i].Status, maxMeeting, K[i].Meeting)
	}
	fmt.Println(borderLine)

	// Footer dengan total data kehadiran tercatat
	footerWidth := maxNim + maxSubj + maxStatus + 6
	fmt.Printf("| %-*s | %-*d |\n", footerWidth, "Total Data Kehadiran Tercatat", maxMeeting, nK)

	// Bottom border untuk menutup tabel
	bottomBorder := fmt.Sprintf("+%s+%s+",
		strings.Repeat("-", footerWidth+2),
		strings.Repeat("-", maxMeeting+2),
	)
	fmt.Println(bottomBorder)
}

// Subprogram untuk memperbarui data kehadiran mahasiswa
func updateAttendance(K *tabAttendance, nK int, M *tabStudent, nM int) {
	if nK == 0 { // Cek jika belum ada data kehadiran mahasiswa yang tersedia untuk diperbarui
		fmt.Println("\nBelum ada data kehadiran yang tersedia untuk diperbarui.")
		return
	}

	var findStdID, findSubjectCode string
	var findMeeting int
	fmt.Println("\n --- Perbarui Presensi Mahasiswa ---")

	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scanln(&findStdID)

	fmt.Print("Masukkan Kode Mata Kuliah: ")
	fmt.Scanln(&findSubjectCode)

	fmt.Print("Masukkan Pertemuan Ke-: ")
	fmt.Scanln(&findMeeting)

	var attendanceIdx int = logIdxSearch(*K, nK, findStdID, findSubjectCode, findMeeting) // Cari index data kehadiran yang sesuai dengan inputan NIM, Kode MK, dan Pertemuan
	if attendanceIdx == -1 {                                                              // Log presensi tidak ditemukan
		fmt.Println("\nLog presensi tidak ditemukan. Pastikan NIM, Kode MK, dan Pertemuan benar.")
		return
	}

	fmt.Printf("Data kehadiran ditemukan:\nNIM: %s\nKode MK: %s\nPertemuan ke-%d\nStatus Kehadiran: %s\n", K[attendanceIdx].StdID, K[attendanceIdx].SubjectCode, K[attendanceIdx].Meeting, K[attendanceIdx].Status)
	fmt.Println("Masukkan status kehadiran baru: Hadir (H), Izin (I), Sakit (S), Alpa (A)")

	var scanNewStatus, tempStatus string
	fmt.Scanln(&scanNewStatus)

	switch scanNewStatus {
	case "H", "h", "Hadir", "hadir", "HADIR":
		tempStatus = "Hadir"
	case "I", "i", "Izin", "izin", "IZIN":
		tempStatus = "Izin"
	case "S", "s", "Sakit", "sakit", "SAKIT":
		tempStatus = "Sakit"
	case "A", "a", "Alpa", "alpa", "ALPA":
		tempStatus = "Alpa"
	default:
		tempStatus = ""
	}

	if tempStatus == "" {
		fmt.Println("Status tidak valid. Pembaruan presensi dibatalkan.")
		return
	}

	var confirm string
	fmt.Printf("Apakah Anda yakin ingin mengubah presensi %s menjadi %s? (y/n):", findStdID, tempStatus)
	fmt.Scanln(&confirm)

	if confirm != "y" && confirm != "Y" {
		fmt.Println("Pembaruan history presensi dibatalkan.")
		return
	}

	idxMhs := searchByStdID(*M, nM, findStdID, 1)
	if idxMhs != -1 {
		switch K[attendanceIdx].Status {
		case "Hadir":
			M[idxMhs].TH--
		case "Izin":
			M[idxMhs].TI--
		case "Sakit":
			M[idxMhs].TS--
		case "Alpa":
			M[idxMhs].TA--
		}

		switch tempStatus {
		case "Hadir":
			M[idxMhs].TH++
		case "Izin":
			M[idxMhs].TI++
		case "Sakit":
			M[idxMhs].TS++
		case "Alpa":
			M[idxMhs].TA++
		}
	}

	K[attendanceIdx].Status = tempStatus
	fmt.Println("History dan akumulasi presensi kehadiran berhasil diperbarui")
}

// SCHEDULE / Jadwal Kuliah Mahasiswa

// Subprogram untuk menambahkan jadwal mata kuliah
func addSchedule(J *tabSchedule, n *int, reader *bufio.Reader) {
	if *n >= NSCH { // Cek jika kapasitas data jadwal kuliah sudah penuh (Dibalik untuk Early Return)
		fmt.Println("Kapasitas data jadwal sudah penuh.")
		return
	}

	fmt.Println("\n --- Tambah Jadwal Mata Kuliah ---")
	fmt.Print("Masukkan Kode Mata Kuliah         : ")
	J[*n].SubjectCode = readStringWithSpace(reader)

	fmt.Print("Masukkan Nama Mata Kuliah         : ")
	J[*n].SubjectName = readStringWithSpace(reader)

	fmt.Print("Masukkan Kode Dosen               : ")
	J[*n].LectureCode = readStringWithSpace(reader)

	fmt.Print("Masukkan Nama Dosen               : ")
	J[*n].LectureName = readStringWithSpace(reader)

	fmt.Print("Masukkan Kelas                    : ")
	J[*n].Class = readStringWithSpace(reader)

	fmt.Print("Masukkan Hari                     : ")
	J[*n].Day = readStringWithSpace(reader)

	fmt.Print("Masukkan Waktu (cth: 08:00-10:00) : ")
	J[*n].Time = readStringWithSpace(reader)

	fmt.Printf("\nYakin menambahkan jadwal %s? (y/n): ", J[*n].SubjectName)
	confirm := readStringWithSpace(reader)

	if confirm == "y" || confirm == "Y" {
		*n++
		fmt.Println("Jadwal mata kuliah berhasil ditambahkan.")
	} else {
		fmt.Println("Penambahan jadwal dibatalkan.")
	}
}

// Subprogram untuk membaca data jadwal kuliah dengan Tabel Dinamis (Auto-Resize)
func readSchedule(J tabSchedule, n int) {
	if n == 0 {
		fmt.Println("\nBelum ada data jadwal yang tersedia untuk ditampilkan.")
		return
	}

	maxSCode := 7
	maxSName := 11
	maxLCode := 8
	maxLName := 10
	maxClass := 5
	maxDay := 4
	maxTime := 5

	for i := 0; i < n; i++ {
		if len(J[i].SubjectCode) > maxSCode {
			maxSCode = len(J[i].SubjectCode)
		}
		if len(J[i].SubjectName) > maxSName {
			maxSName = len(J[i].SubjectName)
		}
		if len(J[i].LectureCode) > maxLCode {
			maxLCode = len(J[i].LectureCode)
		}
		if len(J[i].LectureName) > maxLName {
			maxLName = len(J[i].LectureName)
		}
		if len(J[i].Class) > maxClass {
			maxClass = len(J[i].Class)
		}
		if len(J[i].Day) > maxDay {
			maxDay = len(J[i].Day)
		}
		if len(J[i].Time) > maxTime {
			maxTime = len(J[i].Time)
		}
	}

	borderLine := fmt.Sprintf("+%s+%s+%s+%s+%s+%s+%s+",
		strings.Repeat("-", maxSCode+2), strings.Repeat("-", maxSName+2),
		strings.Repeat("-", maxLCode+2), strings.Repeat("-", maxLName+2),
		strings.Repeat("-", maxClass+2), strings.Repeat("-", maxDay+2),
		strings.Repeat("-", maxTime+2),
	)

	fmt.Println("\n--- Data Jadwal Kuliah ---")
	fmt.Println(borderLine)
	fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
		maxSCode, "Kode MK", maxSName, "Mata Kuliah", maxLCode, "Kode Dsn",
		maxLName, "Nama Dosen", maxClass, "Kelas", maxDay, "Hari", maxTime, "Waktu")
	fmt.Println(borderLine)

	for i := 0; i < n; i++ {
		fmt.Printf("| %-*s | %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
			maxSCode, J[i].SubjectCode, maxSName, J[i].SubjectName,
			maxLCode, J[i].LectureCode, maxLName, J[i].LectureName,
			maxClass, J[i].Class, maxDay, J[i].Day, maxTime, J[i].Time)
	}
	fmt.Println(borderLine)

	// Footer dengan total jadwal tersedia
	footerWidth := maxSCode + maxSName + maxLCode + maxLName + maxClass + maxDay + 15
	fmt.Printf("| %-*s | %-*d |\n", footerWidth, "Total Jadwal Tersedia", maxTime, n)

	// Bottom border menyesuaikan dengan lebar inner footer
	bottomBorder := fmt.Sprintf("+%s+%s+", strings.Repeat("-", footerWidth+2), strings.Repeat("-", maxTime+2))
	fmt.Println(bottomBorder)
}

// Subprogram untuk memperbarui jadwal kuliah
func updateSchedule(J *tabSchedule, n int, reader *bufio.Reader) {
	if n == 0 { // Cek jika belum ada data jadwal kuliah yang tersedia untuk diperbarui
		fmt.Println("\nBelum ada data jadwal untuk diperbarui.")
		return
	}

	var findCode, findClass string
	fmt.Println("\n--- Perbarui Jadwal Kuliah ---")
	fmt.Print("Masukkan Kode Mata Kuliah: ")
	fmt.Scanln(&findCode)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&findClass)

	// Panggil fungsi pembantu untuk mencari indeks
	idx := scheduleIdxSearch(*J, n, findCode, findClass)
	if idx == -1 { // Jika pencarian menggunakan loop selesai tetapi data dengan NIM input tidak ditemukan
		fmt.Printf("\nJadwal untuk MK %s di Kelas %s tidak ditemukan.\n", findCode, findClass)
		return
	}

	fmt.Printf("\nJadwal ditemukan:\nMK: %s (%s)\nDosen: %s (%s)\nHari/Jam: %s / %s\n",
		J[idx].SubjectName, J[idx].SubjectCode, J[idx].LectureName, J[idx].LectureCode, J[idx].Day, J[idx].Time)
	fmt.Println("\nMasukkan data baru (tekan Enter tanpa mengisi untuk mempertahankan data lama):")

	fmt.Print("Nama Mata Kuliah Baru : ")
	newName := readStringWithSpace(reader)
	if newName != "" {
		J[idx].SubjectName = newName
	}

	fmt.Print("Kode Dosen Baru       : ")
	newLCode := readStringWithSpace(reader)
	if newLCode != "" {
		J[idx].LectureCode = newLCode
	}

	fmt.Print("Nama Dosen Baru       : ")
	newLName := readStringWithSpace(reader)
	if newLName != "" {
		J[idx].LectureName = newLName
	}

	fmt.Print("Hari Baru             : ")
	newDay := readStringWithSpace(reader)
	if newDay != "" {
		J[idx].Day = newDay
	}

	fmt.Print("Waktu Baru            : ")
	newTime := readStringWithSpace(reader)
	if newTime != "" {
		J[idx].Time = newTime
	}

	fmt.Println("\nData jadwal berhasil diperbarui.")
}

// Subprogram untuk menghapus jadwal kuliah (Versi Modular)
func deleteSchedule(J *tabSchedule, n *int) {
	if *n == 0 { // Cek jika belum ada data jadwal kuliah yang tersedia untuk dihapus
		fmt.Println("\nBelum ada data jadwal yang tersedia untuk dihapus.")
		return
	}

	var findCode, findClass string
	fmt.Println("\n--- Hapus Jadwal Kuliah ---")
	fmt.Print("Masukkan Kode Mata Kuliah yang ingin dihapus: ")
	fmt.Scanln(&findCode)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&findClass)

	// Panggil fungsi pembantu untuk mencari indeks
	idx := scheduleIdxSearch(*J, *n, findCode, findClass)
	if idx == -1 {
		fmt.Printf("\nData jadwal dengan Kode MK %s di Kelas %s tidak ditemukan.\n", findCode, findClass)
		return
	}

	fmt.Printf("\nData jadwal ditemukan:\nMK: %s (%s)\nDosen: %s\nHari/Jam: %s / %s\n",
		J[idx].SubjectName, J[idx].SubjectCode, J[idx].LectureName, J[idx].Day, J[idx].Time)
	fmt.Printf("\nApakah Anda yakin untuk menghapus jadwal ini? (y/n): ")
	var confirm string
	fmt.Scanln(&confirm)

	if confirm == "y" || confirm == "Y" {
		// Proses menghapus data jadwal dengan menggeser data setelahnya ke kiri
		for j := idx; j < *n-1; j++ {
			J[j] = J[j+1]
		}
		*n--
		fmt.Println("Data jadwal berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan jadwal dibatalkan.")
	}
}

// SEARCHING

// Subprogram untuk mencari data mahasiswa berdasarkan status presensi
func searchByAttendance(K tabAttendance, nK int, M tabStudent, nM int, status string, chosenAlgorithm int) {
	if nK == 0 { // Cek jika belum ada data kehadiran mahasiswa yang tersedia untuk dicari berdasarkan status presensi
		fmt.Println("\nBelum ada history presensi yang tercatat.")
		return
	}
	var statusTarget string
	switch status {
	case "H", "h", "Hadir", "hadir", "HADIR":
		statusTarget = "Hadir"
	case "I", "i", "Izin", "izin", "IZIN":
		statusTarget = "Izin"
	case "S", "s", "Sakit", "sakit", "SAKIT":
		statusTarget = "Sakit"
	case "A", "a", "Alpa", "alpa", "ALPA":
		statusTarget = "Alpa"
	default:
		statusTarget = ""
	}

	if statusTarget == "" {
		fmt.Println("\nStatus pencarian tidak valid. Format : H/h/Hadir/hadir/HADIR")
	} else {
		fmt.Printf("\n--- Hasil Pencarian Mahasiswa dengan Status: %s ---\n", statusTarget)
		var found bool = false

		if chosenAlgorithm == 1 {
			// Sequential Search
			for i := 0; i < nK; i++ {
				if K[i].Status == statusTarget {
					found = true
					var mhsName, mhsClass string
					var mhsFound bool = false
					for j := 0; j < nM && !mhsFound; j++ {
						if M[j].StdID == K[i].StdID {
							mhsName = M[j].Name
							mhsClass = M[j].Class
							mhsFound = true
						}
					}
					fmt.Printf("- Pertemuan : %d | MK : %s | NIM : %s | Nama : %s | Kelas : %s\n", K[i].Meeting, K[i].SubjectCode, K[i].StdID, mhsName, mhsClass)
				}
			}
		} else if chosenAlgorithm == 2 {
			// Binary Search yang dimodifikasi untuk mencari semua data dengan status yang sama
			fmt.Println("[!] Warning: Mengurutkan data log berdasarkan status secara otomatis...")
			sortLogByAttendance(&K, nK)

			var left, right, foundIdx int = 0, nK - 1, -1
			var mid int
			for left <= right && foundIdx == -1 {
				mid = (left + right) / 2
				if K[mid].Status == statusTarget {
					foundIdx = mid
				} else if K[mid].Status < statusTarget {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}

			// Scan ke kiri dan kanan dari foundIdx untuk menemukan semua data dengan status yang sama
			if foundIdx != -1 {
				found = true
				start := foundIdx
				for start > 0 && K[start-1].Status == statusTarget {
					start--
				}
				end := foundIdx
				for end < nK-1 && K[end+1].Status == statusTarget {
					end++
				}
				for i := start; i <= end; i++ {
					var mhsName, mhsClass string
					var mhsFound bool = false
					for j := 0; j < nM && !mhsFound; j++ {
						if M[j].StdID == K[i].StdID {
							mhsName = M[j].Name
							mhsClass = M[j].Class
							mhsFound = true
						}
					}
					fmt.Printf("- Pertemuan : %d | MK : %s | NIM : %s | Nama : %s | Kelas : %s\n", K[i].Meeting, K[i].SubjectCode, K[i].StdID, mhsName, mhsClass)
				}
			}
		} else {
			fmt.Println("Pilihan algoritma invalid. Silakan coba kembali.")
		}

		if !found {
			fmt.Printf("History mahasiswa dengan status %s tidak ditemukan.\n", statusTarget)
		}
		fmt.Println("---------------------------------------------------------------")
	}
}

// Subprogram untuk mencari data mahasiswa berdasarkan NIM
func searchByStdID(M tabStudent, n int, stdID string, order int) int {
	if n == 0 { // Cek jika belum ada data mahasiswa yang tersedia untuk dicari berdasarkan NIM
		fmt.Println("\nBelum ada data mahasiswa untuk dicari berdasarkan NIM.")
		return -1
	}

	found := false
	idx := -1

	if order == 1 {
		// Sequential Search
		for i := 0; i < n && !found; i++ {
			if M[i].StdID == stdID {
				found = true
				idx = i
			}
		}
	} else if order == 2 {
		// Binary Search
		var left, mid int = 0, 0
		var right int = n - 1
		for left <= right && !found {
			mid = (left + right) / 2
			if M[mid].StdID == stdID {
				found = true
				idx = mid
			} else if M[mid].StdID < stdID {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	} else {
		return -1
	}

	return idx
}

// SORTING

// Subprogram untuk mengurutkan data mahasiswa berdasarkan kategori presensi pilihan
func sortByAttendance(M *tabStudent, n int, metric int, order int, algo int) {
	if algo == 1 {
		// Insertion Sort with Tuple Assignment
		if order == 1 {
			// Asc
			for i := 1; i < n; i++ {
				for j := i; j > 0 && getMetricValue(M[j-1], metric) > getMetricValue(M[j], metric); j-- {
					M[j], M[j-1] = M[j-1], M[j] // Instant swap
				}
			}
		} else if order == 2 {
			// Desc
			for i := 1; i < n; i++ {
				for j := i; j > 0 && getMetricValue(M[j-1], metric) < getMetricValue(M[j], metric); j-- {
					M[j], M[j-1] = M[j-1], M[j] // Instant swap
				}
			}
		}
	} else if algo == 2 {
		// Selection Sort with Tuple Assignment
		if order == 1 {
			// Asc
			for i := 0; i < n-1; i++ {
				minIdx := i
				for j := i + 1; j < n; j++ {
					if getMetricValue(M[j], metric) < getMetricValue(M[minIdx], metric) {
						minIdx = j
					}
				}
				M[i], M[minIdx] = M[minIdx], M[i] // Instant swap
			}
		} else if order == 2 {
			// Desc
			for i := 0; i < n-1; i++ {
				maxIdx := i
				for j := i + 1; j < n; j++ {
					if getMetricValue(M[j], metric) > getMetricValue(M[maxIdx], metric) {
						maxIdx = j
					}
				}
				M[i], M[maxIdx] = M[maxIdx], M[i] // Instant swap
			}
		}
	}
}

// Subprogram untuk mengurutkan data mahasiswa berdasarkan nama
func sortByName(M *tabStudent, n int, order int, algo int) {
	if algo == 1 {
		// Insertion Sort with Tuple Assignment
		if order == 1 {
			// Ascending (A-Z)
			for i := 1; i < n; i++ {
				for j := i; j > 0 && M[j-1].Name > M[j].Name; j-- {
					M[j], M[j-1] = M[j-1], M[j] // Instant swap
				}
			}
		} else if order == 2 {
			// Descending (Z-A)
			for i := 1; i < n; i++ {
				for j := i; j > 0 && M[j-1].Name < M[j].Name; j-- {
					M[j], M[j-1] = M[j-1], M[j] // Instant swap
				}
			}
		}
	} else if algo == 2 {
		// Selection Sort with Tuple Assignment
		if order == 1 {
			// Ascending (A-Z)
			for i := 0; i < n-1; i++ {
				minIdx := i
				for j := i + 1; j < n; j++ {
					if M[j].Name < M[minIdx].Name {
						minIdx = j
					}
				}
				M[i], M[minIdx] = M[minIdx], M[i] // Instant swap
			}
		} else if order == 2 {
			// Descending (Z-A)
			for i := 0; i < n-1; i++ {
				maxIdx := i
				for j := i + 1; j < n; j++ {
					if M[j].Name > M[maxIdx].Name {
						maxIdx = j
					}
				}
				M[i], M[maxIdx] = M[maxIdx], M[i] // Instant swap
			}
		}
	}
}

// STATISTIK PRESENSI

// Subprogram untuk menampilkan statistik persentase kehadiran per kelas, progres pertemuan, dan daftar mahasiswa dengan jumlah alpa terbanyak
func attendanceStatistics(M tabStudent, nM int, K tabAttendance, nK int) {
	if nM == 0 { // Guard Clause (Early Return)
		fmt.Println("\nBelum ada data mahasiswa untuk dihitung statistiknya.")
		return
	}

	var totalHadir, totalIzin, totalSakit, totalAlpa int = 0, 0, 0, 0
	var maxAlpaIdx int = 0
	var maxMeeting int = 0 // Variabel untuk melacak pertemuan terjauh

	// 1. Hitung akumulasi dari data mahasiswa
	for i := 0; i < nM; i++ {
		totalHadir += M[i].TH
		totalIzin += M[i].TI
		totalSakit += M[i].TS
		totalAlpa += M[i].TA

		if M[i].TA > M[maxAlpaIdx].TA {
			maxAlpaIdx = i
		}
	}

	// 2. Hitung progres pertemuan terjauh dari data log presensi
	for i := 0; i < nK; i++ {
		if K[i].Meeting > maxMeeting {
			maxMeeting = K[i].Meeting
		}
	}

	totalPencatatan := totalHadir + totalIzin + totalSakit + totalAlpa // Total keseluruhan entri data

	// 3. Tampilkan Dashboard Statistik
	fmt.Println("\n--- Ringkasan Statistik Kehadiran Mahasiswa ---")
	fmt.Printf("Total Pencatatan Presensi : %d entri data\n", totalPencatatan)
	fmt.Printf("Progres Pertemuan Terjauh : %d / 16 Pertemuan\n", maxMeeting)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("Total Seluruh Hadir : %d kali\n", totalHadir)
	fmt.Printf("Total Seluruh Izin  : %d kali\n", totalIzin)
	fmt.Printf("Total Seluruh Sakit : %d kali\n", totalSakit)
	fmt.Printf("Total Seluruh Alpa  : %d kali\n", totalAlpa)
	fmt.Println("-----------------------------------------------")

	// 4. Tampilkan Peringatan Alpa
	if M[maxAlpaIdx].TA > 0 {
		fmt.Println(" [!] PERINGATAN: Mahasiswa dengan Alpa Terbanyak")
		fmt.Printf("\tNama  : %s\n", M[maxAlpaIdx].Name)
		fmt.Printf("\tNIM   : %s\n", M[maxAlpaIdx].StdID)
		fmt.Printf("\tKelas : %s\n", M[maxAlpaIdx].Class)
		fmt.Printf("\tJumlah: %d Alpa\n", M[maxAlpaIdx].TA)
	} else {
		fmt.Println("Luar Biasa! Tidak ada mahasiswa yang memiliki catatan Alpa.")
	}
}

// SUBPROGRAM TAMBAHAN UNTUK PEMBANTU

// Subprogram pembantu untuk membaca input teks yang mengandung spasi
func readStringWithSpace(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Subprogram pembantu untuk mencari index mahasiswa berdasarkan NIM menggunakan algoritma Sequential Search (untuk keperluan update dan delete)
func studentIdxSearch(M tabStudent, n int, nim string) int {
	for i := 0; i < n; i++ {
		if M[i].StdID == nim {
			return i
		}
	}
	return -1
}

// Subprogram pembantu untuk mencari index data kehadiran berdasarkan NIM, Kode MK, dan Pertemuan menggunakan algoritma Sequential Search (untuk keperluan update)
func logIdxSearch(K tabAttendance, nK int, stdID string, subjectCode string, meeting int) int {
	for i := 0; i < nK; i++ {
		if K[i].StdID == stdID && K[i].SubjectCode == subjectCode && K[i].Meeting == meeting {
			return i
		}
	}
	return -1
}

// Subprogram pembantu untuk mencari index jadwal kuliah berdasarkan Kode Mata Kuliah dan Kelas menggunakan algoritma Sequential Search (untuk keperluan update dan delete)
func scheduleIdxSearch(J tabSchedule, n int, subjectCode string, class string) int {
	for i := 0; i < n; i++ {
		if J[i].SubjectCode == subjectCode && J[i].Class == class {
			return i
		}
	}
	return -1
}

// Subprogram pembantu untuk func searchByStdID dengan algoritma Binary Search (Insertion Sort)
func sortByStdID(M *tabStudent, n int) {
	for i := 1; i < n; i++ {
		for j := i; j > 0 && M[j-1].StdID > M[j].StdID; j-- {
			M[j], M[j-1] = M[j-1], M[j] // Instant swap
		}
	}
}

// Subprogram pembantu untuk func searchByAttendance dengan algoritma Binary Search (Insertion Sort)
func sortLogByAttendance(K *tabAttendance, nK int) {
	for i := 1; i < nK; i++ {
		for j := i; j > 0 && K[j-1].Status > K[j].Status; j-- {
			K[j], K[j-1] = K[j-1], K[j] // Instant swap
		}
	}
}

// Subprogram pembantu untuk mengambil nilai metrik presensi secara dinamis
func getMetricValue(s Student, metric int) int {
	switch metric {
	case 1:
		return s.TH // Total Hadir
	case 2:
		return s.TI // Total Izin
	case 3:
		return s.TS // Total Sakit
	case 4:
		return s.TA // Total Alpa
	default:
		return 0
	}
}

// MAIN LOGIC APLIKASI
func main() {
	var std tabStudent
	var sch tabSchedule
	var att tabAttendance
	var nMhs, nSch, nAtt int = 0, 0, 0
	var mainOpt, subOpt int
	reader := bufio.NewReader(os.Stdin) // Untuk membaca input string dengan spasi

	// Pemanggilan data dummy untuk testing
	injectDummyData(&std, &nMhs, &sch, &nSch, &att, &nAtt)

	for {
		clearScreen()                                                        // Bersihkan layar sebelum menampilkan menu utama
		printAppHeader(true, "Layanan Menu Aplikasi", nMhs, nSch, att, nAtt) // Memanggil Header Mode Utama (true)

		mainOpt = -1 // Reset pilihan menu utama setiap loop

		// Menampilkan daftar layanan menu utama
		fmt.Println("\t1. Kelola Data Mahasiswa")
		fmt.Println("\t2. Kelola Presensi Mahasiswa")
		fmt.Println("\t3. Kelola Jadwal Kuliah")
		fmt.Println("\t4. Cari Data Mahasiswa")
		fmt.Println("\t5. Urutkan Data Mahasiswa")
		fmt.Println("\t6. Statistik Kehadiran Mahasiswa")
		fmt.Println("\t0. Keluar Aplikasi")
		fmt.Print("Pilih menu layanan [0-6]: ")
		fmt.Scanln(&mainOpt)

		switch mainOpt {
		case 1: // Kelola Data Mahasiswa
			// Infinite loop sampai user memilih angka 0 untuk keluar dari sub-menu 1
			for {
				clearScreen()
				printAppHeader(false, "Menu Kelola Data Mahasiswa", nMhs, nSch, att, nAtt) // Memanggil Header Mode Sub-Menu (false)

				subOpt = -1 // Reset pilihan sub-menu setiap loop

				// Menampilkan daftar layanan menu sub-menu 1
				fmt.Println("\t1. Tambah Data Mahasiswa")
				fmt.Println("\t2. Lihat Data Mahasiswa")
				fmt.Println("\t3. Perbarui Data Mahasiswa")
				fmt.Println("\t4. Hapus Data Mahasiswa")
				fmt.Println("\t0. Kembali ke Menu Utama")
				fmt.Print("Pilih layanan menu [0-4]: ")
				fmt.Scanln(&subOpt)

				// Tombol keluar dari sub-menu
				if subOpt == 0 {
					break
				}

				switch subOpt {
				case 1: // Tambah Data Mahasiswa
					addMhs(&std, &nMhs, reader)
					pauseTerminal()
				case 2: // Lihat Data Mahasiswa
					readMhs(std, nMhs)
					pauseTerminal()
				case 3: // Perbarui Data Mahasiswa
					updateMhs(&std, nMhs, reader)
					pauseTerminal()
				case 4: // Hapus Data Mahasiswa
					deleteMhs(&std, &nMhs)
					pauseTerminal()
				default: // Menu invalid
					fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
					pauseTerminal()
				}
			}

		case 2: // Kelola Presensi Mahasiswa
			// Infinite loop sampai user memilih angka 0 untuk keluar dari sub-menu 2
			for {
				clearScreen()
				printAppHeader(false, "Menu Kelola Presensi Mahasiswa", nMhs, nSch, att, nAtt) // Memanggil Header Mode Sub-Menu (false)

				subOpt = -1 // Reset pilihan sub-menu setiap loop

				// Menampilkan daftar layanan menu sub-menu 2
				fmt.Println("\t1. Catat Presensi Baru")
				fmt.Println("\t2. Lihat Riwayat Presensi")
				fmt.Println("\t3. Perbarui Data Presensi")
				fmt.Println("\t0. Kembali ke Menu Utama")
				fmt.Print("Pilih layanan menu [0-3]: ")
				fmt.Scanln(&subOpt)

				// Tombol keluar dari sub-menu
				if subOpt == 0 {
					break
				}

				switch subOpt {
				case 1: // Catat Presensi Baru
					markAttendance(&std, nMhs, sch, nSch, &att, &nAtt)
					pauseTerminal()
				case 2: // Lihat Riwayat Presensi
					readAttendance(att, nAtt)
					pauseTerminal()
				case 3: // Perbarui Data Presensi
					updateAttendance(&att, nAtt, &std, nMhs)
					pauseTerminal()
				default: // Menu invalid
					fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
					pauseTerminal()
				}
			}

		case 3: // Kelola Jadwal Kuliah
			// Infinite loop sampai user memilih angka 0 untuk keluar dari sub-menu 3
			for {
				clearScreen()
				printAppHeader(false, "Menu Kelola Jadwal Kuliah", nMhs, nSch, att, nAtt) // Memanggil Header Mode Sub-Menu (false)

				subOpt = -1 // Reset pilihan sub-menu setiap loop

				// Menampilkan daftar layanan menu sub-menu 3
				fmt.Println("\t1. Tambah Jadwal Mata Kuliah")
				fmt.Println("\t2. Lihat Jadwal Kuliah")
				fmt.Println("\t3. Perbarui Jadwal Kuliah")
				fmt.Println("\t4. Hapus Jadwal Kuliah")
				fmt.Println("\t0. Kembali ke Menu Utama")
				fmt.Print("Pilih layanan menu [0-3]: ")
				fmt.Scanln(&subOpt)

				// Tombol keluar dari sub-menu
				if subOpt == 0 {
					break
				}

				switch subOpt {
				case 1: // Tambah Jadwal Mata Kuliah
					addSchedule(&sch, &nSch, reader)
					pauseTerminal()
				case 2: // Lihat Jadwal Kuliah
					readSchedule(sch, nSch)
					pauseTerminal()
				case 3: // Perbarui Jadwal Kuliah
					updateSchedule(&sch, nSch, reader)
					pauseTerminal()
				case 4: // Hapus Jadwal Kuliah
					deleteSchedule(&sch, &nSch)
					pauseTerminal()
				default: // Menu invalid
					fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
					pauseTerminal()
				}
			}

		case 4: // Cari Data Mahasiswa
			// Infinite loop sampai user memilih angka 0 untuk keluar dari sub-menu 4
			for {
				clearScreen()
				printAppHeader(false, "Menu Cari Data Mahasiswa", nMhs, nSch, att, nAtt) // Memanggil Header Mode Sub-Menu (false)

				subOpt = -1 // Reset pilihan sub-menu setiap loop

				// Menampilkan daftar layanan menu sub-menu 4
				fmt.Println("\t1. Cari berdasarkan Status Presensi")
				fmt.Println("\t2. Cari berdasarkan NIM")
				fmt.Println("\t0. Kembali ke Menu Utama")
				fmt.Print("Pilih layanan menu [0-2]: ")
				fmt.Scanln(&subOpt)

				// Tombol keluar dari sub-menu
				if subOpt == 0 {
					break
				}

				switch subOpt {
				case 1: // Cari berdasarkan Status Presensi
					var findStatus, statusTarget string
					fmt.Print("Masukkan status presensi (Hadir/Izin/Sakit/Alpa): ")
					fmt.Scanln(&findStatus)

					// Pilih status target berdasarkan input user dengan berbagai format yang valid (case-insensitive)
					switch findStatus {
					case "H", "h", "Hadir", "hadir", "HADIR": // Valid input untuk status Hadir
						statusTarget = "Hadir"
					case "I", "i", "Izin", "izin", "IZIN": // Valid input untuk status Izin
						statusTarget = "Izin"
					case "S", "s", "Sakit", "sakit", "SAKIT": // Valid input untuk status Sakit
						statusTarget = "Sakit"
					case "A", "a", "Alpa", "alpa", "ALPA": // Valid input untuk status Alpa
						statusTarget = "Alpa"
					default: // Input invalid
						statusTarget = ""
					}

					// Guard Clause untuk memeriksa validitas input status sebelum melanjutkan ke menu algoritma pencarian
					if statusTarget == "" {
						fmt.Println("\n[!] Status pencarian tidak valid. Format : H/h/Hadir/hadir/HADIR")
						pauseTerminal()
						continue // Langsung melompati menu algoritma dan kembali ke sub-menu 4
					}

					// Pilih algoritma pencarian
					var chooseAlgorithm int
					for {

						chooseAlgorithm = -1 // Reset pilihan algoritma setiap loop

						// Menu algoritma pencarian
						fmt.Println("\nAlgoritma Pencarian:")
						fmt.Println("\t1. Sequential Search")
						fmt.Println("\t2. Binary Search") // Dimodifikasi untuk mencari semua data dengan status yang sama
						fmt.Println("\t0. Batalkan Pencarian")
						fmt.Print("Pilih algoritma [0-2]: ")
						fmt.Scanln(&chooseAlgorithm)

						if chooseAlgorithm == 0 { // Tombol keluar dari menu algoritma pencarian
							fmt.Println("Pencarian dibatalkan. Kembali ke menu sebelumnya.")
							break
						} else if chooseAlgorithm == 1 || chooseAlgorithm == 2 { // Validasi input algoritma pencarian
							searchByAttendance(att, nAtt, std, nMhs, statusTarget, chooseAlgorithm) // Cari data mahasiswa berdasarkan status presensi dengan algoritma yang dipilih
							break
						} else { // Input invalid untuk pilihan algoritma
							fmt.Println("\n[!] Pilihan algoritma invalid. Silakan pilih angka 0, 1, atau 2.")
							continue
						}
					}
					pauseTerminal()

				case 2: // Cari berdasarkan NIM
					var findStdID string
					var chooseAlgorithm int
					fmt.Print("Masukkan NIM mahasiswa yang dicari: ")
					fmt.Scanln(&findStdID)
					for {

						chooseAlgorithm = -1 // Reset pilihan algoritma setiap loop

						// Menu algoritma pencarian
						fmt.Println("\nAlgoritma Pencarian:")
						fmt.Println("\t1. Sequential Search")
						fmt.Println("\t2. Binary Search")
						fmt.Println("\t0. Batalkan Pencarian")
						fmt.Print("Pilih algoritma [0-2]: ")
						fmt.Scanln(&chooseAlgorithm)

						if chooseAlgorithm == 0 { // Tombol keluar dari menu algoritma pencarian
							fmt.Println("Pencarian dibatalkan. Kembali ke menu sebelumnya.")
							break
						} else if chooseAlgorithm == 1 || chooseAlgorithm == 2 { // Validasi input algoritma pencarian
							if chooseAlgorithm == 2 { // Binary Search, data diurutkan terlebih dahulu berdasarkan NIM
								sortByStdID(&std, nMhs) // Mengurutkan data mahasiswa berdasarkan NIM sebelum melakukan Binary Search
							}
							index := searchByStdID(std, nMhs, findStdID, chooseAlgorithm) // Cari data mahasiswa berdasarkan NIM dengan algoritma yang dipilih
							if index != -1 {                                              // Data mahasiswa ditemukan
								fmt.Printf("\nData mahasiswa ditemukan:\nNama: %s\nNIM: %s\nKelas: %s\n", std[index].Name, std[index].StdID, std[index].Class)
								fmt.Printf("Presensi ===> Hadir: %d | Izin: %d | Sakit: %d | Alpa: %d\n", std[index].TH, std[index].TI, std[index].TS, std[index].TA)
							} else { // Data mahasiswa tidak ditemukan
								fmt.Printf("\nData mahasiswa dengan NIM %s tidak ditemukan.\n", findStdID)
							}
							break
						} else { // Input invalid untuk pilihan algoritma
							fmt.Println("\n[!] Pilihan algoritma invalid. Silakan pilih angka 0, 1, atau 2.")
							continue
						}
					}
					pauseTerminal()

				default: // Menu invalid
					fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
					pauseTerminal()
				}
			}

		case 5: // Urutkan Data Mahasiswa
			// Infinite loop sampai user memilih angka 0 untuk keluar dari sub-menu 5
			for {
				clearScreen()
				printAppHeader(false, "Menu Urutkan Data Mahasiswa", nMhs, nSch, att, nAtt) // Memanggil Header Mode Sub-Menu (false)

				subOpt = -1 // Reset pilihan sub-menu setiap loop

				// Menampilkan daftar layanan menu sub-menu 5
				fmt.Println("\t1. Urutkan berdasarkan Kategori Presensi")
				fmt.Println("\t2. Urutkan berdasarkan Nama")
				fmt.Println("\t3. Lihat Data Mahasiswa")
				fmt.Println("\t0. Kembali ke Menu Utama")
				fmt.Print("Pilih layanan menu [0-3]: ")
				fmt.Scanln(&subOpt)

				// Tombol keluar dari sub-menu
				if subOpt == 0 {
					break
				}

				switch subOpt {
				case 1: // Urut berdasarkan Kategori Presensi
					var metric, order, algo int
					for {

						metric = -1 // Reset pilihan kategori presensi setiap loop
						order = -1  // Reset pilihan arah urutan setiap loop
						algo = -1   // Reset pilihan algoritma setiap loop

						// Menu pilihan kategori presensi untuk diurutkan
						fmt.Println("\nPilih Kategori Presensi yang ingin diurutkan:")
						fmt.Println("\t1. Total Hadir (TH)")
						fmt.Println("\t2. Total Izin (TI)")
						fmt.Println("\t3. Total Sakit (TS)")
						fmt.Println("\t4. Total Alpa (TA)")
						fmt.Println("\t0. Batalkan Pengurutan")
						fmt.Print("Pilihan [0-4]: ")
						fmt.Scanln(&metric)

						if metric == 0 { // Tombol keluar dari menu kategori presensi
							fmt.Println("Pengurutan dibatalkan.")
							break
						} else if metric < 1 || metric > 4 { // Input invalid untuk pilihan kategori presensi
							fmt.Println("\n[!] Kategori invalid. Silakan ulangi.")
							continue
						}

						// Menu pilihan arah urutan
						fmt.Println("\nPilih Arah Urutan:")
						fmt.Println("\t1. Ascending (Terkecil ke Terbesar)")
						fmt.Println("\t2. Descending (Terbesar ke Terkecil)")
						fmt.Println("\t0. Batalkan Pengurutan")
						fmt.Print("Pilihan [0-2]: ")
						fmt.Scanln(&order)

						if order == 0 { // Tombol keluar dari menu arah urutan
							fmt.Println("Pengurutan dibatalkan.")
							break
						} else if order < 1 || order > 2 { // Input invalid untuk pilihan arah urutan
							fmt.Println("\n[!] Arah urutan invalid. Silakan ulangi.")
							continue
						}

						// Menu pilihan algoritma sorting
						fmt.Println("\nPilih Algoritma Sorting:")
						fmt.Println("\t1. Insertion Sort")
						fmt.Println("\t2. Selection Sort")
						fmt.Println("\t0. Batalkan Pengurutan")
						fmt.Print("Pilihan [0-2]: ")
						fmt.Scanln(&algo)

						if algo == 0 { // Tombol keluar dari menu algoritma sorting
							fmt.Println("Pengurutan dibatalkan.")
							break
						} else if algo < 1 || algo > 2 { // Input invalid untuk pilihan algoritma sorting
							fmt.Println("\n[!] Pilihan algoritma invalid. Silakan ulangi.")
							continue
						}

						sortByAttendance(&std, nMhs, metric, order, algo) // Urut data mahasiswa berdasarkan kategori presensi, arah urutan, dan algoritma yang dipilih
						fmt.Println("\nData mahasiswa berhasil diurutkan sesuai metrik presensi yang dipilih!")
						break
					}
					pauseTerminal()

				case 2: // Urut berdasarkan Nama
					var order, algo int
					for {

						order = -1 // Reset pilihan arah urutan setiap loop
						algo = -1  // Reset pilihan algoritma setiap loop

						// Menu pilihan arah urutan
						fmt.Println("\nPilih Arah Urutan Nama:")
						fmt.Println("\t1. Ascending (A-Z)")
						fmt.Println("\t2. Descending (Z-A)")
						fmt.Println("\t0. Batalkan Pengurutan")
						fmt.Print("Pilihan [0-2]: ")
						fmt.Scanln(&order)

						if order == 0 { // Tombol keluar dari menu arah urutan
							fmt.Println("Pengurutan dibatalkan.")
							break
						} else if order < 1 || order > 2 { // Input invalid untuk pilihan arah urutan
							fmt.Println("\n[!] Arah urutan invalid. Silakan ulangi.")
							continue
						}

						// Menu pilihan algoritma sorting
						fmt.Println("\nPilih Algoritma Sorting:")
						fmt.Println("\t1. Insertion Sort")
						fmt.Println("\t2. Selection Sort")
						fmt.Println("\t0. Batalkan Pengurutan")
						fmt.Print("Pilihan [0-2]: ")
						fmt.Scanln(&algo)

						if algo == 0 { // Tombol keluar dari menu algoritma sorting
							fmt.Println("Pengurutan dibatalkan.")
							break
						} else if algo < 1 || algo > 2 { // Input invalid untuk pilihan algoritma sorting
							fmt.Println("\n[!] Pilihan algoritma invalid. Silakan ulangi.")
							continue
						}

						sortByName(&std, nMhs, order, algo) // Urut data mahasiswa berdasarkan nama dengan arah urutan dan algoritma yang dipilih
						fmt.Println("\nData mahasiswa berhasil diurutkan berdasarkan nama!")
						break
					}
					pauseTerminal()

				case 3: // Lihat Data Mahasiswa
					readMhs(std, nMhs)
					pauseTerminal()

				default: // Menu invalid
					fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
					pauseTerminal()
				}
			}

		case 6: // Statistik Kehadiran Mahasiswa
			clearScreen()
			printAppHeader(false, "Menu Statistik Presensi Mahasiswa", nMhs, nSch, att, nAtt) // Memanggil Header Mode Sub-Menu (false)
			attendanceStatistics(std, nMhs, att, nAtt)                                        // Tampilkan statistik kehadiran mahasiswa
			pauseTerminal()

		case 0: // Keluar Aplikasi
			fmt.Println("\nTerima kasih telah menggunakan Aplikasi SiPresensi. Sayonara!")
			return

		default: // Menu invalid
			fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
			pauseTerminal()
		}
	}
}
