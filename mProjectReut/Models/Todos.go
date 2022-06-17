package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mProjectReut/Config"
	"mProjectReut/Objects"
)

func GetAllChoreTasks(todo *[]TableChore) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetAllHomeWorkTasks(todo *[]TableHomeWork) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTodos(todo *[]TableTasks) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}
func GetAllTodosPerOwnerAndStatusHW(todos *[]TableHomeWork, ownerId string, status Objects.Status) (err error) {

	if err = Config.DB.Where("owner_Id = ?  AND status = ?", ownerId, status).Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTodosPerOwnerAndStatusCH(todos *[]TableChore, ownerId string, status Objects.Status) (err error) {

	if err = Config.DB.Where("owner_Id = ? AND status = ?", ownerId, status).Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTodosPerOwnerHW(todos *[]TableHomeWork, ownerId string) (err error) {

	if err = Config.DB.Where("owner_Id = ?", ownerId).Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTodosPerOwnerChore(todos *[]TableChore, ownerId string) (err error) {

	if err = Config.DB.Where("owner_Id = ?", ownerId).Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTodosPerOwner(todos *[]TableTasks, ownerId string) (err error) {

	if err = Config.DB.Where("owner_Id = ?", ownerId).Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func GetAllPersons(person *[]TablePersons) (err error) {
	if err = Config.DB.Find(person).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *TableTasks) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil

}

func CreateAChore(todo *TableChore) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil

}

func CreateAHomeWork(todo *TableHomeWork) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil

}

func CreateAPerson(person *TablePersons) (err error) {
	if err = Config.DB.Create(person).Error; err != nil {
		return err
	}
	return nil

}

func GetAChoreTask(todo *TableChore, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetAHomeWorkTask(todo *TableHomeWork, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *TableTasks, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetAPerson(p *TablePersons, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(p).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAChoreTask(todo *TableChore, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

func UpdateAHomeWorkTask(todo *TableHomeWork, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

func UpdateATodo(todo *TableTasks, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

func UpdateAPerson(p *TablePersons, id string) (err error) {
	fmt.Println(p)
	Config.DB.Save(p)
	return nil
}

func DeleteAHomeWork(todo *TableHomeWork, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}

func DeleteAChore(todo *TableChore, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}

func DeleteATodo(todo *TableTasks, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}

func DeleteAPerson(p *TablePersons, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(p)
	return nil
}
