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

type RespPegawaiGetAll struct {
	ID        int    `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	IdPegawai string `gorm:"type:varchar(100);not null" json:"id_pegawai"`
	Nama      string `gorm:"type:varchar(255);not null" json:"nama"`
	Email     string `gorm:"unique;type:varchar(100)" json:"email"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
	Aktif     bool   `json:"aktif"`

	DaftarRole []Role `json:"daftar_role" gorm:"-"`
	// Pembimbing       []DataSiswa        `gorm:"foreignKey:FKIdPembimbing;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	// Fasilitator      []DataSiswa        `gorm:"foreignKey:FKIdFasilitator;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

type UpdatePegawaiReq struct {
	ID      interface{}            `json:"id" binding:"required"`
	Payload map[string]interface{} `json:"payload" binding:"required"`
}

type DeletePegawaiReq struct {
	ID interface{} `json:"id" binding:"required"`
}

func GetPegawai() ([]Pegawai, error) {
	var pegawaiModel []Pegawai

	rows := DB.Database.Find(&pegawaiModel)

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

func UpdateSinglePegawai(id int, payload map[string]interface{}) error {

	var pegawai Pegawai
	result := DB.Database.First(&pegawai, id)
	if err := result.Error; err != nil {
		return errors.New("pegawai dengan ID tersebut tidak ditemukan")
	}

	if result.RowsAffected == 0 {
		return errors.New("tidak ada pegawai yang diupdate")
	}

	if err := DB.Database.Model(&pegawai).Updates(payload).Error; err != nil {
		return err
	}
	return nil

}

func UpdateMultiplePegawai(ids []int, payload map[string]interface{}) error {

	result := DB.Database.Model(&Pegawai{}).Where("id IN ?", ids).Updates(payload)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("tidak ada pegawai yang diupdate")
	}

	return nil
}

func UpdateAktifPegawai(id int, aktif bool) error {
	result := DB.Database.Model(&Pegawai{}).Where("id = ?", id).Update("aktif", aktif)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("tidak ada pegawai yang diupdate")
	}

	return nil

}
