package dto

import (
	"encoding/json"
	"io"
	// "github.com/d90ares/gloisp-web/internal/advertising-register/core/domain"
)

type AdvertisingRequest struct {
	Title       string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Logo        string `json:"logo"`
	Banner      string `json:"banner"`
	Phone1      string `json:"phone1"`
	Phone2      string `json:"phone2"`
	Email       string `json:"email"`
	Site        string `json:"site"`
	WhatsApp    string `json:"whatsapp"`
	Facebook    string `json:"facebook"`
	Instagram   string `json:"instagram"`
	Twitter     string `json:"twitte"`
}

func FromJSONCreateADVRegisterReq(body io.Reader) (*AdvertisingRequest, error) {
	createAdvertisingRequest := AdvertisingRequest{}
	if err := json.NewDecoder(body).Decode(&createAdvertisingRequest); err != nil {
		return nil, err
	}
	return &createAdvertisingRequest, nil
}
