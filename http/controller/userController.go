package controller

import (
	"log"
	"management-asset/model"
	"management-asset/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service service.Services
}

func (cnt *Controller) Registration(c *gin.Context) {
	var param model.User
	if err := c.BindJSON(&param); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}

	data, err := cnt.Service.Register(param)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}

func (cnt *Controller) Login(c *gin.Context) {
	var param model.User
	if err := c.BindJSON(&param); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}

	data, err := cnt.Service.Login(param.Username, param.Password)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}
func (cnt *Controller) FindAll(c *gin.Context) {
	var param model.GetAll
	if err := c.Bind(&param); err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}

	data, err := cnt.Service.FindAll(param.SortType, param.SortBy, param.Search, param.From, param.Limit)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}

func (cnt *Controller) FindByID(c *gin.Context) {
	id := c.Param("id")

	data, err := cnt.Service.FindByID(id)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}

func (cnt *Controller) Update(c *gin.Context) {
	var param model.User
	if err := c.BindJSON(&param); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}
	param.ID = c.Param("id")

	data, err := cnt.Service.Update(param)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}

func (cnt *Controller) Delete(c *gin.Context) {
	id := c.Param("id")

	data, err := cnt.Service.Delete(id)
	if err != nil {
		c.JSON(200, gin.H{"status": false, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success", "data": data})
}
