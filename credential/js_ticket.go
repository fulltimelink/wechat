package credential

// JsTicketHandle js ticket获取
type JsTicketHandle interface {
	// GetTicket 获取ticket
	GetTicket(accessToken string) (ticket string, err error)
}

type MustCacheJsTicketHandle interface {
	MustCacheTicket(accessToken string) (expireIn int64, err error)
}
