package cmd

type Response struct {
	FrameCount string `json:"framecount"`
	Error      string `json:"error"`
	Result     []struct {
		Anilist    string `json:"anilist"`
		FileName   string `json:"filename"`
		Episode    int    `json:"episode"`
		From       int    `json:"from"`
		To         int    `json:"to"`
		Similariti int    `json:"smililariti"`
		Video      string `json:"video"`
		Image      string `json:"image"`
	}
}
