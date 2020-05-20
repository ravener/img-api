package routes

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pollen5/img-api/utils"
	"image"
	"net/http"
)

var bobrossTemplate image.Image

func ImageBobross(w http.ResponseWriter, r *http.Request) {
	file := r.FormValue("avatar")

	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Missing 'avatar' query string.\"}"))
		return
	}

	img, err := utils.GetImage(file)

	if err != nil {
		utils.JSON(w, 400, map[string]interface{}{
			"message": err.Error(),
		})

		return
	}

	img = imaging.Resize(img, 428, 360, imaging.Box)

	ctx := gg.NewContext(600, 775)

	ctx.Rotate(0.03)
	ctx.DrawImage(img, 27, 71)
	ctx.Rotate(-0.03) // Undo rotation

	ctx.DrawImage(bobrossTemplate, 0, 0)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/bobross.png")

	if err != nil {
		panic(err)
	}

	bobrossTemplate = img
}
