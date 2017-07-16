// reflector.go - simple HTTP reflector.
//
// To the extent possible under law, Ivan Markin has waived all copyright
// and related or neighboring rights to reflector, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/nogoegst/tlspin/http"
)

func main() {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Transport, err = tlspinhttp.NewTransport(pubkey)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", proxy)
	SetupEnvironment()
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
