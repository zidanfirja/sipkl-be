package Models

import "time"

type Industri struct {
	ID   int    `json:"id" gorm:"primaryKey;type:int"`
	Nama string `json:"nama" gorm:"type:varchar(255)"`

	DataSiswa []DataSiswa `gorm:"foreignKey:FKIdIndustri;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}
