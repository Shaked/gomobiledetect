packages: 
	go get code.google.com/p/go.tools/cmd/cover
	go get github.com/axw/gocov/gocov
	go get gopkg.in/matm/v1/gocov-html
	go get github.com/modocache/gover
	go get github.com/mattn/goveralls

test: packages
		go test -v 

cover: packages
	rm -rf ./cover.*
	touch cover.json
	gocov test . -v >> cover.json; 
	gocov-html cover.json >> cover.html; 

travis: packages
	rm -rf gover.coverprofile
	rm -rf profile.cov
	go test -covermode=count -coverprofile=profile.cov .; \
	$(HOME)/gopath/bin/gover 
	$(HOME)/gopath/bin/goveralls -repotoken Rcel00b0hm7Bd3oy5x2XQvhVNY0Eckbry -coverprofile=gover.coverprofile -service travis-ci


doc:
	godoc . 1> manual.txt
