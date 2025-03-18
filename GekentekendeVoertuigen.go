package rdw

import (
	"fmt"
	"github.com/leapforce-libraries/go_rdw/types"
	go_types "github.com/leapforce-libraries/go_types"
	"net/http"
	"net/url"
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// GekentekendeVoertuigen stores GekentekendeVoertuigen from Service
type GekentekendeVoertuigen struct {
	Kenteken                               string                  `json:"kenteken"`
	Voertuigsoort                          string                  `json:"voertuigsoort"`
	Merk                                   string                  `json:"merk"`
	Handelsbenaming                        string                  `json:"handelsbenaming"`
	VervaldatumApk                         string                  `json:"vervaldatum_apk"`
	DatumTenaamstelling                    string                  `json:"datum_tenaamstelling"`
	BrutoBpm                               *go_types.Int64String   `json:"bruto_bpm"`
	Inrichting                             string                  `json:"inrichting"`
	AantalZitplaatsen                      *go_types.Int64String   `json:"aantal_zitplaatsen"`
	EersteKleur                            string                  `json:"eerste_kleur"`
	TweedeKleur                            string                  `json:"tweede_kleur"`
	AantalCilinders                        *go_types.Int64String   `json:"aantal_cilinders"`
	Cilinderinhoud                         *go_types.Int64String   `json:"cilinderinhoud"`
	MassaLedigVoertuig                     *go_types.Int64String   `json:"massa_ledig_voertuig"`
	ToegestaneMaximumMassaVoertuig         *go_types.Int64String   `json:"toegestane_maximum_massa_voertuig"`
	MassaRijklaar                          *go_types.Int64String   `json:"massa_rijklaar"`
	MaximumMassaTrekkenOngeremd            *go_types.Int64String   `json:"maximum_massa_trekken_ongeremd"`
	MaximumTrekkenMassaGeremd              *go_types.Int64String   `json:"maximum_trekken_massa_geremd"`
	DatumEersteToelating                   string                  `json:"datum_eerste_toelating"`
	DatumEersteTenaamstellingInNederland   string                  `json:"datum_eerste_tenaamstelling_in_nederland"`
	WachtOpKeuren                          string                  `json:"wacht_op_keuren"`
	Catalogusprijs                         string                  `json:"catalogusprijs"`
	WamVerzekerd                           string                  `json:"wam_verzekerd"`
	AantalDeuren                           *go_types.Int64String   `json:"aantal_deuren"`
	AantalWielen                           *go_types.Int64String   `json:"aantal_wielen"`
	EuropeseVoertuigcategorie              string                  `json:"europese_voertuigcategorie"`
	PlaatsChassisnummer                    string                  `json:"plaats_chassisnummer"`
	TechnischeMaxMassaVoertuig             *go_types.Int64String   `json:"technische_max_massa_voertuig"`
	Type                                   string                  `json:"type"`
	Typegoedkeuringsnummer                 string                  `json:"typegoedkeuringsnummer"`
	Variant                                string                  `json:"variant"`
	Uitvoering                             string                  `json:"uitvoering"`
	VolgnummerWijzigingEuTypegoedkeuring   string                  `json:"volgnummer_wijziging_eu_typegoedkeuring"`
	VermogenMassarijklaar                  *go_types.Float64String `json:"vermogen_massarijklaar"`
	Wielbasis                              *go_types.Int64String   `json:"wielbasis"`
	ExportIndicator                        string                  `json:"export_indicator"`
	OpenstaandeTerugroepactieIndicator     string                  `json:"openstaande_terugroepactie_indicator"`
	TaxiIndicator                          string                  `json:"taxi_indicator"`
	MaximumMassaSamenstelling              *go_types.Int64String   `json:"maximum_massa_samenstelling"`
	AantalRolstoelplaatsen                 *go_types.Int64String   `json:"aantal_rolstoelplaatsen"`
	JaarLaatsteRegistratieTellerstand      *go_types.Int64String   `json:"jaar_laatste_registratie_tellerstand"`
	Tellerstandoordeel                     string                  `json:"tellerstandoordeel"`
	CodeToelichtingTellerstandoordeel      string                  `json:"code_toelichting_tellerstandoordeel"`
	TenaamstellenMogelijk                  string                  `json:"tenaamstellen_mogelijk"`
	VervaldatumApkDt                       types.DateString        `json:"vervaldatum_apk_dt"`
	DatumTenaamstellingDt                  types.DateString        `json:"datum_tenaamstelling_dt"`
	DatumEersteToelatingDt                 types.DateString        `json:"datum_eerste_toelating_dt"`
	DatumEersteTenaamstellingInNederlandDt types.DateString        `json:"datum_eerste_tenaamstelling_in_nederland_dt"`
	Zuinigheidsclassificatie               string                  `json:"zuinigheidsclassificatie"`
}

type GetGekentekendeVoertuigenConfig struct {
	Kenteken string
}

// GetGekentekendeVoertuigen returns all gekentekendeVoertuigenBrandstof
func (service *Service) GetGekentekendeVoertuigen(config *GetGekentekendeVoertuigenConfig) (*GekentekendeVoertuigen, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetGekentekendeVoertuigenConfig is nil")
	}

	values := url.Values{}
	values.Set("kenteken", config.Kenteken)

	path := fmt.Sprintf("%v.json?%s", DataIdentifierGekentekendeVoertuigen, values.Encode())
	gekentekendeVoertuigenBrandstof := GekentekendeVoertuigen{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(path),
		ResponseModel: &gekentekendeVoertuigenBrandstof,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &gekentekendeVoertuigenBrandstof, nil
}

type ListGekentekendeVoertuigenConfig struct {
	Limit  *int
	Offset *int
	Order  string
}

// GetGekentekendeVoertuigen returns all gekentekendeVoertuigenBrandstof
func (service *Service) ListGekentekendeVoertuigen(config *ListGekentekendeVoertuigenConfig) (*[]GekentekendeVoertuigen, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetGekentekendeVoertuigenConfig is nil")
	}

	values := url.Values{}
	values.Set("order", config.Order)
	if config.Limit != nil {
		values.Set("limit", strconv.Itoa(*config.Limit))
	}
	if config.Offset != nil {
		values.Set("offset", strconv.Itoa(*config.Offset))
	}

	path := fmt.Sprintf("%v.json?%s", DataIdentifierGekentekendeVoertuigen, values.Encode())
	response := []GekentekendeVoertuigen{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(path),
		ResponseModel: &response,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &response, nil
}
