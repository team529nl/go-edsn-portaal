package portaal

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type SearchMeteringPointsRequestEnvelope_AddressText string

type SearchMeteringPointsRequestEnvelope_Amount int32

type SearchMeteringPointsRequestEnvelope_BAGIDType string

type SearchMeteringPointsRequestEnvelope_BuildingNumber int32

type SearchMeteringPointsRequestEnvelope_CountryCode string

type SearchMeteringPointsRequestEnvelope_DateTime time.Time

type SearchMeteringPointsRequestEnvelope_EnergyConnectionPhysicalStatusType string

type SearchMeteringPointsRequestEnvelope_EnergyConnectionSubtypeType string

type SearchMeteringPointsRequestEnvelope_EnergyConnectionType string

type SearchMeteringPointsRequestEnvelope_EnergyDeliveryStatusType string

type SearchMeteringPointsRequestEnvelope_EnergyProductType string

type SearchMeteringPointsRequestEnvelope_ExBuildingNr string

type SearchMeteringPointsRequestEnvelope_GeoCoordinateMeasureType float64

type SearchMeteringPointsRequestEnvelope_GLNEANCode string

type SearchMeteringPointsRequestEnvelope_GSRNEANCode string

type SearchMeteringPointsRequestEnvelope_Identifier string

type SearchMeteringPointsRequestEnvelope_IndicationType string

type SearchMeteringPointsRequestEnvelope_Indicator string

type SearchMeteringPointsRequestEnvelope_LocationDescriptionType string

type SearchMeteringPointsRequestEnvelope_MarketSegmentType string

type SearchMeteringPointsRequestEnvelope_Numeric float64

type SearchMeteringPointsRequestEnvelope_Text string

type SearchMeteringPointsRequestEnvelope_TNTIDType string

type SearchMeteringPointsRequestEnvelope_YesNoNaType string

type SearchMeteringPointsRequestEnvelope_ZIPCode string

type SearchMeteringPointsRequestEnvelope struct {
	EDSNBusinessDocumentHeader *SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader `xml:"EDSNBusinessDocumentHeader,omitempty"`

	CAR_Content *SearchMeteringPointsRequestEnvelope_CC `xml:"CAR_Content,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_CC struct {
	CAR_MeteringPoint *SearchMeteringPointsRequestEnvelope_CC_CMP `xml:"CAR_MeteringPoint,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_CC_CMP struct {
	EANID *SearchMeteringPointsRequestEnvelope_GSRNEANCode `xml:"EANID,omitempty"`

	GridArea *SearchMeteringPointsRequestEnvelope_GSRNEANCode `xml:"GridArea,omitempty"`

	EDSN_AddressExtended *SearchMeteringPointsRequestEnvelope_Address `xml:"EDSN_AddressExtended,omitempty"`

	MarketSegment *SearchMeteringPointsRequestEnvelope_MarketSegmentType `xml:"MarketSegment,omitempty"`

	LocationDescription *SearchMeteringPointsRequestEnvelope_LocationDescriptionType `xml:"LocationDescription,omitempty"`

	ProductType *SearchMeteringPointsRequestEnvelope_EnergyProductType `xml:"ProductType,omitempty"`

	Result *SearchMeteringPointsRequestEnvelope_CC_CMP_RES `xml:"Result,omitempty"`

	GridOperator_Company *SearchMeteringPointsRequestEnvelope_CC_CMP_GOC `xml:"GridOperator_Company,omitempty"`

	MPCommercialCharacteristics *SearchMeteringPointsRequestEnvelope_CC_CMP_MPCC `xml:"MPCommercialCharacteristics,omitempty"`

	MPPhysicalCharacteristics *SearchMeteringPointsRequestEnvelope_CC_CMP_MPPC `xml:"MPPhysicalCharacteristics,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_CC_CMP_GOC struct {
	ID *SearchMeteringPointsRequestEnvelope_GLNEANCode `xml:"ID,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_CC_CMP_MPCC struct {
	DeterminationComplex *SearchMeteringPointsRequestEnvelope_YesNoNaType `xml:"DeterminationComplex,omitempty"`

	Residential *SearchMeteringPointsRequestEnvelope_YesNoNaType `xml:"Residential,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_CC_CMP_MPPC struct {
	Appliance *SearchMeteringPointsRequestEnvelope_IndicationType `xml:"Appliance,omitempty"`

	EnergyDeliveryStatus *SearchMeteringPointsRequestEnvelope_EnergyDeliveryStatusType `xml:"EnergyDeliveryStatus,omitempty"`

	PhysicalStatus *SearchMeteringPointsRequestEnvelope_EnergyConnectionPhysicalStatusType `xml:"PhysicalStatus,omitempty"`

	Subtype *SearchMeteringPointsRequestEnvelope_EnergyConnectionSubtypeType `xml:"Subtype,omitempty"`

	Type *SearchMeteringPointsRequestEnvelope_EnergyConnectionType `xml:"Type,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_CC_CMP_RES struct {
	ResultSet *SearchMeteringPointsRequestEnvelope_Amount `xml:"ResultSet,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader struct {
	ContentHash *SearchMeteringPointsRequestEnvelope_Text `xml:"ContentHash,omitempty"`

	ConversationID *SearchMeteringPointsRequestEnvelope_Identifier `xml:"ConversationID,omitempty"`

	CorrelationID *SearchMeteringPointsRequestEnvelope_Identifier `xml:"CorrelationID,omitempty"`

	CreationTimestamp *SearchMeteringPointsRequestEnvelope_DateTime `xml:"CreationTimestamp,omitempty"`

	DocumentID *SearchMeteringPointsRequestEnvelope_Text `xml:"DocumentID,omitempty"`

	ExpiresAt *SearchMeteringPointsRequestEnvelope_DateTime `xml:"ExpiresAt,omitempty"`

	MessageID *SearchMeteringPointsRequestEnvelope_Identifier `xml:"MessageID,omitempty"`

	ProcessTypeID *SearchMeteringPointsRequestEnvelope_Text `xml:"ProcessTypeID,omitempty"`

	RepeatedRequest *SearchMeteringPointsRequestEnvelope_Identifier `xml:"RepeatedRequest,omitempty"`

	TestRequest *SearchMeteringPointsRequestEnvelope_Indicator `xml:"TestRequest,omitempty"`

	Destination *SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Destination `xml:"Destination,omitempty"`

	Manifest *SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Manifest `xml:"Manifest,omitempty"`

	Source *SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Source `xml:"Source,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Destination struct {
	Receiver *SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Destination_Receiver `xml:"Receiver,omitempty"`

	Service *SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Destination_Service `xml:"Service,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Destination_Receiver struct {
	Authority *SearchMeteringPointsRequestEnvelope_Text `xml:"Authority,omitempty"`

	ContactTypeIdentifier *SearchMeteringPointsRequestEnvelope_Text `xml:"ContactTypeIdentifier,omitempty"`

	ReceiverID *SearchMeteringPointsRequestEnvelope_Identifier `xml:"ReceiverID,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Destination_Service struct {
	ServiceMethod *SearchMeteringPointsRequestEnvelope_Text `xml:"ServiceMethod,omitempty"`

	ServiceName *SearchMeteringPointsRequestEnvelope_Text `xml:"ServiceName,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Manifest struct {
	NumberofItems *SearchMeteringPointsRequestEnvelope_Numeric `xml:"NumberofItems,omitempty"`

	ManifestItem []*SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Manifest_ManifestItem `xml:"ManifestItem,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Manifest_ManifestItem struct {
	Description *SearchMeteringPointsRequestEnvelope_Text `xml:"Description,omitempty"`

	LanguageCode *SearchMeteringPointsRequestEnvelope_Text `xml:"LanguageCode,omitempty"`

	MimeTypeQualifierCode *SearchMeteringPointsRequestEnvelope_Text `xml:"MimeTypeQualifierCode,omitempty"`

	UniformResourceIdentifier *SearchMeteringPointsRequestEnvelope_Text `xml:"UniformResourceIdentifier,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_EDSNBusinessDocumentHeader_Source struct {
	Authority *SearchMeteringPointsRequestEnvelope_Text `xml:"Authority,omitempty"`

	ContactTypeIdentifier *SearchMeteringPointsRequestEnvelope_Text `xml:"ContactTypeIdentifier,omitempty"`

	SenderID *SearchMeteringPointsRequestEnvelope_Identifier `xml:"SenderID,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_Address struct {
	BAG *SearchMeteringPointsRequestEnvelope_BAGType `xml:"BAG,omitempty"`

	StreetName *SearchMeteringPointsRequestEnvelope_AddressText `xml:"StreetName,omitempty"`

	BuildingNr *SearchMeteringPointsRequestEnvelope_BuildingNumber `xml:"BuildingNr,omitempty"`

	ExBuildingNr *SearchMeteringPointsRequestEnvelope_ExBuildingNr `xml:"ExBuildingNr,omitempty"`

	ZIPCode *SearchMeteringPointsRequestEnvelope_ZIPCode `xml:"ZIPCode,omitempty"`

	CityName *SearchMeteringPointsRequestEnvelope_AddressText `xml:"CityName,omitempty"`

	Country *SearchMeteringPointsRequestEnvelope_CountryCode `xml:"Country,omitempty"`

	EDSN_GeographicalCoordinate *SearchMeteringPointsRequestEnvelope_GeographicalCoordinate `xml:"EDSN_GeographicalCoordinate,omitempty"`

	TNTID *SearchMeteringPointsRequestEnvelope_TNTIDType `xml:"TNTID,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_BAGType struct {
	BAGID *SearchMeteringPointsRequestEnvelope_BAGIDType `xml:"BAGID,omitempty"`

	BAGBuildingID *SearchMeteringPointsRequestEnvelope_BAGIDType `xml:"BAGBuildingID,omitempty"`
}

type SearchMeteringPointsRequestEnvelope_GeographicalCoordinate struct {
	Latitude *SearchMeteringPointsRequestEnvelope_GeoCoordinateMeasureType `xml:"Latitude,omitempty"`

	Longitude *SearchMeteringPointsRequestEnvelope_GeoCoordinateMeasureType `xml:"Longitude,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_AddressText string

type SearchMeteringPointsResponseEnvelope_Amount int32

type SearchMeteringPointsResponseEnvelope_BuildingNumber int32

type SearchMeteringPointsResponseEnvelope_CountryCode string

type SearchMeteringPointsResponseEnvelope_DateTime time.Time

type SearchMeteringPointsResponseEnvelope_EnergyConnectionSubtypeType string

type SearchMeteringPointsResponseEnvelope_EnergyConnectionType string

type SearchMeteringPointsResponseEnvelope_ExBuildingNr string

type SearchMeteringPointsResponseEnvelope_GSRNEANCode string

type SearchMeteringPointsResponseEnvelope_Identifier string

type SearchMeteringPointsResponseEnvelope_IndicationType string

type SearchMeteringPointsResponseEnvelope_Indicator string

type SearchMeteringPointsResponseEnvelope_LocationDescriptionType string

type SearchMeteringPointsResponseEnvelope_Numeric float64

type SearchMeteringPointsResponseEnvelope_Text string

type SearchMeteringPointsResponseEnvelope_YesNoType string

type SearchMeteringPointsResponseEnvelope_ZIPCode string

type SearchMeteringPointsResponseEnvelope_RejectionReasonCARCode string

type SearchMeteringPointsResponseEnvelope struct {
	EDSNBusinessDocumentHeader *SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader `xml:"EDSNBusinessDocumentHeader,omitempty"`

	CAR_Content *SearchMeteringPointsResponseEnvelope_CC `xml:"CAR_Content,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_CC struct {
	Result *SearchMeteringPointsResponseEnvelope_CC_RES `xml:"Result,omitempty"`

	CAR_Rejection *SearchMeteringPointsResponseEnvelope_CC_CRJ `xml:"CAR_Rejection,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_CC_CRJ struct {
	Rejection []*SearchMeteringPointsResponseEnvelope_RejectionCARType `xml:"Rejection,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_CC_RES struct {
	MaxResults *SearchMeteringPointsResponseEnvelope_YesNoType `xml:"MaxResults,omitempty"`

	ResultSet *SearchMeteringPointsResponseEnvelope_Amount `xml:"ResultSet,omitempty"`

	CAR_MeteringPoint []*SearchMeteringPointsResponseEnvelope_CC_RES_CMP `xml:"CAR_MeteringPoint,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_CC_RES_CMP struct {
	EANID *SearchMeteringPointsResponseEnvelope_GSRNEANCode `xml:"EANID,omitempty"`

	EDSN_Address *SearchMeteringPointsResponseEnvelope_SimpleAddressType `xml:"EDSN_Address,omitempty"`

	LocationDescription *SearchMeteringPointsResponseEnvelope_LocationDescriptionType `xml:"LocationDescription,omitempty"`

	MPPhysicalCharacteristics *SearchMeteringPointsResponseEnvelope_CC_RES_CMP_MPPC `xml:"MPPhysicalCharacteristics,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_CC_RES_CMP_MPPC struct {
	Appliance *SearchMeteringPointsResponseEnvelope_IndicationType `xml:"Appliance,omitempty"`

	Subtype *SearchMeteringPointsResponseEnvelope_EnergyConnectionSubtypeType `xml:"Subtype,omitempty"`

	Type *SearchMeteringPointsResponseEnvelope_EnergyConnectionType `xml:"Type,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader struct {
	ContentHash *SearchMeteringPointsResponseEnvelope_Text `xml:"ContentHash,omitempty"`

	ConversationID *SearchMeteringPointsResponseEnvelope_Identifier `xml:"ConversationID,omitempty"`

	CorrelationID *SearchMeteringPointsResponseEnvelope_Identifier `xml:"CorrelationID,omitempty"`

	CreationTimestamp *SearchMeteringPointsResponseEnvelope_DateTime `xml:"CreationTimestamp,omitempty"`

	DocumentID *SearchMeteringPointsResponseEnvelope_Text `xml:"DocumentID,omitempty"`

	ExpiresAt *SearchMeteringPointsResponseEnvelope_DateTime `xml:"ExpiresAt,omitempty"`

	MessageID *SearchMeteringPointsResponseEnvelope_Identifier `xml:"MessageID,omitempty"`

	ProcessTypeID *SearchMeteringPointsResponseEnvelope_Text `xml:"ProcessTypeID,omitempty"`

	RepeatedRequest *SearchMeteringPointsResponseEnvelope_Identifier `xml:"RepeatedRequest,omitempty"`

	TestRequest *SearchMeteringPointsResponseEnvelope_Indicator `xml:"TestRequest,omitempty"`

	Destination *SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Destination `xml:"Destination,omitempty"`

	Manifest *SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Manifest `xml:"Manifest,omitempty"`

	Source *SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Source `xml:"Source,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Destination struct {
	Receiver *SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Destination_Receiver `xml:"Receiver,omitempty"`

	Service *SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Destination_Service `xml:"Service,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Destination_Receiver struct {
	Authority *SearchMeteringPointsResponseEnvelope_Text `xml:"Authority,omitempty"`

	ContactTypeIdentifier *SearchMeteringPointsResponseEnvelope_Text `xml:"ContactTypeIdentifier,omitempty"`

	ReceiverID *SearchMeteringPointsResponseEnvelope_Identifier `xml:"ReceiverID,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Destination_Service struct {
	ServiceMethod *SearchMeteringPointsResponseEnvelope_Text `xml:"ServiceMethod,omitempty"`

	ServiceName *SearchMeteringPointsResponseEnvelope_Text `xml:"ServiceName,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Manifest struct {
	NumberofItems *SearchMeteringPointsResponseEnvelope_Numeric `xml:"NumberofItems,omitempty"`

	ManifestItem []*SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Manifest_ManifestItem `xml:"ManifestItem,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Manifest_ManifestItem struct {
	Description *SearchMeteringPointsResponseEnvelope_Text `xml:"Description,omitempty"`

	LanguageCode *SearchMeteringPointsResponseEnvelope_Text `xml:"LanguageCode,omitempty"`

	MimeTypeQualifierCode *SearchMeteringPointsResponseEnvelope_Text `xml:"MimeTypeQualifierCode,omitempty"`

	UniformResourceIdentifier *SearchMeteringPointsResponseEnvelope_Text `xml:"UniformResourceIdentifier,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_EDSNBusinessDocumentHeader_Source struct {
	Authority *SearchMeteringPointsResponseEnvelope_Text `xml:"Authority,omitempty"`

	ContactTypeIdentifier *SearchMeteringPointsResponseEnvelope_Text `xml:"ContactTypeIdentifier,omitempty"`

	SenderID *SearchMeteringPointsResponseEnvelope_Identifier `xml:"SenderID,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_RejectionCARType struct {
	RejectionCode *SearchMeteringPointsResponseEnvelope_RejectionReasonCARCode `xml:"RejectionCode,omitempty"`

	RejectionText *SearchMeteringPointsResponseEnvelope_Text `xml:"RejectionText,omitempty"`
}

type SearchMeteringPointsResponseEnvelope_SimpleAddressType struct {
	StreetName *SearchMeteringPointsResponseEnvelope_AddressText `xml:"StreetName,omitempty"`

	BuildingNr *SearchMeteringPointsResponseEnvelope_BuildingNumber `xml:"BuildingNr,omitempty"`

	ExBuildingNr *SearchMeteringPointsResponseEnvelope_ExBuildingNr `xml:"ExBuildingNr,omitempty"`

	ZIPCode *SearchMeteringPointsResponseEnvelope_ZIPCode `xml:"ZIPCode,omitempty"`

	CityName *SearchMeteringPointsResponseEnvelope_AddressText `xml:"CityName,omitempty"`

	Country *SearchMeteringPointsResponseEnvelope_CountryCode `xml:"Country,omitempty"`
}

// Nummer (getal).
type SOAPFault_Numeric float64

// Tekst (string).
type SOAPFault_Text string

type SOAPFault struct {
	ErrorCode *SOAPFault_Numeric `xml:"ErrorCode,omitempty"`

	ErrorText *SOAPFault_Text `xml:"ErrorText,omitempty"`

	ErrorDetails *SOAPFault_Text `xml:"ErrorDetails,omitempty"`
}

type SearchMeteringPointsPortType interface {

	// Error can be either of the following types:
	//
	//   - SOAPFault

	SearchMeteringPointsRequest(request *SearchMeteringPointsRequestEnvelope) (*SearchMeteringPointsResponseEnvelope, error)
}

type searchMeteringPointsPortType struct {
	client *soap.Client
}

func NewSearchMeteringPointsPortType(client *soap.Client) SearchMeteringPointsPortType {
	return &searchMeteringPointsPortType{
		client: client,
	}
}

func (service *searchMeteringPointsPortType) SearchMeteringPointsRequest(request *SearchMeteringPointsRequestEnvelope) (*SearchMeteringPointsResponseEnvelope, error) {
	response := new(SearchMeteringPointsResponseEnvelope)
	err := service.client.Call("urn:SearchMeteringPoints", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
