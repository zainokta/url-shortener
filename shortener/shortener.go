package shortener

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Engine *gin.Engine
}

func New() App {
	engine := gin.Default()

	initRouter(engine)

	return App{
		Engine: engine,
	}
}
