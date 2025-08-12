package main

import (
	"fmt"
	"net/http"

	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

func main() {
	// Create uploads folder storage
	store := filestore.FileStore{
		Path: "./uploads",
	}

	// Compose storage backend
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	// Create TUS handler
	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/files/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	// Log when uploads complete
	go func() {
		for {
			event := <-handler.CompleteUploads
			fmt.Printf("Upload %s finished\n", event.Upload.ID)
		}
	}()

	// Wrap handler with CORS middleware
	withCORS := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Upload-Offset, Upload-Length, Tus-Resumable, Upload-Metadata")
			w.Header().Set("Access-Control-Expose-Headers", "Upload-Offset, Upload-Length, Tus-Resumable, Upload-Metadata, Location")

			// Handle preflight
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			h.ServeHTTP(w, r)
		})
	}

	// Serve
	http.Handle("/files/", http.StripPrefix("/files/", withCORS(handler)))

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(fmt.Errorf("Unable to listen: %s", err))
	}
}
