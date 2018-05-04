.PHONY: kowalad cplusplus

Library_Name = kowalad.so

kowalad:
		go build -i -o $@ -v -tags '$(BUILD_TAGS)' $(BUILD_FLAGS) ./cmd/kowalad
		@echo "Compilation is over."
		@echo "Run \"build/bin/kowalad -h\" to view available commands."

$(Library_Name): 
		go build -buildmode=c-archive -gcflags=-shared -asmflags=-shared -installsuffix=_shared -a -o $@ ./cplusplus
		@echo "Dynamic library generation is over."

cplusplus: $(Library_Name)


