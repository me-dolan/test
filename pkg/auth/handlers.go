package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/me-dolan/test/internal/handlers"
)

type TokenHandler struct {
	T *Tokens
}

func NewHandler(T *Tokens) handlers.Handler {
	return &TokenHandler{T: T}
}

func (th *TokenHandler) Register(router *gin.Engine) {
	router.GET("/login/:guid", th.generateTokens)
}

func (th *TokenHandler) generateTokens(c *gin.Context) {
	guid := c.Param("guid")
	at, u, err := th.T.generateTokens(guid)
	if err != nil {
		c.AbortWithStatusJSON(500, "server err")
		return
	}
	err = th.T.creatDb(u, at)
	fmt.Println(err)
	if err != nil {
		c.AbortWithStatusJSON(500, "server err")
		return
	}
	c.JSON(200, at)
}
