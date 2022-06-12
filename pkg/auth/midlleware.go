package auth

import "github.com/gin-gonic/gin"

func (th *TokenHandler) RefreshMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken := ctx.Param("refresh")
		guid := ctx.Param("guid")
		if refreshToken == "" {
			ctx.AbortWithStatusJSON(401, "Invalid Api token")
			return
		}
		ok, err := th.T.checkDb(refreshToken, guid)
		if err != nil {
			ctx.AbortWithStatusJSON(500, "Server err")
			return
		}

		if ok {
			next(ctx)
		} else {
			ctx.AbortWithStatusJSON(401, "Invalid Api token")
			return
		}
	}
}
