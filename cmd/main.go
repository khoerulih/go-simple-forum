package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/khoerulih/go-simple-forum/internal/configs"
	"github.com/khoerulih/go-simple-forum/internal/handler/memberships"
	membershipRepo "github.com/khoerulih/go-simple-forum/internal/repository/memberships"
	"github.com/khoerulih/go-simple-forum/pkg/internalsql"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database: ", err)
	}

	_ = membershipRepo.NewRepository(db)

	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
