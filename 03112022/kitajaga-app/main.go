package main

import (
    "net/http"
	"encoding/json"
	"encoding/base64"
	"log"
	"strings"
)

func protected(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
    userInfo := strings.TrimSpace(req.Header.Get("X-Userinfo"))
	if userInfo == "" {
		w.WriteHeader(http.StatusUnauthorized)
		resp["message"] = "No Userinfo header"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	decodedData, err := base64.StdEncoding.DecodeString(userInfo)
	if err != nil {
		log.Fatalf("Error happened in decoding base64. Err: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(decodedData)
	return
}


func main() {

    http.HandleFunc("/", protected)

    http.ListenAndServe(":8383", nil)
}