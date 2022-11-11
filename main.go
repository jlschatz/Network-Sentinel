package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"sentinel/service"

	"log"
	"net/http"
	"time"
)

func main() {

	svc := service.NewService()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	r.Use(c.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		jsonStr, _ := json.Marshal(svc.Interfaces)
		w.Write(jsonStr)
	})

	log.Println("Server listening on port: 8080")
	http.ListenAndServe(":8080", r)

}

//func main() {
//	// Get a list of all interfaces.
//	ifaces, err := net.Interfaces()
//	if err != nil {
//		panic(err)
//	}
//
//	var wg sync.WaitGroup
//	for _, iface := range ifaces {
//		wg.Add(1)
//		// Start up a scan on each interface.
//		go func(iface net.Interface) {
//			defer wg.Done()
//			if err := scan(&iface); err != nil {
//				log.Printf("interface %v: %v", iface.Name, err)
//			}
//		}(iface)
//	}
//	// Wait for all interfaces' scans to complete.  They'll try to run
//	// forever, but will stop on an error, so if we get past this Wait
//	// it means all attempts to write have failed.
//	wg.Wait()
//}
