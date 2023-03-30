package main

import (
	"fmt"
	"sample-api/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test2"
	u.Email = "test2@example.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()

	/*
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		fmt.Println("削除しました")
		fmt.Println(u)
	*/

	/*
		user, _ := models.GetUser(4)
		todos, _ := user.GetTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
		t, _ := models.GetTodo(1)
		t.Content = "Update Todo"
		t.UpdateTodo()
	*/

	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}
