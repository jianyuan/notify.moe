package fullpage

import (
	"encoding/json"
	"io/ioutil"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

var organization map[string]interface{}
var organizationString string

func init() {
	// Load structured data
	organizationBytes, _ := ioutil.ReadFile("organization.json")
	json.Unmarshal(organizationBytes, &organization)
	organizationBytes, _ = json.Marshal(organization)
	organizationString = string(organizationBytes)
}

// Render layout.
func Render(ctx *aero.Context, content string) string {
	user := utils.GetUser(ctx)
	openGraph, _ := ctx.Data.(*arn.OpenGraph)

	// Make output order deterministic to profit from Aero caching.
	// To do this, we need to create slices and sort the tags.
	var meta []string
	var tags []string

	if openGraph != nil {
		for name := range openGraph.Meta {
			meta = append(meta, name)
		}

		sort.Strings(meta)

		for name := range openGraph.Tags {
			tags = append(tags, name)
		}

		sort.Strings(tags)
	}

	return components.Layout(ctx.App, ctx, user, openGraph, meta, tags, organizationString, content)
}
