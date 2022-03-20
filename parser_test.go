package parser

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	a := assert.New(t)

	req, err := Parse(`curl 'https://example1.com/cnt/' \
  -H 'authority: kraken.rambler.ru' \
  -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'content-type: text/plain;charset=UTF-8' \
  -H 'accept: */*' \
  -H 'origin: https://example1.com' \
  -H 'sec-fetch-site: cross-site' \
  -H 'sec-fetch-mode: no-cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://example1.com/' \
  -H 'accept-language: en-GB,en;q=0.9' \
  -H 'cookie: ruid=1CIAACKCNWK5WwOwAd+0+QB=' \
  --data-raw 'et=act&pid=289463&rid=1647675031.429-2085465145&tid=t1.289463.1550425499.1647673888894&v=2.0.4&exp=exp_bot%2Csplit_a%2Cexp_ping%2Cyes&ct=web&aduid=9d54a630-75d3-42ad-a5b0-9ddd7685a562&aduidsc=example.com&meta=%7B%22time%22%3A19%2C%22screens%22%3A%5B5%2C5%2C4%2C14%2C14%2C0%2C0%2C0%2C0%2C0%5D%7D&eid=8312750512215387&stid=1537878548_1647673888894&sn=1&sen=12&en=12&fid=pA8AAENKs1dAz24vAbTsVgA%3D&fip=pA8AAENKs1d82cPmAbpJPAA%3D' \
  --compressed`)
	if err != nil {
		t.Error(err)
		return
	}
	a.Equal(req.URL.String(), "https://example1.com/cnt/")
	a.Equal(req.Header.Get("authority"), "kraken.rambler.ru")
	a.Equal(req.Header.Get("sec-ch-ua"), "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\"")
	a.Equal(req.Header.Get("sec-ch-ua-mobile"), "?0")
	a.Equal(req.Header.Get("content-type"), "text/plain;charset=UTF-8")
	a.Equal(req.Header.Get("origin"), "https://example1.com")
	a.Equal(req.Header.Get("referer"), "https://example1.com/")

	b, _ := ioutil.ReadAll(req.Body)
	a.Equal(string(b), "et=act&pid=289463&rid=1647675031.429-2085465145&tid=t1.289463.1550425499.1647673888894&v=2.0.4&exp=exp_bot%2Csplit_a%2Cexp_ping%2Cyes&ct=web&aduid=9d54a630-75d3-42ad-a5b0-9ddd7685a562&aduidsc=example.com&meta=%7B%22time%22%3A19%2C%22screens%22%3A%5B5%2C5%2C4%2C14%2C14%2C0%2C0%2C0%2C0%2C0%5D%7D&eid=8312750512215387&stid=1537878548_1647673888894&sn=1&sen=12&en=12&fid=pA8AAENKs1dAz24vAbTsVgA%3D&fip=pA8AAENKs1d82cPmAbpJPAA%3D")

	req, err = Parse(`curl -X POST 'https://example2.org/2' \
  -H 'Connection: keep-alive' \
  -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'Upgrade-Insecure-Requests: 1' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36' \
  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
  -H 'Sec-Fetch-Site: none' \
  -H 'Sec-Fetch-Mode: navigate' \
  -H 'Sec-Fetch-User: ?1' \
  -H 'Sec-Fetch-Dest: document' \
  -H 'Accept-Language: en-GB,en;q=0.9' \
  --compressed`)
	if err != nil {
		t.Error(err)
		return
	}
	a.Equal(req.URL.String(), "https://example2.org/2")
}