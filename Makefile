.PHONY: clean test build deploy-dev 

default: clean test build deploy-dev

clean:
	@echo "# Run clean"
	rm -rf ./bin ./vendor Gopkg.lock .serverless

test:
	@echo "# Run tests"
	@echo "## Clear testcache"
	go clean -testcache

	@echo "## Do test"
	@for fkt in $(shell find functions -name \*main.go | awk -F"/" '{print $$2}') ; do \
		@echo "### Test functions/$$fkt" ; \
		go test lambda-basic-auth-golang/functions/$$fkt -ginkgo.failFast=true ; \
	done


build:
	@echo "# Run build"
	@make clean

	@echo "## Build GO handler"
	@for fkt in $(shell find functions -name \*main.go | awk -F"/" '{print $$2}') ; do \
		echo "### Build bin functions/$$fkt/main.go -> bin/$$fkt" ; \
		env GOOS=linux go build -ldflags="-s -w" -o bin/$$fkt functions/$$fkt/main.go ; \
	done

deploy-dev:
	@echo "# Deploy on stage dev"
	serverless deploy --verbose --stage dev

deploy-live:
	@echo "# Deploy on stage live"
	serverless deploy --stage live

undeploy:
	@echo "# Undeploy deployment"
	@make clean
	serverless remove -v