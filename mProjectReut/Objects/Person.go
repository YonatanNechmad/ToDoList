package Objects

type Person struct {
	//id                          string	`gorm:"default:uuid_generate_v3()"`
	Name                        string `json:"name" binding:"required"`
	Email                       string `json:"email"  binding:"required,email"`
	FavoriteProgrammingLanguage string `json:"favoriteProgrammingLanguage"  binding:"required"`
}
