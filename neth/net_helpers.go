package neth

import "net/url"

type Neth struct{}

func (n Neth) MaskDSN(dsn string) string {
	u, err := url.Parse(dsn)
	if err != nil {
		return "[invalid DSN]"
	}

	if u.User != nil {
		username := u.User.Username()
		u.User = url.UserPassword(username, "****")
	}
	return u.String()
}
