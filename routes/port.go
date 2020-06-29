package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type singlePort struct {
	SrcIP   string `json:"srcip"`
	SrcPort int    `json:"srcport"`
	DstIP   string `json:"dstip"`
	DstPort int    `json:"dstport"`
	Status  string `json:"status"`
}

func testResponse(port int) singlePort {
	var testResponse singlePort
	testResponse.DstIP = "192.168.0.1"
	testResponse.DstPort = port
	testResponse.Status = "LISTENING"
	testResponse.SrcIP = "10.0.0.10"
	testResponse.SrcPort = 5015
	return testResponse
}

// ReturnSinglePort retorna apenas uma porta
func ReturnSinglePort(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["port"])
	json.NewEncoder(w).Encode(testResponse(key))
}
