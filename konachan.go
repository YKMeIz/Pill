package pill

// Konachan returns the parsed details of a Konachan.com post.
func Konachan(id string) (KonachanInfo, error) {
	k, e := danbooru(konachanBaseURL + id)
	return KonachanInfo(k), e
}
