package main

import (
	"math/rand"
	"os"
	"time"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

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

func main() {
	cfg, err := loadConfig()
	if err != nil || cfg.Port == "" {
		cfg.Port = "8080"
	}

	router := gin.Default()
	rand.Seed(time.Now().UnixNano())
	router.GET("/", serveMeme)

	router.Run(":" + cfg.Port)
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
