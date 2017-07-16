// appengine.go - AppEngine Flex environment bits.
//
// To the extent possible under law, Ivan Markin has waived all copyright
// and related or neighboring rights to reflector, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

// +build appengine

package main

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
	return
}

func SetupEnvironment() {
	http.HandleFunc("/_ah/health", healthCheckHandler)
}
