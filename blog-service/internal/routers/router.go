package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middleware.Translations())
	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.POST("/auth", v1.GetAuth)
	tag := v1.NewTag()
	article := v1.NewAriticle()
	apiv1 := r.Group("/api/v1")
	{
		// 标签相关路由
		tagGroup := apiv1.Group("/tags")
		{
			tagGroup.POST("", tag.Create)
			tagGroup.DELETE("/:id", tag.Delete)
			tagGroup.PUT("/:id", tag.Update)
			tagGroup.PATCH("/:id/state", tag.Update)
			tagGroup.GET("", tag.List)
		}
		// 文章相关路由
		articleGroup := apiv1.Group("/articles")
		{
			articleGroup.POST("", article.Create)
			articleGroup.DELETE("/:id", article.Delete)
			articleGroup.PUT("/:id", article.Update)
			articleGroup.PATCH("/:id/state", article.Update)
			articleGroup.GET("/:id", article.Get)
			articleGroup.GET("", article.List)
		}
	}

	return r
}
