package routes

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pollen5/img-api/utils"
	"image"
	"net/http"
)

var patrickTemplate image.Image

func ImagePatrick(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 240, 282, imaging.Box)

	ctx := gg.NewContext(437, 1024)

	ctx.Rotate(-0.1)
	ctx.DrawImage(img, 51, 394)
	ctx.Rotate(0.1) // Undo rotation
	ctx.DrawImage(patrickTemplate, 0, 0)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/patrick.png")

	if err != nil {
		panic(err)
	}

	patrickTemplate = img
}
