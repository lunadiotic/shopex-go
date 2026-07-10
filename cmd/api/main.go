package main

import (
	"fmt"
	"log"

	"github.com/lunadiotic/shopex-go/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.App.Name)
	fmt.Println(cfg.App.Version)
	fmt.Println(cfg.Server.Port)
}