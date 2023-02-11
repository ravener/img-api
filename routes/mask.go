package routes

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ravener/img-api/utils"
)

var maskTemplate image.Image

// Idk if this was a good idea but it's corona time so why not.
func ImageMask(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 512, 512, imaging.Box)

	ctx := gg.NewContextForImage(img)

	ctx.DrawImage(maskTemplate, 100, 256)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/mask.png")

	if err != nil {
		panic(err)
	}

	maskTemplate = imaging.Resize(img, 300, 250, imaging.Box)
}
