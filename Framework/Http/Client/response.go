package Client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Response struct and methods
//_______________________________________________________________________

// Response struct holds Response values of executed Request.
type Response struct {
	Request     *Request
	RawResponse *http.Response

	body       []byte
	size       int64
	receivedAt time.Time
}

// Body method returns HTTP Response as []byte array for the executed Request.
//
// Note: `Response.Body` might be nil, if `Request.SetOutput` is used.
func (r *Response) Body() []byte {
	if r.RawResponse == nil {
		return []byte{}
	}
	return r.body
}

// Status method returns the HTTP status string for the executed Request.
//
//	Example: 200 OK
func (r *Response) Status() string {
	if r.RawResponse == nil {
		return ""
	}
	return r.RawResponse.Status
}

// StatusCode method returns the HTTP status code for the executed Request.
//
//	Example: 200
func (r *Response) StatusCode() int {
	if r.RawResponse == nil {
		return 0
	}
	return r.RawResponse.StatusCode
}

// Proto method returns the HTTP Response protocol used for the Request.
func (r *Response) Proto() string {
	if r.RawResponse == nil {
		return ""
	}
	return r.RawResponse.Proto
}

// Result method returns the Response value as an object if it has one
func (r *Response) Result() interface{} {
	return r.Request.Result
}

// Error method returns the error object if it has one
func (r *Response) Error() interface{} {
	return r.Request.Error
}

// Header method returns the Response headers
func (r *Response) Header() http.Header {
	if r.RawResponse == nil {
		return http.Header{}
	}
	return r.RawResponse.Header
}

// Cookies method to access all the Response cookies
func (r *Response) Cookies() []*http.Cookie {
	if r.RawResponse == nil {
		return make([]*http.Cookie, 0)
	}
	return r.RawResponse.Cookies()
}

// String method returns the body of the server Response as String.
func (r *Response) String() string {
	if r.body == nil {
		return ""
	}
	return strings.TrimSpace(string(r.body))
}

// Time method returns the time of HTTP Response time that from Request we sent and received a Request.
//
// See `Response.ReceivedAt` to know when client received Response and see `Response.Request.Time` to know
// when client sent a Request.
func (r *Response) Time() time.Duration {
	if r.Request.clientTrace != nil {
		return r.Request.TraceInfo().TotalTime
	}
	return r.receivedAt.Sub(r.Request.Time)
}

// ReceivedAt method returns when Response got received from server for the Request.
func (r *Response) ReceivedAt() time.Time {
	return r.receivedAt
}

// Size method returns the HTTP Response size in bytes. Ya, you can relay on HTTP `Content-Length` header,
// however it won't be good for chucked transfer/compressed Response. Since Resty calculates Response size
// at the client end. You will get actual size of the http Response.
func (r *Response) Size() int64 {
	return r.size
}

// RawBody method exposes the HTTP raw Response body. Use this method in-conjunction with `SetDoNotParseResponse`
// option otherwise you get an error as `read err: http: read on closed Response body`.
//
// Do not forget to close the body, otherwise you might get into connection leaks, no connection reuse.
// Basically you have taken over the control of Response parsing from `Resty`.
func (r *Response) RawBody() io.ReadCloser {
	if r.RawResponse == nil {
		return nil
	}
	return r.RawResponse.Body
}

// IsSuccess method returns true if HTTP status `code >= 200 and <= 299` otherwise false.
func (r *Response) IsSuccess() bool {
	return r.StatusCode() > 199 && r.StatusCode() < 300
}

// IsError method returns true if HTTP status `code >= 400` otherwise false.
func (r *Response) IsError() bool {
	return r.StatusCode() > 399
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Response Unexported methods
//_______________________________________________________________________

func (r *Response) setReceivedAt() {
	r.receivedAt = time.Now()
	if r.Request.clientTrace != nil {
		r.Request.clientTrace.endTime = r.receivedAt
	}
}

func (r *Response) fmtBodyString(sl int64) string {
	if r.body != nil {
		if int64(len(r.body)) > sl {
			return fmt.Sprintf("***** RESPONSE TOO LARGE (size - %d) *****", len(r.body))
		}
		ct := r.Header().Get(hdrContentTypeKey)
		if IsJSONType(ct) {
			out := acquireBuffer()
			defer releaseBuffer(out)
			err := json.Indent(out, r.body, "", "   ")
			if err != nil {
				return fmt.Sprintf("*** Error: Unable to format Response body - \"%s\" ***\n\nLog Body as-is:\n%s", err, r.String())
			}
			return out.String()
		}
		return r.String()
	}

	return "***** NO CONTENT *****"
}
