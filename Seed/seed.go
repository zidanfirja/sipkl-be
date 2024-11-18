package Seed

import (
	"fmt"
	DB "go-gin-mysql/Database"
	"go-gin-mysql/Models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedRole() {
	var count int64

	DB.Database.Model(&Models.Role{}).Count(&count)
	if count > 0 {
		return
	}

	// dummy data
	roles := []*Models.Role{
		{Nama: "Pembimbing", Aktif: true},
		{Nama: "Fasilitator", Aktif: true},
	}

	if err := DB.Database.Session(&gorm.Session{PrepareStmt: false}).Create(&roles).Error; err != nil {
		log.Fatal("Gagal mengisi data dummy:", err)
	}
	fmt.Println("Data dummy berhasil ditambahkan")

}

func SeedIndustri() {
	var count int64
	DB.Database.Model(&Models.Industri{}).Count(&count)
	if count > 0 {
		fmt.Println("Data Industri sudah ada, tidak menambahkan lagi")
		return
	}

	// Data dummy Industri
	industriList := []Models.Industri{
		{Nama: "PT. Maju Sejahtera", Alamat: "Jl. Merdeka No. 1", Jurusan: "TKJ", CreatedAt: time.Now()},
		{Nama: "CV. Karya Mandiri", Alamat: "Jl. Industri No. 45", Jurusan: "RPL", CreatedAt: time.Now()},
		{Nama: "PT. Teknologi Nusantara", Alamat: "Jl. Sudirman No. 22", Jurusan: "TKJ", CreatedAt: time.Now()},
		{Nama: "PT. Bangkit", Alamat: "Jl. Bandung No. 21", Jurusan: "RPL", CreatedAt: time.Now()},
		{Nama: "CV. Teknologi Indah", Alamat: "Jl. Raya No. 88", Jurusan: "TKJ", CreatedAt: time.Now()},
		{Nama: "PT. Sinergi Global", Alamat: "Jl. Merdeka No. 5", Jurusan: "RPL", CreatedAt: time.Now()},
		{Nama: "PT. Inovasi Cerdas", Alamat: "Jl. Soekarno Hatta No. 10", Jurusan: "TKJ", CreatedAt: time.Now()},
		{Nama: "CV. Digital Kreatif", Alamat: "Jl. Pahlawan No. 7", Jurusan: "RPL", CreatedAt: time.Now()},
	}

	// Insert data dummy
	if err := DB.Database.Session(&gorm.Session{PrepareStmt: false}).Create(&industriList).Error; err != nil {
		log.Fatal("Gagal mengisi data Industri:", err)
	}
	fmt.Println("Data Industri berhasil ditambahkan")
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

// Fungsi untuk seed data Pegawai
func SeedPegawai() {
	var count int64
	DB.Database.Model(&Models.Pegawai{}).Count(&count)
	if count > 0 {
		fmt.Println("Data Pegawai sudah ada, tidak menambahkan lagi")
		return
	}

	// Data dummy Pegawai
	pegawaiList := []Models.Pegawai{
		{
			IdPegawai: "1001",
			Nama:      "John Doe",
			Email:     "johndoe@example.com",
			Password:  HashPassword("password123"),
			Aktif:     true,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: "1002",
			Nama:      "Jane Smith",
			Email:     "janesmith@example.com",
			Password:  HashPassword("password123"),
			Aktif:     true,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: "1003",
			Nama:      "David Johnson",
			Email:     "davidj@example.com",
			Password:  HashPassword("password123"),
			Aktif:     false,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: "1004",
			Nama:      "Alice Walker",
			Email:     "alicewalker@example.com",
			Password:  HashPassword("password123"),
			Aktif:     true,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: "1005",
			Nama:      "Bob Brown",
			Email:     "bobbrown@example.com",
			Password:  HashPassword("password123"),
			Aktif:     true,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: "1006",
			Nama:      "Charlie Davis",
			Email:     "charliedavis@example.com",
			Password:  HashPassword("password123"),
			Aktif:     true,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: "1007",
			Nama:      "Eve Adams",
			Email:     "eveadams@example.com",
			Password:  HashPassword("password123"),
			Aktif:     false,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: "1008",
			Nama:      "Grace Lee",
			Email:     "gracelee@example.com",
			Password:  HashPassword("password123"),
			Aktif:     true,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: "1009",
			Nama:      "Henry Miller",
			Email:     "henrymiller@example.com",
			Password:  HashPassword("password123"),
			Aktif:     true,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: "1010",
			Nama:      "Isabella Taylor",
			Email:     "isabellataylor@example.com",
			Password:  HashPassword("password123"),
			Aktif:     false,
			CreatedAt: time.Now(),
		},
	}

	// Insert data dummy
	if err := DB.Database.Session(&gorm.Session{PrepareStmt: false}).Create(&pegawaiList).Error; err != nil {
		log.Fatal("Gagal mengisi data Pegawai:", err)
	}
	fmt.Println("Data Pegawai berhasil ditambahkan")
}

func SeedKonfigurasiRoles() {
	var count int64
	DB.Database.Model(&Models.KonfigurasiRoles{}).Count(&count)
	if count > 0 {
		fmt.Println("Data KonfigurasiRoles sudah ada, tidak menambahkan lagi")
		return
	}

	// Ambil data Pegawai
	var pegawaiList []Models.Pegawai
	if err := DB.Database.Find(&pegawaiList).Error; err != nil {
		log.Fatal("Gagal mengambil data Pegawai:", err)
	}

	// Ambil data Role
	var roleList []Models.Role
	if err := DB.Database.Find(&roleList).Error; err != nil {
		log.Fatal("Gagal mengambil data Role:", err)
	}

	// Pastikan data pegawai dan role ada sebelum menambahkan data ke KonfigurasiRoles
	if len(pegawaiList) == 0 || len(roleList) == 0 {
		log.Fatal("Tidak ada data Pegawai atau Role untuk membuat KonfigurasiRoles")
	}

	// Data dummy KonfigurasiRoles
	konfigurasiRolesList := []Models.KonfigurasiRoles{
		{
			FKIdPegawai: pegawaiList[0].ID, // John Doe
			FKIdRole:    &roleList[0].ID,   // Pembimbing
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[0].ID, // John Doe
			FKIdRole:    &roleList[1].ID,   // Fasilitator
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[1].ID, // Jane Smith
			FKIdRole:    &roleList[1].ID,   // Fasilitator
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[2].ID, // David Johnson
			FKIdRole:    &roleList[0].ID,   // Pembimbing
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[3].ID, // Alice Walker
			FKIdRole:    &roleList[1].ID,   // Fasilitator
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[4].ID, // Bob Brown
			FKIdRole:    &roleList[0].ID,   // Pembimbing
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[5].ID, // Charlie Davis
			FKIdRole:    &roleList[1].ID,   // Fasilitator
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[6].ID, // Eve Adams
			FKIdRole:    &roleList[0].ID,   // Pembimbing
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[6].ID, // Eve Adams
			FKIdRole:    &roleList[1].ID,   // Fasilitator
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[7].ID, // Grace Lee
			FKIdRole:    &roleList[0].ID,   // Pembimbing
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[8].ID, // Henry Miller
			FKIdRole:    &roleList[1].ID,   // Fasilitator
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[9].ID, // Isabella Taylor
			FKIdRole:    &roleList[1].ID,   // Fasilitator
			CreatedAt:   time.Now(),
		},
		// Dosen yang memiliki dua peran: Pembimbing dan Fasilitator
		{
			FKIdPegawai: pegawaiList[4].ID, // Bob Brown
			FKIdRole:    &roleList[1].ID,   // Fasilitator
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[5].ID, // Charlie Davis
			FKIdRole:    &roleList[0].ID,   // Pembimbing
			CreatedAt:   time.Now(),
		},
	}

	// Insert data dummy
	if err := DB.Database.Session(&gorm.Session{PrepareStmt: false}).Create(&konfigurasiRolesList).Error; err != nil {
		log.Fatal("Gagal mengisi data KonfigurasiRoles:", err)
	}
	fmt.Println("Data KonfigurasiRoles berhasil ditambahkan")
}

func SeedDataSiswa() {
	var count int64
	DB.Database.Model(&Models.DataSiswa{}).Count(&count)
	if count > 0 {
		fmt.Println("Data DataSiswa sudah ada, tidak menambahkan lagi")
		return
	}

	// Ambil data Pegawai dan Industri untuk relasi
	var pegawaiList []Models.Pegawai
	if err := DB.Database.Find(&pegawaiList).Error; err != nil {
		log.Fatal("Gagal mengambil data Pegawai:", err)
	}

	var industriList []Models.Industri
	if err := DB.Database.Find(&industriList).Error; err != nil {
		log.Fatal("Gagal mengambil data Industri:", err)
	}

	// Pastikan data pegawai dan industri ada
	if len(pegawaiList) < 2 || len(industriList) == 0 {
		log.Fatal("Tidak cukup data Pegawai atau Industri untuk membuat DataSiswa")
	}
	var waktu_sekarang = time.Now()
	// Data dummy DataSiswa
	dataSiswaList := []Models.DataSiswa{
		{
			NIS:                         "12345",
			Nama:                        "John Doe",
			Kelas:                       "12",
			Jurusan:                     "RPL",
			Rombel:                      "A",
			TanggalMasuk:                &waktu_sekarang,
			TanggalKeluar:               &waktu_sekarang,
			Aktif:                       true,
			Email:                       "johndoe@example.com",
			Password:                    "password123",
			FKIdPembimbing:              pegawaiList[0].ID, // Pembimbing (John Doe)
			FKIdFasilitator:             pegawaiList[1].ID, // Fasilitator (Jane Smith)
			FKIdIndustri:                industriList[0].ID,
			NilaiSoftskillIndustri:      85,
			NilaiSoftskillFasilitator:   90,
			NilaiHardskillIndustri:      88,
			NilaiHardskillPembimbing:    92,
			NilaiKemandirianFasilitator: 89,
			NilaiPengujianPembimbing:    87,
			CreatedAt:                   time.Now(),
		},
		{
			NIS:                         "67890",
			Nama:                        "Jane Smith",
			Kelas:                       "11",
			Jurusan:                     "TKJ",
			Rombel:                      "B",
			TanggalMasuk:                &waktu_sekarang,
			TanggalKeluar:               &waktu_sekarang,
			Aktif:                       true,
			Email:                       "janesmith@example.com",
			Password:                    "password456",
			FKIdPembimbing:              pegawaiList[0].ID, // Pembimbing (John Doe)
			FKIdFasilitator:             pegawaiList[1].ID, // Fasilitator (Jane Smith)
			FKIdIndustri:                industriList[1].ID,
			NilaiSoftskillIndustri:      80,
			NilaiSoftskillFasilitator:   85,
			NilaiHardskillIndustri:      82,
			NilaiHardskillPembimbing:    86,
			NilaiKemandirianFasilitator: 88,
			NilaiPengujianPembimbing:    84,
			CreatedAt:                   time.Now(),
		},
		{
			NIS:                         "28494",
			Nama:                        "Dariam",
			Kelas:                       "12",
			Jurusan:                     "RPL",
			Rombel:                      "B",
			TanggalMasuk:                &waktu_sekarang,
			TanggalKeluar:               &waktu_sekarang,
			Aktif:                       true,
			Email:                       "dariamsa@example.com",
			Password:                    "password456",
			FKIdPembimbing:              pegawaiList[2].ID, // Pembimbing (David Johnson)
			FKIdFasilitator:             pegawaiList[1].ID, // Fasilitator (Jane Smith)
			FKIdIndustri:                industriList[2].ID,
			NilaiSoftskillIndustri:      80,
			NilaiSoftskillFasilitator:   85,
			NilaiHardskillIndustri:      82,
			NilaiHardskillPembimbing:    86,
			NilaiKemandirianFasilitator: 88,
			NilaiPengujianPembimbing:    84,
			CreatedAt:                   time.Now(),
		},
		{
			NIS:                         "45678",
			Nama:                        "Alice Brown",
			Kelas:                       "12",
			Jurusan:                     "TKJ",
			Rombel:                      "A",
			TanggalMasuk:                &waktu_sekarang,
			TanggalKeluar:               &waktu_sekarang,
			Aktif:                       true,
			Email:                       "alicebrown@example.com",
			Password:                    "password789",
			FKIdPembimbing:              pegawaiList[3].ID, // Pembimbing (Alice Walker)
			FKIdFasilitator:             pegawaiList[4].ID, // Fasilitator (Bob Brown)
			FKIdIndustri:                industriList[3].ID,
			NilaiSoftskillIndustri:      82,
			NilaiSoftskillFasilitator:   88,
			NilaiHardskillIndustri:      85,
			NilaiHardskillPembimbing:    90,
			NilaiKemandirianFasilitator: 91,
			NilaiPengujianPembimbing:    88,
			CreatedAt:                   time.Now(),
		},
		{
			NIS:                         "34567",
			Nama:                        "Bob White",
			Kelas:                       "11",
			Jurusan:                     "RPL",
			Rombel:                      "C",
			TanggalMasuk:                &waktu_sekarang,
			TanggalKeluar:               &waktu_sekarang,
			Aktif:                       true,
			Email:                       "bobwhite@example.com",
			Password:                    "password101",
			FKIdPembimbing:              pegawaiList[5].ID, // Pembimbing (Charlie Davis)
			FKIdFasilitator:             pegawaiList[6].ID, // Fasilitator (Eve Adams)
			FKIdIndustri:                industriList[0].ID,
			NilaiSoftskillIndustri:      80,
			NilaiSoftskillFasilitator:   84,
			NilaiHardskillIndustri:      90,
			NilaiHardskillPembimbing:    93,
			NilaiKemandirianFasilitator: 86,
			NilaiPengujianPembimbing:    85,
			CreatedAt:                   time.Now(),
		},
		{
			NIS:                         "56789",
			Nama:                        "Liam Green",
			Kelas:                       "12",
			Jurusan:                     "TKJ",
			Rombel:                      "A",
			TanggalMasuk:                &waktu_sekarang,
			TanggalKeluar:               &waktu_sekarang,
			Aktif:                       true,
			Email:                       "liamgreen@example.com",
			Password:                    "password112",
			FKIdPembimbing:              pegawaiList[6].ID, // Pembimbing (Eve Adams)
			FKIdFasilitator:             pegawaiList[7].ID, // Fasilitator (Grace Lee)
			FKIdIndustri:                industriList[1].ID,
			NilaiSoftskillIndustri:      78,
			NilaiSoftskillFasilitator:   85,
			NilaiHardskillIndustri:      90,
			NilaiHardskillPembimbing:    88,
			NilaiKemandirianFasilitator: 90,
			NilaiPengujianPembimbing:    86,
			CreatedAt:                   time.Now(),
		},
		{
			NIS:                         "98765",
			Nama:                        "Sophia Black",
			Kelas:                       "11",
			Jurusan:                     "RPL",
			Rombel:                      "C",
			TanggalMasuk:                &waktu_sekarang,
			TanggalKeluar:               &waktu_sekarang,
			Aktif:                       true,
			Email:                       "sophiablack@example.com",
			Password:                    "password130",
			FKIdPembimbing:              pegawaiList[4].ID, // Pembimbing (Bob Brown)
			FKIdFasilitator:             pegawaiList[5].ID, // Fasilitator (Charlie Davis)
			FKIdIndustri:                industriList[2].ID,
			NilaiSoftskillIndustri:      84,
			NilaiSoftskillFasilitator:   90,
			NilaiHardskillIndustri:      87,
			NilaiHardskillPembimbing:    93,
			NilaiKemandirianFasilitator: 92,
			NilaiPengujianPembimbing:    89,
			CreatedAt:                   time.Now(),
		},
	}

	// Insert data dummy
	if err := DB.Database.Session(&gorm.Session{PrepareStmt: false}).Create(&dataSiswaList).Error; err != nil {
		log.Fatal("Gagal mengisi data DataSiswa:", err)
	}
	fmt.Println("Data DataSiswa berhasil ditambahkan")
}
