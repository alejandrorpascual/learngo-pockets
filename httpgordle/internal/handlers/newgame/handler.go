package newgame

import "net/http"

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Creating a new game"))
}
