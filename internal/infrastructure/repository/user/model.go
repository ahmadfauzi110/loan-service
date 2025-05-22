package user

type Model struct {
	ID       *int   `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name     string `gorm:"column:name;size:100;not null"`
	UserType string `gorm:"column:user_type;size:100;not null;index"`
	Email    string `gorm:"column:email;size:100;not null;index"`
}

func (Model) TableName() string {
	return "users"
}
