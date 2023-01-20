package main

import (
	"fmt"
	"github.com/valyala/fastjson"
	"net"
	"strings"
)

type RequestSettings struct {
	IP        string `json:"ip"`
	ID        string `json:"id"`
	UA        string `json:"ua"`
	RequestID string `json:"request_id"`
	Language  string `json:"language"`
	SourceID  string `json:"source_id"`
	Referrer  string `json:"referrer"`
	Domain    string `json:"domain"`
}

type PublisherRequest struct {
	IP        net.IP
	ID        string
	UA        string
	RequestID string
	Language  string
	SourceID  string
	Referrer  string
	Domain    string
}

func main() {
	body := []byte(`
{
  "id": "1242560044",
  "site": {
    "id": "301999",
    "domain": "example.com",
    "page": "https://example.com"
  },
  "device": {
    "ua": "Mozilla (Linux Android 10.1) Safari 537.36",
    "ip": "127.0.0.1",
    "js": 1,
    "language": "ru"
  },
  "user": {
    "id": "163625693"
  },
  "at": 1,
  "source":"777",
  "referrer":"888",
  "request":"999"
}
`)
	settingsKeys := RequestSettings{
		IP:        "device.ip",
		ID:        "id",
		UA:        "device.ua",
		RequestID: "request",
		Language:  "device.language",
		SourceID:  "source",
		Referrer:  "referrer",
		Domain:    "site.domain",
	}

	res := parseCustomJSON(body, settingsKeys)

	fmt.Println(res)
}

func parseCustomJSON(body []byte, queryArgKeys RequestSettings) PublisherRequest {
	pr := PublisherRequest{}

	jsonParserPool := fastjson.ParserPool{}
	parser := jsonParserPool.Get()
	defer jsonParserPool.Put(parser)

	data, errParseBytes := parser.ParseBytes(body)
	if errParseBytes != nil {
		fmt.Printf("error parsing custom json, %s", errParseBytes.Error())
		return pr
	}

	var item *fastjson.Value
	item = data.Get()

	if key := queryArgKeys.IP; key != "" {
		if v := string(item.GetStringBytes(strings.Split(key, ".")...)); v != "" {
			pr.IP = net.ParseIP(v)
		}
	}

	if key := queryArgKeys.ID; key != "" {
		pr.ID = string(item.GetStringBytes(strings.Split(key, ".")...))
	}

	if key := queryArgKeys.UA; key != "" {
		pr.UA = string(item.GetStringBytes(strings.Split(key, ".")...))
	}

	if key := queryArgKeys.RequestID; key != "" {
		pr.RequestID = string(item.GetStringBytes(strings.Split(key, ".")...))
	}

	if key := queryArgKeys.Referrer; key != "" {
		pr.Referrer = string(item.GetStringBytes(strings.Split(key, ".")...))
	}

	if key := queryArgKeys.SourceID; key != "" {
		pr.SourceID = string(item.GetStringBytes(strings.Split(key, ".")...))
	}

	if key := queryArgKeys.Language; key != "" {
		pr.Language = string(item.GetStringBytes(strings.Split(key, ".")...))
	}

	if key := queryArgKeys.Domain; key != "" {
		pr.Domain = string(item.GetStringBytes(strings.Split(key, ".")...))
	}

	return pr
}
