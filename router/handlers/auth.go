package handlers

import (
	"log"
	"net/http"
	"nimblestack/views"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := views.Auth().Render(r.Context(), w); err != nil {
			log.Println("Error rendering view:", err)
		}
	}

}
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := views.Auth().Render(r.Context(), w); err != nil {
			log.Println("Error rendering view:", err)
		}
	}
}
