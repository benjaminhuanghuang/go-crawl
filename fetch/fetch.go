package fetch

import (
	"net/http"
	"net/http/httputil"
)

/*
Fetch ...
*/
func Fetch(url string, parse func(contents []byte)) {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	// Set the client to mobile device
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Mobile/15E148")

	// server will redirect "https://coding.imooc.com" to "https://coding.m.imooc.com" if browser is a mobile device
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// fmt.Println("Redirect: ", req)
		return nil
	}}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	parse(s)
}
