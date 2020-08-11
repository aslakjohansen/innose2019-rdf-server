TARGETS = \
	parse-sparql \


all: ${TARGETS}

clean:
	rm ${TARGETS}


../src/sparql/y.go: ../src/sparql/lang.y
	cd ../src/sparql ; goyacc lang.y

parse-sparql: parse-sparql.go ../src/sparql/y.go
	go run parse-sparql.go

