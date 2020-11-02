package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/leandrojmp/go-sysmon/config"
)

// LandingPage index landing
func LandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go-sysmon")
	config.InfoLogger.Print("acess landingPage")
}

// KillSwitch endpoint
func KillSwitch(w http.ResponseWriter, r *http.Request) {
	config.InfoLogger.Print("shutdown received...")
	time.Sleep(20 * time.Second)
	config.InfoLogger.Print("application stopped")
	os.Exit(0)

}
