# Interface

## Fetch time from server

```shell
$ curl -X PUT -d '42' http://localhost:8001/time
{
    "success": true,
    "time": 1589391233.535275
}
```

**Note:** The `42` payload is not used, but the payload needs to be valid json.

## Store model to disk

```shell
$ curl -X PUT -d '42' http://localhost:8001/store
{
    "filename": "../var/model/3.ttl",
    "success": true
}
```

**Note:** The `42` payload is not used, but the payload needs to be valid json.

## Fetch the current namespace mapping

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

## Resolving a query

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
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2646"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb906"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb769"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2394"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1822"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1487"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2730"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb508"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1920"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2398"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb313"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1188"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1795"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1395"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb513"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1608"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2666"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "https://brickschema.org/schema/1.1.0/Brick#Point"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1238"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb721"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb453"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb491"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2541"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb303"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb648"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb4"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2063"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1898"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1988"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb755"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb771"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1776"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1229"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb965"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb484"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb709"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2405"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1227"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb714"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1251"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb210"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2304"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2628"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1834"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2805"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1143"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2522"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2708"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb232"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1273"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2371"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb242"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb149"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb960"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb146"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2768"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2036"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2813"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1364"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2602"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1787"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb903"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2451"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2210"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb416"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2348"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1163"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2264"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb941"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb60"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1739"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2456"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb737"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1629"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2562"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb138"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb101"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb371"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1235"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb752"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2566"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2690"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1656"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1811"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1473"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb579"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1755"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1430"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2656"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1259"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1764"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb398"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb954"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb236"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2625"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2103"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb599"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb37"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2192"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2196"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb922"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2309"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb687"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2424"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1401"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2680"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb990"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2270"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2039"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2512"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2693"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2239"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#label",
            "Sensor"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb538"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2787"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb596"
        ],
        [
            "http://www.w3.org/2002/07/owl#sameAs",
            "https://brickschema.org/schema/1.1.0/Brick#Sensor"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2373"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1051"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2777"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1029"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "http://www.w3.org/2002/07/owl#Thing"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1858"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2756"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1367"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1529"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1150"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1095"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1625"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1865"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb701"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2751"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1910"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1994"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb520"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2669"
        ],
        [
            "http://www.w3.org/1999/02/22-rdf-syntax-ns#type",
            "http://www.w3.org/2002/07/owl#Class"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2631"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2275"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1622"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2356"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb611"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1412"
        ],
        [
            "http://www.w3.org/2002/07/owl#equivalentClass",
            "f7384c45206bb42dba21f93c452e1971eb4"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1096"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1337"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1585"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2091"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1175"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1718"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb271"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb2261"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb175"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb660"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb705"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1278"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1547"
        ],
        [
            "http://www.w3.org/2002/07/owl#equivalentClass",
            "https://brickschema.org/schema/1.1.0/Brick#Sensor"
        ],
        [
            "http://www.w3.org/2000/01/rdf-schema#subClassOf",
            "f7384c45206bb42dba21f93c452e1971eb1480"
        ],
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

## Perform an update

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

## Inspect a query

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
