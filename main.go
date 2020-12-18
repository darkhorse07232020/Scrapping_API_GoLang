package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"
)

/*
===========
=ADS COUNT=
===========
>>>>>REQ
GET https://classifieds.immowebapi.be/search/3/count?countries=BE&isSoldOrRented=false&priceType=MONTHLY_RENTAL_PRICE&propertyTypes=APARTMENT%2CHOUSE&transactionTypes=FOR_RENT HTTP/1.1
Host: classifieds.immowebapi.be
x-api-key: q0H4TxoSZ9arZTAtQMa3AaKYviGiEluEuKjzUAj6
Accept: application/json
User-Agent: Immoweb-iOS/5.19.2
Accept-Language: fr
Accept-Encoding: gzip;q=1.0, compress;q=0.5
Connection: keep-alive

>>>>>>RESP
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 5
Connection: keep-alive
Date: Mon, 22 Jun 2020 16:17:34 GMT
x-amzn-RequestId: da98935b-be8d-41c0-a0d6-d93623e84a89
Content-Language: fr
x-amz-apigw-id: OibwvEWsDoEFpag=
x-count-with-geoPoint: 17861
X-Amzn-Trace-Id: Root=1-5ef0d99e-698ffcebaa613c753dbe11c6;Sampled=0
Via: 1.1 2ec3090d74e200e4acdb2780da3c3c44.cloudfront.net (CloudFront), 1.1 e28c193c96684df9ba36cf3fd8976708.cloudfront.net (CloudFront)
X-Amz-Cf-Pop: FRA2-C1
X-Cache: Hit from cloudfront
X-Amz-Cf-Pop: AMS54-C1
X-Amz-Cf-Id: Zvwn8dSC8uYxcJzNsLONjovFCuTitG_-vAvHhEjc6oOdE_nkLYQH3g==
Age: 99

24540
*/
var wg = &sync.WaitGroup{}
var client = &http.Client{Timeout: 500 * time.Second}
var adStartRange = 0
var adSlice = 10
var adEndRange = 9
var currentPage = 0
var maxPage = 0
var adsCount = 0
var adsIds []int
var requests int = 0
var jsonData AdProperty

func main() {
	start := time.Now()
	fmt.Println("Starting...")

	countAds()

	if adsCount > 0 {
		http.HandleFunc("/getMaxPage", getMaxPage)
		http.HandleFunc("/currentPageData", currentPageData)

		err := http.ListenAndServe(":4000", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
	wg.Wait()
	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())
	fmt.Println("Requests : " + strconv.Itoa(requests))
	time.Sleep(time.Second)
}

func fetchAdProperty(propertyID int) string {
	defer wg.Done()
	url := "https://classifieds.immowebapi.be/search/3/" + strconv.Itoa(propertyID)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Host", "classifieds.immowebapi.be")
	req.Header.Set("x-api-key", "q0H4TxoSZ9arZTAtQMa3AaKYviGiEluEuKjzUAj6")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Immoweb-iOS/5.19.2")
	req.Header.Set("Accept-Language", "fr")
	req.Header.Set("Accept-Encoding", "gzip;q=1.0, compress;q=0.5")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	jsonDataFromHTTP, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var jsonData AdProperty

	err = json.Unmarshal([]byte(jsonDataFromHTTP), &jsonData) // here!

	if err != nil {
		println(err)
		panic(err)
	}
	requests++

	return string([]byte(jsonDataFromHTTP))
}

func countAds() {
	url := "https://classifieds.immowebapi.be/search/3/count?countries=BE&isSoldOrRented=false&priceType=MONTHLY_RENTAL_PRICE&propertyTypes=APARTMENT%2CHOUSE&transactionTypes=FOR_RENT"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Host", "classifieds.immowebapi.be")
	req.Header.Set("x-api-key", "q0H4TxoSZ9arZTAtQMa3AaKYviGiEluEuKjzUAj6")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Immoweb-iOS/5.19.2")
	req.Header.Set("Accept-Language", "fr")
	req.Header.Set("Accept-Encoding", "gzip;q=1.0, compress;q=0.5")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	rawResult, _ := ioutil.ReadAll(resp.Body)

	adsCount, _ = strconv.Atoi(string(rawResult))
	maxPage = int(math.Round(float64(adsCount / adSlice)))
}

func fetchAdsListInRange(j *int) {
	var currentPageNumber = *j
	url := "https://classifieds.immowebapi.be/search/3/query?countries=BE&isSoldOrRented=false&minBedroomCount=2&priceType=MONTHLY_RENTAL_PRICE&propertyTypes=APARTMENT%2CHOUSE&range=" + strconv.Itoa((adStartRange+currentPageNumber)*(adEndRange-adStartRange+1)) + "-" + strconv.Itoa(adEndRange+currentPageNumber*(adEndRange-adStartRange+1)) + "&transactionTypes=FOR_RENT"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Host", "classifieds.immowebapi.be")
	req.Header.Set("x-api-key", "q0H4TxoSZ9arZTAtQMa3AaKYviGiEluEuKjzUAj6")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Immoweb-iOS/5.19.2")
	req.Header.Set("Accept-Language", "fr")
	req.Header.Set("Accept-Encoding", "gzip;q=1.0, compress;q=0.5")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	jsonDataFromHTTP, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var jsonData AdsFromSearchQuery

	err = json.Unmarshal([]byte(jsonDataFromHTTP), &jsonData) // here!

	if err != nil {
		println(err)
		panic(err)
	}
	for i := 0; i < len(jsonData); i++ {
		//fmt.Println(jsonData[i].Property.Title)
		if jsonData[i].Transaction.SoldOrRented.IsSoldOrRented == true {
			//fmt.Println("$***>This property is already sold/rented")
		} else {
			//fmt.Println("$***>Available")
			adsIds = append(adsIds, jsonData[i].ID)
		}
	}
}

func getMaxPage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, strconv.Itoa(maxPage))
}
func currentPageData(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		println("ParseForm() err: %v", err)
		return
	}

	var pageNum int
	var res string = "["
	pageNum, _ = strconv.Atoi(req.FormValue("pageNumber"))
	fetchAdsListInRange(&pageNum)

	for i := 0; i < len(adsIds); i++ {
		wg.Add(1)
		fmt.Println("#ID found :", adsIds[i])
		if i > 0 {
			res = res + ","
		}
		res = res + fetchAdProperty(adsIds[i])
	}
	res = res + "]"
	fmt.Fprintf(w, string(res))
}
