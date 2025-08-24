package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nazgool97/telegram-pro-bot/api/graph/generated"
	"github.com/nazgool97/telegram-pro-bot/api/graph"
	dbmodel "github.com/nazgool97/telegram-pro-bot/api/internal/model"
)

func main() {
	dsn := "postgres://botuser:botpass@postgres:5432/botdb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	if err := db.AutoMigrate(&dbmodel.Flow{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)

	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})
	r.Any("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	log.Println("API started on :8080")
	if err := r.Run(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server: %v", err)
	}
}