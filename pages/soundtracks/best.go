package soundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Best renders the best soundtracks.
func Best(ctx *aero.Context) string {
	// Fetch all eligible tracks
	tracks := fetchAll()

	// Sort the tracks by number of likes
	arn.SortSoundTracksPopularFirst(tracks)

	// Render
	return render(ctx, tracks)
}
