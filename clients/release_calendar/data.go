package release_calendar

import "time"

type ReleaseCalendar struct {
	Type     string `json:"type, omitempty"`
	ListType string `json:"listType, omitempty"`
	URI      string `json:"uri, omitempty"`
	Result   Result `json:"result"`
}

type Result struct {
	NumberOfResults int           `json:"numberOfResults"`
	Took            int           `json:"took"`
	Results         *[]Results    `json:"results"`
	Suggestions     []interface{} `json:"suggestions"`
	DocCounts       struct{}      `json:"docCounts"`
	SortBy          string        `json:"sortBy"`
}

type Results struct {
	Type        string         `json:"type"`
	Description *Description   `json:"description"`
	SearchBoost *[]interface{} `json:"searchBoost"`
	URI         string         `json:"uri"`
}

type Description struct {
	Summary            string    `json:"summary"`
	NextRelease        string    `json:"nextRelease"`
	ReleaseDate        time.Time `json:"releaseDate"`
	Finalised          bool      `json:"finalised"`
	Source             string    `json:"source"`
	Published          bool      `json:"published"`
	Title              string    `json:"title"`
	NationalStatistic  bool      `json:"nationalStatistic"`
	Unit               string    `json:"unit"`
	Contact            *Contact  `json:"contact"`
	ProvisionalDate    string    `json:"provisionalDate"`
	Cancelled          bool      `json:"cancelled"`
	PreUnit            string    `json:"preUnit"`
	CancellationNotice []string  `json:"cancellationNotice"`
}

type Contact struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
}
