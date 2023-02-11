package routes

import (
	"net/http"

	"github.com/cenkalti/dominantcolor"
	"github.com/ravener/img-api/utils"
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
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	color := dominantcolor.Find(img)
	hex := dominantcolor.Hex(color)

	utils.JSON(w, http.StatusOK, map[string]interface{}{
		"hex": hex,
		"rgb": map[string]uint8{"r": color.R, "g": color.G, "b": color.B},
	})
}
