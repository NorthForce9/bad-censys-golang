package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type TooLTT struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result Result `json:"result"`
}
type Services struct {
	Port              int    `json:"port"`
	ServiceName       string `json:"service_name"`
	TransportProtocol string `json:"transport_protocol"`
	Certificate       string `json:"certificate,omitempty"`
}
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Location struct {
	Continent             string      `json:"continent"`
	Country               string      `json:"country"`
	CountryCode           string      `json:"country_code"`
	Timezone              string      `json:"timezone"`
	Coordinates           Coordinates `json:"coordinates"`
	RegisteredCountry     string      `json:"registered_country"`
	RegisteredCountryCode string      `json:"registered_country_code"`
}
type AutonomousSystem struct {
	Asn         int    `json:"asn"`
	Description string `json:"description"`
	BgpPrefix   string `json:"bgp_prefix"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
}
type Hits struct {
	IP string `json:"ip"`
}
type Links struct {
	Next string `json:"next"`
}
type Result struct {
	Query string `json:"query"`
	Total int    `json:"total"`
	Hits  []Hits `json:"hits"`
	Links Links  `json:"links"`
}

// ced0db7a22@emailnax.com
func main() {
	query := os.Args[1]
	// tood next parm are only once try to understdand that ok
	page := os.Args[2] // min 0 max 100
	next := os.Args[3]
	x, err := strconv.ParseInt(next, 32, 64)
	if err != nil {
		log.Panic(err)
	}
	var client http.Client

	url := fmt.Sprintf("https://search.censys.io/api/v2/hosts/search?q=%s&per_page=%s&virtual_hosts=EXCLUDE", query, page) // add page
	fmt.Printf("https://search.censys.io/api/v2/hosts/search?q=%s&per_page=%s&virtual_hosts=EXCLUDE", query, page)         // add page

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic("cannot make reqeust")
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Basic NTI5ZDZmODAtMjI3NS00NDFmLWE3YzMtZTJkZDc4MWJiYWIxOnhtbHdjVGdKOXFGT1ZKYXhJcGhNZHZUSmtkY2tya3Za")
	resp, err := client.Do(req)
	if err != nil {
		log.Panic("cannot make request")
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	// if u want to read the body many time
	// u need to restore
	// reader := io.NopCloser(bytes.NewReader(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}
	var result TooLTT

	if err := json.Unmarshal(bodyBytes, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	if x != 0 {
		for i := 0; int64(i) < x; i++ {
			var client http.Client
			var result2 TooLTT

			fmt.Print()

			url := fmt.Sprintf("https://search.censys.io/api/v2/hosts/search?q=%s&per_page=%s&virtual_hosts=EXCLUDE&cursor=%s", query, page, result.Result.Links.Next) // add page
			fmt.Printf("https://search.censys.io/api/v2/hosts/search?q=%s&per_page=%s&virtual_hosts=EXCLUDE&cursor=%s", query, page, result.Result.Links.Next)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Panic("cannot make reqeust")
			}
			req.Header.Add("accept", "application/json")
			req.Header.Add("Authorization", "Basic NTI5ZDZmODAtMjI3NS00NDFmLWE3YzMtZTJkZDc4MWJiYWIxOnhtbHdjVGdKOXFGT1ZKYXhJcGhNZHZUSmtkY2tya3Za")
			resp, err := client.Do(req)
			if err != nil {
				log.Panic("cannot make request")
			}
			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			// if u want to read the body many time
			// u need to restore
			// reader := io.NopCloser(bytes.NewReader(bodyBytes))
			if err != nil {
				log.Fatal(err)
			}

			if err := json.Unmarshal(bodyBytes, &result2); err != nil { // Parse []byte to the go struct pointer
				fmt.Println("Can not unmarshal JSON")
			}
			for _, rec := range result2.Result.Hits {
				fmt.Println(rec.IP)
			}
		}

	}

}

// cool i make( it, 0)
//try to fix and update the valua of next
