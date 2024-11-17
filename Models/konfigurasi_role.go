package Models

import (
	"errors"
	"fmt"
	DB "go-gin-mysql/Database"
	"time"
)

type KonfigurasiRoles struct {
	ID int `gorm:"type:int;primaryKey;autoIncrement" json:"id"`

	FKIdPegawai int     `json:"fk_id_data_pegawai" gorm:"index;type:int"`                             // foreign key column
	Pegawai     Pegawai `gorm:"foreignKey:FKIdPegawai;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // define foreign key relationship`

	// Foreign key dengan constraint
	FKIdRole *int `json:"fk_id_role" gorm:"type:int;index"`
	Role     Role `gorm:"foreignKey:FKIdRole;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

type ReqAssignRole struct {
	ID      int `json:"id"`
	Payload struct {
		IDRole int  `json:"id_role"`
		Aktif  bool `json:"aktif"`
	} `json:"payload"`
}

type IdRequest struct {
	ID int `json:"id_konfigurasi_role"`
}

func GetRoleByIdPegawai(id int) ([]Role, error) {

	var roles []Role

	rows := DB.Database.
		Table("role").
		Select("role.id, role.nama, role.aktif, role.created_at").
		Joins("JOIN konfigurasi_roles ON konfigurasi_roles.fk_id_role = role.id").
		Joins("JOIN pegawai ON pegawai.id = konfigurasi_roles.fk_id_pegawai").
		Where("pegawai.id = ?", id).
		Find(&roles)

	if rows.Error != nil {
		return nil, rows.Error
	}
	return roles, nil

}

func AddKonfigurasiRole(data *KonfigurasiRoles) error {
	created_at := time.Now()
	data.CreatedAt = created_at

	create := DB.Database.Omit("id").Create(data)
	return create.Error

}

func DeleteRolePegawai(id int) error {
	var konfigurasiRole KonfigurasiRoles
	delete := DB.Database.Where("id = ?", id).Delete(&konfigurasiRole)
	if delete.Error != nil {
		return delete.Error
	}

	if delete.RowsAffected == 0 {
		fmt.Println(delete.Error)
		return errors.New("gagal menghapus jabatan tugas pegawai")
	}

	return nil

}
