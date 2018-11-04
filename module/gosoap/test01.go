package main

import (
"github.com/tiaguinho/gosoap"
)

type GetGeoIPResponse struct {
	GetGeoIPResult GetGeoIPResult
}

type GetGeoIPResult struct {
	ReturnCode        string
	IP                string
	ReturnCodeDetails string
	CountryName       string
	CountryCode       string
}

var (
	r GetGeoIPResponse
)

func main() {
	soap, err := gosoap.SoapClient("http://www.webservicex.net/geoipservice.asmx?WSDL")
	if err != nil {
		panic(err)
	}

	params := gosoap.Params{
		"IPAddress": "8.8.8.8",
	}

	err = soap.Call("GetGeoIP", params)
	if err != nil {
		panic(err)
	}

	soap.Unmarshal(&r)
	if r.GetGeoIPResult.CountryCode != "USA" {
		panic(err)
	}
}