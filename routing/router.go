package routing

import (
	"mdgkb/ankets-server/handlers/auth"
	"mdgkb/ankets-server/handlers/humans"
	"mdgkb/ankets-server/handlers/questions"
	"mdgkb/ankets-server/handlers/researches"
	"mdgkb/ankets-server/handlers/users"
	"mdgkb/ankets-server/handlers/usersresearches"
	authRouter "mdgkb/ankets-server/routing/auth"

	humansRouter "mdgkb/ankets-server/routing/humans"
	usersRouter "mdgkb/ankets-server/routing/users"

	"github.com/gin-gonic/gin"

	questionsRouter "mdgkb/ankets-server/routing/questions"
	researchesRouter "mdgkb/ankets-server/routing/researches"
	usersresearchesRouter "mdgkb/ankets-server/routing/usersresearches"

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

	humans.Init(helper)
	humansRouter.Init(api.Group("/humans"), humans.H)

	usersresearches.Init(helper)
	usersresearchesRouter.Init(api.Group("/users-researches"), usersresearches.H)
	researches.Init(helper)
	researchesRouter.Init(api.Group("/researches"), researches.H)

	questions.Init(helper)
	questionsRouter.Init(api.Group("/questions"), questions.H)
}
