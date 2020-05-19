package routes

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pollen5/img-api/utils"
	"image"
	"net/http"
)

var garbageTemplate image.Image

// Wew, this took me forever to make. Still not perfect though but gets the job done I guess.
func ImageGarbage(w http.ResponseWriter, r *http.Request) {
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

	img = imaging.Resize(img, 260, 250, imaging.Box)

	ctx := gg.NewContext(600, 600)

	ctx.DrawCircle(55+130, 70+130, 130)
	ctx.Clip()
	ctx.DrawImage(img, 55, 70)
	ctx.ResetClip()

	ctx.DrawImage(garbageTemplate, 0, 0)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/garbage.png")

	if err != nil {
		panic(err)
	}

	garbageTemplate = img
}
