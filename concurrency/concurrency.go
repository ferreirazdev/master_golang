package concurrency

import "time"

type verifyWebsite func(string) bool
type result struct {
	string
	bool
}

func VerifyWebsites(vw verifyWebsite, urls []string) map[string]bool {
	results := make(map[string]bool)
	chanResult := make(chan result)

	for _, url := range urls {
		go func(u string) {
			chanResult <- result{u, vw(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-chanResult
		results[result.string] = result.bool
	}

	time.Sleep(2 * time.Second)

	return results
}
