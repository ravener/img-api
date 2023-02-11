package routes

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ravener/img-api/utils"
)

var sacredTemplate image.Image

// Not super satisfied with this but gets the job done.
func ImageSacred(w http.ResponseWriter, r *http.Request) {
	file := r.FormValue("avatar")

	if file == "" {
		utils.Message(w, http.StatusBadRequest, "Missing 'avatar' query string.")
		return
	}

	img, err := utils.GetImage(file)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)

		return
	}

	img = imaging.Resize(img, 239, 284, imaging.Box)

	ctx := gg.NewContextForImage(sacredTemplate)

	ctx.Rotate(0.02)
	ctx.DrawImage(img, 670, 569)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/sacred.png")

	if err != nil {
		panic(err)
	}

	sacredTemplate = img
}
