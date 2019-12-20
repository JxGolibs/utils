package utils

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// func httpDo(method, url string, body io.Reader) (*Request, error)) {
// 	client := &http.Client{
// 		Request: &http.Request{
// 			URL:        u,
// 			Method:     "GET",
// 			ProtoMajor: 1,
// 			ProtoMinor: 1,
// 			Header:     make(map[string][]string),
// 		},
// 	}
// 	client.Header.Set("User-Agent", "NTRIP GoClient")
// 	client.Header.Set("Ntrip-Version", "Ntrip/2.0")

// 	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
// 	if err != nil {
// 		// handle error
// 	}

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("Cookie", "name=anny")

// 	resp, err := client.Do(req)

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		// handle error
// 	}

// 	fmt.Println(string(body))
// }
