package pill

import "testing"

var konachanSamples map[string]KonachanInfo

func init() {
	konachanSamples = make(map[string]KonachanInfo)

	konachanSamples["325127"] = KonachanInfo{
		ID: "325127",
		Tags: []string{
			"brown_hair",
			"flowers",
			"harusame_(user_wawj5773)",
			"long_hair",
			"original",
			"pantyhose",
			"red_eyes",
			"rose",
			"shirt",
			"skirt",
		},
		CreatedAt: 1617449377,
		Source:    "https://www.pixiv.net/en/artworks/88896458",
		MD5:       "fc5558b624f2a0187a326279d674f34c",
		File:      "https://konachan.com/image/fc5558b624f2a0187a326279d674f34c/Konachan.com%20-%20325127%20brown_hair%20flowers%20harusame_%28user_wawj5773%29%20long_hair%20original%20pantyhose%20red_eyes%20rose%20shirt%20skirt.jpg",
		Sample:    "https://konachan.com/sample/fc5558b624f2a0187a326279d674f34c/Konachan.com%20-%20325127%20sample.jpg",
		Preview:   "https://konachan.com/data/preview/fc/55/fc5558b624f2a0187a326279d674f34c.jpg",
	}
}

func TestKonachan(t *testing.T) {
	for k, v := range konachanSamples {
		res, err := Konachan(k)
		if err != nil {
			t.Error("konachan() execute failed:", err)
		}

		compareStructVal(t, res, v)
	}
}
