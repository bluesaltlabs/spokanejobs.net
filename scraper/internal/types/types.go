package types


type ScrapedJob struct {
	JobId 			string `json:"job_id"`
	Title 			string `json:"title"`
	Description	string `json:"description"`
	City 				string `json:"city"`
	State 			string `json:"state"`
	Url 				string `json:"url"`
	// todo: needs more attributes, need to state which ones are optional, etc.
}
