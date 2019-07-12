package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/s14228so/WilicoExtra/GolangAPI/controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
	// router.Run(":3000") // ポートをハードコーディングしたい場合
}

// type Context struct {
// 	writermem responseWriter
// 	Request   *http.Request
// 	Writer    ResponseWriter

// 	Params   Params
// 	handlers HandlersChain
// 	index    int8
// }

func routeGet(c *gin.Context) {
	message := "ok"
	c.String(http.StatusOK, message)
}

func router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	u := r.Group("/users")
	{
		ctrl := controller.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	p := r.Group("/plans")
	{
		ctrl := controller.PlanController{}
		p.GET("", ctrl.Index)
		p.GET("/:id", ctrl.Show)
		p.POST("", ctrl.Create)
		p.PUT("/:id", ctrl.Update)
		p.DELETE("/:id", ctrl.Delete)
	}

	return r
}
