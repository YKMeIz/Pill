package pill

import (
	"testing"
)

var yandeSamples map[string]YandeInfo

func init() {
	yandeSamples = make(map[string]YandeInfo)

	yandeSamples["785598"] = YandeInfo{
		ID: "785598",
		Tags: []string{
			"areola",
			"garter",
			"nakatama_kyou",
			"no_bra",
			"nun",
			"pantyhose",
			"skirt_lift",
		},
		CreatedAt: 1621012408,
		Source:    "クールなシスターの事務的“聖処理”生活～“聖婚生活”によるラブハメおまんこをしましょう～",
		MD5:       "82ee37f4775bc3a5807f1cb08ce775a7",
		File:      "https://files.yande.re/image/82ee37f4775bc3a5807f1cb08ce775a7/yande.re%20785598%20areola%20garter%20nakatama_kyou%20no_bra%20nun%20pantyhose%20skirt_lift.jpg",
		Sample:    "https://files.yande.re/sample/82ee37f4775bc3a5807f1cb08ce775a7/yande.re%20785598%20sample%20areola%20garter%20nakatama_kyou%20no_bra%20nun%20pantyhose%20skirt_lift.jpg",
		Preview:   "https://assets.yande.re/data/preview/82/ee/82ee37f4775bc3a5807f1cb08ce775a7.jpg",
	}
}

func TestYande(t *testing.T) {
	for k, v := range yandeSamples {
		res, err := Yande(k)
		if err != nil {
			t.Error("Yande() execute failed:", err)
		}

		compareStructVal(t, res, v)
	}
}
