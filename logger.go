package webutil

import (
    "log"
    "net/http"
)

type LoggerHandler struct {
    H      http.Handler
    Logger *log.Logger
}

func (lh LoggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    lh.Logger.Printf("%s -> %s %s\n", r.RemoteAddr, r.Method, r.URL)
    lh.H.ServeHTTP(w, r)
}
