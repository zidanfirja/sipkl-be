package Models

import (
	DB "go-gin-mysql/Database"
	"time"
)

type IndustriPembimbingFasil struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type NilaiSiswaPklPembimbing struct {
	NIS                      string `json:"nis"`
	Nama                     string `json:"nama"`
	Kelas                    string `json:"kelas"`
	Jurusan                  string `json:"jurusan"`
	Rombel                   string `json:"rombel"`
	NilaiSoftskillIndustri   int    `json:"nilai_softskill_industri"`
	NilaiHardskillIndustri   int    `json:"nilai_hardskill_industri"`
	NilaiHardskillPembimbing int    `json:"nilai_hardskill_pembimbing"`
	NilaiPengujianPembimbing int    `json:"nilai_pengujian_pembimbing"`
}

type NilaiSiswaPklFasilitator struct {
	NIS                         string `json:"nis"`
	Nama                        string `json:"nama"`
	Kelas                       string `json:"kelas"`
	Jurusan                     string `json:"jurusan"`
	Rombel                      string `json:"rombel"`
	NilaiSoftskillFasilitator   int    `json:"nilai_softskill_fasilitator"`
	NilaiKemandirianFasilitator int    `json:"nilai_kemandirian_fasilitator"`
}

type IndustriForNilai struct {
	ID            int       `json:"id_perusahaan" gorm:"column:id_perusahaan"`
	Nama          string    `json:"nama_perusahaan" gorm:"column:nama_perusahaan"`
	Alamat        string    `json:"alamat_perusahaan" gorm:"column:alamat_perusahaan"`
	TanggalMasuk  time.Time `json:"tanggal_masuk" gorm:"column:tanggal_masuk"`
	TanggalKeluar time.Time `json:"tanggal_keluar" gorm:"column:tanggal_keluar"`
}

func GetIndustriPembimbing(id int) ([]IndustriPembimbingFasil, error) {
	var data []IndustriPembimbingFasil

	query := `SELECT fk_id_industri AS id, industri.nama AS nama
	FROM data_siswa 
	JOIN pegawai ON pegawai.id = data_siswa.fk_id_pembimbing
	JOIN industri ON industri.id = data_siswa.fk_id_industri
	WHERE fk_id_pembimbing = ?
	GROUP BY fk_id_industri`

	rows := DB.Database.Raw(query, id).Scan(&data)
	if rows.Error != nil {
		return nil, rows.Error
	}

	return data, nil
}

func GetIndustriFasilitator(id int) ([]IndustriPembimbingFasil, error) {
	var data []IndustriPembimbingFasil

	query := `SELECT fk_id_industri AS id, industri.nama AS nama
	FROM data_siswa 
	JOIN pegawai ON pegawai.id = data_siswa.fk_id_fasilitator
	JOIN industri ON industri.id = data_siswa.fk_id_industri
	WHERE fk_id_fasilitator = ?
    GROUP BY fk_id_industri`

	rows := DB.Database.Raw(query, id).Scan(&data)
	if rows.Error != nil {
		return nil, rows.Error
	}

	return data, nil
}

func GetIndustri(id_industri int) (IndustriForNilai, error) {
	var dataIndustri IndustriForNilai

	query := `SELECT industri.id AS id_perusahaan, industri.nama AS nama_perusahaan,industri.alamat AS alamat_perusahaan, data_siswa.tanggal_masuk, data_siswa.tanggal_keluar from data_siswa 
    JOIN pegawai on pegawai.id = data_siswa.fk_id_pembimbing
    JOIN industri on industri.id = data_siswa.fk_id_industri
    WHERE fk_id_industri = ?
    LIMIT 1`

	rows := DB.Database.Raw(query, id_industri).Scan(&dataIndustri)
	if rows.Error != nil {
		return dataIndustri, rows.Error
	}

	return dataIndustri, nil

}

func GetNilaiByPemb(id_pembimbing, id_industri int) ([]NilaiSiswaPklPembimbing, error) {

	var nilai []NilaiSiswaPklPembimbing

	query := ` SELECT 
    nis,nama,kelas,jurusan,rombel,
    nilai_softskill_industri,
    nilai_hardskill_industri,
    nilai_hardskill_pembimbing,
    nilai_pengujian_pembimbing
    FROM data_siswa
    WHERE fk_id_pembimbing = ? AND fk_id_industri = ?`

	rows := DB.Database.Raw(query, id_pembimbing, id_industri).Scan(&nilai)
	if rows.Error != nil {
		return nilai, rows.Error
	}

	return nilai, nil

}

func GetNilaiByFasil(id_fasil, id_industri int) ([]NilaiSiswaPklFasilitator, error) {

	var nilai []NilaiSiswaPklFasilitator

	query := `SELECT 
    nis,nama,kelas,jurusan,rombel,
    nilai_softskill_fasilitator,
    nilai_kemandirian_fasilitator
    FROM data_siswa
    WHERE fk_id_fasilitator = ? AND fk_id_industri = ?`

	rows := DB.Database.Raw(query, id_fasil, id_industri).Scan(&nilai)
	if rows.Error != nil {
		return nilai, rows.Error
	}

	return nilai, nil

}
