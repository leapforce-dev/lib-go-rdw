package rdw

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

// KentekenInfo stores KentekenInfo from Service
//
type KentekenInfo struct {
	Kenteken          string                `json:"kenteken"`
	Voertuigsoort     *string               `json:"voertuigsoort"`
	Merk              *string               `json:"merk"`
	Handelsbenaming   *string               `json:"handelsbenaming"`
	BrutoBPM          *go_types.Int64String `json:"bruto_bpm"`
	Inrichting        *string               `json:"inrichting"`
	AantalZitplaatsen *go_types.Int64String `json:"aantal_zitplaatsen"`
	EersteKleur       *string               `json:"eerste_kleur"`
	TweedeKleur       *string               `json:"tweede_kleur"`
	AantalCilinders   *go_types.Int64String `json:"aantal_cilinders"`
	Cilinderinhoud    *go_types.Int64String `json:"cilinderinhoud"`
	Catalogusprijs    *go_types.Int64String `json:"catalogusprijs"`
	WAMVerzekerd      *string               `json:"wam_verzekerd"`
	Lengte            *go_types.Int64String `json:"lengte"`
	Breedte           *go_types.Int64String `json:"breedte"`
}

type GetKentekenInfoConfig struct {
	Kenteken string
}

// GetKentekenInfo returns all kentekenInfo
//
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
