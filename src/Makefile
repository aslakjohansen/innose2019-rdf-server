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
	# embedding python
	go get github.com/sbinet/go-python
	
	# websocket
	go get github.com/gorilla/websocket
	
	# command line interpreter
	go get github.com/peterh/liner
	
	# parser generation
	go get -u golang.org/x/tools/cmd/goyacc
	go get github.com/timtadh/lexmachine
	
	# mqtt
	go get github.com/eclipse/paho.mqtt.golang
	go get github.com/gorilla/websocket
	go get golang.org/x/net/proxy



sparql/y.go: sparql/lang.y
	cd sparql ; goyacc lang.y

../var/model:
	mkdir ../var/model

rdf-client: rdf-client.go
	go build -x rdf-client.go

rdf-server: rdf-server.go */*.go */*/*.go */*/*/*.go
	go build -x rdf-server.go

