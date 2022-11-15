package Request

import "strings"

func (req *Request) HasJson() bool {
	header := req.GetHeader("Content-Type")
	for _, jsonHeader := range []string{"/json", "+json"} {
		if strings.Contains(header, jsonHeader) {
			return true
		}
	}

	return false
}

func (req *Request) HasMultiPartFormData() bool {
	header := req.GetHeader("Content-Type")
	for _, jsonHeader := range []string{"multipart/form-data", "application/x-www-form-urlencoded"} {
		if strings.Contains(header, jsonHeader) {
			return true
		}
	}
	return false
}
