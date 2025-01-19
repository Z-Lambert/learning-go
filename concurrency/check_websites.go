package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// To avoid the memory access issues with closures, 
	// I've passed the url as an argument to the goroutine so that the url memory location is not shared between goroutines.
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}