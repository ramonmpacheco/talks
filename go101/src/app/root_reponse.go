package app

type rootResponse struct {
	Name  string `json:"name"`
	Paths []path `json:"paths"`
}

type path struct {
	Name string `json:"name"`
	Desc string `json:"description"`
	Link string `json:"link"`
}

func NewRootResponse(name string) *rootResponse {
	return &rootResponse{Name: name}
}

func (r *rootResponse) appenPath(name, desc string, links []string) {
	for _, link := range links {
		r.Paths = append(r.Paths, path{Name: name, Desc: desc, Link: link})
	}
}
