package routes

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pollen5/img-api/utils"
	"image"
	"net/http"
)

var pictureTemplate image.Image

func ImagePicture(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 800, 650, imaging.Box)

	ctx := gg.NewContext(909, 1024)

	ctx.Rotate(-0.055)
	// backup: 27x10
	ctx.DrawImage(img, 27, 0)
	ctx.Rotate(0.055)                    // Undo rotation
	ctx.DrawImage(pictureTemplate, 0, 0) // Draw template.

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/picture.png")

	if err != nil {
		panic(err)
	}

	pictureTemplate = img
}
