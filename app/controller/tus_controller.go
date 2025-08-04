package controller

import (
	"log"
	"net/http"

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
		// TODO: change optimize file use queue kafka
		for {
			event := <-handler.CompleteUploads
			log.Printf("Upload %s finished\n", event.Upload.ID)
		}
	}()

	return gin.WrapH(http.StripPrefix("/files/", handler))
}
