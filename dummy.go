package main

// Data Mahasiswa (Sengaja diacak posisinya agar fitur Sorting terlihat jelas saat didemokan)
var dummyStudents = []Student{
	{Name: "Joko Anwar", StdID: "130125300010", Class: "IF-49-01", TH: 6, TI: 1, TS: 0, TA: 6},
	{Name: "Citra Lestari", StdID: "130125300003", Class: "IF-49-01", TH: 8, TI: 2, TS: 2, TA: 1},
	{Name: "Indah Permata", StdID: "130125300009", Class: "IF-49-04", TH: 7, TI: 3, TS: 1, TA: 2},
	{Name: "Andi Wijaya", StdID: "130125300001", Class: "IF-49-01", TH: 10, TI: 1, TS: 0, TA: 2},
	{Name: "Fajar Hidayat", StdID: "130125300006", Class: "IF-49-03", TH: 11, TI: 1, TS: 0, TA: 1},
	{Name: "Dina Melati", StdID: "130125300004", Class: "IF-49-02", TH: 5, TI: 0, TS: 0, TA: 8},
	{Name: "Hadi Sucipto", StdID: "130125300008", Class: "IF-49-04", TH: 14, TI: 0, TS: 0, TA: 0},
	{Name: "Budi Santoso", StdID: "130125300002", Class: "IF-49-01", TH: 13, TI: 0, TS: 0, TA: 0},
	{Name: "Gita Gutawa", StdID: "130125300007", Class: "IF-49-03", TH: 9, TI: 0, TS: 3, TA: 1},
	{Name: "Eko Prasetyo", StdID: "130125300005", Class: "IF-49-02", TH: 12, TI: 0, TS: 1, TA: 0},
}

var dummySchedules = []Schedule{
	{SubjectCode: "CS101", SubjectName: "Algoritma dan Pemrograman", LectureCode: "TTG", LectureName: "Dr. Tatang", Class: "IF-49-01", Day: "Senin", Time: "08:00-10:00"},
	{SubjectCode: "CS102", SubjectName: "Struktur Data", LectureCode: "BDI", LectureName: "Prof. Budi", Class: "IF-49-02", Day: "Selasa", Time: "10:00-12:00"},
	{SubjectCode: "CS103", SubjectName: "Basis Data", LectureCode: "RNA", LectureName: "Bu Rina", Class: "IF-49-03", Day: "Rabu", Time: "13:00-15:00"},
}

// Data Log Kehadiran juga ikut diacak agar Binary Search Log (di Menu 3) bisa bekerja melakukan auto-sort
var dummyAttendances = []Attendance{
	{StdID: "130125300005", SubjectCode: "CS102", Status: "Hadir", Meeting: 1},
	{StdID: "130125300002", SubjectCode: "CS101", Status: "Hadir", Meeting: 1},
	{StdID: "130125300008", SubjectCode: "CS103", Status: "Hadir", Meeting: 1},
	{StdID: "130125300010", SubjectCode: "CS101", Status: "Alpa", Meeting: 1},
	{StdID: "130125300001", SubjectCode: "CS101", Status: "Hadir", Meeting: 1},
	{StdID: "130125300007", SubjectCode: "CS103", Status: "Sakit", Meeting: 1},
	{StdID: "130125300004", SubjectCode: "CS102", Status: "Alpa", Meeting: 1},
	{StdID: "130125300006", SubjectCode: "CS103", Status: "Hadir", Meeting: 1},
	{StdID: "130125300003", SubjectCode: "CS101", Status: "Izin", Meeting: 1},
	{StdID: "130125300009", SubjectCode: "CS103", Status: "Izin", Meeting: 1},
}

// Subprogram untuk menyuntikkan data di atas ke dalam aplikasi utamamu
func injectDummyData(M *tabStudent, nM *int, J *tabSchedule, nJ *int, K *tabAttendance, nK *int) {
	// Menuangkan data dummyStudents ke array M
	for i, student := range dummyStudents {
		(*M)[i] = student // <-- Menggunakan tanda kurung pointer
	}
	*nM = len(dummyStudents)

	// Menuangkan data dummySchedules ke array J
	for i, schedule := range dummySchedules {
		(*J)[i] = schedule // <-- Menggunakan tanda kurung pointer
	}
	*nJ = len(dummySchedules)

	// Menuangkan data dummyAttendances ke array K
	for i, attendance := range dummyAttendances {
		(*K)[i] = attendance // <-- Menggunakan tanda kurung pointer
	}
	*nK = len(dummyAttendances)
}
