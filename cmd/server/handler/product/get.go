package product

import (
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (s serviceProduct) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.db.GetProducts()
	if err != nil {
		handler.SetResponse(w, http.StatusInternalServerError, nil, false, err, nil)
	}
	count := len(products)
	handler.SetResponse(w, http.StatusOK, products, true, nil, &count)
}

func (s serviceProduct) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		handler.SetResponse(w, http.StatusBadRequest, nil, false, handler.ErrorInvalidID, nil)
		return
	}

	product, errProduct := s.db.GetProductById(numId)
	if errProduct != nil {
		handler.SetResponse(w, http.StatusNotFound, nil, false, errProduct, nil)
		return
	}

	handler.SetResponse(w, http.StatusOK, product, true, nil, nil)

}
