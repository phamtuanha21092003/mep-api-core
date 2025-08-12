package service

import (
	"github.com/tus/tusd/v2/pkg/filestore"
	tusd "github.com/tus/tusd/v2/pkg/handler"
)

type ITusService interface {
	BuildHandler() (*tusd.UnroutedHandler, error)
}

type tusService struct {
}

func NewTusService() ITusService {
	return &tusService{}
}

func (s *tusService) BuildHandler() (*tusd.UnroutedHandler, error) {
	store := filestore.FileStore{
		Path: "./uploads",
	}

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	config := tusd.Config{
		BasePath:                "/api/v1/files/tus/",
		StoreComposer:           composer,
		NotifyCompleteUploads:   true,
		NotifyTerminatedUploads: true,
		NotifyUploadProgress:    true,
	}

	return tusd.NewUnroutedHandler(config)
}
