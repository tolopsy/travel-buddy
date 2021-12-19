package travel

type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name": P.Name,
		"icon": P.Icon,
		"photos": P.Photos,
		"vicinity": P.Visinity,
		"lat" : P.lat
		"lng":  P.lng
}

type googleResponse struct {
	Results []*Place `json:"results"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"Lat"`
	Lng float64 `json:"Lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}
