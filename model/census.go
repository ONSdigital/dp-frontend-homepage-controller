package model

import (
	"github.com/ONSdigital/dp-renderer/model"
)

// Page contains data re-used for each page type a Data struct for data specific to the page type
type CensusPage struct {
	model.Page
	Data Census `json:"data"`
}

// Census contains data specific to the census hub page
type Census struct {
	EnableCensusTopicSubsection bool `json:"enable_census_topic_subsection`
}
