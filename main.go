package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"text/template"
)

// xml 包通过结构体标签，映射结构体中字段与xml元素和属性之间的关系
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
	// xml文件内容应该是从文件中读取出来的，不是本文重点，因此直接定义成字符串
	xml_data := `<?xml version="1.0" encoding="utf-8"?>
	<protocol>
		<tid_def>
			<tid name="TID_REQ_LOGIN" value="10000" desc = "交易远登陆请求"/>
			<tid name="TID_RSP_LOGIN" value="10001" desc = "交易远登陆应答"/>
			<tid name="TID_REQ_ORDER" value="10002" desc = "交易远登陆应答"/>
		</tid_def>
	</protocol>`

	Protocol := new(protocol)

	err1 := xml.Unmarshal([]byte(xml_data), Protocol)
	fmt.Println(err1)

	fmt.Println(*Protocol)
	// 定义模板: 教会程序如何写字
	t := template.New("temp")
	t, _ = t.Parse(`{{range .Tid_def.Tid}}
{{if ne .Name "TID_REQ_ORDER"}}{{.Name}}		{{.Value}}		 // {{.Dese}}{{end}}{{end}}`)

	// 输出： 根据作业清单，写作业
	t.Execute(os.Stdout, Protocol)
}
