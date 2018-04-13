.PHONY: kowalad all test clean
.PHONY: kowalad-static kowalad-dynamic

GOBIN=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))build/bin
GOLIB=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))build/lib

kowalad:
	go build -i -o $(GOBIN)/kowalad -v -tags '$(BUILD_TAGS)' $(BUILD_FLAGS) ./cmd/kowalad
	@echo "Compilation is over."
	@echo "Run \"build/bin/kowalad -h\" to view available commands."

kowalad-static:
	go build -buildmode=c-shared -gcflags=-shared -asmflags=-shared -installsuffix=_shared -a -o $(GOLIB)/oracle.a ./kcoin/cplusplus
	@echo "Static library generation is over."
	
kowalad-dynamic: 
	go build -buildmode=c-archive -gcflags=-shared -asmflags=-shared -installsuffix=_shared -a -o $(GOLIB)/oracle.so ./kcoin/cplusplus
	@echo "Dynamic library generation is over."

clean:
	rm -rf build
