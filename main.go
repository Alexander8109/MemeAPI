package main

import (
	"math/rand"
	"os"
	"time"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// config handling
type Config struct {
	Port string `json:"port"`
}

func loadConfig() (Config, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

// Main
func main() {
	cfg, err := loadConfig()
	if err != nil || cfg.Port == "" {
		cfg.Port = "8080"
	}

	router := gin.Default()
	rand.Seed(time.Now().UnixNano())

	router.GET("/random", serveRandomMeme)

	router.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		path := "images/" + name

		if _, err := os.Stat(path); err == nil {
			ctx.File(path)
			return
		}
		ctx.JSON(404, gin.H{"error": "Not found"})
	})

	router.Run(":" + cfg.Port)
}

// randomizer
func serveRandomMeme(ctx *gin.Context) {
	files, err := os.ReadDir("./images")
	if err != nil || len(files)==0 {
		ctx.String(500, "No memes found!")
		return
	}

	randomFile := files[rand.Intn(len(files))].Name()

	ctx.File("./images/" + randomFile)
}
