package routes

import (
	"fmt"
	"net/http"

	"github.com/leandrojmp/go-sysmon/config"
)

// LandingPage index landing
func LandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go-sysmon")
	config.InfoLogger.Print("acess landingPage")
}
