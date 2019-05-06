package fetcher

import (
	"bufio"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
)

//模拟网页请求，防止403
func httpRequest(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	client := http.Client{}
	return client.Do(request)
}
func Fetch(url string)([]byte, error) {
	response, err := httpRequest(url)
	if err != nil {
		log.Error("http请求错误，err：", err, "url:", url)
		return nil, err
	}
	defer response.Body.Close()
	reader := bufio.NewReader(response.Body)
	return  ioutil.ReadAll(reader)
}