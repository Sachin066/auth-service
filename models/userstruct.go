package userstruct

import (
	"gorm.io/gorm"
)						

// User struct is used to store user information in the database	
type User struct {
    gorm.Model
    Username string `json:"username" gorm:"unique"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}
