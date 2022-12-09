package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var led_flags map[int]bool
var vibration_flags map[int]bool
var call_flags map[int]bool

type Flag struct {
	Flag bool `json:"flag"`
}

func main() {
	led_flags = map[int]bool{}
	vibration_flags = map[int]bool{}
	call_flags = map[int]bool{}

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/:id/led", GetLED)
	r.POST("/:id/led", SetLED)

	r.GET("/:id/call", GetCall)
	r.POST("/:id/call", SetCall)

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
	c.JSON(http.StatusOK, json)
}

func SetLED(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var flag Flag
	if err := c.ShouldBindJSON(&flag); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	led_flags[id] = flag.Flag

	c.Status(http.StatusOK)
}

func SetCall(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var flag Flag
	if err := c.ShouldBindJSON(&flag); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	call_flags[id] = flag.Flag

	c.Status(http.StatusOK)
}

func GetCall(c *gin.Context) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	json := Flag{call_flags[id]}

	c.JSON(http.StatusOK, json)
}
