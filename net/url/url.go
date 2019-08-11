package url

import (
	"net/url"
)

func FromProto(u URL) url.URL {
	return url.URL{
		Scheme:     u.Scheme,
		Opaque:     u.Opaque,
		User:       userinfo(u.User),
		Host:       u.Host,
		Path:       u.Path,
		RawPath:    u.RawPath,
		ForceQuery: u.ForceQuery,
		RawQuery:   u.RawQuery,
		Fragment:   u.Fragment,
	}
}

func ToProto(u url.URL) URL {
	var user *UserInfo
	if u.User != nil {
		pass, pset := u.User.Password()
		user = &UserInfo{
			Username:    u.User.Username(),
			Password:    pass,
			PasswordSet: pset,
		}
	}

	return URL{
		Scheme:     u.Scheme,
		Opaque:     u.Opaque,
		User:       user,
		Host:       u.Host,
		Path:       u.Path,
		RawPath:    u.RawPath,
		ForceQuery: u.ForceQuery,
		RawQuery:   u.RawQuery,
		Fragment:   u.Fragment,
	}
}

func userinfo(user *UserInfo) *url.Userinfo {
	if user == nil {
		return nil
	}
	if user.Password != "" {
		return url.UserPassword(user.Username, user.Password)
	}

	return url.User(user.Username)
}
