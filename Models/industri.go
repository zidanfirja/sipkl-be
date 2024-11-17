package Models

import (
	"errors"
	DB "go-gin-mysql/Database"
	"time"
)

type Industri struct {
	ID      int    `json:"id" gorm:"primaryKey;type:int"`
	Nama    string `json:"nama" gorm:"type:varchar(255)" binding:"required"`
	Alamat  string `json:"alamat" binding:"required"`
	Jurusan string `json:"jurusan" gorm:"type:varchar(100)"`

	DataSiswa []DataSiswa `gorm:"foreignKey:FKIdIndustri;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

type UpdateIndustriReq struct {
	ID      interface{}            `json:"id" binding:"required"`
	Payload map[string]interface{} `json:"payload" binding:"required"`
}

type DeleteIndustriReq struct {
	ID interface{} `json:"id" binding:"required"`
}

type MultipleIndustri struct {
	Industri []Industri
}

func GetIdustri() ([]Industri, error) {

	var industriModel []Industri
	rows := DB.Database.Find(&industriModel)
	return industriModel, rows.Error

}

func CreateIndustri(industri *Industri) error {

	industri.CreatedAt = time.Now()

	create := DB.Database.Create(&industri)
	if create.Error != nil {
		return create.Error
	}
	return nil
}

func DeleteIndustri(id int) error {
	delete := DB.Database.Where("id = ?", id).Delete(&Industri{})

	if delete.RowsAffected == 0 {
		return errors.New("industri dengan id tersebut tidak ditemukan")
	}

	if delete.Error != nil {
		return delete.Error
	}

	return nil
}

func UpdateSingleIndustri(id int, payload map[string]interface{}) error {

	var industri Industri
	result := DB.Database.First(&industri, id)
	if err := result.Error; err != nil {
		return errors.New("role dengan ID tersebut tidak ditemukan")
	}

	if result.RowsAffected == 0 {
		return errors.New("tidak ada role yang diupdate")
	}

	if err := DB.Database.Model(&industri).Updates(payload).Error; err != nil {
		return err
	}
	return nil

}

// Fungsi untuk mengupdate banyak role berdasarkan array ID
func UpdateMultipleIndustri(ids []int, jurusan string) error {

	payload := map[string]interface{}{
		"jurusan": jurusan,
	}

	// dengan IN
	result := DB.Database.Model(&Industri{}).Where("id IN ?", ids).Updates(payload)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("tidak ada role yang diupdate")
	}

	return nil
}

// func UpdataIndustri(id int, payload map[string]interface{}) error {
// 	update := DB.Database.Where("id = ?", id).Update(&payload)
// }
