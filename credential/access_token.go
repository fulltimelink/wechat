package credential

// AccessTokenHandle AccessToken 接口
type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
}

type MustCacheHandle interface {
	MustCacheAccessToken() (expireIn int64, err error)
}
