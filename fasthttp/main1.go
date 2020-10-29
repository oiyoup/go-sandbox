package main

import (
    "fmt"
    "github.com/rs/zerolog/log"
    "github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
    _, _ = fmt.Fprintf(ctx, "Hi there! RquestURI is %q", ctx.RequestURI())
}
func main() {
    
    if err:= fasthttp.ListenAndServe("127.0.0.1:8081", fastHTTPHandler); err != nil {
        log.Error().
            Msg(err.Error());
    }
}
