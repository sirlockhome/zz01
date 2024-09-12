package model

type UserAuth struct {
	AccessToken  string      `json:"__access_token__"`
	AccessExp    int         `json:"__access_exp__"`
	RefreshToken string      `json:"__refresh_token__"`
	RefreshExp   int         `json:"__refresh_exp__"`
	UserData     interface{} `json:"user"`
}
