package main

import (
	"encoding/xml"
	"fmt"
)

type Protocol struct {
	tid_def Tiddef `xml:"tid_def"`
}

type Tiddef struct {
	tid []Tid `xml:"tid"`
}

type Tid struct {
	name  string `xml:"name,attr"`
	value string `xml:"value,attr"`
	dese  string `xml:"desc,attr"`
}

func main() {
	xml_data := `<?xml version="1.0" encoding="utf-8"?>
	<protocol>
		<tid_def>
			<tid name="TID_REQ_LOGIN" value="10000" desc = "交易远登陆请求">
			<tid name="TID_RSP_LOGIN" value="10001" desc = "交易远登陆请求">
		</tid_def>
	</protocol>`

	protocol := new(Protocol)

	xml.Unmarshal([]byte(xml_data), protocol)
	fmt.Println(*protocol)

}
