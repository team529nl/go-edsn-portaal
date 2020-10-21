package main

import (
	"aqwari.net/xml/wsdlgen"
	"aqwari.net/xml/xsdgen"
	"log"
	"os"
)

func main() {

	log.SetFlags(0)
	var cfg wsdlgen.Config
	cfg.Option(wsdlgen.DefaultOptions...)
	cfg.XSDOption(xsdgen.DefaultOptions...)
	cfg.Option(wsdlgen.LogOutput(log.New(os.Stderr, "", 0)))

	if err := cfg.GenCLI("/path/to/SearchMeteringPoints_2p1/SearchMeteringPoints_2.wsdl", "/path/to/SearchMeteringPoints_2p1/SearchMeteringPointsRequest_2p0.xsd"); err != nil {
		log.Fatal(err)
	}


}
