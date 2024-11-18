package Models

import (
	"fmt"
	DB "go-gin-mysql/Database"
	"time"
)

type DataSiswa struct {
	NIS     string `json:"id" gorm:"primaryKey;type:varchar(50)"`
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
	NIS      string `json:"id" gorm:"primaryKey;type:varchar(50)"`
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
	NIS       string `json:"nis"`
	NamaSiswa string `json:"nama_siswa"`
	Jurusan   string `json:"jurusan"`
	Rombel    string `json:"rombel"`
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
		Select("nis, nama AS nama_siswa, jurusan, rombel").
		Where("fk_id_industri = ?", id).
		Find(&siswa)

	if rows.Error != nil {
		fmt.Println("gagal mengambil data siswa by industri")
		return nil, rows.Error
	}
	return siswa, nil

}

func AddDataPkl(dataSiswa *DataSiswa) error {

	// if dataSiswa.TanggalMasuk.IsZero() {
	// 	dataSiswa.TanggalMasuk = nil
	// }

	dataSiswa.CreatedAt = time.Now()

	create := DB.Database.Create(&dataSiswa)
	if create.Error != nil {
		return create.Error
	}
	return nil
}
