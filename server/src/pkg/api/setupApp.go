package andMapMemoriesAPI

import (
	"net/http"

	andMapMemoriesDB "aykelith/and-map-memories/src/pkg/db"
)

func SetupApp(server *http.ServeMux, pinsTable andMapMemoriesDB.PinsTableHandler) {
	GetPins(server, pinsTable)
}
