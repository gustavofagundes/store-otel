package store

import (
	"fmt"
	"net/http"
)

func Itens(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Laranja, Banana, Laranja")
}
