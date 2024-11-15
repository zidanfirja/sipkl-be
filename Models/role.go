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
	// KonfigurasiRoles []KonfigurasiRoles `gorm:"foreignKey:FKIdRole;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

type DeleteRoleReq struct {
	ID interface{} `json:"id" binding:"required"`
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
	fmt.Println(id)
	delete := DB.Database.Where("id = ?", id).Delete(&Role{})

	if delete.RowsAffected == 0 {
		return errors.New("role dengan id tersebut tidak ditemukan")
	}

	if delete.Error != nil {
		fmt.Println(delete.Error)
		return delete.Error
	}
	return nil
}
