package Models

import (
	DB "go-gin-mysql/Database"
	"time"
)

type Pegawai struct {
	ID        int    `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	IdPegawai string `gorm:"type:varchar(100);not null" json:"id_pegawai"`
	Nama      string `gorm:"type:varchar(255);not null" json:"nama"`
	Email     string `gorm:"unique;type:varchar(100)" json:"email"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
	Aktif     bool   `json:"aktif"`

	KonfigurasiRoles []KonfigurasiRoles `gorm:"foreignKey:FKIdPegawai;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Pembimbing       []DataSiswa        `gorm:"foreignKey:FKIdPembimbing;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Fasilitator      []DataSiswa        `gorm:"foreignKey:FKIdFasilitator;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

func GetPegawai() ([]Pegawai, error) {
	var pegawaiModel []Pegawai

	rows := DB.Database.Find(&pegawaiModel)
	return pegawaiModel, rows.Error
}
