package routes

import (
	"fmt"
	"net/http"
)

// LandingPage index landing
func LandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go-sysmon")
	fmt.Println("access: landingPage")
}
