package models

type ScrapedJob struct {
	Title string
	Description string
	City string
	State string
	Url string
	// todo: needs more attributes, need to state which ones are optional, etc.
}
