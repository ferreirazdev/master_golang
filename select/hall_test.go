package hall

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Defer
// Utilizamos o defer quando queremos que a função seja executada no final de uma função
// mas mantendo essa instrução próxima de onde o servidor foi criado para facilitar a vida das pessoas que forem ler o código futuramente.

// Select
// O que o select te permite fazer é aguardar múltiplos canais. O primeiro a enviar um valor "vence" e o código abaixo do case é executado.

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
