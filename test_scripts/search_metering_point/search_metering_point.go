package main

import (
	"bytes"
	_ "bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	_ "fmt"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	_ "github.com/mitchellh/mapstructure"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"team529.nl/go-edsn-portaal/portaal"
	"team529.nl/go-edsn-portaal/types/soap/car"
	"time"
	_ "time"
)

func main() {

	cert, err := tls.LoadX509KeyPair("/path/to/party/public.pem",
		"/path/to/party/private.key")

	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := tls.Config{
		MinVersion: tls.VersionTLS12,
		InsecureSkipVerify: true,
		Certificates: []tls.Certificate{cert}}

	transport := http.Transport{
		TLSClientConfig: &tlsConfig,
	}

	client := http.Client{
		Transport: &transport,
	}

	input := map[string]interface{}{
		"ContentHash": "2",
		"CorrelationID": "A",
		"ConversationID": "-",
	}

	var header = portaal.SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader{}
	_ = mapstructure.Decode(input, &header)

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	enc.Encode(portaal.SearchMeteringPointsRequestEnvelope{EDSNBusinessDocumentHeader: &header})

	//log.Print(buf.String())

	textXml := strings.NewReader(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"  xmlns:ws="urn:nedu:edsn:service:searchmeteringpointsmp:1:standard" xmlns:ccma1="urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard" xmlns:ccma2="urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard" xmlns:ccma3="urn:edsn:edsn:data:soapfault:1:standard"><soap:Header></soap:Header><soap:Body><ccma1:SearchMeteringPointsRequestEnvelope xmlns:ccma1="urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard" xmlns="urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard"><ccma1:EDSNBusinessDocumentHeader><ccma:CreationTimestamp xmlns:ccma="urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard">2020-01-24T10:09:15Z</ccma:CreationTimestamp><ccma:MessageID xmlns:ccma="urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard">93ca7340-3e91-11ea-a2d6-9ba6d4a65e85</ccma:MessageID><ccma:Destination xmlns:ccma="urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard"><ccma:Receiver><ccma:Authority>EAN.UCC</ccma:Authority><ccma:ContactTypeIdentifier>EDSN</ccma:ContactTypeIdentifier><ccma:ReceiverID>8712423010208</ccma:ReceiverID></ccma:Receiver></ccma:Destination><ccma:Source xmlns:ccma="urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard"><ccma:Authority>EAN.UCC</ccma:Authority><ccma:ContactTypeIdentifier>DDM_M</ccma:ContactTypeIdentifier><ccma:SenderID>1119326115509</ccma:SenderID></ccma:Source></ccma1:EDSNBusinessDocumentHeader><ccma:Portaal_Content xmlns:ccma="urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard"><ccma:Portaal_MeteringPoint><ccma:EDSN_AddressSearch><ccma:BuildingNr>60</ccma:BuildingNr><ccma:ZIPCode>7322JG</ccma:ZIPCode></ccma:EDSN_AddressSearch></ccma:Portaal_MeteringPoint></ccma:Portaal_Content></ccma1:SearchMeteringPointsRequestEnvelope></soap:Body></soap:Envelope>`)
	//log.Print(textXml)


	req, err := http.NewRequest("POST", "https://portaal-opt.edsn.nl/b2b/synchroon/ResponderSearchMeteringPointsMPRespondingActivity", textXml)
	req.Header.Add("SOAPAction", "urn:SearchMeteringPoints")

	//reqDump, _ := httputil.DumpRequest(req, true);
	//
	//log.Println(string(reqDump))

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, _ = httputil.DumpResponse(resp, true)

	var buf2 bytes.Buffer
	enc2 := xml.NewEncoder(&buf2)
	//dec := xml.NewDecoder(os.Stdin)
	var message car.SearchMeteringPointsRequestEnvelope
	input2 := map[string]interface{}{
		"EDSNBusinessDocumentHeader": map[string]interface{}{
			"CreationTimestamp": car.SearchMeteringPointsRequestEnvelopeDateTime(time.Now()),
			"MessageId": uuid.New().String(),
			"Destination": map[string]interface{}{
				"Receiver": map[string]interface{} {
					"Authority": "EAN.UCC",
					"ContactTypeIdentifier": "EDSN",
					"ReceiverID": "8712423010208",
				},
			},
			"Source": map[string]interface{} {
						"Authority": "EAN.UCC",
						"ContactTypeIdentifier": "DDM_M",
						"SenderId": "[ean13]",
			},
		},
		"PortaalContent": map[string]interface{} {
							"PortaalMeteringPoint": map[string]interface{} {
								"EDSNAddressSearch": map[string]interface{} {
									"BuildingNr": 11,
									"ZIPCode": "1111AA",
								},
							},
		},
	}

	err2 := mapstructure.Decode(input2, &message)
	if err2 != nil {
		panic(err)
	}

	enc2.Encode(message)


	soapClient := car.Client{}
	soapClient.HTTPClient = &client

	soapClient.RequestHook = func(req *http.Request) *http.Request {
		data, err := httputil.DumpRequest(req, true)
		_ = data
		if err != nil {
			panic(err)
		}
		//log.Println("REQUEST")
		//log.Println(string(data))
		return req
	}

	soapClient.ResponseHook = func(rsp *http.Response) *http.Response {
		data, err := httputil.DumpResponse(rsp, true)
		_ = data
		if err != nil {
			panic(err)
		}
		//log.Println("RESPONSE")
		log.Println(string(data))
		return rsp
	}

	response, err := soapClient.SearchMeteringPoints(context.TODO(), message)

	if err != nil {
		log.Fatal(err)
	}

	PrettyPrint(response)
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}