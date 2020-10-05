package auth

// User stores the information of a registered user
type User struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Mail     string `json:"email" gorm:"unique;not_null"`
	Password []byte `json:"-" gorm:"not_null"` // Do not marshall
}

// TableName returns the name of the gorm model
func (User) TableName() string {
	return "users"
}
