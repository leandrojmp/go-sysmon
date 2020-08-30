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
		testResponse.SrcIP = localIP
		testResponse.SrcPort = localPort
		testResponse.Status = connectionCode(strings.Split(strings.TrimSpace(line), " ")[3])
		jsonResponse = append(jsonResponse, testResponse)
	}
	fmt.Println("/netstat: ReturnAllPorts")
	json.NewEncoder(w).Encode(jsonResponse)
}

// ReturnSinglePort returns only one port
func ReturnSinglePort(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	port, _ := strconv.Atoi(vars["port"])
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
		if localPort != int64(port) {
			continue
		}
		var remoteIP string = strings.Split(strings.Split(strings.TrimSpace(line), " ")[2], ":")[0]
		remoteIP = convertIP(remoteIP)
		remotePort, _ := strconv.ParseInt(strings.Split(strings.Split(strings.TrimSpace(line), " ")[2], ":")[1], 16, 64)
		var testResponse SinglePort
		testResponse.DstIP = remoteIP
		testResponse.DstPort = remotePort
		testResponse.SrcIP = localIP
		testResponse.SrcPort = localPort
		testResponse.Status = connectionCode(strings.Split(strings.TrimSpace(line), " ")[3])
		jsonResponse = append(jsonResponse, testResponse)
	}
	fmt.Println("/netstat: ReturnSinglePort")
	json.NewEncoder(w).Encode(jsonResponse)
}

func connectionCode(cxcode string) string {
	var connectionState string
	switch cxcode {
	case "01":
		connectionState = "ESTABLISHED"
	case "02":
		connectionState = "SYN_SENT"
	case "03":
		connectionState = "SYN_RECV"
	case "04":
		connectionState = "FIN_WAIT1"
	case "05":
		connectionState = "FIN_WAIT2"
	case "06":
		connectionState = "TIME_WAIT"
	case "07":
		connectionState = "CLOSE"
	case "08":
		connectionState = "CLOSE_WAIT"
	case "09":
		connectionState = "LAST_ACK"
	case "0A":
		connectionState = "LISTENING"
	case "0B":
		connectionState = "CLOSING"
	default:
		connectionState = "UNKOWN_STATE"
	}
	return connectionState
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
