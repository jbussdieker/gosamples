.PHONY: all test install

all: install
install: build
	go install -v libxml libxml/xml libxml/html libxml/element libxml/node
build: test
	go build -v libxml libxml/xml libxml/html libxml/element libxml/node
test:
	#go test -v libxml libxml/xml libxml/html
	go test -v libxml/xml

