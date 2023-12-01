package anime

type Anime struct {
	Title         string         `json:"title"`
	Type          Type           `json:"type"`
	TxtAlt        []string       `json:"txtAlt"`
	Genres        []string       `json:"genres"` // TODO: Create a Genres struct
	Synopsis      string         `json:"synopsis"`
	Cover         string         `json:"cover"`
	Status        Status         `json:"status"`
	RelatedAnimes []RelatedAnime `json:"related_animes"`
	Episodes      []Episode      `json:"episodes"`
}

type Type string

const (
	TV      Type = "Anime"
	OVA     Type = "OVA"
	Movie   Type = "Película"
	Special Type = "Especial"
)

type Status string

const (
	Finished Status = "Finalizado"
	Airing   Status = "En emisión"
	Coming   Status = "Proximamente"
)

type RelatedAnime struct {
	Title    string   `json:"title"`
	Relation Relation `json:"relation"`
}

type Relation string

const (
	Prequel   Relation = "Precuela"
	Secquel   Relation = "Secuela"
	Main      Relation = "Historia principal"
	Alternate Relation = "Historia paralela"
)

type Episode struct {
	Title  string        `json:"title"`
	Videos []VideoServer `json:"urls"`
}

type VideoServer struct {
	Server string `json:"server"`
	Title  string `json:"title"`
	Url    string `json:"url"`
}
