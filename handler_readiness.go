package main

import (
	"net/http"
)

type jsonResponse struct {
	//! In Go, if we put lowercase as the key of the struct, it wont be recognized and returned in the response (by the json.Marshal fn), thats why we use "JSON reflect"
	//? This `json:"someKey"` means that when the response comes, it wont show the key Status, but lowercase status
	Status string `json:"status"`
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, jsonResponse{
		Status: "success",
	})
}
