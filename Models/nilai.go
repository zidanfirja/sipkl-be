package Models

import (
	"fmt"
	DB "go-gin-mysql/Database"
	"strings"
	"time"
)

type IndustriPembimbingFasil struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type NilaiSiswaPkl struct {
	NIS                         string `json:"nis"`
	Nama                        string `json:"nama"`
	Kelas                       string `json:"kelas"`
	Jurusan                     string `json:"jurusan"`
	Rombel                      string `json:"rombel"`
	NilaiSoftskillFasilitator   int    `json:"nilai_softskill_fasilitator"`
	NilaiSoftskillIndustri      int    `json:"nilai_softskill_industri"`
	NilaiHardskillPembimbing    int    `json:"nilai_hardskill_pembimbing"`
	NilaiHardskillIndustri      int    `json:"nilai_hardskill_industri"`
	NilaiKemandirianFasilitator int    `json:"nilai_kemandirian_fasilitator"`
	NilaiPengujianPembimbing    int    `json:"nilai_pengujian_pembimbing"`
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

type ReqUpdateNilaiPembimbing struct {
	NIS                      string  `json:"nis"`
	NilaiSoftskillIndustri   float64 `json:"nilai_softskill_industri"`
	NilaiHardskillIndustri   float64 `json:"nilai_hardskill_industri"`
	NilaiHardskillPembimbing float64 `json:"nilai_hardskill_pembimbing"`
	NilaiPengujianPembimbing float64 `json:"nilai_pengujian_pembimbing"`
}

type ReqUpdateNilaiFasilitator struct {
	NIS                         string  `json:"nis"`
	NilaiSoftskillFasilitator   float64 `json:"nilai_softskilll_fasilitator"`
	NilaiKemandirianFasilitator float64 `json:"nilai_kemandirian_fasilitator"`
}

func GetIndustriPembimbing(id int) ([]IndustriPembimbingFasil, error) {
	var data []IndustriPembimbingFasil

	query := `SELECT fk_id_industri AS id, industri.nama AS nama
	FROM data_siswa 
	JOIN pegawai ON pegawai.id = data_siswa.fk_id_pembimbing
	JOIN industri ON industri.id = data_siswa.fk_id_industri
	WHERE fk_id_pembimbing = ?
	GROUP BY fk_id_industri, industri.nama;`

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
    GROUP BY fk_id_industri,  industri.nama `

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
		return nil, rows.Error
	}

	return nilai, nil

}

func UpdateNilaiPembimbing(data *[]ReqUpdateNilaiPembimbing) error {

	var listNis []string
	var caseNilaiSoftskillIndustri, caseNilaiHardskillIndustri, caseNilaiHardskillPembimbing, caseNilaiPengujianPembimbing string

	for _, dataNilai := range *data {
		listNis = append(listNis, fmt.Sprintf("'%s'", dataNilai.NIS))
		caseNilaiSoftskillIndustri += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiSoftskillIndustri)
		caseNilaiHardskillIndustri += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiHardskillIndustri)
		caseNilaiHardskillPembimbing += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiHardskillPembimbing)
		caseNilaiPengujianPembimbing += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiPengujianPembimbing)
	}

	query := fmt.Sprintf(`
	UPDATE data_siswa
	SET 
	nilai_softskill_industri = CASE nis %s END,
	nilai_hardskill_industri = CASE nis %s END,
	nilai_hardskill_pembimbing = CASE nis %s END,
	nilai_pengujian_pembimbing = CASE nis %s END
	WHERE nis IN (%s);
`, caseNilaiSoftskillIndustri, caseNilaiHardskillIndustri, caseNilaiHardskillPembimbing, caseNilaiPengujianPembimbing, strings.Join(listNis, ", "))

	if err := DB.Database.Exec(query).Error; err != nil {
		return err
	}

	return nil
}

func UpdateNilaiFasilitator(data *[]ReqUpdateNilaiFasilitator) error {

	var listNis []string
	var caseNilaiSoftskillFasilitator, caseNilaiKemandirianFasilitator string

	// NilaiSoftskillFasilitator   float64 `json:"nilai_softskilll_fasilitator"`
	// NilaiKemandirianFasilitator float64 `json:"nilai_kemandirian_fasilitator"`

	for _, dataNilai := range *data {
		listNis = append(listNis, fmt.Sprintf("'%s'", dataNilai.NIS))
		caseNilaiSoftskillFasilitator += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiSoftskillFasilitator)
		caseNilaiKemandirianFasilitator += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiKemandirianFasilitator)

	}

	query := fmt.Sprintf(`
	UPDATE data_siswa
	SET 
	nilai_softskill_fasilitator = CASE nis %s END,
	nilai_kemandirian_fasilitator = CASE nis %s END
	WHERE nis IN (%s);
`, caseNilaiSoftskillFasilitator, caseNilaiKemandirianFasilitator, strings.Join(listNis, ", "))

	if err := DB.Database.Exec(query).Error; err != nil {
		return err
	}

	return nil
}

func GetNilaiWakel(kelas, jurusan, rombel string) ([]NilaiSiswaPkl, error) {
	var dataNilai []NilaiSiswaPkl

	query := `
	SELECT nis, nama, kelas, jurusan, rombel, 
	nilai_softskill_fasilitator,nilai_softskill_industri, 
	nilai_hardskill_pembimbing, nilai_hardskill_industri,
	nilai_kemandirian_fasilitator, nilai_pengujian_pembimbing 
	FROM data_siswa
	WHERE kelas = ? AND jurusan = ? AND rombel = ?`

	rows := DB.Database.Raw(query, kelas, jurusan, rombel).Scan(&dataNilai)
	if rows.Error != nil {
		return nil, rows.Error
	}

	return dataNilai, nil
}
