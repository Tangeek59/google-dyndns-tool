package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const GetIpUrl = "https://api.ipify.org"

func main() {

	user := os.Getenv("DNS_USERNAME")
	password := os.Getenv("DNS_PASSWORD")
	hostname := os.Getenv("HOSTNAME")

	res, err := http.Get(GetIpUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ip, _ := ioutil.ReadAll(res.Body)

	client := &http.Client{}
	URL := buildGoogleUpdateUrl(hostname, string(ip))

	fmt.Println(URL)

	//pass the values to the request's body
	req, _ := http.NewRequest("POST", URL, nil)
	req.SetBasicAuth(user, password)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp)
}

func buildGoogleUpdateUrl(hostname, ip string) string {
	return fmt.Sprintf("https://domains.google.com/nic/update?hostname=%s&%s", hostname, ip)
}
