package pkg //nolint
import (
	"log"
	"net/http"
	"time"
)

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	// headerMap http.Header
	status      int
	wroteHeader bool
	size        int
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b) // write response using original http.ResponseWriter
	rw.size += size                         // capture size
	return size, err
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, status: http.StatusOK}
}

type handler struct {
	next http.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	res := wrapResponseWriter(w)
	h.next.ServeHTTP(res, req)

	cl := req.Header.Get("Content-Length")
	if cl == "" {
		cl = "0"
	}
	myHeader := req.Header.Get("my-header")

	log.Println("my-header", myHeader)
	log.Println("trace-id", res.ResponseWriter.Header().Get("x-trace-id"))
	log.Println("method", req.Method)
	log.Println("uri", req.RequestURI)
	log.Println("status", res.status)
	log.Println("latency", time.Since(start))
	log.Println("bytes_in", cl)
	log.Println("bytes_out", res.size)
	log.Println("user-agent", req.UserAgent())
}

func WithLoggingMiddleware(h http.Handler) http.Handler {
	return &handler{next: h}
}
