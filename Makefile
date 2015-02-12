packages: 
	go get github.com/gorilla/context
	go get golang.org/x/tools/cmd/cover 
	go get github.com/axw/gocov/gocov
	go get gopkg.in/matm/v1/gocov-html
	go get github.com/modocache/gover
	go get github.com/mattn/goveralls

test: 
	go test -v 

bench:
	go test -bench=.

cover: packages
	rm -rf ./cover.*
	touch cover.json
	gocov test . -v >> cover.json; 
	gocov-html cover.json >> cover.html; 

doc:
	godoc . 1> manual.txt
