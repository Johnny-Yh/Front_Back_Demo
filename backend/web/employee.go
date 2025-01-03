package web

import (
	"fmt"
	"net/http"

	"back_go/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func Update(c *gin.Context) {
	employee := model.Employee{}
	if err := c.BindJSON(&employee); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	if err := model.DefaultEmployeeDb.Update(employee); err != nil {
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func Remove(c *gin.Context) {
	id := c.Param("id")
	if err := model.DefaultEmployeeDb.Remove(cast.ToInt32(id)); err != nil {
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, 1)
}

func FindById(c *gin.Context) {
	id := c.Param("id")
	employee, err := model.DefaultEmployeeDb.GetOne(cast.ToInt32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, employee)
}

func FindAll(c *gin.Context) {
	employees, err := model.DefaultEmployeeDb.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, employees)
}

func Save(c *gin.Context) {
	employee := model.Employee{}
	if err := c.BindJSON(&employee); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	if err := model.DefaultEmployeeDb.Save(employee); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, -1)
		return
	}
	c.JSON(http.StatusOK, 1)
}
