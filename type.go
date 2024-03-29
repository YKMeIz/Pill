package pill

const (
	illustBaseURL   = "https://www.pixiv.net/en/artworks/"
	yandeBaseURL    = "https://yande.re/post/show/"
	konachanBaseURL = "https://konachan.com/post/show/"
)

// defaultBrowserHeaders is utilized to pretend http requests are sent from a browser.
var defaultBrowserHeaders = []header{
	{"Accept-Language", "en-US,en"},
	{"User-Agent", "Mozilla/5.0 Firefox"},
}

// header defines http header.
type header struct {
	key,
	val string
}

// PixivInfo describes the details of a Pixiv illustration work.
type PixivInfo struct {
	// Title is illustration title.
	Title string `json:"title"`
	// ID is illustration ID. It comes from url query value of "illust_id".
	ID string `json:"id"`
	// Description is the description texts from illustration page.
	Description string `json:"description"`
	// Tags is the array of tags that illustration has.
	Tags []string `json:"tags"`
	// CreatedAt is unix timestamp of the illustration creation date time.
	CreatedAt int64 `json:"created_at"`
	// Sources is urls of pictures uploaded under the illustration.
	// It is an array because a illustration page may contain multiple pictures.
	Sources []string `json:"sources"`
	// Author is the author of the illustration.
	Author PixivMember `json:"author"`
}

// PixivMember describes the details of a Pixiv member (author).
type PixivMember struct {
	// ID is the Pixiv member ID.
	ID string `json:"id"`
	// Name is the Pixiv member name.
	Name string `json:"name"`
	// Avatar is url of the Pixiv member avatar.
	Avatar string `json:"avatar"`
}

// metaPreloadData is content of HTML tag with id of meta-preload-data.
type metaPreloadData struct {
	Illust struct {
		ID struct {
			Id          string `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			CreateDate  string `json:"createDate"`
			Urls        struct {
				Original string `json:"original"`
			} `json:"urls"`
			Tags struct {
				Tags []struct {
					Tag string `json:"tag"`
				} `json:"tags"`
			} `json:"tags"`
		}
	} `json:"illust"`
	User struct {
		ID struct {
			UserID   string `json:"userID"`
			Name     string `json:"name"`
			ImageBig string `json:"imageBig"`
		}
	} `json:"user"`
}

// YandeInfo describes the details of a post on yande.re.
type YandeInfo DanbooruBase

// KonachanInfo describes the details of a post on konachan.com.
type KonachanInfo DanbooruBase

// DanbooruBase describes the details of a post on danbooru based website.
type DanbooruBase struct {
	// ID is post ID.
	ID string `json:"id"`
	// Tags is the array of tags that post has.
	Tags []string `json:"tags"`
	// CreatedAt is unix timestamp of the post creation date time.
	CreatedAt int64 `json:"created_at"`
	// Source describes where the post is from. It is usually a url.
	Source string `json:"source"`
	// MD5 is md5 hash of the post.
	MD5 string `json:"md5"`
	// File is file url of the post.
	File string `json:"file"`
	// Sample is sample url of the work. It is usually a small version of original file.
	Sample string `json:"sample"`
	// Preview is thumbnail url of the work.
	Preview string `json:"preview"`
}

type postRegisterResp struct {
	Posts []struct {
		// ID is post ID.
		ID int64 `json:"id"`
		// Tags is tags that post has.
		Tags string `json:"tags"`
		// CreatedAt is unix timestamp of the post creation date time.
		CreatedAt int64 `json:"created_at"`
		// Source describes where the post is from. It is usually a url.
		Source string `json:"source"`
		// MD5 is md5 hash of the post.
		MD5 string `json:"md5"`
		// File is file url of the post.
		File string `json:"file_url"`
		// Sample is sample url of the work. It is usually a small version of original file.
		Sample string `json:"sample_url"`
		// Preview is thumbnail url of the work.
		Preview string `json:"preview_url"`
	} `json:"posts"`
}
