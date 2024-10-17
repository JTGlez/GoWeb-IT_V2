package product

import (
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler"
	"net/http"
)

func (s serviceProduct) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.db.GetProducts()
	if err != nil {
		handler.SetResponse(w, http.StatusInternalServerError, nil, false, err, nil)
	}
	count := len(products)
	handler.SetResponse(w, http.StatusOK, products, true, nil, &count)
}
