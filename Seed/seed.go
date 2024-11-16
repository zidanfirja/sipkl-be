package Seed

import (
	"fmt"
	DB "go-gin-mysql/Database"
	"go-gin-mysql/Models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
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
		{Nama: "Manager", Aktif: false},
	}

	if err := DB.Database.Create(&roles).Error; err != nil {
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
	}

	// Insert data dummy
	if err := DB.Database.Create(&industriList).Error; err != nil {
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
			IdPegawai: 1001,
			Nama:      "John Doe",
			Email:     "johndoe@example.com",
			Password:  HashPassword("password123"),
			Aktif:     true,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: 1002,
			Nama:      "Jane Smith",
			Email:     "janesmith@example.com",
			Password:  HashPassword("password123"),
			Aktif:     true,
			CreatedAt: time.Now(),
		},
		{
			IdPegawai: 1003,
			Nama:      "David Johnson",
			Email:     "davidj@example.com",
			Password:  HashPassword("password123"),
			Aktif:     false,
			CreatedAt: time.Now(),
		},
	}

	// Insert data dummy
	if err := DB.Database.Create(&pegawaiList).Error; err != nil {
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
			FKIdPegawai: pegawaiList[0].ID,
			FKIdRole:    &roleList[0].ID,
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[1].ID,
			FKIdRole:    &roleList[1].ID,
			CreatedAt:   time.Now(),
		},
		{
			FKIdPegawai: pegawaiList[2].ID,
			FKIdRole:    &roleList[0].ID,
			CreatedAt:   time.Now(),
		},
	}

	// Insert data dummy
	if err := DB.Database.Create(&konfigurasiRolesList).Error; err != nil {
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

	// Data dummy DataSiswa
	dataSiswaList := []Models.DataSiswa{
		{
			NIS:                         "12345",
			Nama:                        "John Doe",
			Kelas:                       "12",
			Jurusan:                     "RPL",
			Rombel:                      "A",
			TanggalMasuk:                time.Now().AddDate(-1, 0, 0),
			TanggalKeluar:               time.Now().AddDate(0, 6, 0),
			Aktif:                       true,
			Email:                       "johndoe@example.com",
			Password:                    "password123",
			FKIdPembimbing:              pegawaiList[0].ID,
			FKIdFasilitator:             pegawaiList[1].ID,
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
			TanggalMasuk:                time.Now().AddDate(-2, 0, 0),
			TanggalKeluar:               time.Now().AddDate(0, 4, 0),
			Aktif:                       true,
			Email:                       "janesmith@example.com",
			Password:                    "password456",
			FKIdPembimbing:              pegawaiList[0].ID,
			FKIdFasilitator:             pegawaiList[1].ID,
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
			TanggalMasuk:                time.Now().AddDate(-2, 0, 0),
			TanggalKeluar:               time.Now().AddDate(0, 4, 0),
			Aktif:                       true,
			Email:                       "dariamsa@example.com",
			Password:                    "password456",
			FKIdPembimbing:              pegawaiList[2].ID,
			FKIdFasilitator:             pegawaiList[1].ID,
			FKIdIndustri:                industriList[2].ID,
			NilaiSoftskillIndustri:      80,
			NilaiSoftskillFasilitator:   85,
			NilaiHardskillIndustri:      82,
			NilaiHardskillPembimbing:    86,
			NilaiKemandirianFasilitator: 88,
			NilaiPengujianPembimbing:    84,
			CreatedAt:                   time.Now(),
		},
	}

	// Insert data dummy
	if err := DB.Database.Create(&dataSiswaList).Error; err != nil {
		log.Fatal("Gagal mengisi data DataSiswa:", err)
	}
	fmt.Println("Data DataSiswa berhasil ditambahkan")
}
