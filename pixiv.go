package pill

import (
	"encoding/json"
	"regexp"
	"strconv"
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
	content := regexp.MustCompile(`token:.+"storableTags"`).FindString(string(body))
	content += regexp.MustCompile(`user:.+"background"`).FindString(string(body))

	// Extract tags information.
	tags := regexp.MustCompile(`"tag":".+?"`).FindAllString(content, -1)
	for i := range tags {
		tags[i] = tags[i][7 : len(tags[i])-1]
		tags[i] = decodeUnicode(tags[i])
	}

	// Extract illustration title.
	title := regexp.MustCompile(`"title":".+?"`).FindString(content)
	title = decodeUnicode(title[9 : len(title)-1])

	// Extract illustration description.
	description := regexp.MustCompile(`"description":".+?"`).FindString(content)
	description = decodeUnicode(description[15 : len(description)-1])

	// Extract illustration creation date.
	createdAt := regexp.MustCompile(`"createDate":".+?"`).FindString(content)
	// Parse date time.
	t, err := time.Parse(time.RFC3339, createdAt[14:len(createdAt)-1])
	if err != nil {
		return PixivInfo{}, err
	}

	// Extra source url for first original image.
	sourceURL := regexp.MustCompile(`"original":".+?"`).FindString(content)
	// Obtain full list of original images if applicable.
	sources := ping(id, decodeUnicode(sourceURL[12:len(sourceURL)-1]))

	// Extract author ID.
	authorID := regexp.MustCompile(`"userId":".+?"`).FindString(content)
	authorID = authorID[10 : len(authorID)-1]

	// Extract author name.
	authorName := regexp.MustCompile(`"name":".+?"`).FindString(content)
	authorName = decodeUnicode(authorName[8 : len(authorName)-1])

	// Extract author avatar url.
	authorAvatar := regexp.MustCompile(`"imageBig":".+?"`).FindString(content)
	authorAvatar = decodeUnicode(authorAvatar[12 : len(authorAvatar)-1])

	return PixivInfo{
		Title:       title,
		ID:          id,
		Description: description,
		Tags:        tags,
		CreatedAt:   t.Unix(),
		Sources:     sources,
		Author: PixivMember{
			ID:     authorID,
			Name:   authorName,
			Avatar: authorAvatar,
		},
	}, nil

}

// ping is utilized to guess right urls.
func ping(id, source string) []string {
	var (
		res []string
		p   int
	)

	source = source[:len(source)-7]
	// .png or .jpg
	format := source[len(source)-4:]

	for {
		// Pixiv will return 403 forbidden if there is no "Refer" header entity.
		h := append(defaultBrowserHeaders, header{"Referer", illustBaseURL + id + "&page=" + strconv.Itoa(p)})

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
		// TODO handle err
		// log.Println(err)
		return ""
	}

	return tmp.Unicode
}
