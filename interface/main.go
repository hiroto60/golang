package main

import (
	"golang/interface/controller"
	"golang/interface/model"
)

func main() {
	db := model.GetDB()
	todoModel := model.NewTodoModel(db)
	todoController := controller.NewTodoController(todoModel)
	todoController.GetPosts()

	todoModel2 := model.NewTodoModel2(db)
	todoController2 := controller.NewTodoController2(todoModel2)
	todoController2.GetPosts2()
}
