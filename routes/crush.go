package routes

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pollen5/img-api/utils"
	"image"
	"net/http"
)

var crushTemplate image.Image

func ImageCrush(w http.ResponseWriter, r *http.Request) {
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

	file = r.FormValue("target")

	if file == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Missing 'target' query string.\"}"))
		return
	}

	img2, err := utils.GetImage(file)

	if err != nil {
		utils.JSON(w, 400, map[string]interface{}{
			"message": err.Error(),
		})

		return
	}

	img = imaging.Resize(img, 130, 130, imaging.Box)
	img2 = imaging.Resize(img2, 300, 380, imaging.Box)

	ctx := gg.NewContext(600, 873)

	ctx.Rotate(-0.07)
	ctx.DrawImage(img2, 125, 493) // Image 2
	ctx.Rotate(0.07)              // Undo rotation.
	ctx.DrawImage(crushTemplate, 0, 0)

	ctx.DrawCircle(405+65, 50+65, 65)
	ctx.Clip()
	ctx.DrawImage(img, 405, 50) // Image 1

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/crush.png")

	if err != nil {
		panic(err)
	}

	crushTemplate = img
}
