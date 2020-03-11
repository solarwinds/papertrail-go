setup-tools:
	GO111MODULE=off go get -u github.com/mgechev/revive;
	GO111MODULE=off go get -u github.com/kisielk/errcheck;
	GO111MODULE=off go get -u honnef.co/go/tools/cmd/staticcheck;
	GO111MODULE=off go get -u github.com/securego/gosec/cmd/gosec

lint:
	$(GOPATH)/bin/revive -config revive.toml

error_check:
	$(GOPATH)/bin/errcheck ./...

static_check:
	$(GOPATH)/bin/staticcheck -checks all ./...

vet:
	go vet ./...

sec_check:
	$(GOPATH)/bin/gosec ./...

all_checks: vet lint error_check sec_check static_check

tests:
	go test -v ./...