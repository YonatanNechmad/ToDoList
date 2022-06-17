package Models

import (
	myObjects "mProjectReut/Objects"
	"time"
)

type TableTasks struct {
	Id      string           `json:"id"`
	OwnerId string           `json:"ownerId"`
	Status  myObjects.Status `json:"status"`
	Type    string           `json:"type"`
}

type TablePersons struct {
	ID                          string `json:"id" gorm:"primaryKey"`
	Name                        string `json:"name"`
	Email                       string `json:"email" binding:"required,email" gorm:"unique"`
	FavoriteProgrammingLanguage string `json:"favoriteProgrammingLanguage"`
	ActivteTaskCount            uint   `json:"activeTaskCount"`
}

type TableChore struct {
	Id          string           `json:"id"`
	OwnerId     string           `json:"ownerId"`
	Status      myObjects.Status `json:"status"`
	Description string           `json:"description"`
	Size        myObjects.Size   `json:"size"`
	Type        string           `json:"type"`
}

type TableHomeWork struct {
	Id      string           `json:"id"`
	OwnerId string           `json:"ownerId"`
	Status  myObjects.Status `json:"status"`
	Course  string           `json:"course"`
	DueDate time.Time        `json:"dueDate"`
	Details string           `json:"details"`
	Type    string           `json:"type"`
}

func (b *TableTasks) TableName() string {
	return "TableTasks"
}

func (b *TablePersons) TableName() string {
	return "TablePersons"
}

func (b *TableHomeWork) TableName() string {
	return "TableHomeWork"
}

func (b *TableChore) TableName() string {
	return "TableChore"
}
