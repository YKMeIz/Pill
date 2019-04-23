package pill

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Pixiv returns the parsed details of a Pixiv illustration work.
func Pixiv(id string) (PixivInfo, error) {

	// Obtain illustration page.
	content, err := fetch(illustBaseURL+id, defaultBrowserHeaders)
	if err != nil {
		return PixivInfo{}, err
	}

	// Prepare regular expression for extract each entity of details.
	tag := regexp.MustCompile(`<meta name="keywords" content=".+">`).FindString(string(content))
	sources := regexp.MustCompile(`<img src="https://i.pximg.net/[a-zA-Z0-9/\-_.]+"`).FindAllString(string(content), -1)
	description := regexp.MustCompile(`<meta name="description" content=".+">`).FindString(string(content))
	member := regexp.MustCompile(`<a href="member.php\?id=[0-9]+">.+</a></h2><span class="date">`).FindString(string(content))

	author := PixivMember{}

	// Parse Pixiv member ID.
	author.ID = regexp.MustCompile(`member.php\?id=[0-9]+`).FindString(member)
	author.ID = strings.TrimPrefix(author.ID, "member.php?id=")

	// Parse Pixiv member name.
	author.Name = strings.TrimSuffix(member, `</a></h2><span class="date">`)
	nameSplits := strings.Split(author.Name, ">")
	author.Name = nameSplits[len(nameSplits)-1]

	// Parse Pixiv member avatar url.
	author.Avatar = regexp.MustCompile(`img src="http.+" alt=`).FindString(member)
	author.Avatar = strings.TrimPrefix(author.Avatar, `img src="`)
	author.Avatar = strings.TrimSuffix(author.Avatar, `" alt=`)

	// Parse tags of the illustration.
	tag = strings.TrimPrefix(tag, `<meta name="keywords" content="`)
	// Remove useless tag information.
	tag = strings.TrimSuffix(tag, `,イラスト,pixiv,ピクシブ">`)
	tags := strings.Split(tag, ",")

	// Extract urls of pictures from illustration page.
	for _, v := range sources {
		if strings.Contains(v, id) {
			re := regexp.MustCompile(`/img[/0-9]+`)
			// Call ping() to guess right urls.
			sources = ping(id, pximgBaseURL+re.FindString(v))
			break
		}
	}

	// Parse description texts.
	descriptionSplits := strings.Split(description, "「")
	description = regexp.MustCompile(`.+」`).FindString(descriptionSplits[len(descriptionSplits)-1])
	description = strings.TrimSuffix(description, "」")

	// Pixiv utilizes timezone of Japan as date time.
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return PixivInfo{}, err
	}

	// Parse date time.
	t, err := time.ParseInLocation("/2006/01/02/15/04/05/", regexp.MustCompile(`[0-9/]{20,}/`).FindString(sources[0]), tz)
	if err != nil {
		return PixivInfo{}, err
	}

	return PixivInfo{
		ID:          id,
		Description: description,
		Tags:        tags,
		CreatedAt:   t.Unix(),
		Sources:     sources,
		Author:      author,
	}, nil

}

// ping is utilized to guess right urls.
func ping(id, source string) []string {
	var (
		res []string
		wg  sync.WaitGroup
	)

	// Pictures may have file extension of "png" or "jpg"
	for _, v := range []string{".png", ".jpg"} {
		wg.Add(1)

		go func(format string) {

			defer wg.Done()

			p := 0

			for {
				// Pixiv will return 403 forbidden if there is no "Refer" header entity.
				h := append(defaultBrowserHeaders, header{"Referer", illustBaseURL + id + "&page=" + strconv.Itoa(p)})

				_, err := fetch(source+"_p"+strconv.Itoa(p)+format, h)

				if err != nil {
					break
				}

				res = append(res, source+"_p"+strconv.Itoa(p)+format)

				p++
			}
		}(v)
	}

	wg.Wait()

	return res
}
