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
    ip := r.RemoteAddr
    realIp := r.Header.Get("X-Real-Ip")
    if realIp != "" {
        ip = realIp
    }
    lh.Logger.Printf("%s -> %s %s\n", ip, r.Method, r.URL)
    lh.H.ServeHTTP(w, r)
}
