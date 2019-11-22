package representation

type LinkEntityResponse struct {
	LinkID      string    `json:"id"`
	Name        string    `json:"name"`
	Address		string	  `json:"address"`
}

type LinkListResponse struct {
	Items []*LinkEntityResponse `json:"items"`
	Found int                   `json:"found"`
}
