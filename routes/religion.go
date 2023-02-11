package routes

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ravener/img-api/utils"
)

var religionTemplate image.Image

func ImageReligion(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 500, 500, imaging.Box)

	ctx := gg.NewContextForImage(religionTemplate)

	ctx.DrawRoundedRectangle(200, 230, 500, 500, 25)
	ctx.Clip()
	ctx.DrawImage(img, 200, 230)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/religion.png")

	if err != nil {
		panic(err)
	}

	religionTemplate = img
}
