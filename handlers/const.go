package handlers

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var cookieName string
var cookieStore *sessions.CookieStore
var exludedHeaders = map[string]bool{
	"Content-Length": true,
}

func isExluded(h string) bool {
	return exludedHeaders[h]
}

func Init(secret, cookie string) {
	cookieName = cookie

	cookieStore = sessions.NewCookieStore([]byte(secret))
	cookieStore.Options.MaxAge = 0
	cookieStore.Options.HttpOnly = true

	http.DefaultClient.Transport = &http.Transport{Proxy: http.ProxyFromEnvironment}
}

const chunkSize = 1024
