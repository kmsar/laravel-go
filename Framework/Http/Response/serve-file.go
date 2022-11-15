package Response

import (
	"fmt"
	request_file "github.com/kmsar/laravel-go/Framework/Http/Upload"
	"io"
	"net/http"
	"net/url"
	"next-doc/app/application/consts"
	"next-doc/app/application/http-error"
	"next-doc/packages/str"
	"os"
	"path/filepath"
	"strconv"
)

// File writes the specified file into the body stream in an efficient way.
func (res *Response) File(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return http_error.NewHTTPError(404, "file not found")
	}
	defer f.Close()

	fi, _ := f.Stat()
	if fi.IsDir() {
		file = filepath.Join(file, consts.IndexPage)
		f, err = os.Open(file)
		if err != nil {
			return http_error.NewHTTPError(404, "file not found")
		}
		defer f.Close()
		if fi, err = f.Stat(); err != nil {
			return
		}
	}
	http.ServeContent(res.ResponseWriter, res.httpRequest, fi.Name(), fi.ModTime(), f)
	return
}

// FileFromFS writes the specified file from http.FileSystem into the body stream in an efficient way.
func (res *Response) FileFromFS(filepath string, fs http.FileSystem) {
	defer func(old string) {
		res.httpRequest.URL.Path = old
	}(res.httpRequest.URL.Path)

	res.httpRequest.URL.Path = filepath

	http.FileServer(fs).ServeHTTP(res.ResponseWriter, res.httpRequest)
}

// FileAttachment writes the specified file into the body stream in an efficient way
// On the client side, the file will typically be downloaded with the given filename
func (res *Response) FileAttachment(filepath, filename string) {
	if str.IsASCII(filename) {
		res.
			ResponseWriter.
			Header().
			Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	} else {
		res.
			ResponseWriter.
			Header().
			Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(filename))
	}
	http.ServeFile(
		res.ResponseWriter,
		res.httpRequest,
		filepath)
}

//fmt.Println(Response.GetError()) // "panic: something wrong happened"
//Return true if nothing has been written to the Response body yet.
//Response.IsHeaderWritten
//fmt.Println(Response.IsHeaderWritten()) // false
//Response.Redirect("/login") 308
//Response.TemporaryRedirect("/maintenance") 307
//Render

func (res *Response) writeFile(file string, disposition string) (int64, error) {
	if !request_file.FileExists(file) {
		res.SetStatus(http.StatusNotFound)
		return 0, &os.PathError{Op: "open", Path: file, Err: fmt.Errorf("no such file or directory")}
	}

	res.status = http.StatusOK
	mime, size := request_file.GetMIMEType(file)
	header := res.ResponseWriter.Header()
	header.Set("Content-Disposition", disposition)

	if header.Get("Content-Type") == "" {
		header.Set("Content-Type", mime)
	}

	header.Set("Content-Length", strconv.FormatInt(size, 10))

	f, _ := os.Open(file)
	// No need to check for errors, fsutil.FileExists(file) and
	// fsutil.GetMIMEType(file) already handled that.
	defer f.Close()
	return io.Copy(res, f)
}

// FileInline File write a file as an inline element.
// Automatically detects the file MIME type and sets the "Content-Type" header accordingly.
// If the file doesn't exist, respond with status 404 Not Found.
// The given path can be relative or absolute.
//
// If you want the file to be sent as a download ("Content-Disposition: attachment"), use the "Download" function instead.
func (res *Response) FileInline(file string) error {
	_, err := res.writeFile(file, "inline")
	return err
}

// Download write a file as an attachment element.
// Automatically detects the file MIME type and sets the "Content-Type" header accordingly.
// If the file doesn't exist, respond with status 404 Not Found.
// The given path can be relative or absolute.
//
// The "fileName" parameter defines the name the client will see. In other words, it sets the header "Content-Disposition" to
// "attachment; filename="${fileName}""
//
// If you want the file to be sent as an inline element ("Content-Disposition: inline"), use the "File" function instead.
func (res *Response) Download(file string, fileName string) error {
	_, err := res.writeFile(file, fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	return err
}
