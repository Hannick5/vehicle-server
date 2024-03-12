package vehicle

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hannick5/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// var req CreateRequest
	var res string = r.PathValue("id")
	resInt, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		http.Error(rw, "Could not parse Int", http.StatusInternalServerError)
		return
	}

	delVehicle, err := d.store.Vehicle().Delete(
		r.Context(),
		resInt,
	)

	if err != nil {
		http.Error(rw, "Could not delete the vehicle", http.StatusInternalServerError)
		return
	}

	if delVehicle {
		rw.WriteHeader(http.StatusNoContent)
		fmt.Println("Vehicle deleted")
		return
	} else {
		http.Error(rw, "Vehicle not found", http.StatusNotFound)
		// rw.WriteHeader(http.StatusNotFound)
		// fmt.Println("Vehicle not found")
		return
	}

}
