//import "net/http"

package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mProjectReut/Models"
	"mProjectReut/Objects"
	"net/http"
	"time"
)

func AddTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	var p Models.TablePersons
	var t Objects.Task
	//
	//var tHomeWork Objects.HomeWork
	//var t Models.TableTasks
	ownerId := c.Params.ByName("id")
	err := Models.GetAPerson(&p, ownerId)
	if err != nil { // it fail because the person is not exist
		c.Header("Content-Type", "text/plain")
		c.Next()
		c.String(404, "A person with the id "+ownerId+" does not exist")

	} else {
		c.ShouldBindJSON(&t)
		if t.Type == "Chore" {
			var tChore Models.TableChore
			tChore.Id = uuid.New().String()
			tChore.OwnerId = ownerId
			tChore.Status = t.Status
			tChore.Description = t.Description
			tChore.Size = t.Size
			tChore.Type = t.Type
			if t.Course != "" || t.Details != "" || t.DueDate != "" {
				c.Header("Content-Type", "text/plain")
				c.Next()
				c.String(400, "You worte data that belong to ChoreTask")
			} else if t.Description == "" || t.Size == "" {
				c.Header("Content-Type", "text/plain")
				c.Next()
				c.String(400,
					"Required data fields are missing, data makes no sense, or data contains illegal values.")
			} else {
				if tChore.Status.String() == "Active" {
					p.ActivteTaskCount = p.ActivteTaskCount + 1
					err1 := Models.UpdateAPerson(&p, p.ID)
					if err1 != nil {
						c.Header("Content-Type", "text/plain")
						c.Next()
						c.AbortWithStatus(http.StatusNotFound)
					}
				}
				err := Models.CreateAChore(&tChore)
				if err != nil {
					c.Header("Content-Type", "text/plain")
					c.Next()
					c.AbortWithStatus(http.StatusNotFound)
				} else {
					c.Header("Location", fmt.Sprintf("/api/people/%s/tasks/", tChore.OwnerId))
					c.Header("x-Created-Id", tChore.Id)
					c.Abort()
					c.JSON(201, "Task created and assigned successfully")
				}
			}
		} else if t.Type == "HomeWork" {
			var tHomeWork Models.TableHomeWork
			tHomeWork.Id = uuid.New().String()
			tHomeWork.OwnerId = ownerId
			tHomeWork.Status = t.Status
			tHomeWork.Course = t.Course
			tHomeWork.Details = t.Details
			tHomeWork.Type = t.Type
			dueDateT, err := time.Parse("2006-01-02", t.DueDate)
			if err != nil {
				c.Header("Content-Type", "text/plain")
				c.Next()
				c.String(400, "There is a problem with the date of the Task")
			} else {

				tHomeWork.DueDate = dueDateT
				if t.Description != "" || t.Size != "" {
					c.Header("Content-Type", "text/plain")
					c.Next()
					c.String(400, "You worte data that belong to ChoreTask")
				} else if t.Course == "" || t.Details == "" || t.DueDate == "" {
					c.Header("Content-Type", "text/plain")
					c.Next()
					c.String(400,
						"Required data fields are missing, data makes no sense, or data contains illegal values.")
				} else {
					if tHomeWork.Status.String() == "Active" {
						p.ActivteTaskCount = p.ActivteTaskCount + 1
						err1 := Models.UpdateAPerson(&p, p.ID)
						if err1 != nil {
							c.Header("Content-Type", "text/plain")
							c.Next()
							c.AbortWithStatus(http.StatusNotFound)
						}
					}
					err := Models.CreateAHomeWork(&tHomeWork)
					if err != nil {
						c.Header("Content-Type", "text/plain")
						c.Next()
						c.AbortWithStatus(http.StatusNotFound)
					} else {
						c.Header("Location", fmt.Sprintf("/api/people/%s/tasks/", tHomeWork.OwnerId))
						c.Header("x-Created-Id", tHomeWork.Id)
						c.Abort()
						c.JSON(201, "Task created and assigned successfully")
					}
				}
			}
		} else {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.JSON(400, "data contains illegal values")
		}
	}
}

type gettintStatus struct {
	Status string `json:"status"`
}

func GetTasks(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	var p Models.TablePersons
	var statustmp gettintStatus
	c.ShouldBindJSON(&statustmp)
	status := statustmp.Status
	ownerId := c.Params.ByName("id")
	// checking that there is a person:
	err1 := Models.GetAPerson(&p, ownerId)
	if err1 != nil {
		c.Header("Content-Type", "text/plain")
		c.Next()
		c.String(404, "A person with the id "+ownerId+"does not exist")

	} else {

		//status = Objects.CreateStatus(statustmp)
		if status == "Active" || status == "Done" {
			var hm []Models.TableHomeWork
			var ch []Models.TableChore
			errHomeWork := Models.GetAllTodosPerOwnerAndStatusHW(&hm, ownerId, Objects.CreateStatus(status))
			errChore := Models.GetAllTodosPerOwnerAndStatusCH(&ch, ownerId, Objects.CreateStatus(status))
			if errHomeWork != nil || errChore != nil {
				c.AbortWithError(http.StatusNotFound, errHomeWork)
				c.AbortWithError(http.StatusNotFound, errChore)
			} else {
				mixedArray := []interface{}{}
				for i := 0; i < len(ch); i++ {
					mixedArray = append(mixedArray, ch[i])
				}
				for i := 0; i < len(hm); i++ {
					mixedArray = append(mixedArray, hm[i])
				}
				//c.JSON(http.StatusOK, gin.H{"HomeWork Taks": hm, "Chore Tasls": ch})
				c.JSON(http.StatusOK, mixedArray)
			}
		} else { //return everything
			var hm []Models.TableHomeWork
			var ch []Models.TableChore
			errHomeWork := Models.GetAllTodosPerOwnerHW(&hm, ownerId)
			errChore := Models.GetAllTodosPerOwnerChore(&ch, ownerId)
			if errHomeWork != nil || errChore != nil {
				c.AbortWithError(http.StatusNotFound, errHomeWork)
				c.AbortWithError(http.StatusNotFound, errChore)
			} else {
				mixedArray := []interface{}{}
				for i := 0; i < len(ch); i++ {
					mixedArray = append(mixedArray, ch[i])
				}
				for i := 0; i < len(hm); i++ {
					mixedArray = append(mixedArray, hm[i])
				}
				c.JSON(http.StatusOK, mixedArray)
			}

		}
	}
}

func GetPersons(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	var person []Models.TablePersons
	err := Models.GetAllPersons(&person)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, person)
	}
}

func CreateAPerson(c *gin.Context) {
	//var person1 Models.TablePersons
	c.Header("Content-Type", "application/json")
	c.Next()
	var p Objects.Person
	errEmail := c.ShouldBindJSON(&p)
	if errEmail != nil {
		if p.Name == "" || p.FavoriteProgrammingLanguage == "" {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.String(400, "Required data fields are missing")
		} else {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.String(400, "The email "+p.Email+"  is not valid")
		}
	} else {
		var person Models.TablePersons
		person.ID = uuid.New().String()
		person.Email = p.Email
		person.FavoriteProgrammingLanguage = p.FavoriteProgrammingLanguage
		person.Name = p.Name
		person.ActivteTaskCount = 0

		err := Models.CreateAPerson(&person)
		if err != nil {
			c.Header("Content-Type", "text/plain")
			c.Status(400)
			c.Error(fmt.Errorf("Required data fields are missing, data makes no sense, or data contains illegal values"))
			c.String(400, "A person with email "+person.Email+"  already exists")
			c.Abort()

		} else {
			c.Header("Location", fmt.Sprintf("/api/people/%s", person.ID))
			c.Header("x-Created-Id", person.ID)
			c.Abort()
			c.JSON(201, "Person created successfully")
			//c.String(201, "Person created successfully")
		}
	}
}

func GetATodo(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	id := c.Params.ByName("id")
	var todo Models.TableChore
	err := Models.GetAChoreTask(&todo, id)
	if err == nil {
		c.JSON(http.StatusOK, todo)
	} else {
		var todo2 Models.TableHomeWork
		err2 := Models.GetAHomeWorkTask(&todo2, id)
		if err2 != nil {
			///c.AbortWithError(http.StatusNotFound, err2)
			c.JSON(404, "A task with the id"+id+"does not exist")
		} else {
			c.JSON(http.StatusOK, todo2)
		}
	}
}

func GetAStatus(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	id := c.Params.ByName("id")
	var chore Models.TableChore
	var homeWork Models.TableHomeWork
	err := Models.GetAChoreTask(&chore, id)
	if err == nil {
		// we know the task is a chore
		c.JSON(http.StatusOK, Objects.Status.String(chore.Status))
	} else {
		//chekcing if is a homeWork task
		err1 := Models.GetAHomeWorkTask(&homeWork, id)
		if err1 == nil {
			c.JSON(http.StatusOK, Objects.Status.String(homeWork.Status))
		} else {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.JSON(404, "A task with the id"+id+"does not exist")
			//c.AbortWithError(http.StatusNotFound, err1)
		}
	}
}

func GetAOwnerId(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	id := c.Params.ByName("id")
	var chore Models.TableChore
	var homeWork Models.TableHomeWork
	err := Models.GetAChoreTask(&chore, id)
	if err == nil {
		// we know the task is a chore
		c.JSON(http.StatusOK, chore.OwnerId)
	} else {
		//chekcing if is a homeWork task
		err1 := Models.GetAHomeWorkTask(&homeWork, id)
		if err1 == nil {
			c.JSON(http.StatusOK, homeWork.OwnerId)
		} else {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.JSON(404, "A task with the id"+id+"does not exist")
			//	c.AbortWithError(http.StatusNotFound, err1)
		}
	}
}

func ChangeStatus(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	id := c.Params.ByName("id")
	var chore Models.TableChore
	var homeWork Models.TableHomeWork
	//var getInfo settingOwnerStatus
	var statuscurr string
	err := Models.GetAChoreTask(&chore, id)
	if err == nil {
		// we know the task is a chore
		c.BindJSON(&statuscurr)
		if Objects.CreateStatus(statuscurr).String() == "unknown" || Objects.CreateStatus(statuscurr).String() == "" /*-1*/ {
			c.Header("Content-Type", "text/plain")
			c.Next()
			value := "value " + statuscurr + " is not a legal task status"
			c.JSON(400, value)
		} else {
			if Objects.CreateStatus(statuscurr).String() != chore.Status.String() {
				if chore.Status == Objects.CreateStatus("Active") {
					var person Models.TablePersons
					err1 := Models.GetAPerson(&person, chore.OwnerId)
					if err1 == nil {
						person.ActivteTaskCount = person.ActivteTaskCount - 1
						Models.UpdateAPerson(&person, person.ID)
					} else {
						c.Header("Content-Type", "text/plain")
						c.Next()
						c.String(404, "A person with the id"+chore.OwnerId+"does not exist")
					}
				} else {
					var person Models.TablePersons
					err1 := Models.GetAPerson(&person, chore.OwnerId)
					if err1 == nil {
						person.ActivteTaskCount = person.ActivteTaskCount + 1
						Models.UpdateAPerson(&person, person.ID)
					} else {
						c.Header("Content-Type", "text/plain")
						c.Next()
						c.String(404, "A person with the id"+chore.OwnerId+"does not exist")
					}
				}
			}
			chore.Status = Objects.CreateStatus(statuscurr)
			err1 := Models.UpdateAChoreTask(&chore, id)
			if err1 == nil {
				c.JSON(http.StatusNoContent, "Task's status updated successfully")
			} else {
				c.Header("Content-Type", "text/plain")
				c.Next()
				c.String(404, "A task with the id"+id+"does not exist")
			}
		}
	} else {
		err3 := Models.GetAHomeWorkTask(&homeWork, id)
		if err3 == nil {
			// we know the task is a chore
			c.BindJSON(&statuscurr)
			if Objects.CreateStatus(statuscurr).String() == "unknown" || Objects.CreateStatus(statuscurr).String() == "" /*-1*/ {
				value := "value" + statuscurr + "is not a legal task status"
				c.Header("Content-Type", "text/plain")
				c.Next()
				c.String(400, value)
			} else {
				if statuscurr != homeWork.Status.String() {
					if chore.Status.String() == "Active" {
						var person Models.TablePersons
						err1 := Models.GetAPerson(&person, homeWork.OwnerId)
						if err1 == nil {
							person.ActivteTaskCount = person.ActivteTaskCount - 1
							Models.UpdateAPerson(&person, person.ID)
						} else {
							c.Header("Content-Type", "text/plain")
							c.Next()
							c.String(404, "A person with the id"+homeWork.OwnerId+"does not exist")
						}
					} else {
						var person Models.TablePersons
						err1 := Models.GetAPerson(&person, homeWork.OwnerId)
						if err1 == nil {
							person.ActivteTaskCount = person.ActivteTaskCount + 1
							Models.UpdateAPerson(&person, person.ID)
						} else {
							c.Header("Content-Type", "text/plain")
							c.Next()
							c.String(404, "A person with the id"+homeWork.OwnerId+"does not exist")
						}
					}
				}
				homeWork.Status = Objects.CreateStatus(statuscurr)
				err1 := Models.UpdateAHomeWorkTask(&homeWork, id)
				if err1 == nil {
					c.JSON(http.StatusNoContent, "Task's status updated successfully")
				} else {
					c.Header("Content-Type", "text/plain")
					c.Next()
					c.String(404, "A task with the id"+id+"does not exist")
				}
			}
		} else {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.String(404, "A task with the id"+id+"does not exist")

		}

	}
}

func ChangeOwner(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	id := c.Params.ByName("id")
	var chore Models.TableChore
	var homeWork Models.TableHomeWork
	var newOwnerId string
	err := Models.GetAChoreTask(&chore, id)
	if err == nil {
		// we know the task is a chore
		c.BindJSON(&newOwnerId)
		// first check if the new person is exist
		var newPerson Models.TablePersons
		err1 := Models.GetAPerson(&newPerson, newOwnerId)
		if err1 != nil { // its mean this person isnot exist.
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.String(404, "The new Person with the id:"+newOwnerId+"is not exist")
		} else {
			if chore.Status.String() == "Active" {
				var oldPerson Models.TablePersons
				Models.GetAPerson(&oldPerson, chore.OwnerId) /// we didnt put here an error catch because we assume  the old person is exists
				oldPerson.ActivteTaskCount = oldPerson.ActivteTaskCount - 1
				Models.UpdateAPerson(&oldPerson, chore.OwnerId)
				newPerson.ActivteTaskCount = newPerson.ActivteTaskCount + 1
				Models.UpdateAPerson(&newPerson, newPerson.ID)
			}
			chore.OwnerId = newOwnerId
			err2 := Models.UpdateAChoreTask(&chore, id)
			if err2 == nil {
				c.JSON(http.StatusNoContent, "Task owner updated successfully")
			} else {
				c.Header("Content-Type", "text/plain")
				c.Next()
				c.String(404, "A task with the id"+id+"does not exist")
			}
		}
	} else {
		err3 := Models.GetAHomeWorkTask(&homeWork, id)
		if err3 == nil {
			// we know the task is a chore
			c.BindJSON(&newOwnerId)
			// first check if the new person is exist
			var newPerson Models.TablePersons
			err1 := Models.GetAPerson(&newPerson, newOwnerId)
			if err1 != nil { // its mean this person isnot exist.
				c.Header("Content-Type", "text/plain")
				c.Next()
				c.String(400, "The new Person with the id:"+newOwnerId+"is not exist")
			} else {
				if homeWork.Status.String() == "Active" {
					var oldPerson Models.TablePersons
					Models.GetAPerson(&oldPerson, homeWork.OwnerId) /// we didnt put here an error catch because we assume  the old person is exists
					oldPerson.ActivteTaskCount = oldPerson.ActivteTaskCount - 1
					Models.UpdateAPerson(&oldPerson, homeWork.OwnerId)
					newPerson.ActivteTaskCount = newPerson.ActivteTaskCount + 1
					Models.UpdateAPerson(&newPerson, newPerson.ID)
				}
				homeWork.OwnerId = newOwnerId
				err5 := Models.UpdateAHomeWorkTask(&homeWork, id)
				if err5 == nil {
					c.JSON(http.StatusNoContent, "Task owner updated successfully")
				} else {
					c.Header("Content-Type", "text/plain")
					c.Next()
					c.String(404, "A task with the id"+id+"does not exist")
				}
			}
		} else {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.String(404, "A task with the id"+id+"does not exist")
			//c.AbortWithError(http.StatusNotFound, err3)
		}
	}
}

func GetAPerson(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	id := c.Params.ByName("id")
	var p Models.TablePersons
	err := Models.GetAPerson(&p, id)
	if err != nil {
		c.Header("Content-Type", "text/plain")
		c.Next()
		c.JSON(http.StatusNotFound, "A person with the id "+id+"does not exist")

	} else {
		c.JSON(http.StatusOK, p)
	}
}

func UpdateATodo(c *gin.Context) {
	//var task Objects.Task
	c.Header("Content-Type", "application/json")
	c.Next()
	var chore Models.TableChore
	var homeWorkType Objects.HomeWork
	var homeWork Models.TableHomeWork
	id := c.Params.ByName("id")
	err := Models.GetAChoreTask(&chore, id)
	if err == nil {
		statusBeforeUpdate := chore.Status
		c.ShouldBindJSON(&chore)
		if (chore.Status.String() == "Active" || chore.Status.String() == "Done") && statusBeforeUpdate.String() != chore.Status.String() {
			if statusBeforeUpdate == Objects.CreateStatus("Active") {
				var person Models.TablePersons
				err1 := Models.GetAPerson(&person, chore.OwnerId)
				if err1 == nil {
					person.ActivteTaskCount = person.ActivteTaskCount - 1
					Models.UpdateAPerson(&person, person.ID)
				} else {
					c.Header("Content-Type", "text/plain")
					c.Next()
					c.JSON(404, "A person with the id"+chore.OwnerId+"does not exist")
				}
			} else {
				var person Models.TablePersons
				err1 := Models.GetAPerson(&person, chore.OwnerId)
				if err1 == nil {
					person.ActivteTaskCount = person.ActivteTaskCount + 1
					Models.UpdateAPerson(&person, person.ID)
				} else {
					c.Header("Content-Type", "text/plain")
					c.Next()
					c.JSON(404, "A person with the id"+chore.OwnerId+"does not exist")
				}
			}
			/// take the person

		}
		err1 := Models.UpdateAChoreTask(&chore, id)
		if err1 != nil {
			c.AbortWithError(http.StatusNotFound, err1)
		} else {
			c.JSON(http.StatusOK, chore)
		}
	} else {
		///check if is it a homework task

		err2 := Models.GetAHomeWorkTask(&homeWork, id)
		if err2 == nil {
			////we know that we found the id in homeWork task
			statusBeforeUpdate := homeWork.Status
			c.ShouldBindJSON(&homeWorkType) //homeWork inside
			if homeWorkType.DueDate != "" {
				dueDateT, errTime := time.Parse("2006-01-02", homeWorkType.DueDate)
				if errTime != nil {
					c.Header("Content-Type", "text/plain")
					c.Next()
					c.JSON(404, "There was a problem with the date, so the Date hasnt change")
				} else {
					homeWork.DueDate = dueDateT
				}
			}
			if homeWorkType.Details != "" {
				homeWork.Details = homeWorkType.Details
			}

			if homeWorkType.Course != "" {
				homeWork.Course = homeWorkType.Course
			}
			if homeWorkType.Status.String() != "" {
				homeWork.Status = homeWorkType.Status
			}
			if (homeWork.Status == Objects.CreateStatus("Active") || homeWork.Status == Objects.CreateStatus("Done")) && statusBeforeUpdate != homeWork.Status {
				if statusBeforeUpdate == Objects.CreateStatus("Active") {
					var person Models.TablePersons
					err1 := Models.GetAPerson(&person, homeWork.OwnerId)
					if err1 == nil {
						person.ActivteTaskCount = person.ActivteTaskCount - 1
						Models.UpdateAPerson(&person, person.ID)
					} else {
						c.Header("Content-Type", "text/plain")
						c.Next()
						c.JSON(404, "A person with the id"+homeWork.OwnerId+"does not exist")
					}
				} else {
					var person Models.TablePersons
					err1 := Models.GetAPerson(&person, homeWork.OwnerId)
					if err1 == nil {
						person.ActivteTaskCount = person.ActivteTaskCount + 1
						Models.UpdateAPerson(&person, person.ID)
					} else {
						c.Header("Content-Type", "text/plain")
						c.Next()
						c.JSON(404, "A person with the id"+homeWork.OwnerId+"does not exist")
					}
				}
				/// take the person

			}

			err3 := Models.UpdateAHomeWorkTask(&homeWork, id)
			if err3 != nil {
				c.AbortWithError(http.StatusNotFound, err3)
			} else {
				c.JSON(http.StatusOK, homeWork)
			}
		} else {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.JSON(404, "A task with the id"+id+"does not exist")
		}
	}

}

func UpdateAPerson(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Next()
	var p Models.TablePersons
	id := c.Params.ByName("id")
	err := Models.GetAPerson(&p, id)
	if err != nil {
		c.Header("Content-Type", "text/plain")
		c.Next()
		c.String(404, "A person with the id"+id+"does not exist")

	} else {
		errEmail := c.ShouldBindJSON(&p)
		if errEmail != nil {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.String(400, "The email "+p.Email+"  is not valid")
		} else {
			err = Models.UpdateAPerson(&p, id)
			if err != nil {
				c.Header("Content-Type", "text/plain")
				c.Next()
				c.AbortWithError(http.StatusNotFound, err)
			} else {
				c.JSON(http.StatusOK, p)
			}
		}
	}
}

func DeleteATodo(c *gin.Context) {
	var person Models.TablePersons
	var chore Models.TableChore
	var homeWork Models.TableHomeWork
	id := c.Params.ByName("id")
	err := Models.GetAChoreTask(&chore, id)
	if err == nil {
		/// we want to delete a choreTask
		if chore.Status == Objects.CreateStatus("Active") {
			//we want to increase 1 from person activate tasks
			errToGetPerson := Models.GetAPerson(&person, chore.OwnerId)
			if errToGetPerson == nil {
				person.ActivteTaskCount = person.ActivteTaskCount - 1
				Models.UpdateAPerson(&person, chore.OwnerId)
			} else {
				c.AbortWithError(http.StatusNotFound, errToGetPerson)
			}
		}
		//delete the chore Task
		errDeleteChore := Models.DeleteAChore(&chore, id)
		if errDeleteChore == nil {
			c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
		} else {
			c.AbortWithError(http.StatusNotFound, errDeleteChore)
		}
	} else {
		// checking if is a HomeWork Task
		err3 := Models.GetAHomeWorkTask(&homeWork, id)
		if err3 == nil {
			/// we want to delete a choreTask
			if homeWork.Status == Objects.CreateStatus("Active") {
				//we want to increase 1 from person activate tasks
				errToGetPerson := Models.GetAPerson(&person, homeWork.OwnerId)
				if errToGetPerson == nil {
					person.ActivteTaskCount = person.ActivteTaskCount - 1
					Models.UpdateAPerson(&person, homeWork.OwnerId)
				} else {
					c.AbortWithError(http.StatusNotFound, errToGetPerson)
				}
			}
			//delete the homeWork Task
			errDeleteHomeWork := Models.DeleteAHomeWork(&homeWork, id)
			if errDeleteHomeWork == nil {
				c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
			} else {
				c.AbortWithError(http.StatusNotFound, errDeleteHomeWork)
			}
		} else {
			c.Header("Content-Type", "text/plain")
			c.Next()
			c.JSON(404, "A task with the id "+id+"does not exist")

		}
	}
}

func DeleteAPerson(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.Next()
	var p Models.TablePersons
	var chore []Models.TableChore
	var homework []Models.TableHomeWork
	id := c.Params.ByName("id")
	err1 := Models.GetAPerson(&p, id)
	if err1 != nil {
		c.String(404, "A person with the id "+id+" does not exist")
	} else {
		Models.GetAllTodosPerOwnerChore(&chore, id)
		for i := 0; i < len(chore); i++ {
			errdelete := Models.DeleteAChore(&chore[i], chore[i].Id)
			if errdelete != nil {
				c.AbortWithError(404, errdelete)
			}
		}
		Models.GetAllTodosPerOwnerHW(&homework, id)
		for i := 0; i < len(homework); i++ {
			errdelete := Models.DeleteAHomeWork(&homework[i], homework[i].Id)
			if errdelete != nil {
				c.AbortWithError(404, errdelete)
			}
		}

		err := Models.DeleteAPerson(&p, id)
		if err != nil {
			c.AbortWithError(404, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
		}
	}
}
