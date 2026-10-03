SOLC_IMAGE   := ethereum/solc:0.8.30
GETH_VERSION := v1.17.7
OZ_VERSION   := 5.7.0
OZ_DIR       := lib/openzeppelin-contracts
CONTRACTS    := Storage EthSender TokenERC20 MyNFT MyERC1155Token

## test: run the contract and chain tests on go-ethereum's simulated backend
test:
	go test -race ./...

## bindings: compile the contracts and regenerate the Go bindings
bindings: $(OZ_DIR)
	rm -rf build && mkdir -p build bindings
	docker run --rm -u $$(id -u):$$(id -g) -v $(CURDIR):/src -w /src $(SOLC_IMAGE) \
		@openzeppelin/contracts/=$(OZ_DIR)/ --optimize --evm-version cancun \
		--abi --bin --overwrite -o build contracts/*.sol
	for c in $(CONTRACTS); do \
		go run github.com/ethereum/go-ethereum/cmd/abigen@$(GETH_VERSION) \
			--abi build/$$c.abi --bin build/$$c.bin --pkg bindings --type $$c \
			--out bindings/$$(echo $$c | tr A-Z a-z).go || exit 1; \
	done

# OpenZeppelin from the npm registry, checked against the registry's SHA-512.
$(OZ_DIR):
	mkdir -p $(OZ_DIR)
	curl -sfL -o lib/oz.tgz https://registry.npmjs.org/@openzeppelin/contracts/-/contracts-$(OZ_VERSION).tgz
	want=$$(curl -sf https://registry.npmjs.org/@openzeppelin/contracts | python3 -c 'import sys,json; print(json.load(sys.stdin)["versions"]["$(OZ_VERSION)"]["dist"]["integrity"].split("-",1)[1])'); \
	got=$$(openssl dgst -sha512 -binary lib/oz.tgz | base64 -w0); \
	[ "$$want" = "$$got" ] || { echo "OpenZeppelin checksum mismatch"; rm -rf lib; exit 1; }
	tar -xzf lib/oz.tgz -C $(OZ_DIR) --strip-components=1 && rm lib/oz.tgz

## devchain: run a local Anvil chain on :8545 (prints public test accounts)
devchain:
	docker run --rm -p 8545:8545 ghcr.io/foundry-rs/foundry:latest "anvil --host 0.0.0.0"

.PHONY: test bindings devchain
