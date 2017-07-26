package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/miekg/dns"
	"github.com/oschwald/geoip2-golang"
)

type data struct {
	From      string `json:"from"`
	Fw        int    `json:"fw"`
	GroupID   int    `json:"group_id"`
	Lts       int    `json:"lts"`
	MsmID     int    `json:"msm_id"`
	MsmName   string `json:"msm_name"`
	PrbID     int    `json:"prb_id"`
	Resultset []struct {
		Af      int    `json:"af"`
		DstAddr string `json:"dst_addr"`
		Lts     int    `json:"lts"`
		Proto   string `json:"proto"`
		Result  struct {
			ANCOUNT int     `json:"ANCOUNT"`
			ARCOUNT int     `json:"ARCOUNT"`
			ID      int     `json:"ID"`
			NSCOUNT int     `json:"NSCOUNT"`
			QDCOUNT int     `json:"QDCOUNT"`
			Abuf    string  `json:"abuf"`
			Rt      float64 `json:"rt"`
			Size    int     `json:"size"`
		} `json:"result"`
		SrcAddr string `json:"src_addr"`
		Subid   int    `json:"subid"`
		Submax  int    `json:"submax"`
		Time    int    `json:"time"`
	} `json:"resultset"`
	Timestamp int    `json:"timestamp"`
	Type      string `json:"type"`
}

func decode(input string) string {
	data, _ := base64.StdEncoding.DecodeString(input)
	m := &dns.Msg{}
	m.Unpack(data)
	return string(m.String())
}

func queryASN(input string, ASdb *geoip2.Reader) string {
	ip := net.ParseIP(input)
	record, err := ASdb.ASN(ip)
	if err != nil {
		log.Fatal(err)
	}
	return "AS" + strconv.Itoa(int(record.AutonomousSystemNumber)) + " " + record.AutonomousSystemOrganization
}
func queryCountry(input string, Countrydb *geoip2.Reader) string {
	ip := net.ParseIP(input)
	record, err := Countrydb.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	return record.Country.Names["en"]
}

func main() {
	ASdb, err := geoip2.Open("./GeoLite2-ASN.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer ASdb.Close()
	Countrydb, err := geoip2.Open("./GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer Countrydb.Close()
	resp, err := http.Get(os.Args[1])
	if nil != err {
		panic(err)
	}
	var datas []data
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &datas)

	// Loop over structs and display the respective views.
	for p := range datas {
		color.Set(color.FgGreen)
		fmt.Printf("----------------------\n")
		fmt.Printf("IP: %v Country: %v ISP: %v \n", datas[p].From, queryCountry(datas[p].From, Countrydb), queryASN(datas[p].From, ASdb))
		color.Unset()
		fmt.Printf("%v", decode(datas[p].Resultset[0].Result.Abuf))
		fmt.Println(p)
	}

}
