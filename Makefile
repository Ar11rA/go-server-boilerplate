BINARY=engine
test: 
	go test -v ./...

package:
	@dep ensure -v

build: package
	go build -o ${BINARY}

install: 
	go build -o ${BINARY}

test:
	go test -short $$(go list ./... | grep -v /vendor/)

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-srv .
