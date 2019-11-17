package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", echo)
	// 中间件
	http.Handle("/", timeMiddleware(http.HandlerFunc(echo)))
	http.ListenAndServe(":8080", nil)
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	})
}

func echo(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("echo error"))
		return
	}

	writeLen, err := w.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
}

var _ = `
任何方法实现了ServeHTTP，即是一个合法的http.Handler，读到这里你可能会有一些混乱，
我们先来梳理一下http库的Handler，HandlerFunc和ServeHTTP的关系：

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
只要你的handler函数签名是：
func (ResponseWriter, *Request)
那么这个handler和http.HandlerFunc()就有了一致的函数签名，可以将该handler()函数进行类型转换，转为http.HandlerFunc。
而http.HandlerFunc实现了http.Handler这个接口。
在http库需要调用你的handler函数来处理http请求时，会调用HandlerFunc()的ServeHTTP()函数，可见一个请求的基本调用链是这样的：

h = getHandler() => h.ServeHTTP(w, r) => h(w, r)
上面提到的把自定义handler转换为http.HandlerFunc()这个过程是必须的，因为我们的handler没有直接实现ServeHTTP这个接口。上面的代码中我们看到的HandleFunc(注意HandlerFunc和HandleFunc的区别)里也可以看到这个强制转换过程：

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

// 调用

func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	mux.Handle(pattern, HandlerFunc(handler))
}
知道handler是怎么一回事，我们的中间件通过包装handler，再返回一个新的handler就好理解了。

总结一下，我们的中间件要做的事情就是通过一个或多个函数对handler进行包装，返回一个包括了各个中间件逻辑的函数链。我们把上面的包装再做得复杂一些：

customizedHandler = logger(timeout(ratelimit(helloHandler)))
这个函数链在执行过程中的上下文可以用图 5-8来表示。
`
