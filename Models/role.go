package Models

import (
	DB "go-gin-mysql/Database"
	"time"
)

type Role struct {
	ID    int    `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	Nama  string `gorm:"type:varchar(50);not null" json:"nama"`
	Aktif bool   `json:"aktif"`

	// ini untuk migrate db harus di un-comment
	// KonfigurasiRoles []KonfigurasiRoles `gorm:"foreignKey:FKIdRole;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
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
