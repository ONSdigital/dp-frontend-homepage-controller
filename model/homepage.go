package model

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-renderer/v2/model"
)

// Page contains data re-used for each page type a Data struct for data specific to the page type
type Page struct {
	model.Page
	Data Homepage `json:"data"`
}

type HomepageData struct {
	AroundONS       *[]Feature
	EmergencyBanner zebedee.EmergencyBanner
	FeaturedContent *[]Feature
	MainFigures     map[string]*MainFigure
	ReleaseCalendar *ReleaseCalendar
	ServiceMessage  string
}

// Homepage contains data specific to this page type
type Homepage struct {
	MainFigures           map[string]*MainFigure `json:"main_figures"`
	HasMainFigures        bool                   `json:"has_main_figures"`
	ReleaseCalendar       ReleaseCalendar        `json:"release_calendar"`
	Featured              []Feature              `json:"featured"`
	HasFeaturedContent    bool                   `json:"has_featured_content"`
	AroundONS             []Feature              `json:"arounds_ons"`
	FeedbackAPIURL        string                 `json:"feedback_api_url"`
	EnablePreviewSiteTile bool                   `json:"enable_preview_site_tile"`
	PreviewSiteURL        string                 `json:"preview_site_url"`
}

// ReleaseCalendar is data for release calendar block
type ReleaseCalendar struct {
	Releases                         []Release `json:"releases"`
	NumberOfReleases                 int       `json:"number_of_releases"`
	NumberOfOtherReleasesInSevenDays int       `json:"number_of_other_releases"`
}

// Release is the data for an individual release
type Release struct {
	Title       string `json:"title"`
	URI         string `json:"uri"`
	ReleaseDate string `json:"releaseDate"`
}

// MainFigure is the data for an individual timeseries
type MainFigure struct {
	ID               string     `json:"id"`
	Title            string     `json:"title"`
	Summary          string     `json:"summary"`
	Date             string     `json:"date"`
	Figure           string     `json:"figure"`
	Trend            Trend      `json:"trend"`
	ShowTrend        bool       `json:"show_trend"`
	TrendDescription string     `json:"trend_description"`
	Unit             string     `json:"unit"`
	FigureURIs       FigureURIs `json:"figure_uris"`
}

// FigureURIs struct contains URI's to related analysis and data
type FigureURIs struct {
	Analysis string `json:"analysis"`
	Data     string `json:"data"`
}

// Trend contains bool values about the trend of a key figure (up from last month, down from last month, etc.)
type Trend struct {
	IsUp       bool   `json:"is_up"`
	IsDown     bool   `json:"is_down"`
	IsFlat     bool   `json:"is_flat"`
	Difference string `json:"difference"`
	Period     string `json:"period"`
}

// Feature is data for linked content
type Feature struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URI         string `json:"uri"`
	ImageURL    string `json:"image_url"`
	ImageAlt    string `json:"image_alt"`
}
