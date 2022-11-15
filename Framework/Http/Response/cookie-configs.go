package Response

import "net/http"

// CookieConfig SessionCookie contains the configuration settings for session-midelware cookies.
type CookieConfig struct {
	// Name sets the name of the session-midelware cookie. It should not contain
	// whitespace, commas, colons, semicolons, backslashes, the equals sign or
	// control characters as per RFC6265. The default cookie name is "session-midelware".
	// If your application uses two different sessions, you must make sure that
	// the cookie name for each is unique.
	Name string

	// Domain sets the 'Domain' attribute on the session-midelware cookie. By default
	// it will be set to the domain name that the cookie was issued from.
	Domain string

	// HttpOnly sets the 'HttpOnly' attribute on the session-midelware cookie. The
	// default value is true.
	HttpOnly bool

	// Path sets the 'Path' attribute on the session-midelware cookie. The default value
	// is "/". Passing the empty string "" will result in it being set to the
	// path that the cookie was issued from.
	Path string

	// Persist sets whether the session-midelware cookie should be persistent or not
	// (i.e. whether it should be retained after a user closes their browser).
	// The default value is true, which means that the session-midelware cookie will not
	// be destroyed when the user closes their browser and the appropriate
	// 'Expires' and 'MaxAge' values will be added to the session-midelware cookie. If you
	// want to only persist some sessions (rather than all of them), then set this
	// to false and call the RememberMe() method for the specific sessions that you
	// want to persist.
	Persist bool

	// SameSite controls the value of the 'SameSite' attribute on the session-midelware
	// cookie. By default this is set to 'SameSite=Lax'. If you want no SameSite
	// attribute or value in the session-midelware cookie then you should set this to 0.
	SameSite http.SameSite

	// Secure sets the 'Secure' attribute on the session-midelware cookie. The default
	// value is false. It's recommended that you set this to true and serve all
	// requests over HTTPS in production environments.
	// See https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Session_Management_Cheat_Sheet.md#transport-layer-security.
	Secure bool
}
