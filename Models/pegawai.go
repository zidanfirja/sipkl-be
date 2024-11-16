package Models

import (
	"errors"
	"fmt"
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

	KonfigurasiRoles []KonfigurasiRoles `json:"daftar_role" gorm:"foreignKey:FKIdPegawai;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Pembimbing       []DataSiswa        `gorm:"foreignKey:FKIdPembimbing;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Fasilitator      []DataSiswa        `gorm:"foreignKey:FKIdFasilitator;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

type DeletePegawaiReq struct {
	ID interface{} `json:"id" binding:"required"`
}

func GetPegawai() ([]Pegawai, error) {
	var pegawaiModel []Pegawai

	rows := DB.Database.
		Preload("KonfigurasiRoles").
		Preload("KonfigurasiRoles.Role").
		Find(&pegawaiModel)

	return pegawaiModel, rows.Error
}

func CreatePegawai(pegawai *Pegawai) error {

	pegawai.CreatedAt = time.Now()
	createPegawai := DB.Database.Omit("id").Create(pegawai)
	return createPegawai.Error
}

func DeletePegawai(id int) error {

	delete := DB.Database.Where("id = ?", id).Delete(&Pegawai{})

	if delete.RowsAffected == 0 {
		return errors.New("pegawai dengan id tersebut tidak ditemukan")
	}

	if delete.Error != nil {
		fmt.Println(delete.Error)
		return delete.Error
	}
	return nil
}
