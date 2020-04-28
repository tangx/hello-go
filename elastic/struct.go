// Generated by https://quicktype.io
//
// To change quicktype's target language, run command:
//
//   "Set quicktype target language"

package elastic

type ElasticSearchResult struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

type Hits struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type Hit struct {
	Index  string       `json:"_index"`
	Type   string       `json:"_type"`
	ID     string       `json:"_id"`
	Score  float64      `json:"_score"`
	Record DomainRecord `json:"_source"`
}

type DomainRecord struct {
	Domain      string `json:"domain"`
	Locked      bool   `json:"locked"`
	DeletedAt   string `json:"deleted_at"`
	Fdsn        string `json:"fdsn"`
	UpdatedAt   string `json:"updated_at"`
	Version     string `json:"@version"`
	Value       string `json:"value"`
	TTL         int    `json:"ttl"`
	ID          int    `json:"id"`
	Description string `json:"description"`
	Timestamp   string `json:"@timestamp"`
	Rr          string `json:"rr"`
	RecordID    string `json:"record_id"`
	CreatedAt   string `json:"created_at"`
	Status      string `json:"status"`
	Privider    string `json:"privider"`
	Type        string `json:"type"`
	Line        string `json:"line"`
	Priority    int    `json:"priority"`
}

type Total struct {
	Value    int    `json:"value"`
	Relation string `json:"relation"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}
