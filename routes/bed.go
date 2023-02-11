package routes

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/ravener/img-api/utils"
)

var bedTemplate image.Image

func ImageBed(w http.ResponseWriter, r *http.Request) {
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

	file = r.FormValue("target")

	if file == "" {
		utils.Message(w, http.StatusBadRequest, "Missing 'target' query string.")
		return
	}

	img2, err := utils.GetImage(file)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	img = imaging.Resize(img, 100, 100, imaging.Box)
	img2 = imaging.Resize(img2, 70, 70, imaging.Box)

	ctx := gg.NewContextForImage(bedTemplate)

	ctx.DrawImage(img, 25, 100) // Image 1

	img = imaging.Resize(img, 104, 100, imaging.Box)
	ctx.DrawImage(img, 25, 300) // Image 2

	img = imaging.Resize(img, 70, 71, imaging.Box)
	ctx.DrawImage(img, 53, 450) // Image 3

	ctx.DrawImage(img2, 53, 575) // Image 4

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/bed.png")

	if err != nil {
		panic(err)
	}

	bedTemplate = img
}
