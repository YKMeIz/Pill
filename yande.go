package pill

// Yande returns the parsed details of a Yande.re post.
func Yande(id string) (YandeInfo, error) {
	y, e := danbooru(yandeBaseURL + id)
	return YandeInfo(y), e
}
