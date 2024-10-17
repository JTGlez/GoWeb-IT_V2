package product

import (
	"encoding/json"
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler"
	"github.com/JTGlez/GoWeb-IT_V2/internal/models"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

func (s serviceProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product models.ProductResponse
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		handler.SetResponse(w, http.StatusBadRequest, nil, false, err, nil)
		return
	}

	if err := validate.Struct(product); err != nil {
		handler.SetResponse(w, http.StatusBadRequest, nil, false, err, nil)
		return
	}

	dbProduct, errProduct := s.db.CreateProduct(&product)
	if errProduct != nil {
		handler.SetResponse(w, http.StatusInternalServerError, nil, false, errProduct, nil)
		return
	}

	handler.SetResponse(w, http.StatusCreated, dbProduct, true, nil, nil)
}
