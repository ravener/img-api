package routes

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ravener/img-api/utils"
)

var truthTemplate image.Image

func ImageTruth(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 200, 250, imaging.Box)

	ctx := gg.NewContext(539, 1024)

	ctx.Rotate(-0.01)
	ctx.DrawImage(img, 20, 600)
	ctx.Rotate(0.01) // Undo rotation

	ctx.DrawImage(truthTemplate, 0, 0)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/truth.png")

	if err != nil {
		panic(err)
	}

	truthTemplate = img
}
