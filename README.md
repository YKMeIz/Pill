# Pill :pill:

[![Go Reference](https://img.shields.io/badge/go-reference-informational.svg?style=flat-square)](https://pkg.go.dev/github.com/YKMeIz/Pill)
[![Go Report Card](https://goreportcard.com/badge/github.com/YKMeIz/Pill?style=flat-square)](https://goreportcard.com/report/github.com/YKMeIz/Pill)
[![License](https://img.shields.io/github/license/YKMeIz/Pill.svg?color=%232b2b2b&style=flat-square)](https://github.com/YKMeIz/Pill/blob/master/LICENSE)

Pill is a parser utilized to grab illustration information from [Pixiv](https://pixiv.net), [yande.re](https://yande.re), and [Konachan](https://konachan.com).

Pixiv Illustration information contains:
- Illustration Title
- Description
- Tags
- Creation Date & Time (represent in format of unix timestamp)
- Picture URLs
- Author Information
  - Pixiv Member ID
  - Name
  - Avatar URL

Others are various. More details can be found in [GoDoc](https://godoc.org/github.com/YKMeIz/Pill).

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/YKMeIz/Pill"
)

func main() {
	res, err := pill.Pixiv("66917712")

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(res)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

```

The output, in a pretty and human readable JSON format, would be:

```json
{
  "title": "コード:002",
  "id": "66917712",
  "description": "Line drawings converted by exist images.線画は既存の画像で変換されます。",
  "tags": [
    "Code:002",
    "線画"
  ],
  "created_at": 1516617934,
  "sources": [
    "https://i.pximg.net/img-original/img/2018/01/22/19/45/34/66917712_p0.png",
    "https://i.pximg.net/img-original/img/2018/01/22/19/45/34/66917712_p1.png"
  ],
  "author": {
    "id": "16412800",
    "name": "90榣",
    "avatar": "https://i.pximg.net/user-profile/img/2018/01/22/19/30/01/13726842_b73c069f1a20efc265f12c2693fea41d_50.png"
  }
}

```
