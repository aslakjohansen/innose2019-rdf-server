TARGETS = \
	parse-sparql \
	data-dispatch-registration \
	subscription \


all: ${TARGETS}

clean:
	rm ${TARGETS}


../src/sparql/y.go: ../src/sparql/lang.y
	cd ../src/sparql ; goyacc lang.y

parse-sparql: parse-sparql.go ../src/sparql/y.go
	go run parse-sparql.go

data-dispatch-registration: data-dispatch-registration.go
	go run data-dispatch-registration.go

subscription: subscription.go
	go run subscription.go

