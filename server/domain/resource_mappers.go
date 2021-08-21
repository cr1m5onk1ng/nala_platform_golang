package domain

type MappedResource struct {
	Url string `json:"url"`
	// 2 chars language code
	Language     string `json:"language"`
	Difficulty   string `json:"difficulty"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	MediaType    string `json:"media_type"`
	Category     string `json:"category"`
	ThumbnailUrl string `json:"thumbnail_url"`
}
