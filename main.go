package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	db.AutoMigrate(&Todo{})

	todoHandler := TodoHandler{db: db}

	e.GET("/todos", todoHandler.GetTodos)
	e.POST("/todos", todoHandler.CreateTodo)
	e.DELETE("/todos/:id", todoHandler.DeleteTodo)

	e.Logger.Fatal(e.Start(":1323"))
}

type Todo struct {
	gorm.Model
	Description string `json:"Description"`
	Done        bool   `json:"Done"`
}

type TodoHandler struct {
	db *gorm.DB
}

func (t *TodoHandler) GetTodos(c echo.Context) error {
	var todos []Todo
	if err := t.db.Find(&todos).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, todos)
}

func (t *TodoHandler) CreateTodo(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := t.db.Create(&todo).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, todo)
}

func (t *TodoHandler) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	result := t.db.Delete(&Todo{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}
