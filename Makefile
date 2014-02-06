COVERAGE_DIRECTORY = /tmp/coverage/
test: format 
		go test -v 
coverage: format 
		mkdir -p $(COVERAGE_DIRECTORY)
		go test -coverprofile=$(COVERAGE_DIRECTORY)coverage.out && go tool cover -html=$(COVERAGE_DIRECTORY)coverage.out -o $(COVERAGE_DIRECTORY)coverage.html 

format: 
	go fmt

clean: 
		rm -rf ${COVERAGE_DIRECTORY}