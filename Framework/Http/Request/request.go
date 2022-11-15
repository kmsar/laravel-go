package Request

import (
	request_file "github.com/kmsar/laravel-go/Framework/Http/Upload"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"next-doc/app/application/consts"
	"os"
	"strings"
)

type Request struct {
	httpRequest *http.Request
	User        interface{}
}

// SaveUploadedFile uploads the form file to specific dst.
func (req *Request) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func NewRequest(w *http.Request) (r *Request) {
	return &Request{httpRequest: w, User: nil}
}
func (req *Request) GetHttpRequest() *http.Request {
	return req.httpRequest
}
func (req *Request) SetHttpRequest(r *http.Request) {
	req.httpRequest = r
}

func (req *Request) ResolveContentType() string {
	contentType := req.httpRequest.Header.Get("Content-Type")
	if contentType == "" {
		return "text/html"
	}
	return strings.ToLower(strings.TrimSpace(strings.Split(contentType, ";")[0]))
}

// GetPath returns RawPath, if it's empty returns Path from URL
// Difference between RawPath and Path is:
//   - Path is where Request path is stored. Value is stored in decoded form: /%47%6f%2f becomes /Go/.
//   - RawPath is an optional field which only gets set if the default encoding is different from Path.
func (req *Request) GetPath(r *http.Request) string {
	path := r.URL.RawPath
	if path == "" {
		path = r.URL.Path
	}
	return path
}

// Path returns requested path.
//
// The path is valid until returning from RequestHandler.
func (req *Request) Path() string {
	return req.httpRequest.URL.Path
}

// IsAJAX returns if it is a ajax Request
func (req *Request) IsAJAX() bool {
	return strings.Contains(req.GetHeader(consts.HeaderXRequestedWith), "XMLHttpRequest")
}

// Url get Request url
func (req *Request) Url() string {
	return req.httpRequest.URL.String()
}

func (req *Request) Scheme() string {
	// Can't use `r.Request.URL.Scheme`
	// See: https://groups.google.com/forum/#!topic/golang-nuts/pMUkBlQBDF0
	if req.IsTLS() {
		return "https"
	}
	if scheme := req.httpRequest.Header.Get(consts.HeaderXForwardedProto); scheme != "" {
		return scheme
	}
	if scheme := req.httpRequest.Header.Get(consts.HeaderXForwardedProtocol); scheme != "" {
		return scheme
	}
	if ssl := req.httpRequest.Header.Get(consts.HeaderXForwardedSsl); ssl == "on" {
		return "https"
	}
	if scheme := req.httpRequest.Header.Get(consts.HeaderXUrlScheme); scheme != "" {
		return scheme
	}
	return "http"
}
func (req *Request) IsWebSocket() bool {
	upgrade := req.httpRequest.Header.Get(consts.HeaderUpgrade)
	return strings.EqualFold(upgrade, "websocket")
}

// IsTLS returns true if HTTP connection is TLS otherwise false.
func (req *Request) IsTLS() bool {
	return req.httpRequest.TLS != nil
}

// HasFile mimics FormFile method from `http.Request`
//
//	func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
func (req *Request) HasFile(name string) bool {
	_, _, err := req.httpRequest.FormFile(name)
	if err != nil {
		return false
	}
	return true
}

// RealIP ClientIP implements a best effort algorithm to return the real client IP, it parses
// X-Real-IP and X-Forwarded-For in order to work properly with reverse-proxies such us: nginx or haproxy.
// // RealIP returns the first ip from 'X-Forwarded-For' or 'X-Real-IP' header key
// // if not exists data, returns Request.RemoteAddr
// // fixed for #164
func (req *Request) RealIP() string {
	// Fall back to legacy behavior
	if ip := req.httpRequest.Header.Get(consts.HeaderXForwardedFor); ip != "" {
		i := strings.IndexAny(ip, ",")
		if i > 0 {
			return strings.TrimSpace(ip[:i])
		}
		return ip
	}
	if ip := req.httpRequest.Header.Get(consts.HeaderXRealIP); ip != "" {
		return ip
	}
	ra, _, _ := net.SplitHostPort(req.httpRequest.RemoteAddr)
	return ra
}

// IsAjax checks if the Request was made via AJAX,
// the XMLHttpRequest will usually be sent with a X-Requested-With HTTP header.
func (req *Request) IsAjax() bool {
	if req.httpRequest.Header.Get("X-Request-With") != "" {
		return true
	}
	return false
}

func (req *Request) Cookie(name string) (*http.Cookie, error) {
	return req.httpRequest.Cookie(name)
}

func (req *Request) Reset(request *http.Request) {
	req.httpRequest = request
}

// GetHeader returns value from Request headers.
func (req *Request) GetHeader(key string) string {
	return req.httpRequest.Header.Get(key)
}

// IsWebsocket returns true if the Request headers indicate that a websocket
// handshake is being initiated by the client.
func (req *Request) IsWebsocket() bool {
	if strings.Contains(strings.ToLower(req.GetHeader("Connection")), "upgrade") &&
		strings.EqualFold(req.GetHeader("Upgrade"), "websocket") {
		return true
	}
	return false
}

// Method specifies the HTTP method (GET, POST, PUT, etc.).
func (req *Request) Method() string {
	return req.httpRequest.Method
}

// Protocol the protocol used by this Request, "HTTP/1.1" for example.
func (req *Request) Protocol() string {
	return req.httpRequest.Proto
}

// URI specifies the URI being requested.
// Use this if you absolutely need the raw query params, url, etc.
// Otherwise use the provided methods and fields of the "goyave.Request".
func (req *Request) URI() *url.URL {
	return req.httpRequest.URL
}

// Header contains the Request header fields either received
// by the server or to be sent by the client.
// Header names are case-insensitive.
//
// If the raw Request has the following header lines,
//
//	Host: example.com
//	accept-encoding: gzip, deflate
//	Accept-Language: en-us
//	fOO: Bar
//	foo: two
//
// then the header map will look like this:
//
//	Header = map[string][]string{
//		"Accept-Encoding": {"gzip, deflate"},
//		"Accept-Language": {"en-us"},
//		"Foo": {"Bar", "two"},
//	}
func (req *Request) Header() http.Header {
	return req.httpRequest.Header
}

// ContentLength records the length of the associated content.
// The value -1 indicates that the length is unknown.
func (req *Request) ContentLength() int64 {
	return req.httpRequest.ContentLength
}

// Referrer returns the referring URL, if sent in the Request.
func (req *Request) Referrer() string {
	return req.httpRequest.Referer()
}

// UserAgent returns the client's User-Agent, if sent in the Request.
func (req *Request) UserAgent() string {
	return req.httpRequest.UserAgent()
}

// RemoteAddress allows to record the network address that
// sent the Request, usually for logging.
func (req *Request) RemoteAddress() string {
	return req.httpRequest.RemoteAddr
}

// BearerToken extract the auth token from the "Authorization" header.
// Only takes tokens of type "Bearer".
// Returns empty string if no token found or the header is invalid.
func (req *Request) BearerToken() (string, bool) {
	const schema = "Bearer "
	header := req.Header().Get("Authorization")
	if !strings.HasPrefix(header, schema) {
		return "", false
	}
	return strings.TrimSpace(header[len(schema):]), true
}

func (req *Request) Cookies() []*http.Cookie {
	return req.httpRequest.Cookies()
}

// QueryStrings parses RawQuery and returns the corresponding values.
func (req *Request) QueryStrings() url.Values {
	return req.httpRequest.URL.Query()
}

// RawQuery returns the original query string
func (req *Request) RawQuery() string {
	return req.httpRequest.URL.RawQuery
}

// QueryString returns the first value associated with the given key.
func (req *Request) QueryString(key string) string {
	return req.QueryStrings().Get(key)
}

// ExistsQueryKey check is exists from query params with the given key.
func (req *Request) ExistsQueryKey(key string) bool {
	_, isExists := req.QueryStrings()[key]
	return isExists
}

// FormFile get file by form key
func (req *Request) FormFile(key string) (*request_file.UploadFile, error) {
	file, header, err := req.httpRequest.FormFile(key)
	if err != nil {
		return nil, err
	} else {
		return request_file.NewUploadFile(file, header), nil
	}
}

// FormFiles get multi files
// fixed #92
func (req *Request) FormFiles() (map[string]*request_file.UploadFile, error) {
	files := make(map[string]*request_file.UploadFile)

	if req.httpRequest.Form == nil {
		req.httpRequest.ParseForm()
	}

	if req.httpRequest.MultipartForm == nil || req.httpRequest.MultipartForm.File == nil {
		return nil, http.ErrMissingFile
	}
	for key, fileMap := range req.httpRequest.MultipartForm.File {
		if len(fileMap) > 0 {
			file, err := fileMap[0].Open()
			if err == nil {
				files[key] = request_file.NewUploadFile(file, fileMap[0])
			}
		}
	}
	return files, nil
}

func (req *Request) FormValue(name string) string {
	return req.httpRequest.FormValue(name)
}

func (req *Request) FormParams() (url.Values, error) {
	if strings.HasPrefix(req.httpRequest.Header.Get(consts.HeaderContentType), consts.MIMEMultipartForm) {
		if err := req.httpRequest.ParseMultipartForm(consts.DefaultMemory); err != nil {
			return nil, err
		}
	} else {
		if err := req.httpRequest.ParseForm(); err != nil {
			return nil, err
		}
	}
	return req.httpRequest.Form, nil
}

// HasInput checks for the existence of the given
// input name in the inputs sent from a FORM
func (req *Request) HasInput(name string) bool {
	if req.httpRequest.Form == nil {
		req.ParseForm()
	}

	_, ok := req.httpRequest.Form[name]
	return ok
}

// FormValues including both the URL field's query parameters and the POST or PUT form data
func (req *Request) FormValues() map[string][]string {
	req.ParseForm()
	return map[string][]string(req.httpRequest.Form)
}

// PostValues contains the parsed form data from POST, PATCH, or PUT body parameters
func (req *Request) PostValues() map[string][]string {
	req.ParseForm()
	return map[string][]string(req.httpRequest.PostForm)
}

func (req *Request) ParseForm() error {
	if strings.HasPrefix(req.GetHeader(consts.HeaderContentType), consts.MIMEMultipartForm) {
		if err := req.httpRequest.ParseMultipartForm(consts.DefaultMemory); err != nil {
			return err
		}
	} else {
		if err := req.httpRequest.ParseForm(); err != nil {
			return err
		}
	}
	return nil
}

// PostString returns the first value for the named component of the POST
// or PUT Request body. URL query parameters are ignored.
// Deprecated: Use the PostFormValue instead
func (req *Request) PostString(key string) string {
	return req.httpRequest.PostFormValue(key)
}

// RemoteIP RemoteAddr to an "IP" address
func (req *Request) RemoteIP() string {
	host, _, _ := net.SplitHostPort(req.httpRequest.RemoteAddr)
	return host
}

// FullRemoteIP RemoteAddr to an "IP:port" address
func (req *Request) FullRemoteIP() string {
	fullIp := req.httpRequest.RemoteAddr
	return fullIp
}

//// Route returns the current route.
//func (req *Request) Route() *application.Route {
//	return req.route
//}
//
//// Has check if the given field exists in the Request data.
//func (req *Request) Has(field string) bool {
//	_, exists := req.Data[field]
//	return exists
//}
//
//// String get a string field from the Request data.
//// Panics if the field is not a string.
//func (req *Request) String(field string) string {
//	str, ok := req.Data[field].(string)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not a string", field))
//	}
//	return str
//}
//
//// Numeric get a numeric field from the Request data.
//// Panics if the field is not numeric.
//func (req *Request) Numeric(field string) float64 {
//	str, ok := req.Data[field].(float64)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not numeric", field))
//	}
//	return str
//}
//
//// Integer get an integer field from the Request data.
//// Panics if the field is not an integer.
//func (req *Request) Integer(field string) int {
//	str, ok := req.Data[field].(int)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not an integer", field))
//	}
//	return str
//}
//
//// Bool get a bool field from the Request data.
//// Panics if the field is not a bool.
//func (req *Request) Bool(field string) bool {
//	str, ok := req.Data[field].(bool)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not a bool", field))
//	}
//	return str
//}
//
//// File get a file field from the Request data.
//// Panics if the field is not numeric.
//func (req *Request) File(field string) []request_file.File {
//	str, ok := req.Data[field].([]request_file.File)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not a file", field))
//	}
//	return str
//}
//
//// Timezone get a timezone field from the Request data.
//// Panics if the field is not a timezone.
//func (req *Request) Timezone(field string) *time.Location {
//	str, ok := req.Data[field].(*time.Location)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not a timezone", field))
//	}
//	return str
//}
//
//// IP get an IP field from the Request data.
//// Panics if the field is not an IP.
//func (req *Request) IP(field string) net.IP {
//	str, ok := req.Data[field].(net.IP)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not an IP", field))
//	}
//	return str
//}
//
//// URL get a URL field from the Request data.
//// Panics if the field is not a URL.
//func (req *Request) URL(field string) *url.URL {
//	str, ok := req.Data[field].(*url.URL)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not a URL", field))
//	}
//	return str
//}
//
//// UUID get a UUID field from the Request data.
//// Panics if the field is not a UUID.
//func (req *Request) UUID(field string) uuid.UUID {
//	str, ok := req.Data[field].(uuid.UUID)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not an UUID", field))
//	}
//	return str
//}
//
//// Date get a date field from the Request data.
//// Panics if the field is not a date.
//func (req *Request) Date(field string) time.Time {
//	str, ok := req.Data[field].(time.Time)
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not a date", field))
//	}
//	return str
//}
//
//// Object get an object field from the Request data.
//// Panics if the field is not an object.
//func (req *Request) Object(field string) map[string]interface{} {
//	str, ok := req.Data[field].(map[string]interface{})
//	if !ok {
//		panic(fmt.Sprintf("Field \"%s\" is not an object", field))
//	}
//	return str
//}
//
//// ToStruct map the Request data to a struct.
////
////	 type UserInsertRequest struct {
////		 Username string
////		 Email string
////	 }
////	 //...
////	 userInsertRequest := UserInsertRequest{}
////	 if err := Request.ToStruct(&userInsertRequest); err != nil {
////	  panic(err)
////	 }
//func (req *Request) ToStruct(dst interface{}) error {
//	return mergo.Map(dst, req.Data)
//}
//
//func (req *Request) validate() validation.Errors {
//	if req.Rules == nil {
//		return nil
//	}
//
//	extra := map[string]interface{}{
//		"Request": req,
//	}
//	contentType := req.httpRequest.Header.Get("Content-Type")
//	return validation.ValidateWithExtra(req.Data, req.Rules, strings.HasPrefix(contentType, "application/json"), req.Lang, extra)
//}

//RemoteAddress
//ContentLength
//Referrer
//UserAgent
//BearerToken
//Has
//String
//Numeric
//Integer
//Bool
//File
//Timezone
//Date
//IP
//URL
//Object
//Data map[string]interface{}  name=John%20Doe&tags=tag1&tags=tag2
//tags, _ := Request.Data["tags"].([]string)
//Params
///categories/{category}/{product_id}
///categories/3/5
//fmt.Println(Request.Params["category"]) // "3"
//fmt.Println(Request.Params["product_id"]) // "5"

//User is an interface{}

//Only
