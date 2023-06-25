package andMapMemoriesAPI

import (
	"encoding/json"
	"log"
	"net/http"

	andMapMemoriesDB "aykelith/and-map-memories/src/pkg/db"
)

type getPinsReqData struct {
	MinLat float32
	MinLng float32
	MaxLat float32
	MaxLng float32
	ExcludePinsIDs []int32
}

func GetPins(server *http.ServeMux, pinsTable andMapMemoriesDB.PinsTableHandler) {
	server.HandleFunc("/api/pins/get/all", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			http.Error(res, "Not found", http.StatusNotFound);
			return
		}

		// var reqData getPinsReqData;
		// err := json.NewDecoder(req.Body).Decode(&reqData)
		// if err != nil {
		// 	http.Error(res, "Something went wrong", http.StatusInternalServerError);
		// 	return
		// }

		pins, err := pinsTable.GetPinsForMapDisplay()
		if err != nil {
			log.Println(err)
			http.Error(res, "Something went wrong", http.StatusInternalServerError);
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		log.Println(pins);
		pinsData, err := json.Marshal(pins)
		if err != nil {
			log.Println(err)
			http.Error(res, "Something went wrong", http.StatusInternalServerError);
			return
		}

		res.Write(pinsData)
	})
}