package Models

import "time"

type KonfigurasiRoles struct {
	ID int `gorm:"type:int;primaryKey;autoIncrement" json:"id"`

	FKIdPegawai int     `json:"fk_id_data_pegawai" gorm:"index;type:int"`                             // foreign key column
	Pegawai     Pegawai `gorm:"foreignKey:FKIdPegawai;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // define foreign key relationship`

	// Foreign key dengan constraint
	FKIdRole *int `json:"fk_id_role" gorm:"type:int;index"`
	Role     Role `gorm:"foreignKey:FKIdRole;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}

// func (KonfigurasiRoles) TableName() string {
// 	return "konfigurasi_role"
// }
