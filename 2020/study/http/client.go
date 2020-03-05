package main

import(
	"net/http"
	"fmt"
	"net/http/httputil"
)
func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.smzdm.com/", nil)
	if err !=nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1")
	
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	//resp, err := http.Get("https://www.imooc.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	contents , err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contents))
}