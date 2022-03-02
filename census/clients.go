package census

//go:generate moq -out mocks_test.go -pkg census . RenderClient

import (
	"io"

	"github.com/ONSdigital/dp-renderer/model"
)

// RenderClient is an interface with methods required for rendering a template from a page model
type RenderClient interface {
	BuildPage(w io.Writer, pageModel interface{}, templateName string)
	NewBasePageModel() model.Page
}

// Clients contains all the required Clients for frontend homepage controller
type Clients struct {
	Renderer RenderClient
}