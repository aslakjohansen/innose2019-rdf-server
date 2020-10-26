# Building

## Preliminaries

First, install the preliminaries from the main project [README](../README.md).

This project is laid out in an unconventional way that bypasses the typical [Go](https://golang.org) restriction of having to have the source code in the root of the repository. In this repository the source code is in [/src](../src).

To make the work, check out the repository whereever you want. Lets use `~/innose2019-rdf-server` as an example. Set up a Go development environment (`$GOPATH` and `$PATH` according to [this](https://golang.org/doc/gopath_code.html)). Then link `$GOPATH/src/innose2019-rdf-server` to the source code (`~/innose2019-rdf-server/src`, in the example).

On a linux box this is done by:

```shell
$ ln -s ~/innose2019-rdf-server/src $GOPATH/src/innose2019-rdf-server
```

Now install the dependencies:

```shell
innose2019-rdf-server/src$ make deps
```

## Build Command

```shell
innose2019-rdf-server/src$ make rdf-server
...
```
