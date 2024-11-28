package Models

import (
	DB "go-gin-mysql/Database"
	"time"
)

type RiwayatUpdateNilaiPembimbing struct {
	NamaPembimbing           string    `json:"nama_pembimbing"`
	NamaSiswa                string    `json:"nama_siswa"`
	NilaiHardskillPembimbing int       `json:"nilai_hardskill_pembimbing"`
	NilaiPengujianPembimbing int       `json:"nilai_pengujian_pembimbing"`
	NilaiSoftskillIndustri   int       `json:"nilai_softskill_industri"`
	NilaiHardskillIndustri   int       `json:"nilai_hardskill_industri"`
	NamaPerusahaan           string    `json:"nama_perusahaan"`
	UpdatedAtNilaiPembimbing time.Time `json:"waktu"`
}

type RiwayatUpdateNilaiFasilitator struct {
	NamaPembimbing              string    `json:"nama_pembimbing"`
	NamaSiswa                   string    `json:"nama_siswa"`
	NilaiKemandirianFasilitator int       `json:"nilai_kemandirian_fasilitator"`
	NilaiSoftskillFasilitator   int       `json:"nilai_softskill_fasilitator"`
	NamaPerusahaan              string    `json:"nama_perusahaan"`
	UpdatedAtNilaiPembimbing    time.Time `json:"waktu"`
}

func GetRiwayatNilaiPembimbing() ([]RiwayatUpdateNilaiPembimbing, error) {

	var dataRiwayat []RiwayatUpdateNilaiPembimbing

	query := `SELECT pegawai.nama as nama_pembimbing, data_siswa.nama as nama_siswa, nilai_hardskill_pembimbing, nilai_pengujian_pembimbing, nilai_softskill_industri, nilai_hardskill_industri, industri.nama as nama_perusahaan, updated_at_nilai_pembimbing as waktu FROM data_siswa
	JOIN pegawai on pegawai.id = data_siswa.fk_id_pembimbing
	JOIN industri on industri.id = data_siswa.fk_id_industri
	WHERE updated_at_nilai_pembimbing IS NOT NULL`

	rows := DB.Database.Raw(query).Scan(&dataRiwayat)
	return dataRiwayat, rows.Error

}

func GetRiwayatNilaiFasilitator() ([]RiwayatUpdateNilaiFasilitator, error) {

	var dataRiwayat []RiwayatUpdateNilaiFasilitator

	query := `SELECT 
	pegawai.nama as nama_pembimbing, data_siswa.nama as nama_siswa, nilai_kemandirian_fasilitator, nilai_softskill_fasilitator, industri.nama as nama_perusahaan, updated_at_nilai_pembimbing as waktu 
FROM data_siswa
JOIN pegawai ON pegawai.id = data_siswa.fk_id_pembimbing
JOIN industri ON industri.id = data_siswa.fk_id_industri
WHERE updated_at_nilai_pembimbing IS NOT NULL`

	rows := DB.Database.Raw(query).Scan(&dataRiwayat)
	return dataRiwayat, rows.Error

}

func GetJumlahPembimbing() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(kr.id) as jumlah
	FROM konfigurasi_roles kr
	JOIN role r ON r.id = kr.fk_id_role
	WHERE lower(nama) = 'pembimbing'`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetJumlahFasilitator() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(kr.id) as jumlah
	FROM konfigurasi_roles kr
	JOIN role r ON r.id = kr.fk_id_role
	WHERE lower(nama) = 'fasilitator'`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetJumlahHubin() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(kr.id) as jumlah
	FROM konfigurasi_roles kr
	JOIN role r ON r.id = kr.fk_id_role
	WHERE lower(nama) = 'hubin'`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetJumlahPemangku() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(DISTINCT(fk_id_pegawai)) as total_pemangku FROM konfigurasi_roles`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetJumlahSiswaPkl() (int64, error) {
	var jumlah int64
	query := `SELECT COUNT(nis) FROM data_siswa WHERE aktif = true`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return jumlah, err
}

func GetJumlahWakel() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(kr.id) as jumlah
	FROM konfigurasi_roles kr
	JOIN role r ON r.id = kr.fk_id_role
	WHERE lower(nama) = 'wali kelas'`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}
