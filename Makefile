TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=localhost
NAMESPACE=edu
NAME=authz
BINARY=terraform-provider-${NAME}
VERSION=0.3.1
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	mv ${BINARY} ~/.terraform.d/plugins/

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

start-authz:
	docker-compose -f docker_compose/docker-compose.yml up -d

stop-authz:
	docker-compose -f docker_compose/docker-compose.yml down

create-authz-sa:
	./create_authz_service_account.sh

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   
