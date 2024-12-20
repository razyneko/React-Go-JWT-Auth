package models

type User struct {
	Id 			uint	`json:"id"`
	Name 		string	`json:"name"`
	Email 		string 	`json:"email" gorm:"unique"` 	
	Password 	[]byte	`json:"-"`
}
// "-" in password represents we dont want to show password