package handlers

import (
	"log"
	"net/http"

	"github.com/artem-shestakov/go/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != err {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

	w.Write(d)
}
