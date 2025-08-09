package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func parseFloatQuery(r *http.Request, name string) (float64, error) {
	vals := r.URL.Query()
	s := vals.Get(name)
	if s == "" {
		return 0, fmt.Errorf("missing query parameter '%s'", name)
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number for '%s': %v", name, err)
	}
	return f, nil
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func calcHandler(op string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, Response{Error: "only GET is allowed"})
			return
		}

		a, err := parseFloatQuery(r, "a")
		if err != nil {
			writeJSON(w, http.StatusBadRequest, Response{Error: err.Error()})
			return
		}
		b, err := parseFloatQuery(r, "b")
		if err != nil {
			writeJSON(w, http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		var res float64
		switch op {
		case "add":
			res = a + b
		case "sub":
			res = a - b
		case "mul":
			res = a * b
		case "div":
			if b == 0 {
				writeJSON(w, http.StatusBadRequest, Response{Error: "division by zero"})
				return
			}
			res = a / b
		default:
			writeJSON(w, http.StatusBadRequest, Response{Error: "unknown operation"})
			return
		}

		writeJSON(w, http.StatusOK, Response{Result: res})
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", calcHandler("add"))
	mux.HandleFunc("/sub", calcHandler("sub"))
	mux.HandleFunc("/mul", calcHandler("mul"))
	mux.HandleFunc("/div", calcHandler("div"))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	addr := ":8080"
	log.Printf("Starting server on %s...", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}