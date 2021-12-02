package model

// TokenPair used for returning the pair of id and refresh tokens
type TokenPair struct {
	IDToken      string `json: "idToken"`
	RefreshToken string `json: "refreshToken"`
}
