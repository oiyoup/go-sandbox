package main

import (
    "fmt"
    "github.com/valyala/fasthttp"
)

type MyHandler struct {
    foobar string
}

func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
    _, _ = fmt.Fprintf(ctx, "Hello World! Requested path is %q. Foobar is %q",
        ctx.Path(), h.foobar)
}

func main() {
    myHandler := &MyHandler{
        foobar: "good",
    }

    _ = fasthttp.ListenAndServe("127.0.0.1:8080", myHandler.HandleFastHTTP)
}
