package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const HeaderContentType = "Content-Type"
const MimeApplicationJSON = "application/json"

type H map[string]any

// SendJson converts a Go value to JSON and sends it to the client.
func SendJson(w http.ResponseWriter, v any, statusCode int) (ok bool) {
	w.Header().Set(HeaderContentType, MimeApplicationJSON)

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return true
	}

	jsonData, err := json.Marshal(v)
	if err != nil {
		// TODO SECURITY error could have data that could leak data usable for penetration
		// TODO error context should be collected and provided for Ops evaluation
		// TODO error should be securely logged for Ops evaluation
		log.Errorf("SendJson json.Marshal() error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	w.WriteHeader(statusCode)
	if _, err := w.Write(jsonData); err != nil {
		log.Errorf("web.JsonResponse %s", err)
		return false
	}

	return true
}

func SendJsonMsg(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set(HeaderContentType, MimeApplicationJSON)
	log.Errorf("Send Json Error %v", msg)
	w.WriteHeader(statusCode)
	w.Write([]byte(msg))
}

func JsonErrorResponse(w http.ResponseWriter, err error) (ok bool) {
	w.Header().Set(HeaderContentType, MimeApplicationJSON)

	if err != nil {
		log.Errorf("HTTP Hanlder error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return true
	}

	log.Errorf("HTTP Hanlder received nil value error.")
	w.WriteHeader(http.StatusInternalServerError)
	return false
}

func mapToJson() ([]byte, error) {
	data := H{
		"data": "authToken",
	}
	return json.Marshal(data)
}

func mapToJson10() ([]byte, error) {
	data := H{
		"data":   "authToken",
		"data2":  "authToken2",
		"data3":  "authToken3",
		"data4":  "authToken4",
		"data5":  "authToken5",
		"data6":  "authToken6",
		"data7":  "authToken7",
		"data8":  "authToken8",
		"data9":  "authToken9",
		"data10": "authToken10",
	}
	return json.Marshal(data)
}

func structToJson() ([]byte, error) {
	var data = struct {
		Data string `json:"data"`
	}{
		Data: "authToken",
	}
	return json.Marshal(data)
}

func structToJson10() ([]byte, error) {
	var data = struct {
		Data   string `json:"data"`
		Data2  string `json:"data2"`
		Data3  string `json:"data3"`
		Data4  string `json:"data4"`
		Data5  string `json:"data5"`
		Data6  string `json:"data6"`
		Data7  string `json:"data7"`
		Data8  string `json:"data8"`
		Data9  string `json:"data9"`
		Data10 string `json:"data10"`
	}{
		Data:   "authToken",
		Data2:  "authToken2",
		Data3:  "authToken3",
		Data4:  "authToken4",
		Data5:  "authToken5",
		Data6:  "authToken6",
		Data7:  "authToken7",
		Data8:  "authToken8",
		Data9:  "authToken9",
		Data10: "authToken10",
	}
	return json.Marshal(data)
}
