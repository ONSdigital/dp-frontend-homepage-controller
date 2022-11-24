package model

import (
	"github.com/ONSdigital/dp-renderer/model"
)

// Page contains data re-used for each page type a Data struct for data specific to the page type
type CensusPage struct {
	model.Page
	Data Census `json:"data"`
}

// Census is data for the census hub
type Census struct {
	AvailableTopics   []Topics `json:"available_topics"`
	UnavailableTopics []Topics `json:"unavailable_topics"`
}

// Topics is the data for topics
type Topics struct {
	Topic       string `json:"topic"`
	ReleaseDate string `json:"release_date"`
	URL         string `json:"url"`
}
