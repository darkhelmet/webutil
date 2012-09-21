# webutil

Little web utilities for Go

## GzipHandler

Wraps a `http.Handler` to gzip things.

    var handler http.Handler = buildHandler()
    handler = webutil.GzipHandler{handler}
    http.Handle("/", handler)
    
## LogHandler

Basic request logging

    var handler http.Handler = buildHandler()
    logger := log.New(...)
    handler = webutil.LogHandler{handler, logger}
    http.Handle("/", handler)
    
## CanonicalHostHandler

301 redirects back to the canonical host

    var handler http.Handler = buildHandler()
    logger := log.New(...)
    handler = webutil.CanonicalHostHandler{handler, "wayneenterprises.com", "http"}
    http.Handle("/", handler)
