package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MZIchenjl/ooi4/auth"
)

type ServiceHandler struct{}

func (self *ServiceHandler) GetOSAPI(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	loginID := r.Form.Get("login_id")
	password := r.Form.Get("password")
	if loginID != "" && password != "" {
		kancolle := auth.New(loginID, password)
		w.Header().Set("Content-Type", "application/json'")
		err, osapiURL := kancolle.GetOSAPI()
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  0,
				"message": err.Error(),
			})
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    1,
			"osapi_url": osapiURL,
		})
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func (self *ServiceHandler) GetFlash(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	loginID := r.Form.Get("login_id")
	password := r.Form.Get("password")
	if loginID != "" && password != "" {
		kancolle := auth.New(loginID, password)
		w.Header().Set("Content-Type", "application/json'")
		err, entryURL := kancolle.GetEntry()
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  0,
				"message": err.Error(),
			})
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    1,
			"flash_url": entryURL,
		})
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}
