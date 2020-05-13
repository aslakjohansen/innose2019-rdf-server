# RDF Server for a 2019 InnoSE Project

## Requirements

- Python 3.5+
- Python modules:
  - `asyncio`
  - `aiohttp`
  - `json`

## Running

```shell
$ ./rdf-server.py
Syntax: ./rdf-server.py INTERFACE PORT MODEL_DIR
        ./rdf-server.py 0.0.0.0 8001 ../var/model
$ ./rdf-server.py 0.0.0.0 8001 ../var/model
```

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

