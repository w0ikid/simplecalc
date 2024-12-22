package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name       string
		payload    string
		statusCode int
		response   map[string]string
	}{
		// Успешные запросы
		{"Simple Addition", `{"expression": "1+2"}`, http.StatusOK, map[string]string{"result": "3"}},
		{"Simple Multiplication", `{"expression": "2*3"}`, http.StatusOK, map[string]string{"result": "6"}},
		{"Simple Division", `{"expression": "4/2"}`, http.StatusOK, map[string]string{"result": "2"}},
		{"Simple Subtraction", `{"expression": "10-3"}`, http.StatusOK, map[string]string{"result": "7"}},
		{"Mixed Operators", `{"expression": "3+5*2"}`, http.StatusOK, map[string]string{"result": "13"}},
		{"Parentheses Priority", `{"expression": "(3+5)*2"}`, http.StatusOK, map[string]string{"result": "16"}},
		{"Negative Numbers", `{"expression": "-5+3"}`, http.StatusOK, map[string]string{"result": "-2"}},
		{"Negative in Parentheses", `{"expression": "(-5+3)*2"}`, http.StatusOK, map[string]string{"result": "-4"}},
		{"Nested Parentheses", `{"expression": "(1+(2*3))"}`, http.StatusOK, map[string]string{"result": "7"}},
		{"Double Parentheses", `{"expression": "((2+3)*4)"}`, http.StatusOK, map[string]string{"result": "20"}},
		{"Floating Point Numbers", `{"expression": "2.5+3.1"}`, http.StatusOK, map[string]string{"result": "5.6"}},
		{"Division with Floating Point", `{"expression": "5/2"}`, http.StatusOK, map[string]string{"result": "2.5"}},
		{"Spaces in Expression", `{"expression": " 2 + 3 "}`, http.StatusOK, map[string]string{"result": "5"}},

		// Ошибки
		{"Invalid Characters", `{"expression": "2+*"}`, http.StatusUnprocessableEntity, map[string]string{"error": "Expression is not valid"}},
		{"Unmatched Parentheses", `{"expression": "(1+2"}`, http.StatusUnprocessableEntity, map[string]string{"error": "Expression is not valid"}},
		{"Empty Expression", `{"expression": ""}`, http.StatusUnprocessableEntity, map[string]string{"error": "Expression is not valid"}},
		{"Invalid Symbol", `{"expression": "2+@3"}`, http.StatusUnprocessableEntity, map[string]string{"error": "Expression is not valid"}},
		{"Divide by Zero", `{"expression": "10/0"}`, http.StatusUnprocessableEntity, map[string]string{"error": "Expression is not valid"}},

		// Некорректные запросы
		{"Missing Expression Field", `{}`, http.StatusUnprocessableEntity, map[string]string{"error": "Expression is not valid"}},
		{"Invalid JSON", `{"expression": 2+2`, http.StatusUnprocessableEntity, map[string]string{"error": "Expression is not valid"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer([]byte(tt.payload)))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(CalculateHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tt.statusCode)
			}

			var actual map[string]string
			if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
				t.Fatalf("could not parse response body: %v", err)
			}

			if len(actual) != len(tt.response) {
				t.Errorf("handler returned unexpected body: got %v want %v", actual, tt.response)
			}

			for key, value := range tt.response {
				if actual[key] != value {
					t.Errorf("handler returned unexpected body: got %v want %v", actual, tt.response)
				}
			}
		})
	}
}
