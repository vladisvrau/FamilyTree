package rest

import (
	"github.com/vladisvrau/FamilyTree/lib/log"
	"github.com/vladisvrau/FamilyTree/pkg/middleware"
	"github.com/vladisvrau/FamilyTree/pkg/service"

	"net/http"
)

func NewRouter(logger log.Logger, services service.Services) *http.ServeMux {

	mux := http.NewServeMux()
	h := &handler{
		logger:   logger,
		services: services,
	}
	mux.HandleFunc("/health", middleware.LoggingHandler(NewHealthHandler(h).Handle, logger))
	mux.HandleFunc("/person", middleware.LoggingHandler(NewPersonHanlder(h).Handle, logger))
	mux.HandleFunc("/kinship", middleware.LoggingHandler(NewKinship(h).Handle, logger))
	mux.HandleFunc("/tree", middleware.LoggingHandler(NewTree(h).Handle, logger))
	return mux
}

type handler struct {
	logger   log.Logger
	services service.Services
}
