package hall

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHall(t *testing.T) {

	serverA := createServer(11 * time.Millisecond)
	serverB := createServer(12 * time.Millisecond)

	defer serverA.Close()
	defer serverB.Close()

	_, err := Hall(serverA.URL, serverB.URL)

	if err == nil {
		t.Error("expect an error, but don't")
	}
}

func createServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
