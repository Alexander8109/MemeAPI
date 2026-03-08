package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	rand.Seed(time.Now().UnixNano())
	router.GET("/", serveMeme)

	router.Run(":25566")
}

func serveMeme(ctx *gin.Context) {
	files, err := os.ReadDir("./images")
	if err != nil || len(files)==0 {
		ctx.String(500, "No memes found!")
		return
	}

	randomFile := files[rand.Intn(len(files))].Name()

	ctx.File("./images/" + randomFile)
}
