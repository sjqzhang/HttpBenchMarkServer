package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var PORT *string = flag.String("p", "8880", "Server port")

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	size := 1
	var sleep int64
	sleep = 0
	if v, ok := r.Form["size"]; ok {
		var err error
		if size, err = strconv.Atoi(v[0]); err != nil {

		}
	}
	if v, ok := r.Form["sleep"]; ok {
		var err error
		if sleep, err = strconv.ParseInt(v[0], 10, 64); err != nil {

		}
	}
	var out []byte

	if sleep > 0 {
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}

	for i := 0; i < size*1024; i++ {
		out = append(out, 65)
	}

	w.Write(out)
}

func main() {

	flag.Parse()

	http.HandleFunc("/", IndexHandler)

	addr := fmt.Sprintf("0.0.0.0:%s", *PORT)
	fmt.Println(addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err)
	}

}
