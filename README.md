# RDF Server for a 2019 InnoSE Project

This server is a [Go](https://golang.org) program that wraps some [Python](https://www.python.org) code for access to [rdflib](https://rdflib.readthedocs.io).

Index:
- [Requirements](#requirements)
- [Building](doc/building.md)
- [Running](doc/running.md)
- [Testing](doc/testing.md)

## Requirements

- The `python2.7-dev` package (as it is a dependency for go-python)
- The following python 2.7 modules:
  - `rdflib` (both 4.2.2 and 5.0.0 have been tested)
  - `requests` (for some reason rdflib makes use of it without depending on it)

For testing:
- `mosquitto-clients` (manual MQTT publication)

