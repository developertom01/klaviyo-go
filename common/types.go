package common

type Links struct {
	Self     *string `json:"self,omitempty"`
	First    *string `json:"first,omitempty"`
	Last     *string `json:"last,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Next     *string `json:"next,omitempty"`
}
