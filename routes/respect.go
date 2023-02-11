package routes

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ravener/img-api/utils"
)

var respectTemplate image.Image

func ImageRespect(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 60, 90, imaging.Box)

	ctx := gg.NewContext(720, 405)

	ctx.Rotate(-0.15)
	ctx.DrawImage(img, 110, 64)
	ctx.Rotate(0.15) // Undo rotation
	ctx.DrawImage(respectTemplate, 0, 0)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/respect.png")

	if err != nil {
		panic(err)
	}

	respectTemplate = img
}
