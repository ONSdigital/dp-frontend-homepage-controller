package model

import (
	"github.com/ONSdigital/dp-renderer/v2/model"
)

// Page contains data re-used for each page type a Data struct for data specific to the page type
type CensusPage struct {
	model.Page
	Data Census `json:"data"`
}

// Census is data for the census hub
type Census struct {
	AvailableTopics             []Topics `json:"available_topics"`
	UnavailableTopics           []Topics `json:"unavailable_topics"`
	EnableCensusTopicSubsection bool     `json:"enable_census_topic_subsection"`
	EnableGetDataCard           bool     `json:"enable_get_data_card"`
	EnableCustomDataset         bool     `json:"enable_custom_dataset"`
	DatasetFinderEnabled        bool     `json:"dataset_finder_enabled"`
	CensusSearchTopicID         string   `json:"census_search_topic_id"`
	GetCensusDataURLQuery       string   `json:"get_census_data_url_query"`
	EnableFeedbackAPI           bool     `json:"enable_feedback_api"`
	FeedbackAPIURL              string   `json:"feedback_api_url"`
}

// Topics is the data for topics
type Topics struct {
	Topic       string `json:"topic"`
	ReleaseDate string `json:"release_date"`
	URL         string `json:"url"`
	ID          string `json:"id"`
}
