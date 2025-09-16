package ds

type Users struct {
	UserID uint `json:"id" gorm:"primaryKey"`
	Login string `json:"login" gorm:"type:varchar(100);unique"`
	Password string `json:"-" gorm:"type:varchar(100)"`
	IsAdmin bool `json:"is_admin"`
}