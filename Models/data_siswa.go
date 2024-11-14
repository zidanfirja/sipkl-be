package Models

import "time"

type DataSiswa struct {
	NIS     string `json:"id" gorm:"primaryKey;type:varchar(50)"`
	Nama    string `json:"name" gorm:"type:varchar(255)"`
	Kelas   string `json:"kelas" gorm:"type:varchar(255)"`
	Jurusan string `json:"jurusan" gorm:"type:varchar(255)"`
	Rombel  string `json:"rombel" gorm:"type:varchar(255)"`

	TanggalMasuk  time.Time `json:"tanggal_masuk" gorm:"date"`
	TanggalKeluar time.Time `json:"tanggal_keluar" gorm:"date"`

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
