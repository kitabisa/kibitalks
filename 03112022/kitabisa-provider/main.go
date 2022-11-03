package main

import (
    "net/http"
	"encoding/json"
	"log"
)

func auth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
    authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		resp["message"] = "Invalid token"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp["sub"] = "123456"
	resp["user_id"] = "123"
	resp["first_name"] = "Bherly"
	resp["last_name"] = "Novrandy"
	resp["email"] = "bherly@kitabisa.com"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}


func main() {

    http.HandleFunc("/auth", auth)

    http.ListenAndServe(":8181", nil)
}