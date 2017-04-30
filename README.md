# filter

json data filter by string

## Before  --> filter string --> After

* Before

{
  "id": true,
  "test": {
    "a": false,
    "b": 123
  },
  "url": "https://lnx.cm",
  "object": {
    "site": "leenanxi.com",
    "content": {
      "name": "linx",
      "key": "youku"
    },
    "attr": [
      {
        "name": "leenanxi",
        "key": "youku"
      }
    ],
    "attachments": [
      "ttt",
      {
        "url": "https",
        "name": "110908"
      },
      {
        "url": "http",
        "name": "0000"
      }
    ]
  }
}

* filter string

```
test/a,id,object(site,attachments/url,content/name)
```

* After

```json
{
	"id": true,
	"object": {
		"attachments": [
			null,
			{
				"url": "https"
			},
			{
				"url": "http"
			}
		],
		"content": {
			"name": "linx"
		},
		"site": "leenanxi.com"
	},
	"test": {
		"a": false
	}
}
```

## Sample Code

```go
package main

import (
	"strings"
	"encoding/json"
	"github.com/ti/filter"
	"log"
)

func main() {
	jsonString := `{
  "id": true,
  "test": {
    "a": false,
    "b": 123
  },
  "url": "https://lnx.cm",
  "object": {
    "site": "leenanxi.com",
    "content": {
      "name": "linx",
      "key": "youku"
    },
    "attr": [
      {
        "name": "leenanxi",
        "key": "youku"
      }
    ],
    "attachments": [
      "ttt",
      {
        "url": "https",
        "name": "110908"
      },
      {
        "url": "http",
        "name": "0000"
      }
    ]
  }
}`

	filterString := "test/a,id,object(site,attachments/url,content/name)"
	dec := json.NewDecoder(strings.NewReader(jsonString))
	dec.UseNumber()
	var data interface{}
	if err := dec.Decode(&data); err != nil {
		panic(err)
	}
	err := filter.Filter(&data, filterString)
	if err != nil {
		log.Panic(err)
	}
	bytes, _ := json.MarshalIndent(data, "", "\t")
	log.Println("The Filter Result：", string(bytes))

}
```