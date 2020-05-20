package routes

import (
	"github.com/fogleman/gg"
	"image"
	"net/http"
)

var dipshitTemplate image.Image

func ImageDipshit(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")

	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Missing 'text' query string.\"}"))
		return
	}

	if len(text) > 33 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Text must not be longer than 33 characters\"}"))
		return
	}

	ctx := gg.NewContextForImage(dipshitTemplate)

	err := ctx.LoadFontFace("assets/Roboto-Regular.ttf", 20)

	if err != nil {
		panic(err)
	}

	ctx.SetRGB(0, 0, 0)
	ctx.DrawString(text, 140, 76)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/dipshit.png")

	if err != nil {
		panic(err)
	}

	dipshitTemplate = img
}
