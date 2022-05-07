package model

type Admin struct {
	Id        int    `gorm:"type:int(11) unsigned auto_increment" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Email     string `gorm:"unique;not null;comment:邮箱" json:"email"`
	Password  string `json:"-"`
	Avatar    string `json:"avatar"`
	CreatedAt int    `gorm:"type:int(10)" json:"created_at"`
	UpdatedAt int    `gorm:"type:int(10)" json:"updated_at"`
}

func (a Admin) TableName() string {
	return "admin"
}
