package product

import (
	"encoding/json"
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler"
	"github.com/JTGlez/GoWeb-IT_V2/internal/models"
	"net/http"
)

func (s controllerProduct) PutProduct(w http.ResponseWriter, r *http.Request) {

	var putProduct models.ProductResponse
	if err := json.NewDecoder(r.Body).Decode(&putProduct); err != nil {
		handler.SetResponse(w, http.StatusBadRequest, nil, false, err, nil)
		return
	}

	if err := validate.Struct(putProduct); err != nil {
		handler.SetResponse(w, http.StatusBadRequest, nil, false, err, nil)
		return
	}

	updatedProduct, errProduct := s.productSvc.PutProduct(&putProduct)
	if errProduct != nil {
		handler.SetResponse(w, http.StatusNotFound, nil, false, errProduct, nil)
		return
	}

	handler.SetResponse(w, http.StatusOK, updatedProduct, true, nil, nil)
}
