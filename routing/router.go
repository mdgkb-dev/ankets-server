package routing

import (
	"mdgkb/ankets-server/handlers/auth"
	"mdgkb/ankets-server/handlers/users"
	authRouter "mdgkb/ankets-server/routing/auth"

	usersRouter "mdgkb/ankets-server/routing/users"

	//"mdgkb/ankets-server/handlers/representative"
	"github.com/gin-gonic/gin"

	helperPack "github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/middleware"
	baseRouter "github.com/pro-assistance/pro-assister/routing"
)

func Init(r *gin.Engine, helper *helperPack.Helper) {
	m := middleware.CreateMiddleware(helper)
	api, apiNoToken := baseRouter.Init(r, helper)
	// api.Use(m.InjectClaims())
	api.Use(m.InjectFTSP())
	auth.Init(helper)
	authRouter.Init(apiNoToken.Group("/auth"), auth.H)

	users.Init(helper)
	usersRouter.Init(api.Group("/users"), users.H)
}
