# Testing

## Send valid data over MQTT

```shell
$ mosquitto_pub -t "test" -m "{\"time\": 1234.5678, \"value\": 42.56}"
```

## Send invalid data over MQTT

```shell
$ mosquitto_pub -t "test" -m "blah"
```

## SparQL Parsing

```shell
$ cd test
$ make parse-sparql
go run parse-sparql.go
Case: garbage
[TOKENS]
 - 57354 "garbage" 0 (1, 1)-(1, 7)
[PARSE] Error parsing: syntax error

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57352 "?var3" 45 (1, 46)-(1, 50)
 - 57361 "." 51 (1, 52)-(1, 52)
 - 57358 "}" 53 (1, 54)-(1, 54)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 #
?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 41 (2, 1)-(2, 5)
 - 57352 "?var3" 47 (2, 7)-(2, 11)
 - 57361 "." 53 (2, 13)-(2, 13)
 - 57358 "}" 55 (2, 15)-(2, 15)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
...
```

### Data Dispatching

```shell
$ cd test
$ make data-dispatch-registration
go run data-dispatch-registration.go
Registering:
~~~~~~~~~~~
Dispatcher
 - a: 4 entries
 - b: 3 entries
 - c: 2 entries
 - d: 1 entries

Unregistering (not registered):
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Dispatcher
 - c: 2 entries
 - d: 1 entries
 - a: 4 entries
 - b: 3 entries

Unregistering (beginning):
~~~~~~~~~~~~~~~~~~~~~~~~~~
Dispatcher
 - b: 3 entries
 - c: 2 entries
 - d: 1 entries
 - a: 3 entries

Unregistering (middle):
~~~~~~~~~~~~~~~~~~~~~~~
Dispatcher
 - a: 2 entries
 - b: 3 entries
 - c: 2 entries
 - d: 1 entries

Unregistering (end):
~~~~~~~~~~~~~~~~~~~~
Dispatcher
 - c: 2 entries
 - d: 1 entries
 - a: 2 entries
 - b: 2 entries
```

## Subscription Management

```shell
aslak@thera:~/vcs/git/innose2019-rdf-server/src$ ./rdf-client 127.0.0.1 8001
Connecting to ws://127.0.0.1:8001/websocket
>> {"command": "subscribe", "id": "2", "query": "SELECT ?obj WHERE {brick:Sensor rdfs:subClassOf ?obj .}"}
{
    "id": "2",
    "response": "subscribed"
}
>> {"command": "subscribe", "id": "3", "query": "SELECT ?obj WHERE {brick:Sensor rdfs:subClassOf ?obj .}"}
{
    "id": "3",
    "response": "subscribed"
}
>> {"command": "subscribe", "id": "4", "query": "SELECT ?obj WHERE {brick:Sensor rdfs:subClassOf ?obj .}"}
{
    "id": "4",
    "response": "subscribed"
}
>> {"command": "subscriptions", "id": "5"}
{
    "id": "5",
    "subscriptions": [
    "4",
    "2",
    "3"
    ]
}
>> {"command": "unsubscribe", "id": "6", "subscription": "3"}
{
    "id": "6",
    "response": "unsubscribed"
}
>> {"command": "subscriptions", "id": "7"}
{
    "id": "7",
    "subscriptions": [
    "4",
    "2"
    ]
}
>> 
```
