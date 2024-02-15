package main

import (
	"log"

	"antinolabs/handler"

	mysql "antinolabs/mysqldb"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := mysql.ConnectDb()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	mysql.DB = db

	router := gin.Default()

	router.GET("/v1/getblogpostsinfo", handler.GetblogpostsInfoHandler)
	router.GET("/v1/listblogposts", handler.GetListblogposts)

	router.PUT("/v1/createblogposts", handler.CreateblogpostsRecord)
	router.PUT("/v1/updateblogposts", handler.UpdateblogpostsRecord)
	router.PUT("/v1/deleteblogposts", handler.DeleteblogpostsRecord)

	router.Run("localhost:8080")
}
