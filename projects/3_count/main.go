package main

// некоторые импорты нужны для проверки
import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv" // вдруг понадобиться вам ;)
)

var counter int

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"count": counter})

	case http.MethodPost:
		var req map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Ошибка при парсинге запроса", http.StatusBadRequest)
			return
		}
		count, ok := req["count"].(float64)
		if !ok {
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}
		counter += int(count)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Счетчик увеличен на %d Счетчик: %d", int(count), int(counter))

	default:
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	}
}
func main() {
	http.HandleFunc("/count", handler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}