package travel

import "strings"

// kind of journey
type journeyKind struct {
	Name       string
	PlaceTypes []string
}

func (j journeyKind) Public() interface{} {
	return map[string]interface{}{
		"name":    j.Name,
		"journey": strings.Join(j.PlaceTypes, "|"),
	}
}

var Journeys = []interface{}{
	journeyKind{Name: "Romantic", PlaceTypes: []string{"park", "bar", "movie_theatre", "restaurant", "florist"}},
	journeyKind{Name: "Shopping", PlaceTypes: []string{"departmental_store", "cafe", "clothing_store", "shoe_store", "jewelry_store"}},
	journeyKind{Name: "Night out", PlaceTypes: []string{"bar", "casino", "food", "bar", "night_club", "bar", "bar", "hospital"}},
	journeyKind{Name: "Culture", PlaceTypes: []string{"museum", "cafe", "cemetery", "library", "art_gallery"}},
	journeyKind{Name: "Pamper", PlaceTypes: []string{"hair_care", "beauty_salon", "cafe", "spa"}},
}
