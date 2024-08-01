package routing

import (
	"mdgkb/ankets-server/handlers/auth"
	"mdgkb/ankets-server/handlers/fields"
	"mdgkb/ankets-server/handlers/formfills"
	"mdgkb/ankets-server/handlers/forms"
	"mdgkb/ankets-server/handlers/formsections"
	"mdgkb/ankets-server/handlers/humans"
	"mdgkb/ankets-server/handlers/users"
	authRouter "mdgkb/ankets-server/routing/auth"

	humansRouter "mdgkb/ankets-server/routing/humans"
	usersRouter "mdgkb/ankets-server/routing/users"

	"github.com/gin-gonic/gin"

	fieldsRouter "mdgkb/ankets-server/routing/fields"
	formFillsRouter "mdgkb/ankets-server/routing/formfills"
	formsRouter "mdgkb/ankets-server/routing/forms"
	formsectionsRouter "mdgkb/ankets-server/routing/formsections"

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

	formsections.Init(helper)
	formsectionsRouter.Init(api.Group("/form-sections"), formsections.H)

	humans.Init(helper)
	humansRouter.Init(api.Group("/humans"), humans.H)

	forms.Init(helper)
	formsRouter.Init(api.Group("/forms"), forms.H)

	fields.Init(helper)
	fieldsRouter.Init(api.Group("/fields"), fields.H)

	formfills.Init(helper)
	formFillsRouter.Init(api.Group("/form-fills"), formfills.H)
}
