package Models

import (
	"errors"
	DB "go-gin-mysql/Database"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Userdata struct {
	ID        int    `json:"id"`
	IdPegawai string `json:"id_pegawai"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
}

type UserInfoOAuth struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type ClaimsUser struct {
	User        Userdata   `json:"userdata"`
	CurrentRole DataRole   `json:"current_role"`
	DaftarRole  []DataRole `json:"daftar_role"`
	jwt.RegisteredClaims
}

func AuthenticateUser(user *Credential) (*Pegawai, error) {

	var pegawai Pegawai

	email := user.Email
	password := user.Password

	rows := DB.Database.Where("email = ?", email).First(&pegawai)
	if rows.Error != nil {
		return nil, errors.New("data tidak ditemukan")
	}

	if rows.RowsAffected == 0 {
		return nil, errors.New("data tidak ditemukan")
	}

	errComparePass := bcrypt.CompareHashAndPassword([]byte(pegawai.Password), []byte(password))
	if errComparePass != nil {
		return nil, errors.New("data tidak ditemukan")
	}

	return &pegawai, nil
}

func AuthenticateUserCekEmail(user *Credential) (*Pegawai, error) {
	var pegawai Pegawai

	email := user.Email

	rows := DB.Database.Where("email = ?", email).First(&pegawai)
	if rows.Error != nil {
		return nil, errors.New("data tidak ditemukan")
	}

	if rows.RowsAffected == 0 {
		return nil, errors.New("data tidak ditemukan")
	}

	return &pegawai, nil
}
