# RDF Server for a 2019 InnoSE Project

## Requirements

- Python 3.5+
- Python modules:
  - `asyncio`
  - `aiohttp`
  - `json`

## Running

```shell
innose2019-rdf-server/src$ ./rdf-server.py 
Syntax: ./rdf-server.py INTERFACE PORT MODEL_DIR NAMESPACE
        ./rdf-server.py 0.0.0.0 8001 ../var/model http://ses.sdk.dk/junk/test#
innose2019-rdf-server/src$ ./rdf-server.py 0.0.0.0 8001 ../var/model http://ses.sdk.dk/junk/test#
STATUS: Loading model "../var/model/3.ttl".
STATUS: Listening on 0.0.0.0:8001
^C
STATUS: Exiting ...
innose2019-rdf-server/src$ 

```

**Note:** The current version of Brick is quite large and thus takes a considerable amount of time to load/store.

## Interface

### Fetch time from server

```shell
$ curl -X PUT -d '42' http://localhost:8001/time
{
    "success": true,
    "time": 1589391233.535275
}
```

**Note:** The `42` payload is not used, but the payload needs to be valid json.

### Fetch store model to disk

```shell
$ curl -X PUT -d '42' http://localhost:8001/store
{
    "filename": "../var/model/3.ttl",
    "success": true
}
```

