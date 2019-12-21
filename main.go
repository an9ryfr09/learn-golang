package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"sync"
)

type requestUrl struct {
	name   string
	url    string
	method string
}

func init() {
	runtime.GOMAXPROCS(6)
}

func checkErr(receiver chan<- []byte, err error) bool {
	if err != nil {
		receiver <- []byte(err.Error())
		return false
	}
	return true
}

func request(wg *sync.WaitGroup, receiver chan<- []byte, u requestUrl) {
	client := &http.Client{}
	var check bool

	wg.Add(1)
	request, err := http.NewRequest(u.method, u.url, nil)
	if check = checkErr(receiver, err); !check {
		return
	}

	resp, err := client.Do(request)
	if check = checkErr(receiver, err); !check {
		return
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if check = checkErr(receiver, err); !check {
		return
	}

	receiver <- contents
}

func writeFile(file *os.File, wg *sync.WaitGroup, reader <-chan []byte, u requestUrl) {
	defer wg.Done()
	file, _ = os.Create(fmt.Sprintf("result/%s.html", u.name))
	file.WriteString(string(<-reader))
}

func main() {
	var wg sync.WaitGroup
	var file *os.File
	channel := make(chan []byte)

	urls := []requestUrl{
		{
			name:   "速美官网pc端总站",
			url:    "https://www.sumeihome.com/",
			method: "GET",
		},
		{
			name:   "速美官网pc端北京分站",
			url:    "https://bj.sumeihome.com/",
			method: "GET",
		},
		{
			name:   "速美官网h5端总站",
			url:    "https://m.sumeihome.com/",
			method: "GET",
		},
		{
			name:   "速美官网h5端北京分站",
			url:    "https://bj.m.sumeihome.com/",
			method: "GET",
		},
	}

	for _, u := range urls {
		go request(&wg, channel, u)
		go writeFile(file, &wg, channel, u)
	}

	wg.Wait()
}
