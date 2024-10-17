package product

import "net/http"

func (s serviceProduct) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Implement me!!!"))
}
