package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// SinglePort estrutura da resposta
type SinglePort struct {
	SrcIP   string `json:"srcip"`
	SrcPort int64  `json:"srcport"`
	DstIP   string `json:"dstip"`
	DstPort int64  `json:"dstport"`
	Status  string `json:"status"`
}

func testResponse(port int64) SinglePort {
	var testResponse SinglePort
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
	json.NewEncoder(w).Encode(testResponse(int64(key)))
}

// ReturnAllPorts retorna todas as portas
func ReturnAllPorts(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("/proc/net/tcp")
	if err != nil {
		fmt.Println("error:", err)
	}
	lines := strings.Split(string(content), "\n")

	lines = lines[1 : len(lines)-1]

	var jsonResponse []SinglePort

	for _, line := range lines {
		var localIP string = strings.Split(strings.Split(strings.TrimSpace(line), " ")[1], ":")[0]
		localIP = convertIP(localIP)
		localPort, _ := strconv.ParseInt(strings.Split(strings.Split(strings.TrimSpace(line), " ")[1], ":")[1], 16, 64)
		var remoteIP string = strings.Split(strings.Split(strings.TrimSpace(line), " ")[2], ":")[0]
		remoteIP = convertIP(remoteIP)
		remotePort, _ := strconv.ParseInt(strings.Split(strings.Split(strings.TrimSpace(line), " ")[2], ":")[1], 16, 64)
		var testResponse SinglePort
		testResponse.DstIP = remoteIP
		testResponse.DstPort = remotePort
		testResponse.Status = "LISTENING"
		testResponse.SrcIP = localIP
		testResponse.SrcPort = localPort
		jsonResponse = append(jsonResponse, testResponse)
	}
	fmt.Println("/netstat: ReturnAllPorts")
	json.NewEncoder(w).Encode(jsonResponse)
}

func convertIP(addr string) string {
	octetA, _ := strconv.ParseInt(addr[6:8], 16, 64)
	octetB, _ := strconv.ParseInt(addr[4:6], 16, 64)
	octetC, _ := strconv.ParseInt(addr[2:4], 16, 64)
	octetD, _ := strconv.ParseInt(addr[0:2], 16, 64)

	ipOctets := []string{string(strconv.Itoa(int(octetA))),
		string(strconv.Itoa(int(octetB))),
		string(strconv.Itoa(int(octetC))),
		string(strconv.Itoa(int(octetD)))}

	ipAddr := strings.Join(ipOctets, ".")

	return ipAddr
}
