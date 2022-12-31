package entities

type User struct {
	ID       int    `gorm:"primaryKey" json:"id" form:"id"`
	Login    string `json:"login" form:"login"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
}

func (User) TableName() string {
	return "users"
}
