package product

import (
	"encoding/json"
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler"
	"github.com/JTGlez/GoWeb-IT_V2/internal/models"
	"net/http"
)

func (s controllerProduct) PatchProduct(w http.ResponseWriter, r *http.Request) {

	var patchProduct models.ProductResponse
	if err := json.NewDecoder(r.Body).Decode(&patchProduct); err != nil {
		handler.SetResponse(w, http.StatusBadRequest, nil, false, err, nil)
		return
	}

	if patchProduct.CodeValue == "" {
		handler.SetResponse(w, http.StatusBadRequest, nil, false, ErrorMissingCodeValue, nil)
		return
	}

	updatedProduct, errProduct := s.productSvc.PatchProduct(&patchProduct)
	if errProduct != nil {
		handler.SetResponse(w, http.StatusNotFound, nil, false, errProduct, nil)
		return
	}

	handler.SetResponse(w, http.StatusOK, updatedProduct, true, nil, nil)
}
