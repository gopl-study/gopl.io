package main

import (
	"fmt"
	"net/http"

	"study/params"
)

func main() {

}

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // 기본 값 지정
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ... 나머지 처리기 ...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}
