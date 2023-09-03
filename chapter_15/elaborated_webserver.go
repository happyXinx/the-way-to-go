package main

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// hello world 的计数器
var helloRequests = expvar.NewInt("hello-requests")

// flags:
var webroot = flag.String("root", "/Users/happyxin", "web root directory")

// simple flag server
var booleanflag = flag.Bool("boolean", true, "another flag for testing")

type Counter struct {
	n int
}

// 一个通道
type Chan chan int

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(Logger))
	http.Handle("/go/hello", http.HandlerFunc(HelloServer))
	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)
	http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot))))
	http.Handle("/flags", http.HandlerFunc(FlagServer))
	http.Handle("/args", http.HandlerFunc(ArgServer))
	http.Handle("/chan", ChanCreate())
	http.Handle("/date", http.HandlerFunc(DateServer))

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}

func Logger(w http.ResponseWriter, req *http.Request) {
	log.Print(req.URL.String())
	w.WriteHeader(404)
	w.Write([]byte("oops"))
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	helloRequests.Add(1)
	io.WriteString(w, "hello, world!\n")
}

// 通过这个方法满足 expvar.Var 接口，所以就可以直接导出它。
func (ctr *Counter) String() string { return fmt.Sprintf("%v", ctr.n) }

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctr.n++
	case "POST":
		buf := new(bytes.Buffer)
		io.Copy(buf, r.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprint(w, "counter reset\n")
		}
	}
	fmt.Fprint(w, "counter = %d\n", ctr.n)
}

func FlagServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Flags:\n")
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != f.DefValue {
			fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
		} else {
			fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
		}
	})
}

// 简单参数服务器
func ArgServer(w http.ResponseWriter, req *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, " ")
	}
}

func ChanCreate() Chan {
	c := make(Chan)
	go func(c Chan) {
		for x := 0; ; x++ {
			c <- x
		}
	}(c)
	return c
}

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("channel send #%d\n", <-ch))
}

// 执行一个程序，输出重定向
func DateServer(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(rw, "pipe: %s\n", err)
		return
	}

	p, err := os.StartProcess("/bin/date", []string{"date"}, &os.ProcAttr{Files: []*os.File{nil, w, w}})
	defer r.Close()
	w.Close()
	if err != nil {
		fmt.Fprintf(rw, "fork/exec: %s\n", err)
		return
	}
	defer p.Release()
	io.Copy(rw, r)
	wait, err := p.Wait()
	if err != nil {
		fmt.Fprintf(rw, "wait: %s\n", err)
		return
	}
	if !wait.Exited() {
		fmt.Fprintf(rw, "date: %v\n", wait)
		return
	}
}
