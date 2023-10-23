package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/michelpessoa/goexpertmichel/9-Apis/internal/dto"
	"github.com/michelpessoa/goexpertmichel/9-Apis/internal/entity"
	"github.com/michelpessoa/goexpertmichel/9-Apis/internal/infra/database"
)


type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Esta maneira não é recomendada pois deve-se ter uma camada de service
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
