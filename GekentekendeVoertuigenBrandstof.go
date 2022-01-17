package rdw

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	go_types "github.com/leapforce-libraries/go_types"
)

// GekentekendeVoertuigenBrandstof stores GekentekendeVoertuigenBrandstof from Service
//
type GekentekendeVoertuigenBrandstof struct {
	Kenteken                       string                  `json:"kenteken"`
	BrandstofVolgnummer            *go_types.Int64String   `json:"brandstof_volgnummer"`
	BrandstofOmschrijving          *go_types.Int64String   `json:"brandstof_omschrijving"`
	BrandstofverbruikBuiten        *go_types.Float64String `json:"brandstofverbruik_buiten"`
	BrandstofverbruikGecombineerd  *go_types.Float64String `json:"brandstofverbruik_gecombineerd"`
	BrandstofverbruikStad          *go_types.Float64String `json:"brandstofverbruik_stad"`
	CO2UitstootGecombineerd        *go_types.Int64String   `json:"co2_uitstoot_gecombineerd"`
	GeluidsniveauStationair        *go_types.Int64String   `json:"geluidsniveau_stationair"`
	EmissiecodeOmschrijving        *go_types.Int64String   `json:"emissiecode_omschrijving"`
	MilieuklasseEGGoedkeuringLicht *string                 `json:"milieuklasse_eg_goedkeuring_licht"`
	NettoMaximumVermogen           *go_types.Float64String `json:"nettomaximumvermogen"`
	ToerentalGeluidsniveau         *go_types.Int64String   `json:"toerental_geluidsniveau"`
}

type GetGekentekendeVoertuigenBrandstofConfig struct {
	Kenteken string
}

// GetGekentekendeVoertuigenBrandstof returns all gekentekendeVoertuigenBrandstof
//
func (service *Service) GetGekentekendeVoertuigenBrandstof(config *GetGekentekendeVoertuigenBrandstofConfig) (*GekentekendeVoertuigenBrandstof, *errortools.Error) {
	if config != nil {
		return nil, errortools.ErrorMessage("GetGekentekendeVoertuigenBrandstofConfig is nil")
	}

	values := url.Values{}
	values.Set("kenteken", config.Kenteken)

	path := fmt.Sprintf("%v.json?%s", DataIdentifierGekentekendeVoertuigenBrandstof, values.Encode())
	gekentekendeVoertuigenBrandstof := GekentekendeVoertuigenBrandstof{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url(path),
		ResponseModel: &gekentekendeVoertuigenBrandstof,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &gekentekendeVoertuigenBrandstof, nil
}
