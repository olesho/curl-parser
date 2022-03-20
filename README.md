# curl-parser
Parse CURL request string to Golang http.Request

## Example
In Google Chrome developer console `Inspect -> Network` choose any request. Right-click, select `Copy -> Copy as cURL`. Use the copied content as an input for `parser.Parse` function.

```
import (
	"fmt"
	"github.com/olesho/curl-parser"
	"io/ioutil"
	"net/http"
)

func main() {
	r, err := parser.Parse(`curl 'https://httpbin.org/post' \
		-X 'POST' \
		-H 'Accept: application/json' \
		-H 'Origin: https://httpbin.org' \
		-H 'Referer: https://httpbin.org/' \
		-H 'Accept-Encoding: gzip, deflate, br' \
		-H 'Host: httpbin.org' \
		-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15' \
		-H 'Content-Length: 0' \
		-H 'Accept-Language: en-GB,en;q=0.9' \
		-H 'Connection: keep-alive'`)
	if err != nil { panic(err) }
	resp, err := http.DefaultClient.Do(r)
	if err != nil { panic(err) }
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }

	fmt.Println(string(b))
}
```