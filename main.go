package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pollen5/img-api/routes"
	"log"
	"net/http"
)

var port = flag.Int("p", 3030, "Change the port to listen to.")
var addr = flag.String("h", "127.0.0.1", "")
var secret = flag.String("s", "", "Set a password")
var dev = flag.Bool("d", false, "Start in development mode (disable browser cache)")

func main() {
	flag.Parse()

	router := chi.NewRouter()

	// Setup middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	if *secret != "" {
		router.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Localhost requests are immune to Authorization.
				if r.Host == fmt.Sprintf("127.0.0.1:%d", *port) || r.Host == fmt.Sprintf("localhost:%d", *port) {
					next.ServeHTTP(w, r)
					return
				}

				auth := r.Header.Get("Authorization")

				if auth != *secret {
					w.WriteHeader(401)
					w.Header().Set("Content-Type", "application/json")
					w.Write([]byte("{\"message\": \"Unauthorized\"}"))
					return
				}

				next.ServeHTTP(w, r)
			})
		})
	}

	if *dev {
		router.Use(middleware.NoCache)
		router.Mount("/debug", middleware.Profiler())
		log.Print("Starting in development mode. (Browser cache will be disabled and profiler will be mounted.)")
	}

	// Initialize all routes.
	router.Get("/ping", routes.Ping)
	router.Get("/stats", routes.Stats)

	// Image routes.
	router.Get("/religion", routes.ImageReligion)
	router.Get("/beautiful", routes.ImageBeautiful)
	router.Get("/fear", routes.ImageFear)
	router.Get("/sacred", routes.ImageSacred)
	router.Get("/painting", routes.ImagePainting)
	router.Get("/color", routes.ImageColor)
	router.Get("/delete", routes.ImageDelete)
	router.Get("/garbage", routes.ImageGarbage)
	router.Get("/tom", routes.ImageTom)
	router.Get("/bed", routes.ImageBed)
	router.Get("/crush", routes.ImageCrush)
	router.Get("/patrick", routes.ImagePatrick)
	router.Get("/respect", routes.ImageRespect)
	router.Get("/dipshit", routes.ImageDipshit)
	router.Get("/picture", routes.ImagePicture)
	router.Get("/tweet", routes.ImageTweet)
	router.Get("/truth", routes.ImageTruth)
	router.Get("/bobross", routes.ImageBobross)
	router.Get("/mask", routes.ImageMask)
	router.Get("/father", routes.ImageFather)
	router.Get("/achievement", routes.ImageAchievement)

	// Non-Image routes.
	router.Get("/dominantColor", routes.DominantColor)

	// Start the server.
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *addr, *port), router))
}
