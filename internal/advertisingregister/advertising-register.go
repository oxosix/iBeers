package advertisingregister

import (
	"time"

	"github.com/google/uuid"
)

type Advertising struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Logo        string    `json:"logo"`
	Banner      string    `json:"banner"`
	Phone1      string    `json:"phone1"`
	Phone2      string    `json:"phone2"`
	Email       string    `json:"email"`
	Site        string    `json:"site"`
	WhatsApp    string    `json:"whatsapp"`
	Facebook    string    `json:"facebook"`
	Instagram   string    `json:"instagram"`
	Twitter     string    `json:"twitter"`
	CreatedAt   time.Time `json:"created_at"`
}
