package api

import (
	"fmt"
	"net/http"
	"time"
	db "todoapp/db/sqlc"
	"todoapp/token"

	"github.com/gin-gonic/gin"
)

const (
	CookieTime = time.Hour * 1
)

type server struct {
	router     *gin.Engine
	store      *db.Queries
	tokenMaker token.Maker
}

func NewGinServer(store *db.Queries) error {
	//Define JWT secret key
	tokenMaker, err := token.NewJWTMaker("12345678901234567890123456789012")
	if err != nil {
		return fmt.Errorf("cannot create token maker: %w", err)
	}
	//Initialize gin and server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	server := &server{
		router:     router,
		store:      store,
		tokenMaker: tokenMaker,
	}
	//Set front-end js and css
	server.router.LoadHTMLGlob("web/*")
	server.router.Static("/view", "./view")
	server.router.POST("/users", server.createUser)
	server.router.POST("/login", server.loginUser)
	server.router.POST("/logout", server.Logout)
	server.router.GET("/refresh", server.RefreshCookie)
	//Set middleware Token verification
	authRoutes := server.router.Group("/").Use(authMiddlewareJWT(server.tokenMaker))
	authRoutes.POST("/action", server.createAction)
	authRoutes.PUT("/action", server.completedAction)
	authRoutes.DELETE("/action", server.deletedAction)
	authRoutes.GET("/action", server.listAction)
	authRoutes.POST("/subaction", server.createSubAction)
	authRoutes.PUT("/subaction", server.completedSubAction)
	authRoutes.DELETE("/subaction", server.deletSubedAction)
	//Set front-end home page
	server.router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,
			"index.html", nil)
	})
	//Run server
	return server.router.Run()

}
