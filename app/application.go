package app

import (
	"github.com/gin-gonic/gin"
	"github.com/leslesnoa/bookstore_oauth-api/client/cassandra"
	"github.com/leslesnoa/bookstore_oauth-api/domain/access_token"
	athttp "github.com/leslesnoa/bookstore_oauth-api/http/access_token/access_token_http"
	"github.com/leslesnoa/bookstore_oauth-api/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	// dbRepository := db.NewRepository()
	// atService := access_token.NewService(dbRepository)
	atService := access_token.NewService(db.NewRepository())
	atHandler := athttp.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8082")
}
