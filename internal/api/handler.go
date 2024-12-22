package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"github.com/w0ikid/simplecalc/internal/service"
	pkgErrors "github.com/w0ikid/simplecalc/pkg/errors"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		return
	}

	result, err := service.Calculate(req.Expression)
	if err != nil {
		if errors.Is(err, pkgErrors.ErrInvalidExpression) || errors.Is(err, pkgErrors.ErrInvalidCharacter) || errors.Is(err, pkgErrors.ErrDivideByZero) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Error: "Internal server error"})
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Result: result})
}
