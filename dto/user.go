package dto

type UserRegister struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdatePassword struct {
	PasswordLama string `json:"passwordlama" binding:"required"`
	PasswordBaru string `json:"passwordbaru" binding:"required"`
}
