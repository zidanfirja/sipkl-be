package Models

import (
	"errors"
	"fmt"
	DB "go-gin-mysql/Database"
	"time"
)

type Role struct {
	ID    int    `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Nama  string `gorm:"type:varchar(50);not null" json:"nama" binding:"required"`
	Aktif bool   `json:"aktif" binding:"required"`

	// ini untuk migrate db harus di un-comment
	KonfigurasiRoles []KonfigurasiRoles `gorm:"foreignKey:FKIdRole;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

type DataRole struct {
	IDRole   int    `json:"id_role"`
	NamaRole string `json:"nama_role"`
}

type DeleteRoleReq struct {
	ID interface{} `json:"id" binding:"required"`
}

type UpdateRoleReq struct {
	ID      interface{}            `json:"id" binding:"required"`
	Payload map[string]interface{} `json:"payload" binding:"required"`
}

type RespGetRoles struct {
	IDKonRole int       `json:"id_konfigurasi_role" gorm:"column:id_konfigurasi_role"`
	IDRole    int       `json:"id" gorm:"column:role_id"`
	Nama      string    `json:"nama" gorm:"column:nama"`
	Aktif     bool      `json:"aktif" gorm:"column:aktif"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func GetRoles() ([]Role, error) {

	var roleModel []Role
	rows := DB.Database.Find(&roleModel)
	return roleModel, rows.Error

}

func CreateRole(role *Role) error {

	role.CreatedAt = time.Now()

	createRole := DB.Database.Omit("id").Create(&role)

	if createRole.Error != nil {
		return createRole.Error
	}
	return nil

}

func DeleteRole(id int) error {

	delete := DB.Database.Where("id = ?", id).Delete(&Role{})

	if delete.RowsAffected == 0 {
		return errors.New("gagal delete, role dengan id tersebut tidak ditemukan")
	}

	if delete.Error != nil {
		fmt.Println(delete.Error)
		return delete.Error
	}
	return nil
}

func UpdateSingleRole(id int, payload map[string]interface{}) error {

	var role Role
	result := DB.Database.First(&role, id)
	if err := result.Error; err != nil {
		return errors.New("gagal upadate, role dengan ID tersebut tidak ditemukan")
	}

	if result.RowsAffected == 0 {
		return errors.New("tidak ada role yang diupdate")
	}

	if err := DB.Database.Model(&role).Updates(payload).Error; err != nil {
		return err
	}
	return nil

}

// Fungsi untuk mengupdate banyak role berdasarkan array ID
func UpdateMultipleRoles(ids []int, payload map[string]interface{}) error {

	result := DB.Database.Model(&Role{}).Where("id IN ?", ids).Updates(payload)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("tidak ada role yang diupdate")
	}

	return nil
}
