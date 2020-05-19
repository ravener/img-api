package routes

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pollen5/img-api/utils"
	"image"
	"net/http"
)

var fearTemplate image.Image

func ImageFear(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 250, 250, imaging.Box)

	ctx := gg.NewContextForImage(fearTemplate)

	ctx.DrawImage(img, 260, 516)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/fear.png")

	if err != nil {
		panic(err)
	}

	fearTemplate = img
}
