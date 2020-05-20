package routes

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pollen5/img-api/utils"
	"image"
	"net/http"
)

var fatherTemplate image.Image

func ImageFather(w http.ResponseWriter, r *http.Request) {
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

	text := r.FormValue("text")

	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Missing 'text' query string.\"}"))
		return
	}

	if len(text) > 42 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Text must not be longer than 42 characters\"}"))
		return
	}

	img = imaging.Resize(img, 180, 180, imaging.Box)
	img2 := imaging.Resize(img, 120, 120, imaging.Box)

	ctx := gg.NewContextForImage(fatherTemplate)

	err = ctx.LoadFontFace("assets/Roboto-Regular.ttf", 30)

	if err != nil {
		panic(err)
	}

	ctx.DrawCircle(370+60, 240+60, 60)
	ctx.Clip()
	ctx.DrawImage(img2, 370, 240)
	ctx.ResetClip()

	ctx.DrawCircle(20+90, 500+90, 90)
	ctx.Clip()
	ctx.DrawImage(img, 20, 500)

	ctx.ResetClip()
	ctx.SetRGB(0, 0, 0)

	y := float64(450)
	for _, s := range utils.WordWrap(text, 21) {
		ctx.DrawString(s, 5, y)
		y += 30
	}

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/father.png")

	if err != nil {
		panic(err)
	}

	fatherTemplate = img
}
