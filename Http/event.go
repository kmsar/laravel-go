package Http

type ServeClosed struct {
}

func (this *ServeClosed) Event() string {
	return "HTTP_SERVE_CLOSED"
}
