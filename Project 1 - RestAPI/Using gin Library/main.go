package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json: "id"`
	Item      string `json: "item"`
	Completed bool   `json: "completed"`
}

var todos = []todo{
	{ID: "1", Item: "Download VSC", Completed: false},
	{ID: "2", Item: "Bring Packages", Completed: false},
	{ID: "3", Item: "Learn Go", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(context *gin.Context) {
	id := context.Param("id")
	for _, todo := range todos {
		if todo.ID == id {
			context.IndentedJSON(http.StatusOK, todo)
			return
		}
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
}

func updateTodoByID(context *gin.Context) {
	id := context.Param("id")
	var updatedTodo todo
	if err := context.BindJSON(&updatedTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	for index, todo := range todos {
		if todo.ID == id {
			todos[index] = updatedTodo
			context.IndentedJSON(http.StatusOK, updatedTodo)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func deleteTodoByID(context *gin.Context){
	id := context.Param("id")
	for index, todo := range todos{
		if todo.ID == id{
			todos = append(todos[:index], todos[index+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		} 
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/list", getTodos)
	router.POST("/add", addTodos)
	router.PATCH("/update/:id", updateTodoByID)
	router.PUT("/update/:id", updateTodoByID)
	router.GET("/list/:id", getTodoByID)
	router.DELETE("/delete/:id", deleteTodoByID)
	router.Run("localhost:9090")
}
