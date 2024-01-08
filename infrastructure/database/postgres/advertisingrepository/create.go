package advertisingrepository

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/d90ares/gloisp-web/internal/advertisingregister/core/domain"
	"github.com/d90ares/gloisp-web/internal/advertisingregister/core/dto"
)

func (repository repository) Create(ar *dto.AdvertisingRequest) (*domain.Advertising, error) {
	ctx := context.Background()

	category, err := strconv.Atoi(ar.Category)

	if err != nil {
		return nil, err
	}

	product := domain.Advertising{
		ID:          domain.NewID(),
		Title:       ar.Title,
		Description: ar.Description,
		Category:    domain.CategoryType(category),
		Logo:        domain.Image{ID: domain.NewID(), ImagePath: domain.ImagePath(category)},
		Banner:      domain.Image{ID: domain.NewID(), ImagePath: GetPath()},
		Phone1:      ar.Phone1,
		Phone2:      ar.Phone2,
		Email:       ar.Email,
		Site:        ar.Site,
		WhatsApp:    ar.WhatsApp,
		Facebook:    ar.Facebook,
		Instagram:   ar.Instagram,
		Twitter:     ar.Twitter,
		CreatedAt:   time.Now(),
	}

	leng := reflect.ValueOf(product)

	placeholders := make([]string, leng.Len())

	for i := 0; i < leng.NumField(); i++ {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf(`INSERT INTO product (
		id, name, description, category, logo, banner, phone1,
		phone2, email, site, whatsapp, facebook, instagram, twitter, created_at
		) VALUES (%s)
		returning *`, strings.Join(placeholders, ","))

	errr := repository.db.QueryRow(
		ctx,
		query,
	).Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Category,
		&product.Logo,
		&product.Banner,
	)

	if errr != nil {
		return nil, err
	}

	return &product, nil
}
