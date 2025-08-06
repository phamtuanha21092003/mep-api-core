package controller

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/phamtuanha21092003/mep-api-core/app/service"
)

type ITusController interface {
	Handler() gin.HandlerFunc
}

type tusController struct {
	tusSer service.ITusService
}

func NewTusController(tusSer service.ITusService) ITusController {
	return &tusController{tusSer: tusSer}
}

func (contr *tusController) Handler() gin.HandlerFunc {
	handler, err := contr.tusSer.BuildHandler()
	if err != nil {
		panic("failed to build TUS handler: " + err.Error())
	}

	go func() {
		for {
			select {
			case event, ok := <-handler.CompleteUploads:
				if !ok {
					log.Println("✅ CompleteUploads channel closed")
					return
				}

				log.Printf("✅ Upload %s completed", event.Upload.ID)

				// Extract and decode filename
				base64Filename, exists := event.Upload.MetaData["filename"]
				if !exists {
					log.Printf("⚠️  Missing filename in metadata for upload %s", event.Upload.ID)
					continue
				}

				decodedName, err := base64.StdEncoding.DecodeString(base64Filename)
				if err != nil {
					log.Printf("❌ Error decoding filename: %v", err)
					continue
				}

				// Move file to filename
				oldPath := filepath.Join("./uploads", event.Upload.ID)
				newPath := filepath.Join("./uploads", string(decodedName))

				err = os.Rename(oldPath, newPath)
				if err != nil {
					log.Printf("❌ Error renaming file %s -> %s: %v", oldPath, newPath, err)
					continue
				}

				log.Printf("📦 File saved as: %s", newPath)

			case event, ok := <-handler.TerminatedUploads:
				if !ok {
					log.Println("⚠️ TerminatedUploads channel closed")
					return
				}
				log.Printf("🛑 Upload %s terminated", event.Upload.ID)

			case event, ok := <-handler.UploadProgress:
				if !ok {
					log.Println("⚠️ UploadProgress channel closed")
					return
				}
				log.Printf("📶 Upload progress: ID=%s, Offset=%d", event.Upload.ID, event.Upload.Size)
			}
		}
	}()

	return gin.WrapH(http.StripPrefix("/api/v1/files/tus/", handler))
}
