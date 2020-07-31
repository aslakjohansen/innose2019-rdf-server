TARGETS = \
	rdf-server \


all: ${TARGETS}
	PYTHONPATH=. ./rdf-server

clean:
	touch ${TARGETS}
	rm    ${TARGETS}

mrproper:
	touch d~ __pycache__
	rm -R *~ __pycache__

deps:
	go get github.com/sbinet/go-python


rdf-server: rdf-server.go logic.go
	go build -x rdf-server.go logic.go

