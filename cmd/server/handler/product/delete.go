package product

import (
	"github.com/JTGlez/GoWeb-IT_V2/cmd/server/handler"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s controllerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	codeValue := chi.URLParam(r, "code_value")

	if err := s.productSvc.DeleteProduct(codeValue); err != nil {
		handler.SetResponse(w, http.StatusNotFound, nil, false, err, nil)
		return
	}

	handler.SetResponse(w, http.StatusOK, nil, true, nil, nil)

}
