package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func handleAMD(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n ACTION %s [%s] \n", r.URL.Path, r.Method)

	body, _ := io.ReadAll(r.Body)
	if len(body) > 0 {
		fmt.Printf("BODY: %s\n", string(body))
	}

	var response string
	contentType := "application/json"

	switch r.URL.Path {
	// register
	case "/register.php":
		contentType = "text/plain"
		response = "Success"
		fmt.Println("Ok")

	case "/login.php":
		// login
		response = `{"status":"success","user_id":"1337","username":"anonymous","country":"Singapore"}`
		fmt.Println("Ok")

	// Motherboard
	case "/RedeemVoucher.php":
		contentType = "text/plain"
		if r.Method == "POST" {
			response = "Entry Inserted"
			fmt.Println("Ok")
		} else {
			response = "Ready"
		}

	// Balance
	case "/checkvoucher.php":
		contentType = "text/plain"
		response = "1,1,1,10,50,100"
		fmt.Println("Ok")
	// Checker
	case "/CheckExpireVoucher.php", "/checkevent.php":
		contentType = "text/plain"
		response = "Valid"
		fmt.Println("Ok")

	default:
		response = `{"status":"success"}`
		fmt.Println("Ok")
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strings.TrimSpace(response)))
	fmt.Printf("SENT [%s]: %s\n", contentType, response)
}

func main() {
	http.HandleFunc("/", handleAMD)
	fmt.Println("AMD Ryzen Hunt Event Server. Emulator API, Backend for AMD Ryzen Hunt Event.")
	fmt.Println("Made with love by zenlond")
	log.Fatal(http.ListenAndServe(":80", nil))
}
