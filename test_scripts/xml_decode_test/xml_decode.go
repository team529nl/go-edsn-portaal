package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Foo struct {
	XMLName struct{} `xml:"nl.test Foo"`
	Data string `xml:",chardata"`
}

func (f Foo) String() string {
	return fmt.Sprint(f.Data)
}

type soapEnvelope struct {
	XMLName struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  []byte   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    struct {
		Message interface{} `xml:",any,omitempty"`
	} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type soapResponseEnvelope struct {
	XMLName struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  []byte   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    struct {
		Data []byte `xml:",innerxml"`
	} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

func main() {
	rawXML := []byte(`
<?xml version="1.0" encoding="UTF-8"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"><SOAP-ENV:Body>
<ns0:Foo xmlns:ns0="nl.test">Bar</ns0:Foo>
</SOAP-ENV:Body></SOAP-ENV:Envelope>
`)

//	rawXML := []byte(`
//<ns0:Foo ns0="nl.test">Bar</ns0:Foo>
//`)

	var envelope soapEnvelope

	//foo0 := Foo{
	//	Data:    "Test",
	//}

	//envelope.Body.Message = foo0
	//var buf bytes.Buffer
	//enc := xml.NewEncoder(&buf)
	//if err := enc.Encode(envelope); err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println(buf.String())

	//var responseEnvelope soapResponseEnvelope
	var foo1 Foo
	envelope.Body.Message = &foo1
	err := xml.Unmarshal(rawXML, &envelope)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("foo: %#v\n", x)

	//var foo1 Foo
	//err = xml.Unmarshal(responseEnvelope.Body.Data, &foo1)

	if err != nil {
		log.Fatal(err)
	}


	log.Println(envelope.Body)
}
