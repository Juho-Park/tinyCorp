package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func data() {
	key := "Yv7QQQXGIMq0mnH3OiED5PaYCNaNTbB9ZELNP5F4emTeO%2BzfF%2Fm%2FR8KRzVhE%2F%2F5kn%2F2%2BysZv3RjbtPjJ7XE4uw%3D%3D"
	params := url.Values{}
	params.Set("ServiceKey", key)
	params.Add("numOfRows", "10")
	params.Add("pageNo", "1")
	params.Add("sidoName", "서울")
	params.Add("searchCondition", "DAILY")

	url := "http://openapi.airkorea.or.kr/openapi/services/rest/ArpltnInforInqireSvc/getCtprvnMesureSidoLIst"

	resp, err := http.Get(url + "?" + params.Encode())

	// client := &http.Client{}
	// req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Set("ServiceKey", key)

	//resp, err := http.Get("https://www.google.co.kr/search?q=%EC%84%9C%EC%9A%B8+%EB%82%A0%EC%94%A8&oq=%EC%84%9C%EC%9A%B8+%EB%82%A0%EC%94%A8&aqs=chrome..69i57j0l5.7347j0j7&sourceid=chrome&ie=UTF-8")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
}
