package routes

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

var uptime int64

// XXX: Keep this version updated.
const VERSION = "1.0.3"

// Stats returns some information about the API Server process, like how much memory it uses.
func Stats(w http.ResponseWriter, r *http.Request) {
	stats := &runtime.MemStats{}

	if r.FormValue("noStats") == "true" {
		stats = nil
	} else {
		runtime.ReadMemStats(stats)
	}

	bytes, err := json.Marshal(map[string]interface{}{
		"version":    VERSION,
		"stats":      stats,
		"uptime":     time.Now().Unix() - uptime,
		"goroutines": runtime.NumGoroutine(),
	})

	if err != nil {
		panic(err)
	}

	w.Write(bytes)
}

func init() {
	uptime = time.Now().Unix()
}
