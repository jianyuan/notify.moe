package user

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/utils"
)

// Get redirects /+ to /+UserName
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in", nil)
	}

	if user.Nick == "" {
		return ctx.Error(http.StatusInternalServerError, "User did not set a nickname", nil)
	}

	return profile.Profile(ctx, user)
}
