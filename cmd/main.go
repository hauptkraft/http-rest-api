package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/varangXI/go-gin-api/pkg/books"
	"github.com/varangXI/go-gin-api/pkg/common/db"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
