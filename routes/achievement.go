package routes

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pollen5/img-api/utils"
	"image"
	"net/http"
)

var achievementTemplate image.Image

func ImageAchievement(w http.ResponseWriter, r *http.Request) {
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

	if len(text) > 21 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\": \"Text must not be longer than 21 characters\"}"))
		return
	}

	img = imaging.Resize(img, 40, 40, imaging.Box)

	ctx := gg.NewContextForImage(achievementTemplate)

	ctx.DrawCircle(10+20, 10+20, 20)
	ctx.Clip()
	ctx.DrawImage(img, 10, 10)

	ctx.ResetClip()

	err = ctx.LoadFontFace("assets/Mojangles.ttf", 20)

	if err != nil {
		panic(err)
	}

	ctx.SetRGB(1, 1, 1)
	ctx.DrawString(text, 59, 50)

	// Signal the response type.
	w.Header().Set("Content-Type", "image/png")
	// Send
	ctx.EncodePNG(w)
}

func init() {
	img, err := gg.LoadPNG("assets/achievement.png")

	if err != nil {
		panic(err)
	}

	achievementTemplate = img
}
