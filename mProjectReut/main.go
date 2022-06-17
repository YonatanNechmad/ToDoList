package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mProjectReut/Config"
	"mProjectReut/Models"
	"mProjectReut/Routes"
	"net/http"
)

var err error

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	fmt.Println("main :  before do the open")
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	fmt.Println("main :  after do the open")
	if err != nil {
		fmt.Println("statuse: ", err)
	}

	fmt.Println("main :  it wasnt error")
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.TablePersons{})
	Config.DB.AutoMigrate(&Models.TableChore{})
	Config.DB.AutoMigrate(&Models.TableHomeWork{})

	r := Routes.SetupRouter()

	// running
	r.Run(":8080")
}
