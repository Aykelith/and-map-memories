package andMapMemoriesDB

import (
	"database/sql"
	"math"
)

const pinsTableName = "pins";

type Pin struct {
	ID int64
	Name string
	PinType string
	Lat float64
	Lng float64
	Country string
	County string
	LocationName string
	LocationAddress string
}

type PinsTableHandler struct {
	db *sql.DB
}

func calculateCoordonateFromDB(integral int32, fractional int64) float64 {
	return float64(integral) + (float64(fractional) / math.Pow10(int(math.Log(float64(fractional)) * math.Log10E + 1)))
}

func CreatePinsTableHandler(db *sql.DB) PinsTableHandler {
	return PinsTableHandler{db}
}

func (e PinsTableHandler) GetPinsForMapDisplay() ([]Pin, error) {
	rows, err := e.db.Query("SELECT id,name,pin_type,lat_integral,lat_fractional,lng_integral,lng_fractional FROM " + pinsTableName);
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pins := make([]Pin, 0);

	for rows.Next() {
		var latIntegral int32
		var latFractional int64
		var lngIntegral int32
		var lngFractional int64

		var pin Pin;

		err = rows.Scan(&pin.ID, &pin.Name, &pin.PinType, &latIntegral, &latFractional, &lngIntegral, &lngFractional);
		if err != nil {
            return nil, err
        }

		pin.Lat = calculateCoordonateFromDB(latIntegral, latFractional)
		pin.Lng = calculateCoordonateFromDB(lngIntegral, lngFractional)

		pins = append(pins, pin)
	}

	return pins, nil
}

func (e PinsTableHandler) GetPinAddress(id int64) (Pin, error) {
	row := e.db.QueryRow("SELECT country,county,locationName,locationAddress FROM " + pinsTableName + " WHERE id=?", id);

	var pin Pin;

	err := row.Scan(&pin.Country, &pin.County, &pin.LocationName, &pin.LocationAddress);
	if err != nil {
		return Pin{}, err
	}

	return pin, nil
}