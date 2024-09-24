package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gopkg.in/validator.v2"
	"hkn-be/config"
	"hkn-be/handlers"
	"hkn-be/infras/db"
	"hkn-be/infras/jwt_infra"
	"hkn-be/infras/mail"
	"hkn-be/infras/redis"
	"hkn-be/repositories"
	"hkn-be/router"
	"hkn-be/services"
	"log"
)

//import "gorm.io/gen"

func main() {
	//init config
	cfg := config.InitConfig()

	migrate(&cfg.DB)

	//init database
	gormDb, err := db.InitGorm(&cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}

	////generate models
	//g := gen.NewGenerator(
	//	gen.Config{
	//		OutPath: "gen_models", // Output directory for generated files
	//	},
	//)
	//g.UseDB(gormDb)
	//
	//// Generate structs from all tables of the current database
	//g.ApplyBasic(g.GenerateAllTable()...)
	//// Execute the code generation
	//g.Execute()

	//init redis
	redisServer := redis.NewRedisServer(&cfg.Redis)
	rdb, err := redisServer.Connect(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	//init mail service
	mailService := mail.NewMailServiceInterface(&cfg.Mailer)

	//init jwt service
	jwtService := jwt_infra.NewJwt(&cfg.JwtConfig, rdb)

	validator.SetPrintJSON(true)

	//
	ctxRepositories := repositories.InitRepositories(gormDb)
	ctxServices := services.InitServices(ctxRepositories, cfg.Server, jwtService, rdb, mailService)
	ctxHandlers := handlers.InitHandlers(ctxServices)

	r := router.InitRouter(ctxHandlers, jwtService)

	r.Logger.Fatal(r.Start(cfg.Server.Port))
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func migrate(cfg *config.DBConfig) {
	connString := fmt.Sprintf(
		"host=%s port=5432 dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Name, cfg.Username, cfg.Password,
	)

	database, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(database, "migrations"); err != nil {
		log.Println(err.Error())
	}
}
