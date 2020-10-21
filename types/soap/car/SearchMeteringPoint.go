// Package car
//
package car

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SOAPFault struct {
	ErrorCode    float64 `xml:"urn:edsn:edsn:data:soapfault:1:standard ErrorCode"`
	ErrorText    string  `xml:"urn:edsn:edsn:data:soapfault:1:standard ErrorText"`
	ErrorDetails string  `xml:"urn:edsn:edsn:data:soapfault:1:standard ErrorDetails,omitempty"`
}

type SearchMeteringPointsRequestEnvelope struct {
	XMLName xml.Name `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard SearchMeteringPointsRequestEnvelope"`
	EDSNBusinessDocumentHeader SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeader `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard EDSNBusinessDocumentHeader"`
	PortaalContent             SearchMeteringPointsRequestEnvelopePC                         `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Portaal_Content"`
}

// Must be at least 1 items long
type SearchMeteringPointsRequestEnvelopeAddressText string

// Must match the pattern [0-9]*
type SearchMeteringPointsRequestEnvelopeBAGIDType string

type SearchMeteringPointsRequestEnvelopeBAGType struct {
	BAGID         SearchMeteringPointsRequestEnvelopeBAGIDType `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard BAGID,omitempty"`
	BAGBuildingID SearchMeteringPointsRequestEnvelopeBAGIDType `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard BAGBuildingID,omitempty"`
}

// Must be at least 2 items long
type SearchMeteringPointsRequestEnvelopeCountryCode string

type SearchMeteringPointsRequestEnvelopeDateTime time.Time

func (t *SearchMeteringPointsRequestEnvelopeDateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t SearchMeteringPointsRequestEnvelopeDateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeader struct {
	ContentHash       string                                                                   `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ContentHash,omitempty"`
	ConversationID    string                                                                   `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ConversationID,omitempty"`
	CorrelationID     string                                                                   `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard CorrelationID,omitempty"`
	CreationTimestamp SearchMeteringPointsRequestEnvelopeDateTime                              `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard CreationTimestamp"`
	DocumentID        string                                                                   `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard DocumentID,omitempty"`
	ExpiresAt         *SearchMeteringPointsRequestEnvelopeDateTime                              `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ExpiresAt,omitempty"`
	MessageID         string                                                                   `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard MessageID"`
	ProcessTypeID     string                                                                   `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ProcessTypeID,omitempty"`
	RepeatedRequest   string                                                                   `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard RepeatedRequest,omitempty"`
	TestRequest       string                                                                   `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard TestRequest,omitempty"`
	Destination       SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderDestination `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Destination"`
	Manifest          *SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderManifest    `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Manifest,omitempty"`
	Source            SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderSource      `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Source"`
}

type SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderDestination struct {
	Receiver SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Receiver"`
	Service  SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderDestinationService  `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Service,omitempty"`
}

type SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver struct {
	Authority             string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ContactTypeIdentifier,omitempty"`
	ReceiverID            string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ReceiverID"`
}

type SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderDestinationService struct {
	ServiceMethod string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ServiceMethod,omitempty"`
	ServiceName   string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ServiceName,omitempty"`
}

type SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderManifest struct {
	NumberofItems float64                                                                             `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard NumberofItems"`
	ManifestItem  []SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ManifestItem"`
}

type SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem struct {
	Description               string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Description,omitempty"`
	LanguageCode              string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard LanguageCode,omitempty"`
	MimeTypeQualifierCode     string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard MimeTypeQualifierCode"`
	UniformResourceIdentifier string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard UniformResourceIdentifier"`
}

type SearchMeteringPointsRequestEnvelopeEDSNBusinessDocumentHeaderSource struct {
	Authority             string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ContactTypeIdentifier,omitempty"`
	SenderID              string `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard SenderID"`
}

// May be one of ELK, GAS
type SearchMeteringPointsRequestEnvelopeEnergyProductPortaalTypeCode string

// Must be at least 1 items long
type SearchMeteringPointsRequestEnvelopeExBuildingNr string

// Must match the pattern [0-9]{18}
type SearchMeteringPointsRequestEnvelopeGSRNEANCode string

// Must be at least 1 items long
type SearchMeteringPointsRequestEnvelopeLocationDescriptionType string

type SearchMeteringPointsRequestEnvelopeMPAddressRequestType struct {
	BAG          *SearchMeteringPointsRequestEnvelopeBAGType      `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard BAG,omitempty"`
	StreetName   SearchMeteringPointsRequestEnvelopeAddressText  `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard StreetName,omitempty"`
	BuildingNr   int                                             `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard BuildingNr,omitempty"`
	ExBuildingNr SearchMeteringPointsRequestEnvelopeExBuildingNr `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ExBuildingNr,omitempty"`
	ZIPCode      SearchMeteringPointsRequestEnvelopeZIPCode      `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ZIPCode,omitempty"`
	CityName     SearchMeteringPointsRequestEnvelopeAddressText  `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard CityName,omitempty"`
	Country      SearchMeteringPointsRequestEnvelopeCountryCode  `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Country,omitempty"`
}

// May be one of ART, GVB, KVB
type SearchMeteringPointsRequestEnvelopeMarketSegmentCode string

type SearchMeteringPointsRequestEnvelopePC struct {
	PortaalMeteringPoint SearchMeteringPointsRequestEnvelopePCPMP `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard Portaal_MeteringPoint"`
}

type SearchMeteringPointsRequestEnvelopePCPMP struct {
	EANID               SearchMeteringPointsRequestEnvelopeGSRNEANCode                  `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard EANID,omitempty"`
	EDSNAddressSearch   SearchMeteringPointsRequestEnvelopeMPAddressRequestType         `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard EDSN_AddressSearch,omitempty"`
	LocationDescription SearchMeteringPointsRequestEnvelopeLocationDescriptionType      `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard LocationDescription,omitempty"`
	MarketSegment       SearchMeteringPointsRequestEnvelopeMarketSegmentCode            `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard MarketSegment,omitempty"`
	ProductType         SearchMeteringPointsRequestEnvelopeEnergyProductPortaalTypeCode `xml:"urn:nedu:edsn:data:searchmeteringpointsrequest:1:standard ProductType,omitempty"`
}

// Must be at least 1 items long
type SearchMeteringPointsRequestEnvelopeZIPCode string

type SearchMeteringPointsResponseEnvelope struct {
	XMLName struct{} `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard SearchMeteringPointsResponseEnvelope"`
	EDSNBusinessDocumentHeader SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeader `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard EDSNBusinessDocumentHeader"`
	PortaalContent             SearchMeteringPointsResponseEnvelopePC                         `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Portaal_Content"`
}

// Must be at least 1 items long
type SearchMeteringPointsResponseEnvelopeAddressText string

// Must match the pattern [0-9]*
type SearchMeteringPointsResponseEnvelopeBAGIDType string

type SearchMeteringPointsResponseEnvelopeBAGType struct {
	BAGID         SearchMeteringPointsResponseEnvelopeBAGIDType `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard BAGID"`
	BAGBuildingID SearchMeteringPointsResponseEnvelopeBAGIDType `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard BAGBuildingID,omitempty"`
}

// Must be at least 2 items long
type SearchMeteringPointsResponseEnvelopeCountryCode string

type SearchMeteringPointsResponseEnvelopeDateTime time.Time

func (t *SearchMeteringPointsResponseEnvelopeDateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t SearchMeteringPointsResponseEnvelopeDateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeader struct {
	ContentHash       string                                                                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ContentHash,omitempty"`
	ConversationID    string                                                                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ConversationID,omitempty"`
	CorrelationID     string                                                                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard CorrelationID,omitempty"`
	CreationTimestamp SearchMeteringPointsResponseEnvelopeDateTime                              `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard CreationTimestamp"`
	DocumentID        string                                                                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard DocumentID,omitempty"`
	ExpiresAt         SearchMeteringPointsResponseEnvelopeDateTime                              `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ExpiresAt,omitempty"`
	MessageID         string                                                                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard MessageID"`
	ProcessTypeID     string                                                                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ProcessTypeID,omitempty"`
	RepeatedRequest   string                                                                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard RepeatedRequest,omitempty"`
	TestRequest       string                                                                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard TestRequest,omitempty"`
	Destination       SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderDestination `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Destination"`
	Manifest          SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderManifest    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Manifest,omitempty"`
	Source            SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderSource      `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Source"`
}

type SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderDestination struct {
	Receiver SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Receiver"`
	Service  SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderDestinationService  `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Service,omitempty"`
}

type SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderDestinationReceiver struct {
	Authority             string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ContactTypeIdentifier,omitempty"`
	ReceiverID            string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ReceiverID"`
}

type SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderDestinationService struct {
	ServiceMethod string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ServiceMethod,omitempty"`
	ServiceName   string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ServiceName,omitempty"`
}

type SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderManifest struct {
	NumberofItems float64                                                                              `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard NumberofItems"`
	ManifestItem  []SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ManifestItem"`
}

type SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderManifestManifestItem struct {
	Description               string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Description,omitempty"`
	LanguageCode              string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard LanguageCode,omitempty"`
	MimeTypeQualifierCode     string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard MimeTypeQualifierCode"`
	UniformResourceIdentifier string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard UniformResourceIdentifier"`
}

type SearchMeteringPointsResponseEnvelopeEDSNBusinessDocumentHeaderSource struct {
	Authority             string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Authority,omitempty"`
	ContactTypeIdentifier string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ContactTypeIdentifier,omitempty"`
	SenderID              string `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard SenderID"`
}

// May be no more than 3 items long
type SearchMeteringPointsResponseEnvelopeEnergyAllocationMethodType string

// May be no more than 3 items long
type SearchMeteringPointsResponseEnvelopeEnergyConnectionSubtypeType string

// May be no more than 3 items long
type SearchMeteringPointsResponseEnvelopeEnergyFlowDirectionType string

// Must be at least 1 items long
type SearchMeteringPointsResponseEnvelopeEnergyMeterIDType string

// May be no more than 3 items long
type SearchMeteringPointsResponseEnvelopeEnergyMeteringMethodType string

// May be no more than 3 items long
type SearchMeteringPointsResponseEnvelopeEnergyProductType string

// May be no more than 3 items long
type SearchMeteringPointsResponseEnvelopeEnergyUsageProfileType string

// Must be at least 1 items long
type SearchMeteringPointsResponseEnvelopeExBuildingNr string

// Must match the pattern [0-9]{13}
type SearchMeteringPointsResponseEnvelopeGLNEANCode string

// Must match the pattern [0-9]{18}
type SearchMeteringPointsResponseEnvelopeGSRNEANCode string

// Must match the pattern [0-1][0-9]
type SearchMeteringPointsResponseEnvelopeInvoiceMonthType string

// Must be at least 1 items long
type SearchMeteringPointsResponseEnvelopeLocationDescriptionType string

type SearchMeteringPointsResponseEnvelopeMPAddressResponseType struct {
	BAG          SearchMeteringPointsResponseEnvelopeBAGType      `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard BAG,omitempty"`
	StreetName   SearchMeteringPointsResponseEnvelopeAddressText  `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard StreetName,omitempty"`
	BuildingNr   int                                              `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard BuildingNr,omitempty"`
	ExBuildingNr SearchMeteringPointsResponseEnvelopeExBuildingNr `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ExBuildingNr,omitempty"`
	ZIPCode      SearchMeteringPointsResponseEnvelopeZIPCode      `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ZIPCode,omitempty"`
	CityName     SearchMeteringPointsResponseEnvelopeAddressText  `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard CityName"`
	Country      SearchMeteringPointsResponseEnvelopeCountryCode  `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Country"`
}

// May be no more than 3 items long
type SearchMeteringPointsResponseEnvelopeMarketSegmentType string

type SearchMeteringPointsResponseEnvelopePC struct {
	Result           SearchMeteringPointsResponseEnvelopePCResult `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Result"`
	PortaalRejection SearchMeteringPointsResponseEnvelopePCPR     `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Portaal_Rejection"`
}

type SearchMeteringPointsResponseEnvelopePCPR struct {
	Rejection []SearchMeteringPointsResponseEnvelopeRejectionPortaalType `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Rejection"`
}

type SearchMeteringPointsResponseEnvelopePCResult struct {
	ReachedMaxResult     int                                               `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ReachedMaxResult"`
	PortaalMeteringPoint []SearchMeteringPointsResponseEnvelopePCResultPMP `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Portaal_MeteringPoint,omitempty"`
}

type SearchMeteringPointsResponseEnvelopePCResultPMP struct {
	EANID                       SearchMeteringPointsResponseEnvelopeGSRNEANCode                          `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard EANID"`
	EDSNAddressSearch           SearchMeteringPointsResponseEnvelopeMPAddressResponseType                `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard EDSN_AddressSearch,omitempty"`
	GridArea                    SearchMeteringPointsResponseEnvelopeGSRNEANCode                          `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard GridArea,omitempty"`
	LocationDescription         SearchMeteringPointsResponseEnvelopeLocationDescriptionType              `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard LocationDescription,omitempty"`
	MarketSegment               SearchMeteringPointsResponseEnvelopeMarketSegmentType                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard MarketSegment"`
	ProductType                 SearchMeteringPointsResponseEnvelopeEnergyProductType                    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ProductType"`
	GridOperatorCompany         SearchMeteringPointsResponseEnvelopePCResultPMPGridOperatorCompany       `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard GridOperator_Company"`
	MPCommercialCharacteristics SearchMeteringPointsResponseEnvelopePCResultPMPMPCC                      `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard MPCommercialCharacteristics,omitempty"`
	PortaalEnergyMeter          []SearchMeteringPointsResponseEnvelopePCResultPMPPortaalEnergyMeter      `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Portaal_EnergyMeter,omitempty"`
	MPPhysicalCharacteristics   SearchMeteringPointsResponseEnvelopePCResultPMPMPPhysicalCharacteristics `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard MPPhysicalCharacteristics,omitempty"`
	MeteringPointGroup          SearchMeteringPointsResponseEnvelopePCResultPMPMeteringPointGroup        `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard MeteringPointGroup,omitempty"`
}

type SearchMeteringPointsResponseEnvelopePCResultPMPGridOperatorCompany struct {
	ID   SearchMeteringPointsResponseEnvelopeGLNEANCode `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ID"`
	Name string                                         `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Name,omitempty"`
}

type SearchMeteringPointsResponseEnvelopePCResultPMPMPCC struct {
	DeterminationComplex SearchMeteringPointsResponseEnvelopeYesNoNaType `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard DeterminationComplex"`
	Residential          SearchMeteringPointsResponseEnvelopeYesNoNaType `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Residential"`
}

type SearchMeteringPointsResponseEnvelopePCResultPMPMPPhysicalCharacteristics struct {
	AllocationMethod    SearchMeteringPointsResponseEnvelopeEnergyAllocationMethodType  `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard AllocationMethod,omitempty"`
	ContractedCapacity  int                                                             `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ContractedCapacity,omitempty"`
	EnergyFlowDirection SearchMeteringPointsResponseEnvelopeEnergyFlowDirectionType     `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard EnergyFlowDirection,omitempty"`
	InvoiceMonth        SearchMeteringPointsResponseEnvelopeInvoiceMonthType            `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard InvoiceMonth,omitempty"`
	MeteringMethod      SearchMeteringPointsResponseEnvelopeEnergyMeteringMethodType    `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard MeteringMethod,omitempty"`
	PhysicalCapacity    SearchMeteringPointsResponseEnvelopePhysicalCapacityType        `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard PhysicalCapacity,omitempty"`
	ProfileCategory     SearchMeteringPointsResponseEnvelopeEnergyUsageProfileType      `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ProfileCategory,omitempty"`
	Subtype             SearchMeteringPointsResponseEnvelopeEnergyConnectionSubtypeType `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard Subtype,omitempty"`
}

type SearchMeteringPointsResponseEnvelopePCResultPMPMeteringPointGroup struct {
	PAP SearchMeteringPointsResponseEnvelopePCResultPMPMeteringPointGroupPAP `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard PAP"`
}

type SearchMeteringPointsResponseEnvelopePCResultPMPMeteringPointGroupPAP struct {
	EANID SearchMeteringPointsResponseEnvelopeGSRNEANCode `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard EANID"`
}

type SearchMeteringPointsResponseEnvelopePCResultPMPPortaalEnergyMeter struct {
	ID SearchMeteringPointsResponseEnvelopeEnergyMeterIDType `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard ID,omitempty"`
}

// Must be at least 2 items long
type SearchMeteringPointsResponseEnvelopePhysicalCapacityType string

type SearchMeteringPointsResponseEnvelopeRejectionPortaalType struct {
	RejectionCode SearchMeteringPointsResponseEnvelopeRejectionReasonType `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard RejectionCode"`
	RejectionText string                                                  `xml:"urn:nedu:edsn:data:searchmeteringpointsresponse:1:standard RejectionText,omitempty"`
}

// May be no more than 3 items long
type SearchMeteringPointsResponseEnvelopeRejectionReasonType string

// Must be at least 1 items long
type SearchMeteringPointsResponseEnvelopeYesNoNaType string

// Must be at least 1 items long
type SearchMeteringPointsResponseEnvelopeZIPCode string

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type Client struct {
	HTTPClient   *http.Client
	ResponseHook func(*http.Response) *http.Response
	RequestHook  func(*http.Request) *http.Request
}
type soapEnvelope struct {
	XMLName struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  []byte   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    struct {
		Message interface{} `xml:",any,omitempty"`
		Fault   *struct {
			String string `xml:"faultstring,omitempty"`
			Code   string `xml:"faultcode,omitempty"`
			Detail string `xml:"detail,omitempty"`
		} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
	} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

func (c *Client) do(ctx context.Context, method, uri, action string, in, out interface{}) error {
	var body io.Reader
	var envelope soapEnvelope
	if method == "POST" || method == "PUT" {
		var buf bytes.Buffer
		envelope.Body.Message = in
		enc := xml.NewEncoder(&buf)
		if err := enc.Encode(envelope); err != nil {
			return err
		}
		if err := enc.Flush(); err != nil {
			return err
		}
		body = &buf
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}
	req.Header.Set("SOAPAction", action)
	req = req.WithContext(ctx)
	if c.RequestHook != nil {
		req = c.RequestHook(req)
	}
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	rsp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	if c.ResponseHook != nil {
		rsp = c.ResponseHook(rsp)
	}
	dec := xml.NewDecoder(rsp.Body)

	envelope.Body.Message = out
	if err := dec.Decode(&envelope); err != nil {
		return err
	}

	if envelope.Body.Fault != nil {
		return fmt.Errorf("%s: %s", envelope.Body.Fault.Code, envelope.Body.Fault.String)
	}
	return nil
}
func (c *Client) SearchMeteringPoints(ctx context.Context, part1 SearchMeteringPointsRequestEnvelope) (SearchMeteringPointsResponseEnvelope, error) {
	var response SearchMeteringPointsResponseEnvelope
	err := c.do(ctx, "POST", "https://portaal-opt.edsn.nl/b2b/synchroon/ResponderSearchMeteringPointsMPRespondingActivity", "urn:SearchMeteringPoints", &part1, &response)
	return response, err
}