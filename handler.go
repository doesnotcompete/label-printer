package main

import (
	"encoding/json"
	"net/http"
)

func printHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		setStatus(w, http.StatusBadRequest, "bad request")
		return
	}

	err = req.Validate()
	if err != nil {
		setStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	err = PrintTag(req.Text, req.CodeText, req.CodeType)
	if err != nil {
		setStatus(w, http.StatusInternalServerError, err.Error())
		return
	}

	setStatus(w, http.StatusOK, "tag printed")
}

func setStatus(w http.ResponseWriter, code int, msg string) error {
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(&Response{Status: msg})
}
