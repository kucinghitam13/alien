package baak

type JadkulList struct {
	Total   int      `json:"total,omitempty"`
	Jadkuls []Jadkul `json:"jadkuls,omitempty"`
}

type Jadkul struct {
	Class    string `json:"class,omitempty"`
	Day      string `json:"day,omitempty"`
	Matkul   string `json:"matkul,omitempty"`
	Period   string `json:"period,omitempty"`
	Room     string `json:"room,omitempty"`
	Lecturer string `json:"lecturer,omitempty"`
}
