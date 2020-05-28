package release_calendar

import (
	"context"
	"encoding/json"
	"fmt"
	healthcheck "github.com/ONSdigital/dp-api-clients-go/health"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	rchttp "github.com/ONSdigital/dp-rchttp"
	"github.com/ONSdigital/log.go/log"
	"io/ioutil"
	"net/http"
)

const service = "Babbage"

// Client represents a babbage client
type Client struct {
	cli rchttp.Clienter
	url string
}

// ErrInvalidBabbageResponse is returned when the babbage service does not respond with a status 200
type ErrInvalidBabbageResponse struct {
	responseCode int
}

// Error should be called by the user to print out the stringified version of the error
func (e ErrInvalidBabbageResponse) Error() string {
	return fmt.Sprintf("invalid response from babbage service - status %d", e.responseCode)
}

// Code returns the status code received from babbage if an error is returned
func (e ErrInvalidBabbageResponse) Code() int {
	return e.responseCode
}

// New creates a new instance of Client with a given babbage url
func New(babbageURL string) *Client {
	hcClient := healthcheck.NewClient(service, babbageURL)

	return &Client{
		cli: hcClient.Client,
		url: babbageURL,
	}
}

// Checker calls babbage health endpoint and returns a check object to the caller.
func (c *Client) Checker(ctx context.Context, check *health.CheckState) error {
	hcClient := healthcheck.Client{
		Client: c.cli,
		URL:    c.url,
		Name:   service,
	}

	return hcClient.Checker(ctx, check)
}

func (c *Client) GetReleaseCalendar(ctx context.Context, userAccessToken, fromDay, fromMonth, fromYear string) (rc ReleaseCalendar, err error) {
	reqURL := fmt.Sprintf(c.url+"/releasecalendar/data?fromDateDay=%s&fromDateMonth=%s&fromDateYear=%s)", fromDay, fromMonth, fromYear)
	resp, err := c.get(ctx, reqURL)
	if err != nil {
		return rc, err
	}
	defer closeResponseBody(ctx, resp)

	if resp.StatusCode != http.StatusOK {
		return rc, ErrInvalidBabbageResponse{resp.StatusCode}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &rc)
	return
}

func (c *Client) get(ctx context.Context, uri string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return c.cli.Do(ctx, req)
}

// closeResponseBody closes the response body and logs an error containing the context if unsuccessful
func closeResponseBody(ctx context.Context, resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		log.Event(ctx, "error closing http response body", log.ERROR, log.Error(err))
	}
}
