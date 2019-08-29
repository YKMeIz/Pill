package pill

import (
	"reflect"
	"testing"
)

var samples map[string]PixivInfo

func init() {
	samples = make(map[string]PixivInfo)

	samples["66917712"] = PixivInfo{
		Title:       "コード:002",
		ID:          "66917712",
		Description: "Line drawings converted by exist images.<br />線画は既存の画像で変換されます。",
		Tags: []string{
			"Code:002",
			"線画",
		},
		CreatedAt: 1516617934,
		Sources: []string{
			"https://i.pximg.net/img-original/img/2018/01/22/19/45/34/66917712_p0.png",
			"https://i.pximg.net/img-original/img/2018/01/22/19/45/34/66917712_p1.png",
		},
		Author: PixivMember{
			ID:     "16412800",
			Name:   "90榣",
			Avatar: "https://i.pximg.net/user-profile/img/2018/01/22/19/30/01/13726842_b73c069f1a20efc265f12c2693fea41d_170.png",
		},
	}

	samples["66917851"] = PixivInfo{
		Title:       "コード:002と魚",
		ID:          "66917851",
		Description: "Line drawings converted by exist images.<br />線画は既存の画像で変換されます。",
		Tags: []string{
			"Code:002",
			"線画",
		},
		CreatedAt: 1516618445,
		Sources: []string{
			"https://i.pximg.net/img-original/img/2018/01/22/19/54/05/66917851_p0.png",
		},
		Author: PixivMember{
			ID:     "16412800",
			Name:   "90榣",
			Avatar: "https://i.pximg.net/user-profile/img/2018/01/22/19/30/01/13726842_b73c069f1a20efc265f12c2693fea41d_170.png",
		},
	}
}

func TestPixiv(t *testing.T) {
	for k, v := range samples {
		res, err := Pixiv(k)
		if err != nil {
			t.Error("Pixiv() execute failed:", err)
		}

		if reflect.DeepEqual(res, v) {
			continue
		}

		resVal := reflect.ValueOf(res)
		val := reflect.ValueOf(v)

		for i := 0; i < val.NumField(); i++ {
			if !reflect.DeepEqual(val.Field(i).Interface(), resVal.Field(i).Interface()) {
				t.Error("Error on", val.Type().Field(i).Name, ": expect", val.Field(i).Interface(), ", get", resVal.Field(i).Interface())
			}
		}
	}
}

func TestFetchInPixiv(t *testing.T) {
	p, _ := Pixiv("")

	if !reflect.DeepEqual(p, PixivInfo{}) {
		t.Error(`Pixiv() returned result while it shouldn't'.`)
	}
}
