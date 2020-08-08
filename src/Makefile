TARGETS = \
	rdf-client \
	rdf-server \


all: ${TARGETS} ../var/model
	PYTHONPATH=. ./rdf-server 0.0.0.0 8001 ../var/model ../var/ontologies

clean:
	touch ${TARGETS}
	rm    ${TARGETS}

mrproper:
	touch d~ __pycache__
	rm -R *~ __pycache__

deps:
	go get github.com/sbinet/go-python
	go get github.com/gorilla/websocket
	go get github.com/peterh/liner


../var/model:
	mkdir ../var/model

rdf-client: rdf-client.go
	go build -x rdf-client.go

rdf-server: rdf-server.go
	go build -x rdf-server.go

