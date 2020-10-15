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

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 #comment
?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 48 (2, 1)-(2, 5)
 - 57352 "?var3" 54 (2, 7)-(2, 11)
 - 57361 "." 60 (2, 13)-(2, 13)
 - 57358 "}" 62 (2, 15)-(2, 15)
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

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2) ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57367 "(" 39 (1, 40)-(1, 40)
 - 57352 "?var2" 40 (1, 41)-(1, 45)
 - 57368 ")" 45 (1, 46)-(1, 46)
 - 57352 "?var3" 47 (1, 48)-(1, 52)
 - 57361 "." 53 (1, 54)-(1, 54)
 - 57358 "}" 55 (1, 56)-(1, 56)
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

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2/?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57362 "/" 44 (1, 45)-(1, 45)
 - 57352 "?var2" 45 (1, 46)-(1, 50)
 - 57352 "?var3" 51 (1, 52)-(1, 56)
 - 57361 "." 57 (1, 58)-(1, 58)
 - 57358 "}" 59 (1, 60)-(1, 60)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (sequence "/" (var "?var2") (var "?var2")) (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2/?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2/?var2 ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2|?var3 ?var4 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57363 "|" 44 (1, 45)-(1, 45)
 - 57352 "?var3" 45 (1, 46)-(1, 50)
 - 57352 "?var4" 51 (1, 52)-(1, 56)
 - 57361 "." 57 (1, 58)-(1, 58)
 - 57358 "}" 59 (1, 60)-(1, 60)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (choice "|" (var "?var2") (var "?var3")) (var "?var4")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3) ?var4 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3) ?var4 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2|?var3) ?var4 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57367 "(" 39 (1, 40)-(1, 40)
 - 57352 "?var2" 40 (1, 41)-(1, 45)
 - 57363 "|" 45 (1, 46)-(1, 46)
 - 57352 "?var3" 46 (1, 47)-(1, 51)
 - 57368 ")" 51 (1, 52)-(1, 52)
 - 57352 "?var4" 53 (1, 54)-(1, 58)
 - 57361 "." 59 (1, 60)-(1, 60)
 - 57358 "}" 61 (1, 62)-(1, 62)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (choice "|" (var "?var2") (var "?var3")) (var "?var4")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3) ?var4 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3) ?var4 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2+ ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57364 "+" 44 (1, 45)-(1, 45)
 - 57352 "?var3" 46 (1, 47)-(1, 51)
 - 57361 "." 52 (1, 53)-(1, 53)
 - 57358 "}" 54 (1, 55)-(1, 55)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (one-or-more "+" (var "?var2")) (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2+ ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2+ ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2++ ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57364 "+" 44 (1, 45)-(1, 45)
 - 57364 "+" 45 (1, 46)-(1, 46)
 - 57352 "?var3" 47 (1, 48)-(1, 52)
 - 57361 "." 53 (1, 54)-(1, 54)
 - 57358 "}" 55 (1, 56)-(1, 56)
[PARSE] Error parsing: syntax error

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2* ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57352 "?var2" 39 (1, 40)-(1, 44)
 - 57365 "*" 44 (1, 45)-(1, 45)
 - 57352 "?var3" 46 (1, 47)-(1, 51)
 - 57361 "." 52 (1, 53)-(1, 53)
 - 57358 "}" 54 (1, 55)-(1, 55)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (zero-or-more "*" (var "?var2")) (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2* ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2* ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2|?var3)/?var4 ?var5 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57352 "?var1" 33 (1, 34)-(1, 38)
 - 57367 "(" 39 (1, 40)-(1, 40)
 - 57352 "?var2" 40 (1, 41)-(1, 45)
 - 57363 "|" 45 (1, 46)-(1, 46)
 - 57352 "?var3" 46 (1, 47)-(1, 51)
 - 57368 ")" 51 (1, 52)-(1, 52)
 - 57362 "/" 52 (1, 53)-(1, 53)
 - 57352 "?var4" 53 (1, 54)-(1, 58)
 - 57352 "?var5" 59 (1, 60)-(1, 64)
 - 57361 "." 65 (1, 66)-(1, 66)
 - 57358 "}" 67 (1, 68)-(1, 68)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (sequence "/" (choice "|" (var "?var2") (var "?var3")) (var "?var4")) (var "?var5")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3)/?var4 ?var5 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 (?var2|?var3)/?var4 ?var5 .
}

Case: SELECT ?var1 ?var2 WHERE { { ?var1 ?var1 ?var2 } UNION { ?var1 ?var2 ?var2 } . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57357 "{" 27 (1, 28)-(1, 28)
 - 57352 "?var1" 29 (1, 30)-(1, 34)
 - 57352 "?var1" 35 (1, 36)-(1, 40)
 - 57352 "?var2" 41 (1, 42)-(1, 46)
 - 57358 "}" 47 (1, 48)-(1, 48)
 - 57348 "UNION" 49 (1, 50)-(1, 54)
 - 57357 "{" 55 (1, 56)-(1, 56)
 - 57352 "?var1" 57 (1, 58)-(1, 62)
 - 57352 "?var2" 63 (1, 64)-(1, 68)
 - 57352 "?var2" 69 (1, 70)-(1, 74)
 - 57358 "}" 75 (1, 76)-(1, 76)
 - 57361 "." 77 (1, 78)-(1, 78)
 - 57358 "}" 79 (1, 80)-(1, 80)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "{" (union "UNION" (restriction "?var1" (var "?var1") (var "?var1") (var "?var2")) (restriction "?var1" (var "?var1") (var "?var2") (var "?var2"))))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    {
        ?var1 ?var1 ?var2 .
    } UNION {
        ?var1 ?var2 ?var2 .
    } .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    {
        ?var1 ?var1 ?var2 .
    } UNION {
        ?var1 ?var2 ?var2 .
    } .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 < http://www.google.com#test > . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57359 "<" 39 (1, 40)-(1, 40)
 - 57353 "http://www.google.com#test" 41 (1, 42)-(1, 67)
 - 57360 ">" 68 (1, 69)-(1, 69)
 - 57361 "." 70 (1, 71)-(1, 71)
 - 57358 "}" 72 (1, 73)-(1, 73)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (uri "http://www.google.com#test")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 http://www.google.com#test .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 http://www.google.com#test .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 <http://www.google.com#test> . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57359 "<" 39 (1, 40)-(1, 40)
 - 57353 "http://www.google.com#test" 40 (1, 41)-(1, 66)
 - 57360 ">" 66 (1, 67)-(1, 67)
 - 57361 "." 68 (1, 69)-(1, 69)
 - 57358 "}" 70 (1, 71)-(1, 71)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (uri "http://www.google.com#test")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 http://www.google.com#test .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 http://www.google.com#test .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 "" . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57355 "\"\"" 39 (1, 40)-(1, 41)
 - 57361 "." 42 (1, 43)-(1, 43)
 - 57358 "}" 44 (1, 45)-(1, 45)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (string "\"\"")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "" .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "" .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 "a b" . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57355 "\"a b\"" 39 (1, 40)-(1, 44)
 - 57361 "." 45 (1, 46)-(1, 46)
 - 57358 "}" 47 (1, 48)-(1, 48)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (string "\"a b\"")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "a b" .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "a b" .
}

Case: SELECT ?var1 ?var2 WHERE { ?var1 ?var2 "a \" b" . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57347 "WHERE" 19 (1, 20)-(1, 24)
 - 57357 "{" 25 (1, 26)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57355 "\"a \\\" b\"" 39 (1, 40)-(1, 47)
 - 57361 "." 48 (1, 49)-(1, 49)
 - 57358 "}" 50 (1, 51)-(1, 51)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (string "\"a \\\" b\"")))))
[NORM]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "a \" b" .
}
[RESPARQL]
SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 "a \" b" .
}

Case: PREFIX a : <http://b> SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 9 (1, 10)-(1, 10)
 - 57359 "<" 11 (1, 12)-(1, 12)
 - 57353 "http://b" 12 (1, 13)-(1, 20)
 - 57360 ">" 20 (1, 21)-(1, 21)
 - 57346 "SELECT" 22 (1, 23)-(1, 28)
 - 57352 "?var1" 29 (1, 30)-(1, 34)
 - 57352 "?var2" 35 (1, 36)-(1, 40)
 - 57347 "WHERE" 41 (1, 42)-(1, 46)
 - 57357 "{" 47 (1, 48)-(1, 48)
 - 57352 "?var1" 49 (1, 50)-(1, 54)
 - 57352 "?var2" 55 (1, 56)-(1, 60)
 - 57352 "?var3" 61 (1, 62)-(1, 66)
 - 57361 "." 67 (1, 68)-(1, 68)
 - 57358 "}" 69 (1, 70)-(1, 70)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "PREFIX") (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57346 "SELECT" 20 (1, 21)-(1, 26)
 - 57352 "?var1" 27 (1, 28)-(1, 32)
 - 57352 "?var2" 33 (1, 34)-(1, 38)
 - 57347 "WHERE" 39 (1, 40)-(1, 44)
 - 57357 "{" 45 (1, 46)-(1, 46)
 - 57352 "?var1" 47 (1, 48)-(1, 52)
 - 57352 "?var2" 53 (1, 54)-(1, 58)
 - 57352 "?var3" 59 (1, 60)-(1, 64)
 - 57361 "." 65 (1, 66)-(1, 66)
 - 57358 "}" 67 (1, 68)-(1, 68)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "PREFIX") (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> PREFIX c:<http://d> SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57351 "PREFIX" 20 (1, 21)-(1, 26)
 - 57354 "c" 27 (1, 28)-(1, 28)
 - 57366 ":" 28 (1, 29)-(1, 29)
 - 57359 "<" 29 (1, 30)-(1, 30)
 - 57353 "http://d" 30 (1, 31)-(1, 38)
 - 57360 ">" 38 (1, 39)-(1, 39)
 - 57346 "SELECT" 40 (1, 41)-(1, 46)
 - 57352 "?var1" 47 (1, 48)-(1, 52)
 - 57352 "?var2" 53 (1, 54)-(1, 58)
 - 57347 "WHERE" 59 (1, 60)-(1, 64)
 - 57357 "{" 65 (1, 66)-(1, 66)
 - 57352 "?var1" 67 (1, 68)-(1, 72)
 - 57352 "?var2" 73 (1, 74)-(1, 78)
 - 57352 "?var3" 79 (1, 80)-(1, 84)
 - 57361 "." 85 (1, 86)-(1, 86)
 - 57358 "}" 87 (1, 88)-(1, 88)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b")) (prefix "PREFIX" (id "c") (uri "http://d"))) (list "PREFIX") (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>
PREFIX c: <http://d>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>
PREFIX c: <http://d>

SELECT ?var1 ?var2
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: SELECT ?var1 ?var2 ?var3 WHERE { a:b ?var2 ?var3 . }
[TOKENS]
 - 57346 "SELECT" 0 (1, 1)-(1, 6)
 - 57352 "?var1" 7 (1, 8)-(1, 12)
 - 57352 "?var2" 13 (1, 14)-(1, 18)
 - 57352 "?var3" 19 (1, 20)-(1, 24)
 - 57347 "WHERE" 25 (1, 26)-(1, 30)
 - 57357 "{" 31 (1, 32)-(1, 32)
 - 57354 "a" 33 (1, 34)-(1, 34)
 - 57366 ":" 34 (1, 35)-(1, 35)
 - 57354 "b" 35 (1, 36)-(1, 36)
 - 57352 "?var2" 37 (1, 38)-(1, 42)
 - 57352 "?var3" 43 (1, 44)-(1, 48)
 - 57361 "." 49 (1, 50)-(1, 50)
 - 57358 "}" 51 (1, 52)-(1, 52)
[PARSE] (query "SELECT" (list "SELECT") (list "SELECT") (list "SELECT") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "a" (restriction "a" (prefixed ":" (id "a") (id "b")) (var "?var2") (var "?var3")))))
[NORM]
SELECT ?var1 ?var2 ?var3
WHERE {
    a:b ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    a:b ?var2 ?var3 .
}

Case: DATA ?var1 ?var3 SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57349 "DATA" 0 (1, 1)-(1, 4)
 - 57352 "?var1" 5 (1, 6)-(1, 10)
 - 57352 "?var3" 11 (1, 12)-(1, 16)
 - 57346 "SELECT" 17 (1, 18)-(1, 23)
 - 57352 "?var1" 24 (1, 25)-(1, 29)
 - 57352 "?var2" 30 (1, 31)-(1, 35)
 - 57352 "?var3" 36 (1, 37)-(1, 41)
 - 57347 "WHERE" 42 (1, 43)-(1, 47)
 - 57357 "{" 48 (1, 49)-(1, 49)
 - 57352 "?var1" 50 (1, 51)-(1, 55)
 - 57352 "?var2" 56 (1, 57)-(1, 61)
 - 57352 "?var3" 62 (1, 63)-(1, 67)
 - 57361 "." 68 (1, 69)-(1, 69)
 - 57358 "}" 70 (1, 71)-(1, 71)
[PARSE] (query "DATA" (list "DATA") (list "?var1" (var "?var1") (var "?var3")) (list "DATA") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
DATA
    ?var1
    ?var3
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> DATA ?var1 ?var3 SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57349 "DATA" 20 (1, 21)-(1, 24)
 - 57352 "?var1" 25 (1, 26)-(1, 30)
 - 57352 "?var3" 31 (1, 32)-(1, 36)
 - 57346 "SELECT" 37 (1, 38)-(1, 43)
 - 57352 "?var1" 44 (1, 45)-(1, 49)
 - 57352 "?var2" 50 (1, 51)-(1, 55)
 - 57352 "?var3" 56 (1, 57)-(1, 61)
 - 57347 "WHERE" 62 (1, 63)-(1, 67)
 - 57357 "{" 68 (1, 69)-(1, 69)
 - 57352 "?var1" 70 (1, 71)-(1, 75)
 - 57352 "?var2" 76 (1, 77)-(1, 81)
 - 57352 "?var3" 82 (1, 83)-(1, 87)
 - 57361 "." 88 (1, 89)-(1, 89)
 - 57358 "}" 90 (1, 91)-(1, 91)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "?var1" (var "?var1") (var "?var3")) (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

DATA
    ?var1
    ?var3
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> DATA ?var1 ?var3 SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57349 "DATA" 20 (1, 21)-(1, 24)
 - 57352 "?var1" 25 (1, 26)-(1, 30)
 - 57352 "?var3" 31 (1, 32)-(1, 36)
 - 57346 "SELECT" 37 (1, 38)-(1, 43)
 - 57352 "?var1" 44 (1, 45)-(1, 49)
 - 57352 "?var2" 50 (1, 51)-(1, 55)
 - 57352 "?var3" 56 (1, 57)-(1, 61)
 - 57347 "WHERE" 62 (1, 63)-(1, 67)
 - 57357 "{" 68 (1, 69)-(1, 69)
 - 57352 "?var1" 70 (1, 71)-(1, 75)
 - 57352 "?var2" 76 (1, 77)-(1, 81)
 - 57352 "?var3" 82 (1, 83)-(1, 87)
 - 57361 "." 88 (1, 89)-(1, 89)
 - 57358 "}" 90 (1, 91)-(1, 91)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "?var1" (var "?var1") (var "?var3")) (list "PREFIX") (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

DATA
    ?var1
    ?var3
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}

Case: PREFIX a:<http://b> DATA ?var1 ?var3 UNITS mod:temp->unit:degc mod:dist->unit:m SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }
[TOKENS]
 - 57351 "PREFIX" 0 (1, 1)-(1, 6)
 - 57354 "a" 7 (1, 8)-(1, 8)
 - 57366 ":" 8 (1, 9)-(1, 9)
 - 57359 "<" 9 (1, 10)-(1, 10)
 - 57353 "http://b" 10 (1, 11)-(1, 18)
 - 57360 ">" 18 (1, 19)-(1, 19)
 - 57349 "DATA" 20 (1, 21)-(1, 24)
 - 57352 "?var1" 25 (1, 26)-(1, 30)
 - 57352 "?var3" 31 (1, 32)-(1, 36)
 - 57350 "UNITS" 37 (1, 38)-(1, 42)
 - 57354 "mod" 43 (1, 44)-(1, 46)
 - 57366 ":" 46 (1, 47)-(1, 47)
 - 57354 "temp" 47 (1, 48)-(1, 51)
 - 57356 "->" 51 (1, 52)-(1, 53)
 - 57354 "unit" 53 (1, 54)-(1, 57)
 - 57366 ":" 57 (1, 58)-(1, 58)
 - 57354 "degc" 58 (1, 59)-(1, 62)
 - 57354 "mod" 63 (1, 64)-(1, 66)
 - 57366 ":" 66 (1, 67)-(1, 67)
 - 57354 "dist" 67 (1, 68)-(1, 71)
 - 57356 "->" 71 (1, 72)-(1, 73)
 - 57354 "unit" 73 (1, 74)-(1, 77)
 - 57366 ":" 77 (1, 78)-(1, 78)
 - 57354 "m" 78 (1, 79)-(1, 79)
 - 57346 "SELECT" 80 (1, 81)-(1, 86)
 - 57352 "?var1" 87 (1, 88)-(1, 92)
 - 57352 "?var2" 93 (1, 94)-(1, 98)
 - 57352 "?var3" 99 (1, 100)-(1, 104)
 - 57347 "WHERE" 105 (1, 106)-(1, 110)
 - 57357 "{" 111 (1, 112)-(1, 112)
 - 57352 "?var1" 113 (1, 114)-(1, 118)
 - 57352 "?var2" 119 (1, 120)-(1, 124)
 - 57352 "?var3" 125 (1, 126)-(1, 130)
 - 57361 "." 131 (1, 132)-(1, 132)
 - 57358 "}" 133 (1, 134)-(1, 134)
[PARSE] (query "PREFIX" (list "PREFIX" (prefix "PREFIX" (id "a") (uri "http://b"))) (list "?var1" (var "?var1") (var "?var3")) (list "mod" (mapping "->" (prefixed ":" (id "mod") (id "temp")) (prefixed ":" (id "unit") (id "degc"))) (mapping "->" (prefixed ":" (id "mod") (id "dist")) (prefixed ":" (id "unit") (id "m")))) (select "SELECT" (list "?var1" (var "?var1") (var "?var2") (var "?var3")) (list "?var1" (restriction "?var1" (var "?var1") (var "?var2") (var "?var3")))))
[NORM]
PREFIX a: <http://b>

DATA
    ?var1
    ?var3
UNITS
    mod:temp -> unit:degc
    mod:dist -> unit:m
SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
[RESPARQL]
PREFIX a: <http://b>

SELECT ?var1 ?var2 ?var3
WHERE {
    ?var1 ?var2 ?var3 .
}
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
