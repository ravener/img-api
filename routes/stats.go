package routes

import (
	"net/http"
	"runtime"
	"time"

	"github.com/ravener/img-api/utils"
)

var uptime int64

// XXX: Keep this version updated.
const VERSION = "1.2.5"

// Stats returns some information about the API Server process, like how much memory it uses.
func Stats(w http.ResponseWriter, r *http.Request) {
	stats := &runtime.MemStats{}

	if r.FormValue("noStats") == "true" {
		stats = nil
	} else {
		runtime.ReadMemStats(stats)
	}

	utils.JSON(w, http.StatusOK, map[string]interface{}{
		"version":    VERSION,
		"stats":      stats,
		"uptime":     time.Now().Unix() - uptime,
		"goroutines": runtime.NumGoroutine(),
	})
}

func init() {
	uptime = time.Now().Unix()
}
