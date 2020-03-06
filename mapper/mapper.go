package mapper

import (
	"context"

	"github.com/ONSdigital/dp-frontend-models/model/homepage"
)

// Homepage maps data to our homepage frontend model
func Homepage(ctx context.Context) homepage.Page {
	var page homepage.Page
	page.Type = "homepage"
	return page
}
