package accounts

import (
	"encoding/json"
	"net/http"


	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)


func (h *handler) GET(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	param := chi.URLParam(r, "username")
	accountrepo := h.app.Dao.Account() // domain/repository の取得	
	user, err := accountrepo.FindByUsername(ctx,param)
	if err != nil {		
		httperror.InternalServerError(w, err)		
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

}