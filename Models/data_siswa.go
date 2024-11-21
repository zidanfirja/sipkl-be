package Models

import (
	"fmt"
	DB "go-gin-mysql/Database"
	"log"
	"strings"
	"time"
)

type DataSiswa struct {
	NIS     string `json:"nis" gorm:"primaryKey;type:varchar(50)"`
	Nama    string `json:"name" gorm:"type:varchar(255)"`
	Kelas   string `json:"kelas" gorm:"type:varchar(255)"`
	Jurusan string `json:"jurusan" gorm:"type:varchar(255)"`
	Rombel  string `json:"rombel" gorm:"type:varchar(255)"`

	TanggalMasuk  *time.Time `json:"tanggal_masuk" gorm:"date"`
	TanggalKeluar *time.Time `json:"tanggal_keluar" gorm:"date"`

	Aktif bool `json:"aktif"`

	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`

	FKIdPembimbing    int     `json:"fk_id_pembimbing"`
	PegawaiPembimbing Pegawai `gorm:"foreignKey:FKIdPembimbing;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	FKIdFasilitator    int     `json:"fk_id_fasilitator"`
	PegawaiFasilitator Pegawai `gorm:"foreignKey:FKIdFasilitator;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	FKIdIndustri int      `json:"fk_id_industri" gorm:"type:int;index"`
	Industri     Industri `gorm:"foreignKey:FKIdIndustri;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	NilaiSoftskillIndustri      int `json:"nilai_softskill_industri"`
	NilaiSoftskillFasilitator   int `json:"nilai_softskill_fasilitator"`
	NilaiHardskillIndustri      int `json:"nilai_hardskill_industri"`
	NilaiHardskillPembimbing    int `json:"nilai_hardskill_pembimbing"`
	NilaiKemandirianFasilitator int `json:"nilai_kemandirian_fasilitator"`
	NilaiPengujianPembimbing    int `json:"nilai_pengujian_pembimbing"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

type ReqAddDataSiswa struct {
	NIS      string `json:"nis" gorm:"primaryKey;type:varchar(50)"`
	Nama     string `json:"nama" gorm:"type:varchar(255)"`
	Kelas    string `json:"kelas" gorm:"type:varchar(255)"`
	Jurusan  string `json:"jurusan" gorm:"type:varchar(255)"`
	Rombel   string `json:"rombel" gorm:"type:varchar(255)"`
	Aktif    bool   `json:"aktif"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`

	FKIdPembimbing  int `json:"fk_id_pembimbing"`
	FKIdFasilitator int `json:"fk_id_fasilitator"`
	FKIdIndustri    int `json:"fk_id_industri" gorm:"type:int;index"`

	TanggalMasuk  string    `json:"tanggal_masuk" gorm:"date"`
	TanggalKeluar string    `json:"tanggal_keluar" gorm:"date"`
	CreatedAt     time.Time `json:"created_at" gorm:"type:timestamp"`
}

type RespDataPkl struct {
	IDPerusahaan    int     `json:"id_perusahaan"`
	NamaPerusahaan  string  `json:"nama_perusahaan"`
	IDPembimbing    int     `json:"id_pembimbing"`
	NamaPembimbing  string  `json:"nama_pembimbing"`
	IDFasilitator   int     `json:"id_fasilitator"`
	NamaFasilitator string  `json:"nama_fasilitator"`
	DaftarSiswa     []Siswa `json:"daftar_siswa" gorm:"-"`
}

type Siswa struct {
	NIS           string    `json:"nis"`
	Nama          string    `json:"nama"`
	Kelas         string    `json:"kelas"`
	TanggalMasuk  time.Time `json:"tanggal_masuk"`
	TanggalKeluar time.Time `json:"tanggal_keluar"`
	Jurusan       string    `json:"jurusan"`
	Rombel        string    `json:"rombel"`
}

type ReqUpdateDataPkl struct {
	NIS     interface{}            `json:"nis" binding:"required"`
	Payload map[string]interface{} `json:"payload" binding:"required"`
}

type UpdatePetugas struct {
	NIS             string `json:"nis"`
	FKIdPembimbing  int    `json:"fk_id_pembimbing"`
	FKIdFasilitator int    `json:"fk_id_fasilitator"`
	FKIdIndustri    int    `json:"fk_id_industri" gorm:"type:int;index"`
	Aktif           bool   `json:"aktif"`
}

func GetDataPkl() ([]RespDataPkl, error) {
	var rows []RespDataPkl
	query := `SELECT i.id AS id_perusahaan, i.nama AS nama_perusahaan, p.id AS id_pembimbing, p.nama AS nama_pembimbing, f.id AS id_fasilitator, f.nama AS nama_fasilitator FROM data_siswa s LEFT JOIN industri i ON i.id = s.fk_id_industri LEFT JOIN pegawai p ON p.id = s.fk_id_pembimbing LEFT JOIN pegawai f ON f.id = s.fk_id_fasilitator GROUP BY i.id, p.id, f.id`

	// Menjalankan query
	result := DB.Database.Raw(query).Scan(&rows)
	if result.Error != nil {
		fmt.Println("Gagal mengambl data:", result.Error)
		return nil, result.Error
	}

	for i := range rows {
		dataSiswa, err := GetSiswaByIndustri(rows[i].IDPerusahaan)
		if err != nil {
			return nil, err
		}
		rows[i].DaftarSiswa = dataSiswa
	}

	return rows, nil
}

func GetSiswaByIndustri(id int) ([]Siswa, error) {
	var siswa []Siswa

	rows := DB.Database.Table("data_siswa").
		Select("nis,nama,kelas,tanggal_masuk,tanggal_keluar,jurusan,rombel").
		Where("fk_id_industri = ?", id).
		Find(&siswa)

	if rows.Error != nil {
		fmt.Println("gagal mengambil data siswa by industri")
		return nil, rows.Error
	}
	return siswa, nil

}

func AddDataPkl(dataSiswa *DataSiswa) error {

	create := DB.Database.Create(&dataSiswa)
	if create.Error != nil {
		return create.Error
	}
	return nil
}
func AddMultipleDataPkl(dataSiswa *[]DataSiswa) error {

	create := DB.Database.Create(&dataSiswa)
	if create.Error != nil {
		return create.Error
	}
	return nil
}

func UpdateTanggalMasuk(nis []string, tanggal_masuk time.Time) error {

	log.Println(nis)

	payload := map[string]interface{}{
		"tanggal_masuk": tanggal_masuk,
	}

	// dengan IN
	result := DB.Database.Model(&DataSiswa{}).Where("nis IN ?", nis).Updates(payload)
	if result.Error != nil {
		return result.Error
	}

	// if result.RowsAffected == 0 {
	// 	return errors.New("tidak ada role yang diupdate")
	// }

	return nil
}

func UpdateTanggalKeluar(nis []string, tanggal_keluar time.Time) error {
	payload := map[string]interface{}{
		"tanggal_keluar": tanggal_keluar,
	}

	// dengan IN
	result := DB.Database.Model(&DataSiswa{}).Where("nis IN ?", nis).Updates(payload)
	if result.Error != nil {
		return result.Error
	}

	// if result.RowsAffected == 0 {
	// 	return errors.New("tidak ada role yang diupdate")
	// }

	return nil
}

func UpdatePengurusPkl(data *[]UpdatePetugas) error {

	var nisList []string
	var caseFkPembimbing, caseFkFasilitator, caseFkIndustri, caseAktif string

	for _, petugas := range *data {
		nisList = append(nisList, fmt.Sprintf("'%s'", petugas.NIS))
		caseFkPembimbing += fmt.Sprintf("WHEN '%s' THEN %d ", petugas.NIS, petugas.FKIdPembimbing)
		caseFkFasilitator += fmt.Sprintf("WHEN '%s' THEN %d ", petugas.NIS, petugas.FKIdFasilitator)
		caseFkIndustri += fmt.Sprintf("WHEN '%s' THEN %d ", petugas.NIS, petugas.FKIdIndustri)
		caseAktif += fmt.Sprintf("WHEN '%s' THEN %t ", petugas.NIS, petugas.Aktif)
	}

	query := fmt.Sprintf(`
	UPDATE data_siswa
	SET 
		fk_id_pembimbing = CASE nis %s END,
		fk_id_fasilitator = CASE nis %s END,
		fk_id_industri = CASE nis %s END,
		aktif = CASE nis %s END
	WHERE nis IN (%s);
`, caseFkPembimbing, caseFkFasilitator, caseFkIndustri, caseAktif, strings.Join(nisList, ", "))

	// Eksekusi query sekali
	if err := DB.Database.Exec(query).Error; err != nil {
		// DB.Database.Rollback() // Rollback jika ada error
		return err
	}
	return nil

	// Commit transaksi
	// if err := DB.Database.Commit().Error; err != nil {
	// 	return errors.New("erro di sini")
	// }

	// return nil

	// Menyusun query batch update
	// query := "UPDATE data_siswa SET fk_id_pembimbing = ?, fk_id_fasilitator = ?, fk_id_industri = ?, aktif = ? WHERE nis = ?"

	// for _, petugas := range *data {
	// 	if err := DB.Database.Exec(query, petugas.FKIdPembimbing, petugas.FKIdFasilitator, petugas.FKIdIndustri, petugas.Aktif, petugas.NIS).Error; err != nil {
	// DB.Database.Rollback() // Rollback jika ada error
	// 		return err
	// 	}
	// }
	// Commit transaksi setelah batch update
	// if err := DB.Database.Commit().Error; err != nil {
	// 	return err
	// }
	// return nil

}

// UPDATE update_petugas SET
// 	tfk_id_pembimbing = CASE nis
// 		WHEN '11ahmad12345' THEN 2
// 		WHEN '12rah67890' THEN 2
// 		WHEN '13SDSJI' THEN 2  END,
// 	tfk_id_fasilitator = CASE nis
// 		WHEN '11ahmad12345' THEN 2
// 		WHEN '12rah67890' THEN 2
// 		WHEN '13SDSJI' THEN 2  END,
// 	tfk_id_industri = CASE nis
// 		WHEN '11ahmad12345' THEN 2
// 		WHEN '12rah67890' THEN 2
// 		WHEN '13SDSJI' THEN 2  END,
// 	taktif = CASE nis
// 	WHEN '11ahmad12345' THEN true
// 	WHEN '12rah67890' THEN true
// 	WHEN '13SDSJI' THEN true  END
// 	WHERE nis IN ('11ahmad12345', '12rah67890', '13SDSJI');
