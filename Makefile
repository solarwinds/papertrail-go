lint:
	go get -u github.com/mgechev/revive; $(GOPATH)/bin/revive -config revive.toml

error_check:
	go get -u github.com/kisielk/errcheck; $(GOPATH)/bin/errcheck ./...

static_check:
	go get -u honnef.co/go/tools/cmd/staticcheck; $(GOPATH)/bin/staticcheck -checks all ./...

vet:
	go vet ./...

sec_check:
	go get github.com/securego/gosec/cmd/gosec; $(GOPATH)/bin/gosec ./...

tests:
	go test -v ./...