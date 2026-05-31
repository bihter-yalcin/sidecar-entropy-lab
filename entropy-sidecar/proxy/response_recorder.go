package proxy

import(
	"bytes"
	"net/http"
)

type responseRecorder struct {
	http.ResponseWriter
	body *bytes.Buffer	
	statusCode int
}

func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{
		ResponseWriter: w,
        body: &bytes.Buffer{},
		statusCode: http.StatusOK, 
	}
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *responseRecorder) Write(body []byte) (int, error) {
	r.body.Write(body)
	return r.ResponseWriter.Write(body)
}
