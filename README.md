reflector
=========

Usage on AppEngine Flex
-----------------------

Go to Cloud Console and run:

```
  $ go get -v -d github.com/nogoegst/reflector
  $ cd gopath/src/github.com/nogoegst/reflector
  $ gcloud app deploy
```

Using another target
--------------------

The target has to run with TLS and HTTP/2 support (e.g. `meek-server` built with Go 1.6+).

Probe the TLS public key to pin:

```
  $ go get -u github.com/nogoegst/tlspin/cmd/...
  $ tlspin-probe meek.bamsoftware.com:443
```

and set it in `config.go` along with other values.

