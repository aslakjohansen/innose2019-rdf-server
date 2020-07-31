TARGETS = \
	rdf-server \


all: ${TARGETS}
	PYTHONPATH=. ./rdf-server

clean:
	rm ${TARGETS}

deps:
	go get github.com/sbinet/go-python


rdf-server: rdf-server.go
	go build -x rdf-server.go

