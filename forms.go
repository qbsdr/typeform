package typeform

import (
	"fmt"
	"net/url"
	"time"
)

// FormsResponse handles responses for /forms endpoint
type FormsResponse struct {
	TotalItems int     `json:"total_items"`
	PageCount  int     `json:"page_count"`
	Items      []*Item `json:"items"`
}

type Settings struct {
	IsPublic        bool   `json:"is_public"`
	IsTrial         bool   `json:"is_trial"`
	Language        string `json:"language"`
	ProgressBar     string `json:"progress_bar"`
	ShowProgressBar bool   `json:"show_progress_bar"`
}

type Href struct {
	Href string `json:"href"`
}

type Item struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	LastUpdatedAt string   `json:"last_updated_at"`
	Settings      Settings `json:"settings"`
	Self          Href     `json:"self"`
	Theme         Href     `json:"theme"`
	Links         struct {
		Display string `json:"display"`
	} `json:"_links"`
}

// Form handles response for /forms/<formID> endpoint
type Form struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Theme    Href     `json:"theme"`
	Settings Settings `json:"settings"`
	Fields   []Field  `json:"fields"`
}

type Field struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Ref         string      `json:"ref"`
	Type        string      `json:"type"`
	Properties  Properties  `json:"properties"`
	Validations Validations `json:"validations"`
}

type FieldChoice struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// FieldProperties define the properties available for all
// question types.
type Properties struct {
	Description string `json:"description"`
	Randomize   bool   `json:"randomize"`
	// true to allow respondents to select more than one answer choice.
	// false to allow respondents to select only one answer choice.
	// Available for multiple_choice and picture_choice types.
	AllowMultipleSelection bool `json:"allow_multiple_selection"`

	// true to include an "Other" option so respondents can enter a different answer choice from those listed.
	// false to limit answer choices to those listed. Available for multiple_choice and picture_choice types.
	AllowOtherChoice bool `json:"allow_other_choice"`

	// Answer choices. Available for dropdown, multiple_choice, and picture_choice types
	Choices []FieldChoice `json:"choices"`

	// true to list answer choices vertically. false to list answer choices horizontally. Available for multiple_choice types.
	VerticalAlignment bool   `json:"vertical_alignment"`
	AlphabeticalOrder bool   `json:"alphabetical_order"`
	Structure         string `json:"structure"`
}

type Validations struct {
	Required  bool `json:"required"`
	MaxLength int  `json:"max_length"`
	MinLength int  `json:"min_length"`
	MaxValue  int  `json:"max_value"`
	MinValue  int  `json:"min_value"`
}

type ResponsesResponse struct {
	TotalItems int            `json:"total_items"`
	PageCount  int            `json:"page_count"`
	Items      []ResponseItem `json:"items"`
}

type ResponseItem struct {
	LandingID   string   `json:"landing_id"`
	Token       string   `json:"token"`
	SubmittedAt string   `json:"submitted_at"`
	Metadata    Metadata `json:"metadata"`
	Answers     []Answer `json:"answers"`
}

type Metadata struct {
	UserAgent string `json:"user_agent"`
	Platform  string `json:"platform"`
	Referer   string `json:"referer"`
	NetworkID string `json:"network_id"`
	Browser   string `json:"browser"`
}

type Answer struct {
	Field AnswerField `json:"field"`
	Type  string      `json:"type"`
	// Shown only for choices type
	Choices AnswerChoices `json:"choices"`
	Choice  AnswerChoice  `json:"choice"`
	// Shown only for date type
	Date    string `json:"date"`
	Text    string `json:"text"`
	Email   string `json:"email"`
	Boolean bool   `json:"boolean"`
}

type AnswerChoices struct {
	Labels []string `json:"labels"`
	Label  string   `json:"label"`
	Other  string   `json:"other"`
}

type AnswerChoice struct {
	Label string `json:"label"`
}

type AnswerField struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Query struct {
	PageSize  string
	Since     time.Time
	Until     time.Time
	After     string
	Before    string
	Completed string
}

func (q *Query) Encode() string {
	v := url.Values{}
	if !q.Since.IsZero() {
		v.Set("since", fmt.Sprintf("%v", q.Since.Unix()))
	}

	if q.Completed != "" {
		v.Set("completed", q.Completed)
	}
	return v.Encode()
}
