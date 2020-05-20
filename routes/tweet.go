package routes

import (
	"github.com/fogleman/gg"
	"github.com/pollen5/img-api/utils"
	"image"
	"net/http"
)

var tweetTemplate image.Image

func ImageTweet(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")

	if text == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Missing 'text' query string.\"}"))
		return
	}

	if len(text) > 165 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Text must not be longer than 165 characters\"}"))
		return
	}

	ctx := gg.NewContextForImage(tweetTemplate)

	err := ctx.LoadFontFace("assets/Roboto-Regular.ttf", 50)

	if err != nil {
		panic(err)
	}

	ctx.SetRGB(0, 0, 0)

	y := float64(170)

	for _, s := range utils.WordWrap(text, 40) {
		ctx.DrawString(s, 35, y)
		y += 50
	}

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/tweet.png")

	if err != nil {
		panic(err)
	}

	tweetTemplate = img
}
