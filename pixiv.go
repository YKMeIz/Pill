package pill

import (
	"encoding/json"
	"html"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Pixiv returns the parsed details of a Pixiv illustration work.
func Pixiv(id string) (PixivInfo, error) {

	// Obtain illustration page.
	body, err := fetch(illustBaseURL+id, defaultBrowserHeaders)
	if err != nil {
		return PixivInfo{}, err
	}

	// Select necessary json data.
	data := regexp.MustCompile(`id="meta-preload-data" content='{.+}'>`).FindString(string(body))
	data = strings.TrimPrefix(data, `id="meta-preload-data" content='`)
	data = strings.TrimSuffix(data, `'>`)
	data = regexp.MustCompile(`"[0-9]+":`).ReplaceAllString(data, `"ID":`)

	// Decode json data.
	var meta metaPreloadData
	if err := json.Unmarshal([]byte(data), &meta); err != nil {
		return PixivInfo{}, err
	}

	// Parse illustration creation date.
	t, err := time.Parse(time.RFC3339, meta.Illust.ID.CreateDate)
	if err != nil {
		return PixivInfo{}, err
	}

	// Extract tags information.
	var tags []string
	for _, v := range meta.Illust.ID.Tags.Tags {
		tags = append(tags, v.Tag)
	}

	return PixivInfo{
		Title: meta.Illust.ID.Title,
		ID:    meta.Illust.ID.Id,
		// Convert description to raw HTML.
		Description: decodeUnicode(meta.Illust.ID.Description),
		Tags:        tags,
		CreatedAt:   t.Unix(),
		// Obtain full list of original images if applicable.
		Sources: ping(meta.Illust.ID.Urls.Original),
		Author: PixivMember{
			ID:     meta.User.ID.UserID,
			Name:   meta.User.ID.Name,
			Avatar: meta.User.ID.ImageBig,
		},
	}, nil

}

// ping is utilized to guess right urls.
func ping(source string) []string {
	var (
		res []string
		p   int
	)
	// .png or .jpg
	format := source[len(source)-4:]
	// base url before _p0.png/jpg
	source = source[:len(source)-7]

	for {
		// Pixiv will return 403 forbidden if there is no "Refer" header entity.
		h := append(defaultBrowserHeaders, header{"Referer", illustBaseURL})

		if !headStatusOk(source+"_p"+strconv.Itoa(p)+format, h) {
			break
		}

		res = append(res, source+"_p"+strconv.Itoa(p)+format)

		p++
	}

	return res
}

// decode HTML generated unicode to UTF-8
func decodeUnicode(s string) string {
	s = `{"Unicode":"` + s + `"}`
	tmp := struct {
		Unicode string
	}{}

	if err := json.Unmarshal([]byte(s), &tmp); err != nil {
		return s
	}

	return html.UnescapeString(tmp.Unicode)
}
