package routes

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ravener/img-api/utils"
)

var beautifulTemplate image.Image

func ImageBeautiful(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 147, 170, imaging.Box)

	ctx := gg.NewContextForImage(beautifulTemplate)

	ctx.DrawImage(img, 433, 44)

	ctx.DrawImage(img, 436, 380)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/beautiful.png")

	if err != nil {
		panic(err)
	}

	beautifulTemplate = img
}
