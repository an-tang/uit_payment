package model

type Partner struct {
	BaseModel
	Name        string `gorm:"column:name" json:"name"`
	Key         string `gorm:"column:key" json:"key"`
	CallbackURL string `gorm:"callback_url" json:"callback_url"`
}

// TableName mapping table name in database
func (Partner) TableName() string {
	return "partners"
}
