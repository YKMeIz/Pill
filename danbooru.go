package pill

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

// danbooru returns the parsed details of a post on danbooru based website.
func danbooru(url string) (DanbooruBase, error) {
	// Obtain illustration page.
	body, err := fetch(url, defaultBrowserHeaders)
	if err != nil {
		return DanbooruBase{}, err
	}

	// Select necessary json data.
	data := regexp.MustCompile(`Post.register_resp\({.+}\); </script>`).FindString(string(body))
	data = strings.TrimPrefix(data, `Post.register_resp(`)
	data = strings.TrimSuffix(data, `); </script>`)

	// Decode json data.
	var meta postRegisterResp
	if err := json.Unmarshal([]byte(data), &meta); err != nil {
		return DanbooruBase{}, err
	}

	// Convert tags string to string array.
	var tags []string
	for _, v := range strings.Split(meta.Posts[0].Tags, " ") {
		tags = append(tags, v)
	}

	return DanbooruBase{
		ID:        strconv.FormatInt(meta.Posts[0].ID, 10),
		Tags:      tags,
		CreatedAt: meta.Posts[0].CreatedAt,
		Source:    formatURL(meta.Posts[0].Source),
		MD5:       meta.Posts[0].MD5,
		File:      formatURL(meta.Posts[0].File),
		Sample:    formatURL(meta.Posts[0].Sample),
		Preview:   formatURL(meta.Posts[0].Preview),
	}, nil
}

func formatURL(url string) string {
	// Trim backslash in url.
	return strings.ReplaceAll(url, "\\/", "/")
}
