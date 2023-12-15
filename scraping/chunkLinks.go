package scraping

// The function `chunkLinks` takes a slice of strings and a chunk size, and returns a 2D slice of
// strings where each inner slice contains a chunk of the original slice.
func chunkLinks(links []string, chunkSize int) [][]string {
	var chunks [][]string = make([][]string, 0, (len(links)+chunkSize-1)/chunkSize)
	for i := 0; i < len(links); i += chunkSize {
		end := i + chunkSize
		if end > len(links) {
			end = len(links)
		}
		chunk := make([]string, 0, chunkSize)
		chunk = append(chunk, links[i:end]...)
		chunks = append(chunks, chunk)
	}
	return chunks
}
