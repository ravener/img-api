package routes

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ravener/img-api/utils"
)

var tomTemplate image.Image

// Just a bit unsatisfied with the image sizes but gets the job done for now.
func ImageTom(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 175, 175, imaging.Box)

	ctx := gg.NewContextForImage(tomTemplate)

	ctx.DrawCircle(224+87, 70+87, 87)
	ctx.Clip()
	ctx.DrawImage(img, 224, 70)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/tom.png")

	if err != nil {
		panic(err)
	}

	tomTemplate = img
}
