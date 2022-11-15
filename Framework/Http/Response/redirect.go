package Response

import (
	"net/http"
	"next-doc/app/application/consts"
)

// Redirect send a permanent redirect Response
func (res *Response) Redirect(url string) {
	http.Redirect(res, res.httpRequest, url, http.StatusPermanentRedirect)
}

// TemporaryRedirect send a temporary redirect Response
func (res *Response) TemporaryRedirect(url string) {
	http.Redirect(res, res.httpRequest, url, http.StatusTemporaryRedirect)
}

func (r *Response) Redirect1(code int, targetUrl string) error {
	r.Header().Set(consts.HeaderCacheControl, "no-cache")
	r.Header().Set(consts.HeaderLocation, targetUrl)
	r.ResponseWriter.WriteHeader(code)
	return nil
}
