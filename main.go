package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"text/template"
)

type protocol struct {
	Tid_def tiddef `xml:"tid_def"`
}

type tiddef struct {
	Tid []tid `xml:"tid"`
}

type tid struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
	Dese  string `xml:"desc,attr"`
}

func main() {
	xml_data := `<?xml version="1.0" encoding="utf-8"?>
	<protocol>
		<tid_def>
			<tid name="TID_REQ_LOGIN" value="10000" desc = "交易远登陆请求"/>
			<tid name="TID_RSP_LOGIN" value="10001" desc = "交易远登陆应答"/>
		</tid_def>
	</protocol>`

	Protocol := new(protocol)

	err1 := xml.Unmarshal([]byte(xml_data), Protocol)
	fmt.Println(err1)

	fmt.Println(*Protocol)

	t := template.New("temp")
	t, _ = t.Parse(`{{range .Tid_def.Tid}}
{{.Name}}
{{end}}`)
	t.Execute(os.Stdout, Protocol)
}
