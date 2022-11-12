package entity

type URL struct {
	Created string `json:"created_at"`
	OriginalURL string `json:"original_url"`
	ShortenURL string `json:"shorten_url"`
	RedirectCount uint32 `json:"redirect_count"`
}