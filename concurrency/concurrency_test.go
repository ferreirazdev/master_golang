package concurrency

import (
	"reflect"
	"testing"
)

// Channels
// Canais são uma estrutura de dados em Go que pode receber e enviar valores.
// Essas operações, junto de seus detalhes, permitem a comunicação entre processos diferentes.

func mockVerifyWebSite(url string) bool {
	// if url == "waat://furhurterwe.geds" {
	// 	return false
	// }

	// return true
	return url != "waat://furhurterwe.geds"
}

func TestVerifyWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expect := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	result := VerifyWebsites(mockVerifyWebSite, websites)

	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("expected '%v', result '%v'", result, expect)
	}
}
