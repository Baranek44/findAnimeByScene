package cmd

type Response struct {
	FrameCount string `json:"framecount"`
	Error      string `json:"error"`
	Result     []struct {
		Anilist    Anilist `json:"anilist"`
		FileName   string  `json:"filename"`
		Episode    int     `json:"episode"`
		From       int     `json:"from"`
		To         int     `json:"to"`
		Similariti int     `json:"smililariti"`
		Video      string  `json:"video"`
		Image      string  `json:"image"`
	}
}

type Anilist struct {
	ID       int    `json:"id"`
	IDMal    int    `json:"idmal"`
	Title    Title  `json:"title"`
	Synonyms string `json:"synonyms"`
	IsAdlut  bool   `json:"isadlut"`
}

type Title struct {
	Native  string `json:"native"`
	Romaji  string `json:"romaji"`
	English string `json:"english"`
}
