package problems

type Problem struct {
	Id             string  `json:"id"`
	Message        string  `json:"message"`
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"httpStatusCode"`
}

func (p Problem) Error() string {
	return p.Message
}
