package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/khoerulih/go-simple-forum/internal/configs"
	"github.com/khoerulih/go-simple-forum/internal/handler/memberships"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFilename("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config: ", err)
	}

	cfg = configs.Get()

	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
