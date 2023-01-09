package main

import (
	"net/http"

	"github.com/vladisvrau/FamilyTree/lib/log"
	"github.com/vladisvrau/FamilyTree/pkg/rest"
	"github.com/vladisvrau/FamilyTree/pkg/service"
)

func main() {
	logger := log.NewLogger(log.Info)
	logger.Info("listening on port 3000")
	http.ListenAndServe(":3000", rest.NewRouter(logger, *service.NewServices()))
}
