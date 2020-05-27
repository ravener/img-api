package routes

import (
	"github.com/pollen5/img-api/utils"
	"github.com/cenkalti/dominantcolor"
	"net/http"
)

func DominantColor(w http.ResponseWriter, r *http.Request) {
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

	color := dominantcolor.Find(img)
	hex := dominantcolor.Hex(color)

	utils.JSON(w, 200, map[string]interface{}{
		"hex": hex,
		"rgb": map[string]uint8{ "r": color.R, "g": color.G, "b": color.B },
	})
}
