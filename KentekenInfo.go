package rdw

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

// KentekenInfo stores KentekenInfo from Service
type KentekenInfo struct {
	Kenteken                               string                `json:"kenteken"`
	Voertuigsoort                          *string               `json:"voertuigsoort"`
	Merk                                   *string               `json:"merk"`
	Handelsbenaming                        *string               `json:"handelsbenaming"`
	VervaldatumApk                         string                `json:"vervaldatum_apk"`
	DatumTenaamstelling                    string                `json:"datum_tenaamstelling"`
	BrutoBpm                               *go_types.Int64String `json:"bruto_bpm"`
	Inrichting                             *string               `json:"inrichting"`
	AantalZitplaatsen                      *go_types.Int64String `json:"aantal_zitplaatsen"`
	EersteKleur                            *string               `json:"eerste_kleur"`
	TweedeKleur                            *string               `json:"tweede_kleur"`
	AantalCilinders                        *go_types.Int64String `json:"aantal_cilinders"`
	Cilinderinhoud                         *go_types.Int64String `json:"cilinderinhoud"`
	DatumEersteToelating                   string                `json:"datum_eerste_toelating"`
	DatumEersteTenaamstellingInNederland   string                `json:"datum_eerste_tenaamstelling_in_nederland"`
	Catalogusprijs                         string                `json:"catalogusprijs"`
	WamVerzekerd                           string                `json:"wam_verzekerd"`
	Lengte                                 *go_types.Int64String `json:"lengte"`
	Breedte                                *go_types.Int64String `json:"breedte"`
	JaarLaatsteRegistratieTellerstand      string                `json:"jaar_laatste_registratie_tellerstand"`
	Tellerstandoordeel                     string                `json:"tellerstandoordeel"`
	CodeToelichtingTellerstandoordeel      string                `json:"code_toelichting_tellerstandoordeel"`
	TenaamstellenMogelijk                  string                `json:"tenaamstellen_mogelijk"`
	VervaldatumApkDt                       string                `json:"vervaldatum_apk_dt"`
	DatumTenaamstellingDt                  string                `json:"datum_tenaamstelling_dt"`
	DatumEersteToelatingDt                 string                `json:"datum_eerste_toelating_dt"`
	DatumEersteTenaamstellingInNederlandDt string                `json:"datum_eerste_tenaamstelling_in_nederland_dt"`
	Zuinigheidsclassificatie               string                `json:"zuinigheidsclassificatie"`
}

type GetKentekenInfoConfig struct {
	Kenteken string
}

// GetKentekenInfo returns all kentekenInfo
func (service *Service) GetKentekenInfo(config *GetKentekenInfoConfig) (*[]KentekenInfo, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetKentekenInfoConfig is nil")
	}

	values := url.Values{}
	values.Set("kenteken", config.Kenteken)

	path := fmt.Sprintf("%v.json?%s", DataIdentifierKentekenInfo, values.Encode())
	kentekenInfo := []KentekenInfo{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(path),
		ResponseModel: &kentekenInfo,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &kentekenInfo, nil
}

type GetKentekenInfosConfig struct {
	Limit  *int
	Offset *int
	Order  string
}

// GetKentekenInfos returns all kentekenInfo
func (service *Service) GetKentekenInfos(config *GetKentekenInfosConfig) (*[]KentekenInfo, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetKentekenInfoConfig is nil")
	}

	values := url.Values{}
	values.Set("order", config.Order)
	if config.Limit != nil {
		values.Set("limit", strconv.Itoa(*config.Limit))
	}
	if config.Offset != nil {
		values.Set("offset", strconv.Itoa(*config.Offset))
	}

	path := fmt.Sprintf("%v.json?%s", DataIdentifierKentekenInfo, values.Encode())
	kentekenInfo := []KentekenInfo{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(path),
		ResponseModel: &kentekenInfo,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &kentekenInfo, nil
}
