package main

//imports for routes
import (
	"net/http"
	"os"
	"server/api/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// router.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		// Add the necessary CORS headers
	// 		if origin := r.Header.Get("Origin"); origin != "" {
	// 			w.Header().Set("Access-Control-Allow-Origin", origin)
	// 			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// 		}

	// 		// Allow preflight requests
	// 		if r.Method == "OPTIONS" {
	// 			w.Header().Set("Access-Control-Allow-Origin", "*")
	// 			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// 			w.WriteHeader(http.StatusOK)
	// 			return
	// 		}
	// 		// Call the next handler
	// 		next.ServeHTTP(w, r)
	// 	})
	// })

	// Get the routes
	r := routes.GetRoutes()
	// Register the routes
	for _, route := range r {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	handler := cors.Default().Handler(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	// Start the server
	http.ListenAndServe("0.0.0.0:"+port, handler)
}
