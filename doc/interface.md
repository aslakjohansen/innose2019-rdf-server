# Interface

## REST

### Fetch time from server

```shell
$ curl -X PUT -d '42' http://localhost:8001/time
{
    "success": true,
    "time": 1589391233.535275
}
```

**Note:** The `42` payload is not used, but the payload needs to be valid json.

### Store model to disk

```shell
$ curl -X PUT -d '42' http://localhost:8001/store
{
    "filename": "../var/model/3.ttl",
    "success": true
}
```

**Note:** The `42` payload is not used, but the payload needs to be valid json.

### Fetch the current namespace mapping

```shell
$ curl -X PUT -d '42' http://localhost:8001/namespaces
{
    "namespaces": {
        "bldg": "https://brickschema.org/schema/1.1.0/ExampleBuilding#",
        "brick": "https://brickschema.org/schema/1.1.0/Brick#",
        "dcterms": "http://purl.org/dc/terms#",
        "n": "http://ses.sdk.dk/junk/test#",
        "owl": "http://www.w3.org/2002/07/owl#",
        "rdf": "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
        "rdfs": "http://www.w3.org/2000/01/rdf-schema#",
        "sdo": "http://schema.org#",
        "skos": "http://www.w3.org/2004/02/skos/core#",
        "tag": "https://brickschema.org/schema/1.1.0/BrickTag#",
        "xml": "http://www.w3.org/XML/1998/namespace",
        "xsd": "http://www.w3.org/2001/XMLSchema#"
    },
    "success": true
}
```

**Note:** The `42` payload is not used, but the payload needs to be valid json.

### Resolving a query

```shell
$ curl -X PUT -d '"SELECT ?pred ?obj WHERE {brick:Sensor ?pred ?obj .}"' http://localhost:8001/query
{
    "resultset": [
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2111"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "https://brickschema.org/schema/1.1.0/Brick#Sensor"
        ],
...
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2150"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1556"
        ]
    ],
    "success": true
}

```

### Perform an update

```shell
$ curl -X PUT -d '"PREFIX brick: <https://brickschema.org/schema/1.1.0/Brick#>\n\nDELETE { brick:Sensor rdfs:subClassOf ?obj .} WHERE {brick:Sensor rdfs:subClassOf ?obj .}"' http://localhost:8001/update
{
    "success": true
}
```

**Note:** rdflib has a few restrictions:
- namespace prefixes must be used
- a where clause must exist
- no slashes in entity names

### Inspect a query

```shell
$ curl -X PUT -d '"! SELECT ?pred ?obj WHERE {?sub ?pred ?obj .}"' http://localhost:8001/inspect
{
    "success": false,
    "error": {
        "type": "lex",
        "data": "Lexer error: could not match text starting at 1:1 failing at 1:2.\n\tunmatched text: \"!\""
    }
}
$ curl -X PUT -d '"SELECT SELECT ?pred ?obj WHERE {?sub ?pred ?obj .}"' http://localhost:8001/inspect
{
    "success": false,
    "error": {
        "type": "parse",
        "data": "syntax error"
    }
}
$ curl -X PUT -d '"SELECT ?pred ?obj WHERE {?sub ?pred ?obj .}"' http://localhost:8001/inspect
{
    "success": true,
    "tokens": [
        "57352 "SELECT" 0 (1, 1)-(1, 6)",
        "57346 "?pred" 7 (1, 8)-(1, 12)",
        "57346 "?obj" 13 (1, 14)-(1, 17)",
        "57353 "WHERE" 18 (1, 19)-(1, 23)",
        "57355 "{" 24 (1, 25)-(1, 25)",
        "57346 "?sub" 25 (1, 26)-(1, 29)",
        "57346 "?pred" 30 (1, 31)-(1, 35)",
        "57346 "?obj" 36 (1, 37)-(1, 40)",
        "57359 "." 41 (1, 42)-(1, 42)",
        "57356 "}" 42 (1, 43)-(1, 43)"
    ],
    "sexp": "(select \"SELECT\" (list \"SELECT\") (list \"?pred\" (var \"?pred\") (var \"?obj\")) (list \"?sub\" (restriction \"?sub\" (var \"?sub\") (var \"?pred\") (var \"?obj\"))))"
}
```

**Note:** This may provide useful insights while debugging.

## WebSocket

There is a dotfile [here](./etc/.rdf-client_history). Copy it to `~/.rdf-client_history`.

Build and run client:
```shell
innose2019-rdf-server/src$ make rdf-client
innose2019-rdf-server/src$ ./rdf-client 127.0.0.1 8001
```

Use the up arrow to cycle through the commands in the history file. Press enter to sent the current command. Replies will be buffered and the buffer flushed to the screen when enter is pressed.
