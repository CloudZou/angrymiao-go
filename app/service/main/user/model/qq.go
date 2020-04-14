package model


type QQErrorInfo struct {
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
type QQTokenResponse struct {
	QQErrorInfo
	Msg string `json:"msg"`
	AccessToken string `json::"access_token"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
type QQOpenIDResponse struct {
	QQErrorInfo
	UnionID string `json:"unionid"`
	OpenID string `json:"openid"`
}

type QQUserInfoResponse struct {
	Ret int `json:"ret"`
	Msg string `json:"msg"`
	Nickname string `json:"nickname"`
	FigureUrl string	`json:"figureurl"`
	FigureUrl_1 string	`json:"figureurl_1"`
	FigureUrl_2 string	`json:"figureurl_2"`
	FigureUrl_qq_1 string	`json:"figure_qq_1"`
	FigureUrl_qq_2 string	`json:"figure_qq_2"`
	Gender string	`json:"gender"`
}

