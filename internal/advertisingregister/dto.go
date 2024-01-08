package advertisingregister

// "github.com/d90ares/gloisp-web/internal/advertising-register/core/domain"

type AdvertisingDTO struct {
	Title       string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Logo        []byte `json:"logo"`
	Banner      []byte `json:"banner"`
	Phone1      string `json:"phone1"`
	Phone2      string `json:"phone2"`
	Email       string `json:"email"`
	Site        string `json:"site"`
	WhatsApp    string `json:"whatsapp"`
	Facebook    string `json:"facebook"`
	Instagram   string `json:"instagram"`
	Twitter     string `json:"twitte"`
}
