package routes

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ravener/img-api/utils"
)

var kaguyaTemplate image.Image

func ImageKaguya(w http.ResponseWriter, r *http.Request) {
	file := r.FormValue("avatar")

	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Missing 'avatar' query string.\"}"))
		return
	}

	img, err := utils.GetImage(file)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)

		return
	}

	img = imaging.Resize(img, 140, 180, imaging.Box)

	ctx := gg.NewContext(640, 832)

	ctx.Rotate(-0.25)
	ctx.DrawImage(img, 125, 410)
	ctx.Rotate(0.25) // Undo rotation
	ctx.DrawImage(kaguyaTemplate, 0, 0)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/kaguya.png")

	if err != nil {
		panic(err)
	}

	kaguyaTemplate = img
}
