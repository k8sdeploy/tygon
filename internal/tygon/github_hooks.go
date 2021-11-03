package tygon

type PingEventConfig struct {
	ContentType string `json:"content_type"`
	InsecureSSL int    `json:"insecure_ssl"`
	Secret      string `json:"secret"`
	URL         string `json:"url"`
}
