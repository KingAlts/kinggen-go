package kinggen

type KingGen struct {
	apiKey string
}

type Alt struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Profile struct {
	Username       string `json:"username"`
	Generated      int    `json:"generated"`
	GeneratedToday int    `json:"generatedToday"`
	Stock          int    `json:"stock"`
}
