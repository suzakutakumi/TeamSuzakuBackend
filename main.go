package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var led_flags map[int]bool
var vibration_flags map[int]bool

type Flag struct {
	Flag bool `json:"flag"`
}

func main() {
	led_flags = map[int]bool{}
	vibration_flags = map[int]bool{}

	r := gin.Default()

	r.GET("/:id/led", GetLED)
	r.POST("/:id/led", SetLED)

	r.GET("/:id/vibration", GetVibration)
	r.POST("/:id/vibration", SetVibration)

	r.GET("/", Index)
	r.Run("0.0.0.0:80")
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "nemui")
}

func GetLED(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	json := Flag{led_flags[id]}
	led_flags[id] = false
	c.JSON(http.StatusOK, json)
}

func SetLED(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	led_flags[id] = true
	c.Status(http.StatusOK)
}

func GetVibration(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	json := Flag{vibration_flags[id]}
	vibration_flags[id] = false
	c.JSON(http.StatusOK, json)
}

func SetVibration(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	vibration_flags[id] = true
	c.Status(http.StatusOK)
}
