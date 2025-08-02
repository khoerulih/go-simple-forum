package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/khoerulih/go-simple-forum/internal/configs"
	"github.com/khoerulih/go-simple-forum/internal/handler/memberships"
	"github.com/khoerulih/go-simple-forum/internal/handler/posts"
	membershipRepo "github.com/khoerulih/go-simple-forum/internal/repository/memberships"
	postRepo "github.com/khoerulih/go-simple-forum/internal/repository/posts"
	membershipSvc "github.com/khoerulih/go-simple-forum/internal/service/memberships"
	postSvc "github.com/khoerulih/go-simple-forum/internal/service/posts"
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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepository := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(cfg, membershipRepository)
	membershipsHandler := memberships.NewHandler(r, membershipService)
	membershipsHandler.RegisterRoute()

	postRepository := postRepo.NewRepository(db)
	postService := postSvc.NewService(cfg, postRepository)
	postsHandler := posts.NewHandler(r, postService)
	postsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
