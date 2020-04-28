package homepage

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-api-clients-go/renderer"
	"github.com/ONSdigital/dp-api-clients-go/zebedee"
	"github.com/ONSdigital/dp-frontend-homepage-controller/mapper"
	"github.com/ONSdigital/log.go/log"
)

// Handler handles requests to homepage endpoint
func Handler(rend renderer.Renderer, zcli *zebedee.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handle(w, req, rend, zcli)
	}
}

func handle(w http.ResponseWriter, req *http.Request, rend renderer.Renderer, zcli *zebedee.Client) {
	ctx := req.Context()

	m := mapper.Homepage(ctx)

	b, err := json.Marshal(m)
	if err != nil {
		log.Event(ctx, "error marsahlling body data to json", log.Error(err))
		http.Error(w, "error marsahlling body data to json", http.StatusBadRequest)
		return
	}

	templateHTML, err := rend.Do("homepage", b)
	if err != nil {
		log.Event(ctx, "error rendering page", log.Error(err))
		http.Error(w, "error rendering page", http.StatusInternalServerError)
		return
	}

	w.Write(templateHTML)
	return
}
