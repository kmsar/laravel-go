package Response

import (
	"bufio"
	"net"
	"net/http"
	"next-doc/app/application/http-error"
)

type Response struct {
	httpRequest    *http.Request //used for redirect
	beforeFuncs    []func()
	afterFuncs     []func()
	ResponseWriter http.ResponseWriter
	status         int
	size           int64
	Committed    bool
	CookieConfig CookieConfig
}

// NewResponse creates a new instance of Response.
func NewResponse(req *http.Request, w http.ResponseWriter, cookieConfig CookieConfig) (r *Response) {
	return &Response{httpRequest: req, ResponseWriter: w, CookieConfig: cookieConfig}
}

// Before registers a function which is called just before the Response is written.
func (res *Response) Before(fn func()) {
	res.beforeFuncs = append(res.beforeFuncs, fn)
}

// After registers a function which is called just after the Response is written.
// If the `Content-Length` is unknown, none of the after function is executed.
func (res *Response) After(fn func()) {
	res.afterFuncs = append(res.afterFuncs, fn)
}

// WriteHeader sends an HTTP Response header with status code. If WriteHeader is
// not called explicitly, the first call to Write will trigger an implicit
// WriteHeader(http.StatusOK). Thus explicit calls to WriteHeader are mainly
// used to send error codes.
func (res *Response) WriteHeader(code int) {
	if res.Committed {
		return
	}
	res.status = code
	for _, fn := range res.beforeFuncs {
		fn()
	}
	res.ResponseWriter.WriteHeader(res.status)
	res.Committed = true
}

// Write writes the data to the connection as part of an HTTP reply.
func (res *Response) Write(b []byte) (n int, err error) {
	if !res.Committed {
		if res.status == 0 {
			res.status = http.StatusOK
		}
		res.WriteHeader(res.status)
	}
	n, err = res.ResponseWriter.Write(b)
	res.size += int64(n)
	for _, fn := range res.afterFuncs {
		fn()
	}
	return
}

// Hijack implements the http.Hijacker interface to allow an HTTP handler to
// take over the connection.
// See [http.Hijacker](https://golang.org/pkg/net/http/#Hijacker)
// Hijack hijacker for http
func (res *Response) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj, ok := res.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, http_error.NewHTTPError(500, "webserver doesn't support hijacking")
	}
	return hj.Hijack()
}

// Flush http.Flusher
// Flush implements the http.Flusher interface to allow an HTTP handler to flush
// buffered data to the client.
// See [http.Flusher](https://golang.org/pkg/net/http/#Flusher)
func (res *Response) Flush() {
	if f, ok := res.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// CloseNotify http.CloseNotifier
func (res *Response) CloseNotify() <-chan bool {
	if cn, ok := res.ResponseWriter.(http.CloseNotifier); ok {
		return cn.CloseNotify()
	}
	return nil
}

func (res *Response) Push(target string, opts *http.PushOptions) error {
	if pusher, ok := res.ResponseWriter.(http.Pusher); ok {
		return pusher.Push(target, opts)
	}
	return http.ErrNotSupported
}

func (res *Response) Reset(r *http.Request, w http.ResponseWriter) {
	res.beforeFuncs = nil
	res.afterFuncs = nil
	res.ResponseWriter = w
	res.httpRequest = r
	res.size = 0
	res.status = http.StatusOK
	res.Committed = false
}

// Status GetStatusCode
func (res *Response) Status() int {
	return res.status
}

func (res *Response) Size() int64 {
	return res.size
}

func (res *Response) SetStatus(code int) {
	res.status = code
}

func (res *Response) SetCookie(cookie *http.Cookie) {
	http.SetCookie(res.ResponseWriter, cookie)
}

// SetHeader is an intelligent shortcut for c.Writer.Header().Set(key, value).
// It writes a header in the Response.
// If value == "", this method removes the header `c.Writer.Header().Del(key)`
func (res *Response) SetHeader(key, value string) {
	if value == "" {
		res.Header().Del(key)
		return
	}
	res.Header().Set(key, value)
}

// Header returns the header map for the writer that will be sent by
// WriteHeader. Changing the header after a call to WriteHeader (or Write) has
// no effect unless the modified headers were declared as trailers by setting
// the "Trailer" header before the call to WriteHeader (see example)
// To suppress implicit Response headers, set their value to nil.
// Example: https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
func (res *Response) Header() http.Header {
	return res.ResponseWriter.Header()
}

//
//// Render a text template with the given data.
//// The template path is relative to the "resources/template" directory.
//func (res *Response) Render(responseCode int, templatePath string, data interface{}) error {
//	tmplt, err := template.ParseFiles(res.getTemplateDirectory() + templatePath)
//	if err != nil {
//		return err
//	}
//
//	var b bytes.Buffer
//	if err := tmplt.Execute(&b, data); err != nil {
//		return err
//	}
//
//	return res.String(responseCode, b.String())
//}
//
//// RenderHTML an HTML template with the given data.
//// The template path is relative to the "resources/template" directory.
//func (res *Response) RenderHTML(responseCode int, templatePath string, data interface{}) error {
//	tmplt, err := htmltemplate.ParseFiles(res.getTemplateDirectory() + templatePath)
//	if err != nil {
//		return err
//	}
//
//	var b bytes.Buffer
//	if err := tmplt.Execute(&b, data); err != nil {
//		return err
//	}
//
//	return res.String(responseCode, b.String())
//}
//
//func (res *Response) getTemplateDirectory() string {
//	sep := string(os.PathSeparator)
//	workingDir, err := os.Getwd()
//	if err != nil {
//		panic(err)
//	}
//	return workingDir + sep + "resources" + sep + "template" + sep
//}
