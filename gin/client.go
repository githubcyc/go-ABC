package main

import (
	"net/http"
	"fmt"
)
var base = "http://127.0.0.1:8080/"

func main(){
    // 调用最基本的GET,并获得返回值
    resp,_ := http.Get(base + "ping/1")
    fmt.Println(resp)

    // 调用最基本的POST,并获得返回值
    // resp,_ = http.Post(base  + "test2", "",strings.NewReader(""))
    // helpRead(resp)
}