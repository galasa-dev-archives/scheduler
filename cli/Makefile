all: bin/galasasch-linux-amd64 bin/galasasch-windows-amd64 bin/galasasch-darwin-amd64 bin/galasasch-linux-s390x


bin/galasasch-linux-amd64 : ./cmd/galasasch/main.go ./pkg/cmd/*.go ./pkg/app/*.go ./pkg/status/*.go ./pkg/util/*.go ./pkg/scheduleryaml/*.go ./pkg/openapi/*.go  
	GOOS=linux GOARCH=amd64 go build -o bin/galasasch-linux-amd64 ./cmd/galasasch

bin/galasasch-windows-amd64 : ./cmd/galasasch/main.go ./pkg/cmd/*.go  ./pkg/app/*.go ./pkg/status/*.go ./pkg/util/*.go ./pkg/scheduleryaml/*.go ./pkg/openapi/*.go  
	GOOS=windows GOARCH=amd64 go build -o bin/galasasch-windows-amd64 ./cmd/galasasch

bin/galasasch-darwin-amd64 : ./cmd/galasasch/main.go ./pkg/cmd/*.go  ./pkg/app/*.go ./pkg/status/*.go ./pkg/util/*.go ./pkg/scheduleryaml/*.go ./pkg/openapi/*.go  
	GOOS=darwin GOARCH=amd64 go build -o bin/galasasch-darwin-amd64 ./cmd/galasasch

bin/galasasch-linux-s390x : ./cmd/galasasch/main.go ./pkg/cmd/*.go  ./pkg/app/*.go ./pkg/status/*.go ./pkg/util/*.go ./pkg/scheduleryaml/*.go ./pkg/openapi/*.go  
	GOOS=linux GOARCH=s390x go build -o bin/galasasch-linux-s390x ./cmd/galasasch


clean:
	rm -rf bin
