package product

import (
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (s controllerProduct) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.productSvc.GetProducts()
	if err != nil {
		handler.SetResponse(w, http.StatusInternalServerError, nil, false, err, nil)
	}
	count := len(products)
	handler.SetResponse(w, http.StatusOK, products, true, nil, &count)
}

func (s controllerProduct) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		handler.SetResponse(w, http.StatusBadRequest, nil, false, ErrorInvalidID, nil)
		return
	}

	product, errProduct := s.productSvc.GetProduct(uint64(numId))
	if errProduct != nil {
		handler.SetResponse(w, http.StatusNotFound, nil, false, errProduct, nil)
		return
	}

	handler.SetResponse(w, http.StatusOK, product, true, nil, nil)

}

func (s controllerProduct) GetProductsByPrice(w http.ResponseWriter, r *http.Request) {
	priceGtStr := r.URL.Query().Get("priceGt")

	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		handler.SetResponse(w, http.StatusBadRequest, nil, false, ErrorInvalidPrice, nil)
		return
	}

	products, errProduct := s.productSvc.GetProductsByPrice(priceGt)
	if errProduct != nil {
		handler.SetResponse(w, http.StatusNotFound, nil, false, errProduct, nil)
		return
	}

	count := len(products)
	handler.SetResponse(w, http.StatusOK, products, true, nil, &count)
}

func (s controllerProduct) GetProductByCodeValue(w http.ResponseWriter, r *http.Request) {
	codeValue := chi.URLParam(r, "code_value")

	product, err := s.productSvc.GetProductByCodeValue(codeValue)
	if err != nil {
		handler.SetResponse(w, http.StatusNotFound, nil, false, err, nil)
		return
	}

	handler.SetResponse(w, http.StatusOK, product, true, nil, nil)
}
