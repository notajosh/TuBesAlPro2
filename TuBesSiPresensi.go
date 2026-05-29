package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// DEKLARASI KONSTANTA DAN TIPE DATA SERTA BENTUKAN DALAM VARIABEL GLOBAL

const NMHS int = 1000
const NSCH int = 200

type Student struct {
	Name, StdID, Class string
	TH, TI, TS, TA     int
}

type Schedule struct {
	SubjectCode, SubjectName, Lecture, Class, Day, Time string
}

type Attendance struct {
	StdID, SubjectCode, Status string
	Meeting                    int
}

type tabStudent [NMHS]Student
type tabSchedule [NSCH]Schedule
type tabAttendance [NMHS]Attendance

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

func pauseTerminal() {
	var enter string
	fmt.Print("Tekan Enter untuk melanjutkan...")
	fmt.Scanln(&enter)
}

// Subprogram pembantu untuk mencetak header aplikasi
func printAppHeader(isMainMenu bool, subMenuTitle string, nMhs int, nSch int) {
	fmt.Println("||==================================================================================================||")

	interiorWidth := 98 // Total ruang kosong di dalam kotak

	if isMainMenu {
		// MODE 1: Header Lengkap untuk Menu Utama
		title := "Selamat datang di Aplikasi SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa"
		padLeft := (interiorWidth - len(title)) / 2
		padRight := interiorWidth - len(title) - padLeft
		fmt.Printf("||%s%s%s||\n", strings.Repeat(" ", padLeft), title, strings.Repeat(" ", padRight))
		fmt.Println("||--------------------------------------------------------------------------------------------------||")

		// Cetak Dashboard
		mhsFilled := int((float64(nMhs) / float64(NMHS)) * 30)
		mhsBar := strings.Repeat("■", mhsFilled) + strings.Repeat(".", 30-mhsFilled)
		mhsText := fmt.Sprintf("Kapasitas Mahasiswa : [%s] %4d/%d Slot Terpakai", mhsBar, nMhs, NMHS)

		// PERBAIKAN DI SINI: Ganti NMHS menjadi NSCH untuk Jadwal
		schFilled := int((float64(nSch) / float64(NSCH)) * 30)
		schBar := strings.Repeat("■", schFilled) + strings.Repeat(".", 30-schFilled)
		schText := fmt.Sprintf("Kapasitas Jadwal    : [%s] %4d/%d Slot Terpakai", schBar, nSch, NSCH)

		fmt.Printf("||    %-90s    ||\n", mhsText)
		fmt.Printf("||    %-90s    ||\n", schText)
		fmt.Println("||--------------------------------------------------------------------------------------------------||")
	} else {
		// MODE 2: Header Ringkas untuk Sub-Menu (Tanpa Selamat Datang & Tanpa Dashboard)
		title := "SiPresensi : Sistem Monitoring Presensi dan Kehadiran Mahasiswa"
		padLeft := (interiorWidth - len(title)) / 2
		padRight := interiorWidth - len(title) - padLeft
		fmt.Printf("||%s%s%s||\n", strings.Repeat(" ", padLeft), title, strings.Repeat(" ", padRight))
		fmt.Println("||--------------------------------------------------------------------------------------------------||")
	}

	// Cetak Nama Menu Pilihan (Selalu Rata Tengah)
	padLeftSub := (interiorWidth - len(subMenuTitle)) / 2
	padRightSub := interiorWidth - len(subMenuTitle) - padLeftSub
	fmt.Printf("||%s%s%s||\n", strings.Repeat(" ", padLeftSub), subMenuTitle, strings.Repeat(" ", padRightSub))
	fmt.Println("||==================================================================================================||")
}

// CRUD DATA MAHASISWA

// Subprogram untuk menambahkan data mahasiswa
func addMhs(M *tabStudent, n *int, reader *bufio.Reader) {
	if *n < NMHS {
		fmt.Println("\n--- Tambah Data Mahasiswa ---")
		fmt.Print("Masukkan Nama Mahasiswa: ")
		M[*n].Name = readStringWithSpace(reader)

		// Validasi inputan NIM agar tepat 12 digit
		for {
			fmt.Print("Masukkan NIM Mahasiswa (Wajib 12 Digit): ")
			tempNIM := readStringWithSpace(reader)

			// Mengecek apakah panjang karakternya tepat 12
			if len(tempNIM) == 12 {
				M[*n].StdID = tempNIM
				break
			} else {
				fmt.Printf("[!] Error: NIM Anda %d digit. Harus tepat 12 digit! Silakan ulangi.\n\n", len(tempNIM))
			}
		}
		fmt.Print("Masukkan Kelas Mahasiswa: ")
		M[*n].Class = readStringWithSpace(reader)

		fmt.Printf("\nApakah Anda yakin menambahkan mahasiswa %s (%s)? (y/n): ", M[*n].Name, M[*n].StdID)
		confirm := readStringWithSpace(reader)
		if confirm == "y" || confirm == "Y" {
			M[*n].TH, M[*n].TI, M[*n].TS, M[*n].TA = 0, 0, 0, 0
			*n++
			fmt.Println("Data mahasiswa berhasil ditambahkan.")
		} else {
			fmt.Println("Penambahan data mahasiswa dibatalkan.")
		}
	} else {
		fmt.Println("Kapasitas data mahasiswa sudah penuh.")
	}
}

// Subprogram untuk membaca data mahasiswa dengan Tabel Dinamis (Auto-Resize)
func readMhs(M tabStudent, n int) {
	if n == 0 {
		fmt.Println("\nBelum ada data mahasiswa yang tersedia untuk ditampilkan.")
	} else {
		fmt.Print("\nApakah Anda yakin ingin melihat data seluruh mahasiswa? (y/n): ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" {
			// Cari teks terpanjang di setiap kolom untuk menentukan lebar kolom yang dinamis
			maxName := 20 // Jatah minimal 20 karakter untuk kolom Nama
			maxNim := 15  // Jatah minimal 15 karakter untuk kolom NIM
			maxClass := 8 // Jatah minimal 8 karakter untuk kolom Kelas

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

			for i := 0; i < n; i++ {
				fmt.Printf("| %-*s | %-*s | %-*s | %-5d | %-5d | %-5d | %-5d |\n",
					maxName, M[i].Name, maxNim, M[i].StdID, maxClass, M[i].Class, M[i].TH, M[i].TI, M[i].TS, M[i].TA)
			}
			fmt.Println(borderLine)

			// Footer dengan total mahasiswa terdaftar
			footerWidth := maxName + maxNim + maxClass + 8
			fmt.Printf("| %-*s | %-29d |\n", footerWidth-1, "Total Mahasiswa Terdaftar", n)

			bottomBorder := fmt.Sprintf("+%s+-------------------------------+", strings.Repeat("-", footerWidth))
			fmt.Println(bottomBorder)

		} else {
			fmt.Println("Menampilkan data mahasiswa dibatalkan.")
		}
	}
}

// Subprogram untuk memperbarui data mahasiswa
func updateMhs(M *tabStudent, n int, reader *bufio.Reader) {
	if n == 0 {
		fmt.Println("\nBelum ada data mahasiswa untuk diperbarui.")
	} else {
		var findStdID string
		found := false
		fmt.Println("\n--- Perbarui Data Mahasiswa ---")
		fmt.Print("Masukkan NIM mahasiswa yang ingin diperbarui: ")
		fmt.Scanln(&findStdID) // Pencarian biasa tetap pakai Scanln tidak apa-apa

		for i := 0; i < n && !found; i++ {
			if M[i].StdID == findStdID {
				found = true
				fmt.Printf("Data ditemukan:\nNama: %s\nNIM: %s\nKelas: %s\n", M[i].Name, M[i].StdID, M[i].Class)
				fmt.Println("Masukkan data baru (tekan Enter untuk mempertahankan data lama):")
				fmt.Print("Nama Mahasiswa: ")
				newName := readStringWithSpace(reader)
				fmt.Print("Kelas Mahasiswa: ")
				newClass := readStringWithSpace(reader)

				fmt.Printf("\nYakin menyimpan perubahan data %s? (y/n): ", M[i].StdID)
				confirm := readStringWithSpace(reader)
				if confirm == "y" || confirm == "Y" {
					if newName != "" {
						M[i].Name = newName
					}
					if newClass != "" {
						M[i].Class = newClass
					}
					fmt.Println("Data mahasiswa berhasil diperbarui.")
				} else {
					fmt.Println("Perubahan dibatalkan.")
				}
			}
		}
		if !found {
			fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", findStdID)
		}
	}
}

// Subprogram untuk menghapus data mahasiswa
func deleteMhs(M *tabStudent, n *int) {
	if *n == 0 {
		fmt.Println("\nBelum ada data mahasiswa yang tersedia untuk dihapus.")
	} else {
		var findStdID string
		var idx int
		found := false
		fmt.Println("\n--- Hapus Data Mahasiswa ---")
		fmt.Print("Masukkan NIM mahasiswa yang ingin dihapus: ")
		fmt.Scanln(&findStdID)
		for i := 0; i < *n && !found; i++ {
			if M[i].StdID == findStdID {
				found = true
				idx = i
				fmt.Printf("Data mahasiswa ditemukan:\nNama: %s\nNIM: %s\nKelas: %s\n", M[i].Name, M[i].StdID, M[i].Class)
				fmt.Printf("\nApakah Anda yakin untuk menghapus data mahasiswa %s? (y/n): ", M[i].StdID)
				var confirm string
				fmt.Scanln(&confirm)
				if confirm == "y" || confirm == "Y" {
					for j := idx; j < *n-1; j++ {
						M[j] = M[j+1]
					}
					*n--
					fmt.Println("Data mahasiswa berhasil dihapus.")
				} else {
					fmt.Println("Penghapusan data mahasiswa dibatalkan.")
				}
			} else {
				fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan. Tidak ada data yang dihapus.\n", findStdID)
			}
		}
	}
}

// PRESENSI MAHASISWA

// Subprogram untuk menandai kehadiran mahasiswa
func markAttendance(M *tabStudent, nMhs int, J tabSchedule, nJ int, K *tabAttendance, nK *int) {
	if nMhs == 0 {
		fmt.Println("\nBelum ada data mahasiswa yang tersedia untuk mencatat presensi.")
	} else if nJ == 0 {
		fmt.Println("\nBelum ada data jadwal yang tersedia untuk mencatat presensi.")
	} else {
		var targetClass, subjectCode string
		var meeting int
		fmt.Println("\n--- Catat Presensi Mahasiswa ---")
		fmt.Print("Masukkan Kelas: ")
		fmt.Scanln(&targetClass)
		fmt.Print("Masukkan Kode Mata Kuliah: ")
		fmt.Scanln(&subjectCode)
		fmt.Print("Masukkan Pertemuan Ke-: ")
		fmt.Scanln(&meeting)
		var scheduleFound bool = false
		var scheduleIdx int
		for i := 0; i < nJ && !scheduleFound; i++ {
			if J[i].Class == targetClass && J[i].SubjectCode == subjectCode {
				scheduleFound = true
				scheduleIdx = i
			}
		}
		if scheduleFound {
			fmt.Printf("Jadwal ditemukan!\nMata Kuliah: %s\nDosen: %s\nHari: %s\nJam: %s\n", J[scheduleIdx].SubjectName, J[scheduleIdx].Lecture, J[scheduleIdx].Day, J[scheduleIdx].Time)
			fmt.Println("Masukkan status kehadiran untuk setiap mahasiswa: Hadir (H), Izin (I), Sakit (S), Alpa (A)")
			var classFound bool = false
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
					if *nK < NMHS {
						K[*nK].StdID = M[i].StdID
						K[*nK].SubjectCode = subjectCode
						K[*nK].Status = statusFull
						K[*nK].Meeting = meeting
						*nK++
					}
				}
			}
			if classFound {
				fmt.Printf("Presensi pertemuan ke-%d untuk kelas %s pada mata kuliah %s berhasil dicatat.\n", meeting, targetClass, J[scheduleIdx].SubjectName)
			} else {
				fmt.Printf("Tidak ditemukan mahasiswa yang terdaftar di kelas %s. Presensi tidak dapat dicatat.\n", targetClass)
			}
		} else {
			fmt.Printf("Jadwal dengan Kelas %s dan Kode Mata Kuliah %s tidak ditemukan.\n", targetClass, subjectCode)
		}
	}
}

// Subprogram untuk membaca data kehadiran mahasiswa
func readAttendance(K tabAttendance, nK int) {
	if nK == 0 {
		fmt.Println("\nBelum ada data kehadiran yang tersedia untuk ditampilkan.")
	} else {
		// Cari teks terpanjang di setiap kolom untuk menentukan lebar kolom yang dinamis
		maxNim := 15    // Lebar minimal kolom NIM
		maxSubj := 15   // Lebar minimal kolom Kode MK
		maxStatus := 9  // Lebar minimal kolom Status
		maxMeeting := 9 // Lebar tetap untuk kolom Pertemuan

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

		for i := 0; i < nK; i++ {
			fmt.Printf("| %-*s | %-*s | %-*s | %-*d |\n",
				maxNim, K[i].StdID, maxSubj, K[i].SubjectCode, maxStatus, K[i].Status, maxMeeting, K[i].Meeting)
		}
		fmt.Println(borderLine)

		// Footer dengan total data kehadiran tercatat
		footerWidth := maxNim + maxSubj + maxStatus + 6
		fmt.Printf("| %-*s | %-*d |\n", footerWidth, "Total Data Kehadiran Tercatat", maxMeeting, nK)

		bottomBorder := fmt.Sprintf("+%s+%s+",
			strings.Repeat("-", footerWidth+2),
			strings.Repeat("-", maxMeeting+2),
		)
		fmt.Println(bottomBorder)
	}
}

// Subprogram untuk memperbarui data kehadiran mahasiswa
func updateAttendance(K *tabAttendance, nK int, M *tabStudent, nM int) {
	if nK == 0 {
		fmt.Println("\nBelum ada data kehadiran yang tersedia untuk diperbarui.")
	} else {
		var findStdID, findSubjectCode string
		var findMeeting int
		fmt.Println("\n --- Perbarui Presensi Mahasiswa ---")
		fmt.Print("Masukkan NIM mahasiswa: ")
		fmt.Scanln(&findStdID)
		fmt.Print("Masukkan Kode Mata Kuliah: ")
		fmt.Scanln(&findSubjectCode)
		fmt.Print("Masukkan Pertemuan Ke-: ")
		fmt.Scanln(&findMeeting)
		var attendanceIdx int = logIdxSearch(*K, nK, findStdID, findSubjectCode, findMeeting)
		if attendanceIdx != -1 {
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
			} else {
				var confirm string
				fmt.Printf("Apakah Anda yakin ingin mengubah presensi %s menjadi %s? (y/n):", findStdID, tempStatus)
				fmt.Scanln(&confirm)
				if confirm == "y" || confirm == "Y" {
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
				} else {
					fmt.Println("Pembaruan history presensi dibatalkan.")
				}
			}
		} else {
			fmt.Println("\nLog presensi tidak ditemukan. Pastikan NIM, Kode MK, dan Pertemuan benar.")
		}
	}
}

// SCHEDULE / Jadwal Kuliah Mahasiswa

// Subprogram untuk menambahkan jadwal mata kuliah
func addSchedule(J *tabSchedule, n *int, reader *bufio.Reader) {
	if *n < NMHS {
		fmt.Println("\n --- Tambah Jadwal Mata Kuliah ---")

		fmt.Print("Masukkan Kode Mata Kuliah         : ")
		J[*n].SubjectCode = readStringWithSpace(reader)

		fmt.Print("Masukkan Nama Mata Kuliah         : ")
		J[*n].SubjectName = readStringWithSpace(reader)

		fmt.Print("Masukkan Nama Dosen (Boleh Spasi) : ")
		J[*n].Lecture = readStringWithSpace(reader)

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
	} else {
		fmt.Println("Kapasitas data jadwal sudah penuh.")
	}
}

// SEARCHING

// Subprogram untuk mencari data mahasiswa berdasarkan status presensi
func searchByAttendance(K tabAttendance, nK int, M tabStudent, nM int, status string, chosenAlgorithm int) {
	if nK == 0 {
		fmt.Println("\nBelum ada history presensi yang tercatat.")
	} else {
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
}

// Subprogram untuk mencari data mahasiswa berdasarkan NIM
func searchByStdID(M tabStudent, n int, stdID string, order int) int {
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
		// Insertion Sort
		if order == 1 {
			// Asc
			for i := 1; i < n; i++ {
				temp := M[i]
				j := i - 1
				for j >= 0 && getMetricValue(M[j], metric) > getMetricValue(temp, metric) {
					M[j+1] = M[j]
					j--
				}
				M[j+1] = temp
			}
		} else if order == 2 {
			// Desc
			for i := 1; i < n; i++ {
				temp := M[i]
				j := i - 1
				for j >= 0 && getMetricValue(M[j], metric) < getMetricValue(temp, metric) {
					M[j+1] = M[j]
					j--
				}
				M[j+1] = temp
			}
		}
	} else if algo == 2 {
		// Selection Sort
		if order == 1 {
			// Asc
			for i := 0; i < n-1; i++ {
				minIdx := i
				for j := i + 1; j < n; j++ {
					if getMetricValue(M[j], metric) < getMetricValue(M[minIdx], metric) {
						minIdx = j
					}
				}
				temp := M[i]
				M[i] = M[minIdx]
				M[minIdx] = temp
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
				temp := M[i]
				M[i] = M[maxIdx]
				M[maxIdx] = temp
			}
		}
	}
}

// Subprogram untuk mengurutkan data mahasiswa berdasarkan nama
func sortByName(M *tabStudent, n int, order int, algo int) {
	if algo == 1 {
		// Insertion Sort
		if order == 1 {
			// Ascending (A-Z)
			for i := 1; i < n; i++ {
				temp := M[i]
				j := i - 1
				for j >= 0 && M[j].Name > temp.Name {
					M[j+1] = M[j]
					j--
				}
				M[j+1] = temp
			}
		} else if order == 2 {
			// Descending (Z-A)
			for i := 1; i < n; i++ {
				temp := M[i]
				j := i - 1
				for j >= 0 && M[j].Name < temp.Name {
					M[j+1] = M[j]
					j--
				}
				M[j+1] = temp
			}
		}
	} else if algo == 2 {
		// Selection Sort
		if order == 1 {
			// Ascending (A-Z)
			for i := 0; i < n-1; i++ {
				minIdx := i
				for j := i + 1; j < n; j++ {
					if M[j].Name < M[minIdx].Name {
						minIdx = j
					}
				}
				temp := M[i]
				M[i] = M[minIdx]
				M[minIdx] = temp
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
				temp := M[i]
				M[i] = M[maxIdx]
				M[maxIdx] = temp
			}
		}
	}
}

// STATISTIK PRESENSI

// Subprogram untuk menampilkan statistik persentase kehadiran per kelas dan daftar mahasiswa dengan jumlah alpa terbanyak
func attendanceStatistics(M tabStudent, n int) {
	if n == 0 {
		fmt.Println("\nBelum ada data mahasiswa untuk dihitung statistiknya.")
	} else {
		var totalHadir, totalIzin, totalSakit, totalAlpa int = 0, 0, 0, 0
		var maxAlpaIdx int = 0

		for i := 0; i < n; i++ {
			totalHadir += M[i].TH
			totalIzin += M[i].TI
			totalSakit += M[i].TS
			totalAlpa += M[i].TA

			if M[i].TA > M[maxAlpaIdx].TA {
				maxAlpaIdx = i
			}
		}

		fmt.Println("\n--- Ringkasan Statistik Kehadiran Mahasiswa ---")
		fmt.Printf("Total Seluruh Hadir : %d pertemuan\n", totalHadir)
		fmt.Printf("Total Seluruh Izin  : %d kali\n", totalIzin)
		fmt.Printf("Total Seluruh Sakit : %d kali\n", totalSakit)
		fmt.Printf("Total Seluruh Alpa  : %d kali\n", totalAlpa)
		fmt.Println("-----------------------------------------------")

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
}

// SUBPROGRAM TAMBAHAN UNTUK PEMBANTU

// Subprogram pembantu untuk membaca input teks yang mengandung spasi
func readStringWithSpace(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Subprogram pembantu untuk mencari data presensi mahasiswa berdasarkan NIM, Kode Mata Kuliah, dan Pertemuan menggunakan algoritma Sequential Search
func logIdxSearch(K tabAttendance, nK int, stdID string, subjectCode string, meeting int) int {
	found := false
	idx := -1
	for i := 0; i < nK && !found; i++ {
		if K[i].StdID == stdID && K[i].SubjectCode == subjectCode && K[i].Meeting == meeting {
			found = true
			idx = i
		}
	}
	return idx
}

// Subprogram pembantu untuk func searchByStdID dengan algoritma Binary Search
func sortByStdID(M *tabStudent, n int) {
	for i := 1; i < n; i++ {
		temp := M[i]
		j := i - 1
		for j >= 0 && M[j].StdID > temp.StdID {
			M[j+1] = M[j]
			j--
		}
		M[j+1] = temp
	}
}

// Subprogram pembantu untuk func searchByAttendance dengan algoritma Binary Search
func sortLogByAttendance(K *tabAttendance, nK int) {
	for i := 1; i < nK; i++ {
		temp := K[i]
		j := i - 1
		for j >= 0 && K[j].Status > temp.Status {
			K[j+1] = K[j]
			j--
		}
		K[j+1] = temp
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
		clearScreen()
		printAppHeader(true, "Layanan Menu Aplikasi", nMhs, nSch) // Memanggil Header Mode Utama (true)

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
		case 1:
			clearScreen()
			printAppHeader(false, "Menu Kelola Data Mahasiswa", nMhs, nSch) // Memanggil Header Mode Ringkas (false)

			fmt.Println("\t1. Tambah Data Mahasiswa")
			fmt.Println("\t2. Lihat Data Mahasiswa")
			fmt.Println("\t3. Perbarui Data Mahasiswa")
			fmt.Println("\t4. Hapus Data Mahasiswa")
			fmt.Println("\t0. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu [0-4]: ")
			fmt.Scanln(&subOpt)
			switch subOpt {
			case 1:
				addMhs(&std, &nMhs, reader)
				pauseTerminal()
			case 2:
				readMhs(std, nMhs)
				pauseTerminal()
			case 3:
				updateMhs(&std, nMhs, reader)
				pauseTerminal()
			case 4:
				deleteMhs(&std, &nMhs)
				pauseTerminal()
			case 0:
				continue
			default:
				fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
				pauseTerminal()
				continue
			}
		case 2:
			clearScreen()
			printAppHeader(false, "Menu Kelola Presensi Mahasiswa", nMhs, nSch)

			fmt.Println("\t1. Catat Presensi Baru")
			fmt.Println("\t2. Lihat Riwayat Presensi")
			fmt.Println("\t3. Perbarui Data Presensi")
			fmt.Println("\t0. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu [0-3]: ")
			fmt.Scanln(&subOpt)
			switch subOpt {
			case 1:
				markAttendance(&std, nMhs, sch, nSch, &att, &nAtt)
				pauseTerminal()
			case 2:
				readAttendance(att, nAtt)
				pauseTerminal()
			case 3:
				updateAttendance(&att, nAtt, &std, nMhs)
				pauseTerminal()
			case 0:
				continue
			default:
				fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
				pauseTerminal()
				continue
			}
		case 3:
			clearScreen()
			printAppHeader(false, "Menu Kelola Jadwal Kuliah", nMhs, nSch)

			fmt.Println("\t1. Tambah Jadwal Mata Kuliah")
			fmt.Println("\t0. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu [0-1]: ")
			fmt.Scanln(&subOpt)
			switch subOpt {
			case 1:
				addSchedule(&sch, &nSch, reader)
				pauseTerminal()
			case 0:
				continue
			default:
				fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
				pauseTerminal()
				continue
			}
		case 4:
			clearScreen()
			printAppHeader(false, "Menu Cari Data Mahasiswa", nMhs, nSch)

			fmt.Println("\t1. Cari berdasarkan Status Presensi")
			fmt.Println("\t2. Cari berdasarkan NIM")
			fmt.Println("\t0. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu [0-2]: ")
			fmt.Scanln(&subOpt)
			switch subOpt {
			case 1:
				var findStatus string
				var chooseAlgorithm int
				fmt.Print("Masukkan status presensi (Hadir/Izin/Sakit/Alpa): ")
				fmt.Scanln(&findStatus)

				for {
					fmt.Println("\nAlgoritma Pencarian:")
					fmt.Println("\t1. Sequential Search")
					fmt.Println("\t2. Binary Search (dimodifikasi untuk duplikasi)")
					fmt.Println("\t0. Batalkan Pencarian")
					fmt.Print("Pilih algoritma [0-2]: ")
					fmt.Scanln(&chooseAlgorithm)

					if chooseAlgorithm == 0 {
						fmt.Println("Pencarian dibatalkan. Kembali ke menu sebelumnya.")
						break
					} else if chooseAlgorithm == 1 || chooseAlgorithm == 2 {
						searchByAttendance(att, nAtt, std, nMhs, findStatus, chooseAlgorithm)
						break
					} else {
						fmt.Println("\n[!] Pilihan algoritma invalid. Silakan pilih angka 0, 1, atau 2.")
						continue
					}
				}
				pauseTerminal()
			case 2:
				var findStdID string
				var chooseAlgorithm int
				fmt.Print("Masukkan NIM mahasiswa yang dicari: ")
				fmt.Scanln(&findStdID)
				for {
					fmt.Println("\nAlgoritma Pencarian:")
					fmt.Println("\t1. Sequential Search")
					fmt.Println("\t2. Binary Search")
					fmt.Println("\t0. Batalkan Pencarian")
					fmt.Print("Pilih algoritma [0-2]: ")
					fmt.Scanln(&chooseAlgorithm)
					if chooseAlgorithm == 0 {
						fmt.Println("Pencarian dibatalkan. Kembali ke menu sebelumnya.")
						break
					} else if chooseAlgorithm == 1 || chooseAlgorithm == 2 {
						if chooseAlgorithm == 2 {
							sortByStdID(&std, nMhs)
						}
						index := searchByStdID(std, nMhs, findStdID, chooseAlgorithm)
						if index != -1 {
							fmt.Printf("\nData mahasiswa ditemukan:\nNama: %s\nNIM: %s\nKelas: %s\n", std[index].Name, std[index].StdID, std[index].Class)
							fmt.Printf("Presensi ===> Hadir: %d | Izin: %d | Sakit: %d | Alpa: %d\n", std[index].TH, std[index].TI, std[index].TS, std[index].TA)
						} else {
							fmt.Printf("\nData mahasiswa dengan NIM %s tidak ditemukan.\n", findStdID)
						}
						break
					} else {
						fmt.Println("\n[!] Pilihan algoritma invalid. Silakan pilih angka 0, 1, atau 2.")
						continue
					}
				}
				pauseTerminal()
			case 0:
				continue
			default:
				fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
				pauseTerminal()
				continue
			}
		case 5:
			clearScreen()
			printAppHeader(false, "Menu Urutkan Data Mahasiswa", nMhs, nSch)

			fmt.Println("\t1. Urutkan berdasarkan Kategori Presensi")
			fmt.Println("\t2. Urutkan berdasarkan Nama")
			fmt.Println("\t3. Lihat Data Mahasiswa")
			fmt.Println("\t0. Kembali ke Menu Utama")
			fmt.Print("Pilih layanan menu [0-3]: ")
			fmt.Scanln(&subOpt)

			switch subOpt {
			case 1:
				var metric, order, algo int
				for {
					fmt.Println("\nPilih Kategori Presensi yang ingin diurutkan:")
					fmt.Println("\t1. Total Hadir (TH)")
					fmt.Println("\t2. Total Izin (TI)")
					fmt.Println("\t3. Total Sakit (TS)")
					fmt.Println("\t4. Total Alpa (TA)")
					fmt.Println("\t0. Batalkan Pengurutan")
					fmt.Print("Pilihan [0-4]: ")
					fmt.Scanln(&metric)

					if metric == 0 {
						fmt.Println("Pengurutan dibatalkan.")
						break
					} else if metric < 1 || metric > 4 {
						fmt.Println("\n[!] Kategori invalid. Silakan ulangi.")
						continue
					}

					fmt.Println("\nPilih Arah Urutan:")
					fmt.Println("\t1. Ascending (Terkecil ke Terbesar)")
					fmt.Println("\t2. Descending (Terbesar ke Terkecil)")
					fmt.Println("\t0. Batalkan Pengurutan")
					fmt.Print("Pilihan [0-2]: ")
					fmt.Scanln(&order)

					if order == 0 {
						fmt.Println("Pengurutan dibatalkan.")
						break
					} else if order < 1 || order > 2 {
						fmt.Println("\n[!] Arah urutan invalid. Silakan ulangi.")
						continue
					}

					fmt.Println("\nPilih Algoritma Sorting:")
					fmt.Println("\t1. Insertion Sort")
					fmt.Println("\t2. Selection Sort")
					fmt.Println("\t0. Batalkan Pengurutan")
					fmt.Print("Pilihan [0-2]: ")
					fmt.Scanln(&algo)

					if algo == 0 {
						fmt.Println("Pengurutan dibatalkan.")
						break
					} else if algo < 1 || algo > 2 {
						fmt.Println("\n[!] Pilihan algoritma invalid. Silakan ulangi.")
						continue
					}

					sortByAttendance(&std, nMhs, metric, order, algo)
					fmt.Println("\nData mahasiswa berhasil diurutkan sesuai metrik presensi yang dipilih!")
					break
				}
				pauseTerminal()

			case 2:
				var order, algo int
				for {
					fmt.Println("\nPilih Arah Urutan Nama:")
					fmt.Println("\t1. Ascending (A-Z)")
					fmt.Println("\t2. Descending (Z-A)")
					fmt.Println("\t0. Batalkan Pengurutan")
					fmt.Print("Pilihan [0-2]: ")
					fmt.Scanln(&order)

					if order == 0 {
						fmt.Println("Pengurutan dibatalkan.")
						break
					} else if order < 1 || order > 2 {
						fmt.Println("\n[!] Arah urutan invalid. Silakan ulangi.")
						continue
					}

					fmt.Println("\nPilih Algoritma Sorting:")
					fmt.Println("\t1. Insertion Sort")
					fmt.Println("\t2. Selection Sort")
					fmt.Println("\t0. Batalkan Pengurutan")
					fmt.Print("Pilihan [0-2]: ")
					fmt.Scanln(&algo)

					if algo == 0 {
						fmt.Println("Pengurutan dibatalkan.")
						break
					} else if algo < 1 || algo > 2 {
						fmt.Println("\n[!] Pilihan algoritma invalid. Silakan ulangi.")
						continue
					}

					sortByName(&std, nMhs, order, algo)
					fmt.Println("\nData mahasiswa berhasil diurutkan berdasarkan nama!")
					break
				}
				pauseTerminal()
			case 3:
				readMhs(std, nMhs)
				pauseTerminal()
			case 0:
				continue
			default:
				fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
				pauseTerminal()
				continue
			}
		case 6:
			clearScreen()
			printAppHeader(false, "Menu Statistik Presensi Mahasiswa", nMhs, nSch)
			attendanceStatistics(std, nMhs)
			pauseTerminal()
		case 0:
			fmt.Println("\nTerima kasih telah menggunakan Aplikasi SiPresensi. Sayonara!")
			return
		default:
			fmt.Println("Menu invalid, silakan pilih menu yang tersedia.")
			pauseTerminal()
			continue
		}
	}
}
