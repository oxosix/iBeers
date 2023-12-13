package advertisingrepository

import (
	"context"
	"strconv"
	"time"

	"github.com/d90ares/gloisp-web/internal/advertisingregister/core/domain"
	"github.com/d90ares/gloisp-web/internal/advertisingregister/core/dto"
)

func (repository repository) Create(ar *dto.AdvertisingRequest) (*domain.Advertising, error) {
	ctx := context.Background()

	category, err := strconv.Atoi(ar.Category)

	product := &domain.Advertising{
		ID:          domain.NewID(),
		Title:       ar.Title,
		Description: ar.Description,
		Category:    domain.CategoryType(category),
		Logo: domain.LogoPath{
			ID:        domain.NewID(),
			ImagePath: GetPath(),
		},
		// Banner: domain.Banner{
		// 	ID:        domain.NewID(),
		// 	ImagePath: GetPath(),
		// },
		Phone1:    ar.Phone1,
		Phone2:    ar.Phone2,
		Email:     ar.Email,
		Site:      ar.Site,
		WhatsApp:  ar.WhatsApp,
		Facebook:  ar.Facebook,
		Instagram: ar.Instagram,
		Twitter:   ar.Twitter,
		CreatedAt: time.Now(),
	}

	// placeholders := make([]string, len(product))

	// for i := range product {
	// 	placeholders[i] = fmt.Sprintf("$%d", i+1)
	// }

	// query := fmt.Sprintf(`INSERT INTO product (
	// 	id, name, description, category, logo, banner, phone1,
	// 	phone2, email, site, whatsapp, facebook, instagram, twitter, created_at
	// 	) VALUES (%s)
	// 	returning *`, strings.Join(placeholders, ","))

	err := repository.db.QueryRow(
		ctx,
		``,
	).Scan(
		&product.ID,
		&product.Title,
		&product.Price,
		&product.Description,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
