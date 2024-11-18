package Models

import (
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

func GetRoleByIdPegawai(id int) ([]RespGetRoles, error) {

	var roles []RespGetRoles

	rows := DB.Database.
		Table("konfigurasi_roles").
		Select("konfigurasi_roles.id AS id_konfigurasi_role, role.id AS role_id, role.nama, role.aktif, role.created_at").
		Joins("JOIN role ON konfigurasi_roles.fk_id_role = role.id").
		Joins("JOIN pegawai ON pegawai.id = konfigurasi_roles.fk_id_pegawai").
		Where("pegawai.id = ?", id).
		Scan(&roles)

	if rows.Error != nil {
		return nil, rows.Error
	}

	// log.Println(roles)
	return roles, nil

}

func AddKonfigurasiRole(data *KonfigurasiRoles) error {
	created_at := time.Now()
	data.CreatedAt = created_at

	create := DB.Database.Omit("id").Create(data)
	return create.Error

}

func CekRolePegawai(id_pegawai int, id_role int) bool {
	var data KonfigurasiRoles
	return DB.Database.Where("fk_id_pegawai = ? AND fk_id_role = ?", id_pegawai, id_role).First(&data).Error == nil
}

func DeleteRolePegawai(id int) error {
	var konfigurasiRole KonfigurasiRoles
	delete := DB.Database.Where("id = ?", id).Delete(&konfigurasiRole)
	if delete.Error != nil {
		return delete.Error
	}

	if delete.RowsAffected == 0 {
		fmt.Println(delete.Error)
		return fmt.Errorf("tidak ada data konfigurasi role dengan id %d", id)
	}

	return nil

}
